FROM alpine:3.8

RUN apk add --no-cache ca-certificates

COPY ./the-collector-package /opt/the-collector

WORKDIR /opt/the-collector

ENTRYPOINT ["/opt/the-collector/the-collector"]

EXPOSE 8080