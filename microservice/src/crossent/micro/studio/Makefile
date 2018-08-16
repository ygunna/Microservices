PUBLIC_FILES = $(shell find dist/ -type f)

.PHONY: all clean

all: msxpert

clean:
	rm -rf web/dist/*

bindata:
	go-bindata -pkg studio -prefix "web/dist/" web/dist/...
	go fmt bindata.go

build:
	cd web && npm install && ng build -prod

msxpert: clean build bindata
