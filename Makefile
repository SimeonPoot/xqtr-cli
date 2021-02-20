.PHONY: all build delete

build:
	go build -o ./bin/xqtr

delete:
	rm ./bin/xqtr

version: 
	git rev-parse HEAD