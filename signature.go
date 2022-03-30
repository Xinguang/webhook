package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func calculateSha256Signature(payload interface{}, secret string) string {
	if len(secret) == 0 {
		return ""
	}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		log.Error("prepareWebhooks.JSONPayload: %v", err)
	}
	sig := hmac.New(sha256.New, []byte(secret))
	sig.Write(data)
	return hex.EncodeToString(sig.Sum(nil))
}
