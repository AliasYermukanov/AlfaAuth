package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Configs struct {
	Elastic ElasticConfig `json:"elastic"`
}

type ElasticConfig struct {
	ConnectionUrl []string `json:"connection_url"`
}

var AllConfigs *Configs

func GetConfigs() error {
	var filePath string
	if os.Getenv("config") == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		filePath = pwd + "/src/config/config.json"
	} else {
		filePath = os.Getenv("config")
	}
	file, err := os.Open(filePath)

	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	var configs Configs
	err = decoder.Decode(&configs)

	if err != nil {
		return err
	}
	AllConfigs = &configs
	return nil
}

func Healthchecks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
