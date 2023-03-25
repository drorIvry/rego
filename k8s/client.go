package k8s_client

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/drorivry/matter/dao"
	"github.com/drorivry/matter/models"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func BuildJobName(taskEx models.TaskExecution) string {
	jobName := taskEx.Name + "-" + taskEx.Image + "-" + strconv.FormatUint(uint64(taskEx.ID), 10)

	return jobName
}

func LaunchK8sJob(clientset *kubernetes.Clientset, jobName *string, taskEx *models.TaskExecution) {
	jobs := clientset.BatchV1().Jobs(taskEx.NameSpace)
	var backOffLimit int32 = 0
	var ttlSecondsAfterFinished int32 = int32(taskEx.TtlSecondsAfterFinished)
	var containers []v1.Container = nil
	if len(taskEx.Cmd) > 0 {
		containers = []v1.Container{
			{
				Name:    *jobName,
				Image:   taskEx.Image,
				Command: strings.Split(taskEx.Cmd, " "),
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

	// TODO: also add support for config, env, secrets, serviceaccounts
	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      *jobName,
			Namespace: taskEx.NameSpace,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers:    containers,
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
			BackoffLimit:            &backOffLimit,
			TTLSecondsAfterFinished: &ttlSecondsAfterFinished,
		},
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln("Failed to create K8s job. ", err)
	}

	//print job details
	log.Println("Created K8s job successfully")
}

func ConnectToK8s(kuneConfigPath *string) *kubernetes.Clientset {
	var configPath string = ""

	if *kuneConfigPath != "" {
		configPath = *kuneConfigPath
	} else {
		home, exists := os.LookupEnv("HOME")
		if !exists {
			home = "/root"
		}
		configPath = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatalln(
			"failed to create K8s config ",
			err,
		)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Failed to create K8s clientset ", err)
	}

	return clientset
}

func AbortTask(executionId uint) error {
	execution := dao.GetExecutionById(executionId)
	jobName := BuildJobName(*execution)
	deleteOptions := metav1.DeleteOptions{}

	jobs := initializers.ClientSet.BatchV1().Jobs(execution.NameSpace)
	err := jobs.Delete(context.TODO(), jobName, deleteOptions)
	if err != nil {
		log.Panic("Could not delete job ", jobName)
		return err
	}
	return nil
}
