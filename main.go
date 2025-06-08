package main

import (
	"context"
	"time"

	"hr-system-salary/cmd/rest"
	"hr-system-salary/config"
	appSetup "hr-system-salary/internal/setup"
	"hr-system-salary/pkg/log"
)

func main() {
	// config init
	log.InitZeroLog()
	config.InitConfig()
	// conf := config.GetConfig()

	_, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// app setup init
	setup := appSetup.Init()

	rest.StartServer(setup)
}
