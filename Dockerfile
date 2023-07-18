# PopSenzawa Echo
# (c) 2023 SuperSonic (https://github.com/supersonictw).

FROM golang:alpine AS builder
COPY . /workplace
WORKDIR /workplace
RUN apk add --no-cache make
RUN make && make clean-deps

FROM alpine:latest
ENV GIN_MODE release
COPY --from=builder /workplace/build/echo /workplace/echo
WORKDIR /workplace
ENTRYPOINT /workplace/echo
EXPOSE 8000
