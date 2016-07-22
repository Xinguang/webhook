package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	Hooks []Hook
}

type Hook struct {
	Repo   string // "starboychina/webhook"
	Branch string // "master"
	Shell  string // "docker restart homepage" or "shell.sh"
}

func loadConfig() {
	configData, err := ioutil.ReadFile(*cmdConfigFile)
	if err != nil {
		fmt.Println(*cmdConfigFile + " not found")
		log.Fatal(err)
	}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println(*cmdConfigFile + " invalid")
		log.Fatal(err)
	}
}
