package main

import (
	"fmt"
	"strings"
)

type bitbucketPayload struct {
	// Push
	Push bitbucketPush `json:"push"`
	// Repository
	Repository bitbucketRepository `json:"repository"`
}

// Push is the common Bitbucket Push Sub Entity
type bitbucketPush struct {
	Changes []bitbucketChange `json:"changes"`
}

// Change is the common Bitbucket Change Sub Entity
type bitbucketChange struct {
	New bitbucketChangeData `json:"new"`
}

// ChangeData is the common Bitbucket ChangeData Sub Entity
type bitbucketChangeData struct {
	Name string `json:"name"`
}

// Repository is the common Bitbucket Repository Entity
type bitbucketRepository struct {
	Links bitbucketLinks `json:"links"`
}

// Links is the common Bitbucket Links Sub Entity
type bitbucketLinks struct {
	HTML bitbucketHTML `json:"html"`
}

// HTML is the common Bitbucket HTML Sub Entity
type bitbucketHTML struct {
	HREF string `json:"href"`
}

func init() {
	var payload bitbucketPayload
	route("/bitbucket", &payload, func(hook Hook) bool {
		fmt.Println(payload.Push.Changes[0].New.Name)
		return strings.TrimRight(hook.Repo, "/") == payload.Repository.Links.HTML.HREF &&
			payload.Push.Changes[0].New.Name == hook.Branch
	})
}
