package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func route(pattern string,
	payload interface{},
	handler func(*http.Request, Hook) bool) {

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
			if handler(req, hook) {
				msg = fmt.Sprintf(executed, hook.Repo, hook.Branch)
				log.Infof("success: %s\n", msg)
				msg += "`" + hook.Shell + "`"
				executeShell(hook.Shell)
			}
		}

		if len(msg) < 1 {
			log.Info(nothingToDo)
			msg = nothingToDo
		}
		fmt.Fprintf(w, msg)
	})
}
