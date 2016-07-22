#!/bin/sh

# go run cert/generate_cert.go --host blog.kansea.com --org blog.kansea.com

GOOS=linux GOARCH=amd64 go build -o webhook

docker rmi starboychina/webhook 
docker build -t starboychina/webhook .
