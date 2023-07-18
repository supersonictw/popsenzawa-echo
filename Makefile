WORKPLACE := $(shell pwd)

all: echo

echo:
	@mkdir -p $(WORKPLACE)/build
	cd cmd/echo && go build -o $(WORKPLACE)/build

.PHONY: clean dev

clean: clean-deps
	rm -rf $(WORKPLACE)/build

clean-deps:
	go clean -cache

dev:
	@go install github.com/cosmtrek/air@latest
	cd $(WORKPLACE) && air -c .air.toml
