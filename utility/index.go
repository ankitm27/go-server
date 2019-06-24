package utility

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Configuration struct {
	DatabaseUrl string
}

// var filename = "/go-server/utility/config.json"

func GetConfig() Configuration {
	absPath, _ := filepath.Abs("./utility/config.json")

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
