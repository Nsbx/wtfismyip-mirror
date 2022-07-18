FROM alpine:edge
MAINTAINER Clint Ruoho clint@wtfismyip.com
RUN apk update
RUN apk add --no-cache coturn
COPY turnserver.conf /etc/turnserver.conf
ENTRYPOINT [ "/usr/bin/turnserver" ]
