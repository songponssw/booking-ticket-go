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
		grpc-gateway/proto/grpc_gateway.proto

generate-%:
	mkdir $*/proto/golang
	protoc \
		--proto_path=$*/proto \
		--go_out=$*/proto/golang \
		--go_opt=paths=source_relative \
		--go-grpc_out=$*/proto/golang \
		--go-grpc_opt=paths=source_relative \
		./$*/proto/$*.proto


clean-%:
	rm -rf $*/proto/golang


generate-all: $(addprefix generate-,$(SERVICES))
