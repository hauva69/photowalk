all:
	go build -mod=vendor

install:
	cp photowalk /usr/local/bin/
