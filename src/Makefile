.PHONY: gen
gen: ## Compile the proto file.
	protoc --go_out=./pkg/proto --go-grpc_out=./pkg/proto pkg/proto/cipher/*.proto

.PHONY: build-server
build-server: ## Build server.
	go build -race -o bin/server ./server/cmd/...

.PHONY: build-client
build-client: ## Build client.
	go build -race -o bin/client ./client/cmd/...

.PHONY: build
build: gen build-server build-client

.PHONY: grpcui
grpcui: ## connect with grpcui client
	grpcui -port 8088 -plaintext localhost:8080
