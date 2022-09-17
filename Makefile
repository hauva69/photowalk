all:
	go build -mod=vendor

install:
	install -m 755 photowalk /usr/local/bin/
