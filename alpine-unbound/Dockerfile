FROM alpine:latest

MAINTAINER Clint Ruoho clint@wtfismyip.com


RUN apk add --no-cache unbound
COPY unbound.conf /etc/unbound/unbound.conf
ENTRYPOINT [ "/usr/sbin/unbound", "-d" ]
