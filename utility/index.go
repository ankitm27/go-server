package utility

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Configuration struct {
	DatabaseUrl string
	BackendUrl  string
	Port        string
	PubSubUrl   string
	PubSubPort  string
	SocketPort  int
}

// var filename = "/go-server/utility/config.json"

func GetConfig() Configuration {
	var absPath string
	var _ string
	if os.Getenv("Environemt") == "production" {
		absPath, _ = filepath.Abs("./utility/config.json")
	} else {
		absPath, _ = filepath.Abs("./utility/development.json")
	}
	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println("There is some problem in getting config, Please try after some time", err)
	}
	var configuration Configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("There is some problem in getting the config, Please try after some time1", err)
	}
	return configuration
}
