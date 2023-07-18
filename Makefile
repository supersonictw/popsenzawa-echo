WORKPLACE := $(shell pwd)

all: echo

echo:
	@mkdir -p $(WORKPLACE)/build
	cd cmd/echo && go build -o $(WORKPLACE)/build

.PHONY: clean dev

clean:
	rm -rf $(WORKPLACE)/build

dev:
	@go install github.com/cosmtrek/air@latest
	cd $(WORKPLACE) && air -c .air.toml
