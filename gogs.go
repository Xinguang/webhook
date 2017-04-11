package main

import (
	"strings"
)

type gogsPayload struct {
	// ref
	Ref string `json:"ref"` // "refs/heads/master"
	// repository
	Repo struct {
		URL string `json:"clone_url"` // "https://try.gogs.io/starboychina/webhook"
	} `json:"repository"`
}

func init() {
	var payload gogsPayload
	route("/gogs", &payload, func(hook Hook) bool {
		return strings.TrimRight(hook.Repo, "/") == payload.Repo.URL &&
			payload.Ref == "refs/heads/"+hook.Branch
	})
}
