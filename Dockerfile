FROM alpine:3.8

COPY ./the-collector-package /opt/the-collector

WORKDIR /opt/the-collector

ENTRYPOINT ["/opt/the-collector/the-collector"]

EXPOSE 8080