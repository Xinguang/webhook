package main

import (
	"io"
	"log"
	"os"
)

var logger *log.Logger

func setLog() {
	logHandler, err := os.OpenFile(*cmdLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("cannot write log -> " + *cmdLogFile)
	}

	logger = log.New(io.MultiWriter(logHandler, os.Stdout),
		"",
		log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetFlags(5)
}
