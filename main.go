package main

import (
	"flag"
	"sync"

	"github.com/drorivry/rego/config"
	"github.com/drorivry/rego/initializers"
	k8s_client "github.com/drorivry/rego/k8s"
	"github.com/drorivry/rego/poller"
	"github.com/drorivry/rego/tasker"
	"github.com/gin-gonic/gin"
)

func runServer(server *gin.Engine) {
	server.Run()
}

func init() {
	config.InitConfig()
	initializers.InitDBConnection(config.DB_URL)
}

//	@title			Matter
//	@version		1.0
//	@description	Schedualing workloads made easy.

//	@contact.name	XXXX
//	@contact.email	XXX@gmail.com

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit/

//	@host		localhost:3000
//	@BasePath	/api/v1
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
