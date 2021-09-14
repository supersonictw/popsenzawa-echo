# PopCat Echo
# (c) 2021 SuperSonic (https://github.com/supersonictw).

FROM golang:alpine

WORKDIR /app
ADD . /app

ENV GIN_MODE release
ENV PUBLISH_ADDRESS ":8013"

RUN cd /app/cmd/echo && go build
RUN go clean -cache

ENTRYPOINT /app/cmd/echo/echo

EXPOSE 8013
