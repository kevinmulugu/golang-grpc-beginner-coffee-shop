OUT_DIR = coffeshop_protos
PROTO_FILE = coffee_shop.proto

.PHONY: all
all: generate

.PHONY: generate
generate:
	mkdir -p $(OUT_DIR)
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)