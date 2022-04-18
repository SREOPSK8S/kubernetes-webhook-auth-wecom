package main

import  "github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/router"

func main() {
	router := router.GetRouter()
	router.Run(":8443")
}
