package main

import (
	"os"

	cmd "github.com/badboyd/krakend-cobra"
	viper "github.com/badboyd/krakend-viper"
	"github.com/badboyd/lura/config"
	"github.com/badboyd/lura/logging"
	"github.com/badboyd/lura/proxy"
	"github.com/badboyd/lura/router/gin"
)

func main() {

	cmd.Execute(viper.New(), func(serviceConfig config.ServiceConfig) {
		logger, _ := logging.NewLogger("DEBUG", os.Stdout, "")
		gin.DefaultFactory(proxy.DefaultFactory(logger), logger).New().Run(serviceConfig)
	})

}
