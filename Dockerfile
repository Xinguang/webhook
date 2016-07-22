FROM docker:1.11.2-git

COPY webhook /usr/bin/webhook

ENTRYPOINT ["/usr/bin/webhook", "--config", "/config.json"]
