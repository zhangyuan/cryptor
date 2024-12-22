build:
	go build -ldflags "-s -w"

clean:
	rm -rf cryptor
	rm -rf bin/cryptor-*

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/cryptor_darwin-amd64 -ldflags "-s -w"
	GOOS=darwin GOARCH=arm64 go build -o bin/cryptor_darwin-arm64 -ldflags "-s -w"

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/cryptor_linux-amd64 -ldflags "-s -w"

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/cryptor_windows-amd64 -ldflags "-s -w"

build-all: clean build-macos build-linux build-windows