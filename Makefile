metadataproto:
	protoc --go_out=. --go-grpc_out=. grpc_server/metadata/metadata.proto

.PHONY: proto