package main

import (
	"context"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/config"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/router"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/repository"
)

func main() {
	// load config
	config.InitAndLoad()
	// create table
	conn := repository.DataClient{}
	client := conn.GetConnection()
	err := conn.CreateSchema(context.TODO(),client)
	if err != nil {
		return
	}
	router := router.GetRouter()
	router.Run(":8443")
}
