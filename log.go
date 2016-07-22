package main

import (
	"log"
	"os"
)

func setLog() {
	logHandler, err := os.OpenFile(*cmdLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("cannot write log")
	}
	log.SetOutput(logHandler)
	log.SetFlags(5)
}
