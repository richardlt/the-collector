FROM alpine:3.3

MAINTAINER Victor Vieux <victor@docker.com>

COPY ./the-collector.zip /opt/the-collector.zip

RUN unzip /opt/the-collector.zip -d /opt

EXPOSE 8080

CMD ["/opt/the-collector", "start"]
