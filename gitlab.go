package main

import (
	"fmt"
	"strings"
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
	route("/gitlab", &payload, func(hook Hook) bool {
		fmt.Println(payload.Ref)
		fmt.Println(strings.TrimRight(hook.Repo, "/"))
		return strings.Contains(payload.Repository.GitURL, strings.TrimRight(hook.Repo, "/")) &&
			payload.Ref == "refs/heads/"+hook.Branch
	})
}
