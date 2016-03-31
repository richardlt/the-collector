FROM alpine:3.3

MAINTAINER Richard LE TERRIER <richard.le.terrier@gmail.com>

COPY ./the-collector.zip /opt/the-collector.zip

RUN unzip /opt/the-collector.zip -d /opt && chmod +x /opt/the-collector

EXPOSE 8080

CMD ["/opt/the-collector", "start"]
