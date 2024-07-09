IMAGE_NAME=x3platform/x3util
VERSION=1.0.0
#TAG=${NAME}::${VERSION}

.PHONY: build
build:
	# Linux AMD64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -o bin/dbinit-${VERSION}.linux-amd64
	# Linux ARM64
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0  go build -o bin/dbinit-${VERSION}.linux-arm64
	# Windows  AMD64
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0  go build -o bin/dbinit-${VERSION}.windows-amd64.exe
	# Mac AMD64
	# 达梦不支持 Mac 系统
	#GOOS=darwin GOARCH=amd64 CGO_ENABLED=0  go build -o bin/dbinit-${VERSION}.darwin-amd64


.PHONY: test
test:
	go test x3platform.com/x3util/pkg...