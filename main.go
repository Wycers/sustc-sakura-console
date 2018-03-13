package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"os"
	"io"
	"os/signal"
	"syscall"
	"net/http"

	"github.com/wycers/sustc-sakura-console/log"
	"github.com/wycers/sustc-sakura-console/util"
	"github.com/wycers/sustc-sakura-console/controller"
	"github.com/wycers/sustc-sakura-console/service"
)

var logger *log.Logger

func init() {
	rand.Seed(time.Now().Unix())

	log.SetLevel("Trace")
	logger = log.NewLogger(os.Stdout)

	util.LoadConfig()

	if util.Config.RuntimeMode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

}

func main() {
	service.ConnectDB()
	//service.Upgrade.Perform()

	router := controller.Routes()
	server := &http.Server{
		Addr:    "0.0.0.0:" + util.Config.Port,
		Handler: router,
	}
	handleSignal(server)

	logger.Infof("Sakura (v%s) is running [%s]", util.Version, util.Config.Server)

	if util.Config.RuntimeMode == "dev" {
		server.ListenAndServe()
	} else {
		server.ListenAndServeTLS("214502860120160.crt", "214502860120160.key")
	}
}

func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logger.Infof("got signal [%s], exiting sakura now", s)
		if err := server.Close(); nil != err {
			logger.Errorf("server close failed: ", err)
		}

		service.DisconnectDB()

		logger.Infof("Sakura exited")
		os.Exit(0)
	}()
}