package main

import (
	"fmt"
	"monopoly-server/httpserver"
	"monopoly-server/logger"
	"monopoly-server/settings"
)

func recoverFunc() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main()  {
	fmt.Println("hello world")

	logger.NewLogger()
	defer logger.Sync()

	settings.NewConfig()

	httpserver.InitializeWebServer()

	go func() {
		defer recoverFunc()

		if err := httpserver.RunWebServer();err != nil {
			logger.Fatal("webServer runtime error: %v",err)
		}
	}()
}
