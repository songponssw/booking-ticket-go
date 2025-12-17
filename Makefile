SERVICES := search order payment
.PHONY: generate-all

PROTO_OUT := proto/golang
generate-%:
	protoc \
		--proto_path=$*/proto \
		--go_out=$*/proto/golang \
		--go_opt=paths=source_relative \
		--go-grpc_out=$*/proto/golang \
		--go-grpc_opt=paths=source_relative \
		./$*/proto/$*.proto


generate-all: $(addprefix generate-,$(SERVICES))
