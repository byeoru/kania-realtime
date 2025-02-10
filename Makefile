metadataproto:
	protoc --go_out=. --go-grpc_out=. grpc_server/metadata/metadata.proto

updatesproto:
	protoc --go_out=. --go-grpc_out=. grpc_server/updates/updates.proto

.PHONY: proto updatesproto