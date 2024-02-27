package main

import (
	"fmt"
	"log"
	"monopoly-server/logger"
	"monopoly-server/httpserver"
	"monopoly-server/settings"
)

func main()  {
	fmt.Println("hello world")

	logger.NewLogger()
	settings.NewConfig()

	httpserver.InitializeWebServer()

	go func() {
		if err := httpserver.RunWebServer();err != nil {
			log.Fatal(err)
		}
	}()
}
