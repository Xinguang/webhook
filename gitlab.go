package main

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type gitlabPayload struct {
	// ref
	Ref string `json:"ref"` // "refs/heads/master"
	// Repository
	Repository gitlabRepository `json:"repository"`
}

// Repository is the common gitlab Repository Entity
type gitlabRepository struct {
	GitURL string `json:"git_http_url"`
}

func init() {
	var payload gitlabPayload
	route("/gitlab", &payload, func(req *http.Request, hook Hook) bool {
		if len(hook.Token) > 0 && hook.Token != req.Header.Get("X-Gitlab-Token") {
			return false
		}
		log.Info(payload.Ref)
		log.Info(strings.TrimRight(hook.Repo, "/"))
		return strings.Contains(payload.Repository.GitURL, strings.TrimRight(hook.Repo, "/")) &&
			payload.Ref == "refs/heads/"+hook.Branch
	})
}
