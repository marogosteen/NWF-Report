package config

import (
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Account          string
	ConnectionString string
	ContainerName    string
	Key              string
	Port             string
}

var Config ConfigList

func init() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	cfg, err := ini.Load("config.ini")
	if err == nil {
		blobsection := cfg.Section("azureblob")
		Config = ConfigList{
			Account:          blobsection.Key("Account").String(),
			ConnectionString: blobsection.Key("ConnectionString").String(),
			ContainerName:    blobsection.Key("ContainerName").String(),
			Key:              blobsection.Key("Key").String(),
			Port:             port,
		}
	} else {
		Config = ConfigList{
			Account:          os.Getenv("Account"),
			ConnectionString: os.Getenv("ConnectionString"),
			ContainerName:    os.Getenv("ContainerName"),
			Key:              os.Getenv("Key"),
			Port:             port,
		}
	}
}
