FROM docker:latest

ENV GOPATH /go

RUN apk add --no-cache \
		git \
		openssh-client go g++ && \
    rm -rf /var/cache/apk/* && \
    go get -u github.com/starboychina/webhook && \
    cd /go/src/github.com/starboychina/webhook && \
    ls -la && \
    go build -o /usr/bin/webhook . && \
    webhook -h && \
    apk del go g++ && \
    rm -rf /var/cache/apk/* && \
    rm -rf /go

ENTRYPOINT ["/usr/bin/webhook", "--config", "/config.json"]
