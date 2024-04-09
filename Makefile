PROTOC_VERSION = 26.1
PROTOC_GEN_GO_VERSION = v1.28
PROTOC_GEN_GO_GRPC_VERSION = v1.2

test: unit

unit:


tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b `go env GOPATH`/bin $(GOLANGCI_LINT_VERSION)

	wget https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip
	unzip protoc-$(PROTOC_VERSION)-linux-x86_64.zip bin/protoc
	rm protoc-$(PROTOC_VERSION)-linux-x86_64.zip

	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

proto:
	./bin/protoc services/upload/upload.proto \
		--go_opt=module=github.com/maxtek6/filegrpc-go \
		--go-grpc_opt=module=github.com/maxtek6/filegrpc-go \
		--go_out=services/upload \
		--go-grpc_out=services/upload
	./bin/protoc services/download/download.proto \
		--go_opt=module=github.com/maxtek6/filegrpc-go \
		--go-grpc_opt=module=github.com/maxtek6/filegrpc-go \
		--go_out=services/download \
		--go-grpc_out=services/download