#!make
GOPATH := $(shell go env GOPATH)

watch-server:  ## 👀 Run all services with hot reload
	go mod vendor; 
	${GOPATH}/bin/air -c .air.toml