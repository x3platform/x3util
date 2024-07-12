PROAGRAM_NAME=x3util
IMAGE_NAME=x3platform/x3util
VERSION=0.0.2
#TAG=${NAME}::${VERSION}

.PHONY: build
build:
	# Linux AMD64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -o bin/${PROAGRAM_NAME}-${VERSION}.linux-amd64
	# Linux ARM64
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0  go build -o bin/${PROAGRAM_NAME}-${VERSION}.linux-arm64
	# Windows  AMD64
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0  go build -o bin/${PROAGRAM_NAME}-${VERSION}.windows-amd64.exe
	# Mac AMD64
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0  go build -o bin/${PROAGRAM_NAME}-${VERSION}.darwin-amd64

.PHONY: test
test:
	go test x3platform.com/x3util/pkg/...

clearn:
	rm -f bin/*