package k8s_client

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func LaunchK8sJob(clientset *kubernetes.Clientset, jobName *string, image *string, cmd *string, namespace *string) {
	jobs := clientset.BatchV1().Jobs("default")
	var backOffLimit int32 = 0
	var ttlSecondsAfterFinished int32 = 10
	var containers []v1.Container = nil
	if len(*cmd) > 0 {
		containers = []v1.Container{
			{
				Name:    *jobName,
				Image:   *image,
				Command: strings.Split(*cmd, " "),
			},
		}
	} else {
		containers = []v1.Container{
			{
				Name:  *jobName,
				Image: *image,
			},
		}
	}

	// TODO: also add support for config, env, secrets, serviceaccounts
	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      *jobName,
			Namespace: *namespace,
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
		log.Fatalln("Failed to create K8s job.")
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
		log.Fatalln("failed to create K8s config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Failed to create K8s clientset")
	}

	return clientset
}
