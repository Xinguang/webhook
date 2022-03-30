package main

import (
	"net/http"
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
	route("/gogs", &payload, func(req *http.Request, hook Hook) bool {
		if len(hook.Token) > 0 && calculateSha256Signature(payload, hook.Token) != req.Header.Get("X-Gogs-Signature") {
			return false
		}
		return strings.TrimRight(hook.Repo, "/") == payload.Repo.URL &&
			payload.Ref == "refs/heads/"+hook.Branch
	})
}
