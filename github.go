package main

import (
	"strings"
)

type githubPayload struct {
	// ref
	Ref string `json:"ref"` // "refs/heads/master"
	// repository
	Repo struct {
		URL string `json:"url"` // "https://github.com/starboychina/webhook"
	} `json:"repository"`
}

func init() {
	var payload githubPayload
	route("/github", &payload, func(hook Hook) bool {
		return strings.TrimRight(hook.Repo, "/") == payload.Repo.URL &&
			payload.Ref == "refs/heads/"+hook.Branch
	})
}
