package common

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const configFileName = "./agents.config"

func ReadAgentsConfig() []string {
	configFile, err := os.ReadFile(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	endpoints := strings.Split(strings.Trim(strings.Trim(string(configFile), "["), "]"), " ")
	return endpoints
}

func OverwriteAgentsConfig(endpoints []string) {
	configFile, err := os.Create(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	_, err = configFile.Write([]byte(fmt.Sprintf("%v", endpoints)))
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteAgentsConfig() {
	err := os.Remove(configFileName)
	if err != nil {
		log.Fatal(err)
	}
}
