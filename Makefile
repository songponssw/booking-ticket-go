SHELL := /bin/bash
PROTOC := protoc
SERVICES := search api_gateway
THIRD_PARTY := third_party
DEBUG_DIR := debug

.PHONY: $(addprefix generate-,$(SERVICES)) $(addprefix clean-,$(SERVICES))

# Fer bash-completion
$(foreach s,$(SERVICES),\
	$(eval generate-$(s): ; @$(MAKE) generate-$(s)-template) \
	$(eval clean-$(s):    ; @$(MAKE) clean-$(s)-template) \
)

generate-%-template:
	mkdir -p $*/proto/golang
	$(PROTOC) \
		-I $*/proto \
		-I $(THIRD_PARTY) \
		--proto_path=$*/proto \
		--go_out=$*/proto/golang --go_opt=paths=source_relative \
		--go-grpc_out=$*/proto/golang --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$*/proto/golang --grpc-gateway_opt=paths=source_relative \
		./$*/proto/$*.proto


clean-%-template:
	rm -rf $*/proto/golang

debug-dev:
	exec -a debug-search go run search/cmd/grpc/main.go >> $(DEBUG_DIR)/search.log 2>&1 & \
	exec -a debug-gateway go run api_gateway/cmd/main.go >> $(DEBUG_DIR)/api_gateway.log 2>&1 & \
	wait

# Make debug-search
# go run search/cmd/grpc/main.go 2>&1 debug/search.log
