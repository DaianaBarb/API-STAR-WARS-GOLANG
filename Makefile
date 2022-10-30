lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yml

test:
	go test --race ./...

build:
	env GOSUMDB=off GOOS=linux GOARCH=amd64 go build -o bin/main cmd/main.go

package: clean build
	cd bin && rm -f main.zip && zip main.zip main