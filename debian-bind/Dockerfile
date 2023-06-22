FROM debian:unstable

MAINTAINER Clint Ruoho clint@wtfismyip.com

RUN apt clean
RUN apt update
RUN apt install -y bind9 vim procps dnsutils 
COPY named.conf /etc/bind/named.conf
COPY start.sh /etc/bind
COPY rndc.key /etc/bind
CMD [ "bash", "-c", "/etc/bind/start.sh" ]
