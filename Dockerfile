FROM scratch

COPY cli-template /usr/local/bin/cli-template

ENTRYPOINT ["/usr/local/bin/cli-template"]