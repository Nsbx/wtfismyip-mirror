FROM alpine:edge

MAINTAINER Clint Ruoho clint@wtfismyip.com
LABEL package.date=2024-02-13

RUN apk update
RUN apk add --no-cache unbound openssl procps-ng
COPY unbound.conf /etc/unbound/unbound.conf
RUN /usr/sbin/unbound-control-setup
ENTRYPOINT [ "/usr/sbin/unbound", "-d" ]
