BINARY_NAME=app

build:
	go build -o .bin/${BINARY_NAME} cmd/server/main.go 
run:
	go build -o .bin/${BINARY_NAME} cmd/server/main.go 
	chmod +x .bin/${BINARY_NAME}
	./.bin/${BINARY_NAME}
clean:
	go clean
	rm .bin/${BINARY_NAME}
lint:
	golangci-lint run
test:
	go test ./...
protos:
	protoc --go_out ./pkg ./proto/auth.proto
	protoc --go-grpc_out ./pkg ./proto/auth.proto	
	protoc --go_out ./pkg ./proto/user.proto
	protoc --go-grpc_out ./pkg ./proto/user.proto