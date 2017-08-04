.PHONY: app

app:
	go build
	env GOOS=linux GOARCH=amd64 go build -o dm-amd64
	env GOOS=linux GOARCH=386 go build -o dm-386

default: app