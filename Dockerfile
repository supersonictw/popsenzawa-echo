# PopCat Echo
# (c) 2023 SuperSonic (https://github.com/supersonictw).

FROM golang:alpine AS builder
COPY . /builder
RUN cd /builder \
    && go build ./cmd/echo \
    && go clean -cache

FROM alpine:latest
ENV GIN_MODE release
COPY --from=builder /builder/build/echo /app/echo
WORKDIR /app
ENTRYPOINT /app/echo
EXPOSE 8000
