build:
	go build -o bin/tils cmd/tils/main.go

install:
	go install ./...

run:
	go run cmd/tils/main.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/tils-${VERSION}-linux-amd64 cmd/tils/main.go
	GOOS=linux GOARCH=386 go build -o bin/tils-${VERSION}-linux-386 cmd/tils/main.go

build_darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/tils-${VERSION}-darwin-amd64 cmd/tils/main.go
	GOOS=darwin GOARCH=386 go build -o bin/tils-${VERSION}-darwin-386 cmd/tils/main.go

build_windows:
	GOOS=windows GOARCH=amd64 go build -o bin/tils-${VERSION}-windows-amd64 cmd/tils/main.go
	GOOS=windows GOARCH=386 go build -o bin/tils-${VERSION}-windows-386 cmd/tils/main.go

build_all: clean build_linux build_darwin build_windows

clean:
	rm -f bin/*
