FROM alpine:edge

MAINTAINER Clint Ruoho clint@wtfismyip.com

RUN apk add --no-cache redis  

ENTRYPOINT [ "/usr/bin/redis-server", "--protected-mode", "no" ]
