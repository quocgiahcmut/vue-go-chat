package main

import (
	"log"

	"github.com/quocgiahcmut/vue-go-chat/api"
	"github.com/quocgiahcmut/vue-go-chat/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	server := api.NewServer(config)
	server.Start(config.RESTServerAddress)
}
