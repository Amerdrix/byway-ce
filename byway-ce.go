package main

import (
	"fmt"

	"github.com/amerdrix/byway-ce/config"
	"github.com/amerdrix/byway-ce/core"
)

func main() {
	fmt.Println("Welcome to byway!")

	config := make(chan *core.Config, 1)

	bywayConfig.ReadConfigFile(config)

	core.Init(bywayConfig.LogConfig(config))
	exit := make(chan bool)
	<-exit
}
