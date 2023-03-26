package main

import (
	"flag"
	"sync"

	"github.com/drorivry/matter/config"
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
	config.InitConfig()
	initializers.InitDBConnection(config.DB_URL)
}

func main() {
	kubeConfigPath := flag.String("kubeConfigPath", "", "The path to the kubeconfig")
	pollInterval := flag.Int("interval", 1, "The polling interval")
	k8s_client.InitK8SClientSet(kubeConfigPath)
	//todo replace that with cobra
	flag.Parse()

	var wg sync.WaitGroup

	server := tasker.GetServer()
	wg.Add(1)
	go runServer(server)

	wg.Add(1)
	go poller.Run(*pollInterval)

	wg.Wait()
}
