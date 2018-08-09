FROM golang:1.10.1-alpine3.7 AS b0
RUN echo http://mirror.yandex.ru/mirrors/alpine/v3.7/main >> /etc/apk/repositories
RUN echo http://mirror.yandex.ru/mirrors/alpine/v3.7/community >> /etc/apk/repositories
RUN apk add --no-cache git ca-certificates
RUN go get -u github.com/llater/cue-site \
              github.com/gorilla/mux
WORKDIR /go/src/github.com/llater/cue-site
RUN go build -v -o /tmp/cue-site .
FROM alpine:3.7
WORKDIR /srv
COPY --from=b0 /tmp/cue-site .
ENTRYPOINT ./cue-site
