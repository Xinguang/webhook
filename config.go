package main

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Hooks []Hook
}

type Hook struct {
	Repo   string // "starboychina/webhook"
	Token  string // "token"
	Branch string // "master"
	Shell  string // "docker restart homepage" or "shell.sh"
}

func loadConfig() {
	configData, err := ioutil.ReadFile(*cmdConfigFile)
	if err != nil {
		log.Info(*cmdConfigFile + " not found")
		log.Fatal(err)
	}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Info(*cmdConfigFile + " invalid")
		log.Fatal(err)
	}
}
