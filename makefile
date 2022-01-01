.PHONY: fmt test vet lint build clean


fmt:
	go fmt ./...

test:
	go test ./... -v

vet:
	go vet ./...

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  protos/greeting.proto

build:
	go build -o build/server/greetServer server/main.go
	go build -o build/client/greetClient client/main.go

clean: fmt vet build