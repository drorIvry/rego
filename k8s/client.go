package k8s_client

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/config"
	"github.com/drorivry/rego/dao"
	"github.com/drorivry/rego/models"
	"github.com/google/uuid"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

var ClientSet *kubernetes.Clientset

func BuildJobName(taskEx *models.TaskExecution) string {
	jobName := ""
	if taskEx.Name != "" {
		jobName += taskEx.Name + "-"
	}

	jobName += taskEx.Image + "-" + taskEx.ID.String()
	jobName = strings.Replace(jobName, ":", "-", -1)
	jobName = strings.Replace(jobName, ".", "-", -1)
	jobName = strings.Replace(jobName, "_", "-", -1)
	return jobName
}

func InitK8SClientSet(kubeConfigPath *string) {
	if config.IN_CLUSTER {
		ClientSet = ConnectToK8SInCluster()
	} else {
		ClientSet = ConnectToK8s(kubeConfigPath)
	}
}

func LaunchK8sJob(
	jobName *string,
	taskEx *models.TaskExecution,
) {
	var metadata batchv1.JobSpec
	if taskEx.Metadata != nil {
		metadataErr := json.Unmarshal([]byte(taskEx.Metadata), &metadata)
		if metadataErr != nil {
			log.Error().Err(metadataErr).Msg("Error parsing metadata")
			return
		}
	}
	jobs := ClientSet.BatchV1().Jobs(taskEx.Namespace)
	var containers []v1.Container = nil

	if len(taskEx.Cmd) > 0 {
		containers = []v1.Container{
			{
				Name:    *jobName,
				Image:   taskEx.Image,
				Command: taskEx.Cmd,
			},
		}
	} else {
		containers = []v1.Container{
			{
				Name:  *jobName,
				Image: taskEx.Image,
			},
		}
	}
	metadata.Template = v1.PodTemplateSpec{
		Spec: v1.PodSpec{
			Containers:    containers,
			RestartPolicy: v1.RestartPolicyNever,
		},
	}

	// TODO: also add support for config, env, secrets, serviceaccounts
	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      *jobName,
			Namespace: taskEx.Namespace,
		},
		Spec: metadata,
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create K8s job")
		return
	}

	//print job details
	log.Info().Str("execution_id", taskEx.ID.String()).Msg("Created K8s job successfully")
}

func ConnectToK8s(kubeConfigPath *string) *kubernetes.Clientset {
	var configPath string = ""

	if *kubeConfigPath != "" {
		configPath = *kubeConfigPath
	} else {
		home, exists := os.LookupEnv("HOME")
		if !exists {
			home = "/root"
		}
		configPath = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to create K8s config")
		return nil
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create K8s clientset")
		return nil
	}

	return clientset
}

func ConnectToK8SInCluster() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get K8s config")
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create K8s clientset")
	}
	return clientset

}

func AbortTask(executionId uuid.UUID) error {
	execution, err := dao.GetExecutionById(executionId)
	if err != nil {
		log.Error().Err(err).Msg("Execution id wasn't found")
		return err
	}
	jobName := BuildJobName(execution)
	deleteOptions := metav1.DeleteOptions{}
	var zero int64 = 0
	bg := metav1.DeletePropagationBackground
	deleteOptions.GracePeriodSeconds = &zero
	deleteOptions.PropagationPolicy = &bg

	jobs := ClientSet.BatchV1().Jobs(execution.Namespace)
	err = jobs.Delete(context.TODO(), jobName, deleteOptions)
	if err != nil {
		log.Error().Err(err).Str("job_name", jobName).Msg("Could not delete job")
		return err
	}
	return nil
}

func GetJobStatus(executionId uuid.UUID) (models.Status, error) {
	execution, err := dao.GetExecutionById(executionId)
	if err != nil {
		return models.PENDING, err
	}

	jobName := BuildJobName(execution)
	job, err := ClientSet.BatchV1().Jobs(
		execution.Namespace,
	).Get(
		context.TODO(),
		jobName,
		metav1.GetOptions{},
	)

	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return models.PROC_ERROR, err
		}
	}

	if job.Status.Active == 0 && job.Status.Succeeded == 0 && job.Status.Failed == 0 {
		return models.PENDING, nil
	}

	if job.Status.Active > 0 {
		return models.RUNNING, nil
	}

	if job.Status.Succeeded > 0 {
		return models.SUCCESS, nil // Job ran successfully
	}

	return models.APP_ERROR, nil
}
