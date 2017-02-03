package main

import (
	"fmt"
	"log"
	"net/http"
)

func route(pattern string,
	payload interface{},
	handler func(Hook) bool) {

	log.Printf(pattern)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		err := unmarshal(req.Body, &payload)
		if err != nil {
			fmt.Fprintf(w, something)
			return
		}

		loadConfig()

		msg := ""
		for _, hook := range config.Hooks {
			if handler(hook) {
				msg = fmt.Sprintf(executed, hook.Repo, hook.Branch)
				logger.Printf("success: %s\n", msg)
				msg += "`" + hook.Shell + "`"
				executeShell(hook.Shell)
			}
		}

		if len(msg) < 1 {
			logger.Printf(nothingToDo)
			msg = nothingToDo
		}
		fmt.Fprintf(w, msg)
	})
}
