package main

import (
	"os"

	cmd "github.com/betprophet1/krakend-cobra"
	viper "github.com/betprophet1/krakend-viper"
	"github.com/betprophet1/lura/config"
	"github.com/betprophet1/lura/logging"
	"github.com/betprophet1/lura/proxy"
	"github.com/betprophet1/lura/router/gin"
)

func main() {

	cmd.Execute(viper.New(), func(serviceConfig config.ServiceConfig) {
		logger, _ := logging.NewLogger("DEBUG", os.Stdout, "")
		gin.DefaultFactory(proxy.DefaultFactory(logger), logger).New().Run(serviceConfig)
	})

}
