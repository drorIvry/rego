package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/config"
	"github.com/drorivry/rego/initializers"
	k8s_client "github.com/drorivry/rego/k8s"
	"github.com/drorivry/rego/poller"
	"github.com/drorivry/rego/tasker"
)

func runServer(server *http.Server, wg *sync.WaitGroup) {
	server.ListenAndServe()
	defer wg.Done()
}

func runPoller(p *poller.Poller, wg *sync.WaitGroup) {
	p.Run()
	defer wg.Done()
}

func handleInterrupt(c chan os.Signal, server *http.Server, p *poller.Poller) {
	<-c
	log.Info().Msg("Received Ctrl+C")
	log.Info().Msg("Shutting down api server")
	server.Shutdown(context.Background())
	log.Info().Msg("Shutting down poller")
	p.Shutdown()
}

func init() {
	initializers.LoadEnvVars()
	config.InitConfig()
	initializers.InitDBConnection()
}

//	@title			Rego
//	@version		1.0
//	@description	Schedualing workloads made easy.

//	@contact.name	XXXX
//	@contact.email	XXX@gmail.com

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit/

// @host		localhost:3000
// @BasePath	/api/v1
func main() {
	kubeConfigPath := flag.String("kubeConfigPath", "", "The path to the kubeconfig")
	pollInterval := flag.Int("interval", 1, "The polling interval")
	k8s_client.InitK8SClientSet(kubeConfigPath)
	//todo replace that with cobra
	flag.Parse()

	var wg sync.WaitGroup

	wg.Add(1)
	server := tasker.GetServer(config.SERVER_PORT)
	log.Println("Starting server on port " + strconv.Itoa(config.SERVER_PORT))
	go runServer(server, &wg)

	wg.Add(1)
	p := poller.NewPoller(*pollInterval)
	go runPoller(p, &wg)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go handleInterrupt(c, server, p)

	wg.Wait()
}
