package bywayConfig

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/amerdrix/byway-ce/core"
)

// LogConfig - intercepts a chan *core.Config and logs the value. Returns a new chan *core.Config
func LogConfig(input chan *core.Config) chan *core.Config {
	configWriter := make(chan *core.Config, 1)
	go func() {
		for {
			table := <-input

			loaded, _ := yaml.Marshal(table)
			fmt.Printf("Configuration:\n%s", loaded)

			configWriter <- table
		}
	}()
	return configWriter
}
