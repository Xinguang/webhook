FROM docker:1.13.0-git

ENV GOPATH /go

RUN apk update && apk add go g++ && \
    rm -rf /var/cache/apk/* && \
    go get -u github.com/starboychina/webhook && \
    cd /go/src/github.com/starboychina/webhook && \
    go build -o /usr/bin/webhook . && \
    apk del go g++ && \
    rm -rf /var/cache/apk/* && \
    rm -rf /go

ENTRYPOINT ["/usr/bin/webhook", "--config", "/config.json"]
