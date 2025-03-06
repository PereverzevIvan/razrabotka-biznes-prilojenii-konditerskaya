PROTO_FOLDER = ./api
# PROTOC_PATH_TO_PROTOS = $(PROTO_FOLDER)/**/**/*.proto # use if have versions; example: ./api/domain/v1/*.proto
PROTOC_PATH_TO_PROTOS = $(PROTO_FOLDER)/**/*.proto # use if have no versions; example: ./api/domain/*.proto
PKG_FOLDER = ./pkg

# install bin dependencies
bin-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

run-swagger:
	docker compose up -d

gen: clean-gen gen-go gen-gateway gen-swagger

gen-go:
	@protoc -I . \
    --go_out $(PKG_FOLDER) --go_opt paths=source_relative \
    --go-grpc_out $(PKG_FOLDER) --go-grpc_opt paths=source_relative \
	$(PROTOC_PATH_TO_PROTOS)

gen-gateway:
	@protoc -I . \
	--grpc-gateway_out $(PKG_FOLDER) --grpc-gateway_opt paths=source_relative \
	$(PROTOC_PATH_TO_PROTOS)

gen-swagger:
	@protoc -I . \
	--openapiv2_out ./swagger \
	--openapiv2_opt logtostderr=true,allow_merge=true,merge_file_name=swagger \
	$(PROTOC_PATH_TO_PROTOS)

clean-gen:
	@rm -rf $(PKG_FOLDER)/*

.PHONY: gen gen-go gen-gw gen-swagger run-swagger clean-gen
