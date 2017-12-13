GOOGLEAPIS := googleapis
RPC_PATH := google/rpc

# BUILD UP PROTOC ARGs
PROTO_PATH := --proto_path=protoc-tmp
MAPPING := Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,
MAPPING := $(MAPPING)Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,
MAPPING := $(MAPPING)Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,
PLUGIN := --gogoslick_out=$(MAPPING):$(GOOGLEAPIS)

STATUS_PROTO := protoc-tmp/$(RPC_PATH)/status.proto
CODE_PROTO := protoc-tmp/$(RPC_PATH)/code.proto
ERR_PROTO := protoc-tmp/$(RPC_PATH)/error_details.proto

PROTOC = protoc-min-version -version=3.0.0

gofmt:
	gofmt -l -s -w .

SHA=c8c975543a134177cc41b64cbbf10b88fe66aa1d
GOOGLEAPIS_URL=https://raw.githubusercontent.com/googleapis/googleapis/$(SHA)

update:
	@dep ensure

	# Install binaries
	@go install github.com/gogo/protobuf/protoc-gen-gogoslick
	@go install github.com/gogo/protobuf/gogoreplace
	@go install github.com/gogo/protobuf/protoc-min-version

	# Record protoc version
	@echo `protoc --version` > protoc.version

	@mkdir -p protoc-tmp

	# GOOGLEAPIS Generation
	@mkdir -p protoc-tmp/$(RPC_PATH)

	## Generate google/rpc/status.pb.go
	@curl -sS $(GOOGLEAPIS_URL)/google/rpc/status.proto -o $(STATUS_PROTO)
	@gogoreplace 'option go_package = "google.golang.org/genproto/googleapis/rpc/status;status";' '' $(STATUS_PROTO)
	@$(PROTOC) $(PROTO_PATH) $(PLUGIN) $(STATUS_PROTO)

	## Generate google/rpc/code.pb.go
	@curl -sS $(GOOGLEAPIS_URL)/google/rpc/code.proto -o $(CODE_PROTO)
	@gogoreplace 'option go_package = "google.golang.org/genproto/googleapis/rpc/code;code";' '' $(CODE_PROTO)
	@$(PROTOC) $(PROTO_PATH) $(PLUGIN) $(CODE_PROTO)

	## Generate google/rpc/error_details.pb.go
	@curl -sS $(GOOGLEAPIS_URL)/google/rpc/error_details.proto -o $(ERR_PROTO)
	@gogoreplace 'option go_package = "google.golang.org/genproto/googleapis/rpc/errdetails;errdetails";' '' $(ERR_PROTO)
	@$(PROTOC) $(PROTO_PATH) $(PLUGIN) $(ERR_PROTO)

	# Format code
	@gofmt -l -s -w .

	# Cleanup
	@rm -rf protoc-tmp

clean:
	rm -rf protoc-tmp