GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto: proto/wallet.proto
	rm -rf api
	mkdir -p api/pb
	protoc --go_out=api/pb --go_opt=paths=source_relative \
        --go-grpc_out=api/pb --go-grpc_opt=paths=source_relative \
        proto/wallet.proto