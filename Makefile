.PHONY: default install test
all: default install test

APPNAME=ip
VERSION=v1.0.0

gosec:
	go get github.com/securego/gosec/cmd/gosec

sec:
	@gosec ./...
	@echo "[OK] Go security check was completed!"

proxy:
	export GOPROXY=https://goproxy.cn

default: proxy
	gofmt -s -w .&&go mod tidy&&go fmt ./...&&revive .&&goimports -w .&&golangci-lint run --enable-all

install: proxy
	go install -ldflags="-s -w" ./...
	ls -lh ~/go/bin/$(APPNAME)
	upx ~/go/bin/$(APPNAME)
	ls -lh ~/go/bin/$(APPNAME)
package: install
	mv ~/go/bin/$(APPNAME) ~/go/bin/$(APPNAME)-$(VERSION)-darwin-amd64
	gzip -f ~/go/bin/$(APPNAME)-$(VERSION)-darwin-amd64
	ls -lh ~/go/bin/$(APPNAME)*

linux:
	GOOS=linux GOARCH=amd64 go install -ldflags="-s -w" ./...
	upx ~/go/bin/linux_amd64/$(APPNAME)

test: proxy
	go test ./...
