.PHONY: fmt test vet lint build clean


fmt:
	go fmt ./...

test:
	go test ./... -v

vet:
	go vet ./...

proto:
	protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative  protos/balance/*.proto

build:
	go build -o build/server/greetServer server/main.go
	go build -o build/client/greetClient client/main.go

clean: fmt vet build