FROM docker:1.13.0-git

#ADD https://github.com/starboychina/webhook/releases/download/0.1.0/webhook /usr/bin/webhook
COPY webhook /usr/bin/webhook
RUN chmod +x /usr/bin/webhook

ENTRYPOINT ["/usr/bin/webhook", "--config", "/config.json"]
