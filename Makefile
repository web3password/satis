run:
	go run cmd/main.go -conf config.yaml

build:
	go build -o cmd/satis cmd/main.go

# Build linux binary on other platforms
build-linux:
	GOOS=linux GOARCH=amd64 go build -o cmd/satis_linux cmd/main.go