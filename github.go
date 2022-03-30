package main

import (
	"net/http"
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
	route("/github", &payload, func(req *http.Request, hook Hook) bool {
		if len(hook.Token) > 0 && calculateSha256Signature(payload, hook.Token) != req.Header.Get("X-Hub-Signature") {
			return false
		}
		return strings.TrimRight(hook.Repo, "/") == payload.Repo.URL &&
			payload.Ref == "refs/heads/"+hook.Branch
	})
}
