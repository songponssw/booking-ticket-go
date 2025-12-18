SERVICES := search order payment
.PHONY: generate-all

PROTO_OUT := proto/golang


generate-gateway:
	mkdir -p grpc-gateway/proto/golang
	protoc \
	  -I grpc-gateway/proto \
		-I . \
	  --go_out=grpc-gateway/proto/golang \
		--go_opt=paths=source_relative \
		--go-grpc_out=grpc-gateway/proto/golang \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=grpc-gateway/proto/golang \
		--grpc-gateway_opt=paths=source_relative \
		grpc-gateway/proto/grpc-gateway.proto

generate-%:
	protoc \
		--proto_path=$*/proto \
		--go_out=$*/proto/golang \
		--go_opt=paths=source_relative \
		--go-grpc_out=$*/proto/golang \
		--go-grpc_opt=paths=source_relative \
		./$*/proto/$*.proto


generate-all: $(addprefix generate-,$(SERVICES))
