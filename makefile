.DEFAULT_GOAL := build
clean:
	go clean -v
	go get -u ./...
build: clean
	go build -v ./...
	go test -v ./...
