package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port             string
	Account          string
	Key              string
	ConnectionString string
	ContainerName    string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	websection := cfg.Section("web")
	blobsection := cfg.Section("azureblob")
	Config = ConfigList{
		Port:             websection.Key("port").String(),
		Account:          blobsection.Key("Account").String(),
		Key:              blobsection.Key("Key").String(),
		ConnectionString: blobsection.Key("ConnectionString").String(),
		ContainerName:    blobsection.Key("ContainerName").String(),
	}
}
