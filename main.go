package main

import (
	"flag"
	"sync"

	"github.com/drorivry/matter/initializers"
	k8s_client "github.com/drorivry/matter/k8s"
	"github.com/drorivry/matter/poller"
	"github.com/drorivry/matter/tasker"
	"github.com/gin-gonic/gin"
)

func runServer(server *gin.Engine) {
	server.Run()
}

func init() {
	initializers.LoadEnvVars()
}

func main() {
	jobName := flag.String("jobname", "test-job", "The name of the job")
	containerImage := flag.String("image", "ubuntu:latest", "Name of the container image")
	entryCommand := flag.String("command", "", "The command to run inside the container")
	namespace := flag.String("namespace", "default", "The job's namespace to deploy to")
	kuneConfigPath := flag.String("kubeConfigPath", "", "The path to the kubeconfig")
	pollInterval := flag.Int("interval", 1, "The polling interval")

	//todo replace that with cobra
	flag.Parse()

	var wg sync.WaitGroup

	server := tasker.GetServer()
	wg.Add(1)
	go runServer(server)

	wg.Add(1)
	go poller.Run(*pollInterval)

	clientset := k8s_client.ConnectToK8s(kuneConfigPath)
	k8s_client.LaunchK8sJob(clientset, jobName, containerImage, entryCommand, namespace)
	wg.Wait()
}
