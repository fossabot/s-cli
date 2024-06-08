FROM alpine
RUN apk update && apk add --update --no-cache ca-certificates
COPY skynxctl /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/skynxctl"]
