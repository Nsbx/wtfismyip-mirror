FROM alpine:edge

MAINTAINER Clint Ruoho clint@wtfismyip.com

RUN apk update

RUN apk add --no-cache redis procps-ng

ENTRYPOINT [ "/usr/bin/redis-server", "--protected-mode", "no" ]
