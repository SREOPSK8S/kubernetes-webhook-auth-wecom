package main

import (
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/config"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/router"
)

func main() {
	config.InitAndLoad()
	router := router.GetRouter()
	router.Run(":8443")
}
