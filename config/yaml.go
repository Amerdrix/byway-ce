package bywayConfig

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/amerdrix/byway-ce/core"
)

// ReadConfigFile - Reads byway config from './byway.yml'
func ReadConfigFile(channel chan *core.Config) {
	configFile, err := ioutil.ReadFile("./byway.yml")
	if err != nil {
		log.Fatal(err)
	}

	newConfig := core.NewConfig()

	err = yaml.Unmarshal(configFile, &newConfig)
	if err != nil {
		log.Fatal(err)
	}

	channel <- newConfig
}
