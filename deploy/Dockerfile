FROM alpine:3.9.2

RUN apk --no-cache add iputils ca-certificates

COPY ./deploy/config.yaml .
COPY ./deploy/app_server_tasks .