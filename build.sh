#!/bin/sh

# go run cert/generate_cert.go --host blog.kansea.com --org blog.kansea.com

# GOOS=linux GOARCH=amd64 go build -o webhook

rm webhook
docker run -it --rm -v $(pwd):/root golang:alpine sh -c '
  apk add --update git
  go get github.com/starboychina/webhook
  cd /go/src/github.com/starboychina/webhook
  go build -o /root/webhook
'

docker rmi starboychina/webhook
docker build -t starboychina/webhook:1.13.0 .
docker build -t starboychina .
