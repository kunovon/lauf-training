FROM alpine:3.18.4

LABEL maintainer="markus"

RUN go build -o laufen main.go

ADD laufen /opt/laufen
