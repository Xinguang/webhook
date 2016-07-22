FROM docker:1.11.2-git

ADD https://github.com/starboychina/webhook/releases/download/0.1.0/webhook /usr/bin/webhook
RUN chmod +x /usr/bin/webhook

ENTRYPOINT ["/usr/bin/webhook", "--config", "/config.json"]
