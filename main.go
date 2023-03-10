package main

import (
	"flag"

	k8s_client "github.com/drorivry/matter/k8s"
)

func main() {
	jobName := flag.String("jobname", "test-job", "The name of the job")
	containerImage := flag.String("image", "ubuntu:latest", "Name of the container image")
	entryCommand := flag.String("command", "ls", "The command to run inside the container")

	flag.Parse()

	clientset := k8s_client.ConnectToK8s()
	k8s_client.LaunchK8sJob(clientset, jobName, containerImage, entryCommand)
}
