FROM golang:1.22.4 as builder

MAINTAINER Clint Ruoho clint@wtfismyip.com

WORKDIR /app
COPY go.mod go.sum ./
COPY wtf.go ./
COPY default.pgo ./

RUN go mod download && CGO_ENABLED=0 GOOS=linux GOAMD64=v4 go build -a -installsuffix cgo -o wtf .

FROM alpine:edge

WORKDIR /app

RUN mkdir -p /usr/local/wtf/GeoIP && apk update && apk upgrade && apk add --no-cache vim procps-ng && addgroup -S appuser && adduser -u 666 -S -G appuser appuser
COPY static/GeoIP /usr/local/wtf/GeoIP
COPY static/static /usr/local/wtf/static
COPY static/docker /docker
COPY static/docker/resolv.conf /etc/resolv.conf
COPY --from=builder /app/wtf .
COPY start.sh .

USER appuser
CMD [ "./start.sh" ]
