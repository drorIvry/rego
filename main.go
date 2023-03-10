package main

import (
	"flag"
	"sync"

	k8s_client "github.com/drorivry/matter/k8s"
	"github.com/drorivry/matter/tasker"
	"github.com/gin-gonic/gin"
)

func runServer(server *gin.Engine) {
	server.Run()
}

func main() {
	jobName := flag.String("jobname", "test-job", "The name of the job")
	containerImage := flag.String("image", "ubuntu:latest", "Name of the container image")
	entryCommand := flag.String("command", "", "The command to run inside the container")
	namespace := flag.String("namespace", "default", "The job's namespace to deploy to")
	kuneConfigPath := flag.String("kubeConfigPath", "", "The path to the kubeconfig")

	var wg sync.WaitGroup

	flag.Parse()

	server := tasker.GetServer()
	wg.Add(1)
	go runServer(server)

	clientset := k8s_client.ConnectToK8s(kuneConfigPath)
	k8s_client.LaunchK8sJob(clientset, jobName, containerImage, entryCommand, namespace)
	wg.Wait()
}
