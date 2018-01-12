GOOGLEAPIS := googleapis

# BUILD UP PROTOC ARGs
PROTO_PATH := --proto_path=${GOOGLEAPIS}

MAPPING := Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,
MAPPING := $(MAPPING)Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,
MAPPING := $(MAPPING)Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,
PLUGIN := --gogoslick_out=$(MAPPING):$(GOOGLEAPIS)

STATUS_PROTO := google/rpc/status.proto
CODE_PROTO := google/rpc/code.proto
ERR_PROTO := google/rpc/error_details.proto

PROTOC = protoc-min-version -version=3.5.0

SHA = c8c975543a134177cc41b64cbbf10b88fe66aa1d
GOOGLEAPIS_URL = https://raw.githubusercontent.com/googleapis/googleapis/$(SHA)

update:
	@dep ensure

	# Install binaries
	@go install github.com/gogo/protobuf/protoc-gen-gogoslick
	@go install github.com/gogo/protobuf/gogoreplace
	@go install github.com/gogo/protobuf/protoc-min-version

	# Record protoc version
	@echo `protoc --version` > protoc.version

	# GOOGLEAPIS Generation
	@mkdir -p googleapis/google/rpc

	## Generate google/rpc/status.pb.go
	@curl -sS $(GOOGLEAPIS_URL)/$(STATUS_PROTO) -o googleapis/$(STATUS_PROTO)
	@gogoreplace 'option go_package = "google.golang.org/genproto/googleapis/rpc/status;status";' '' googleapis/$(STATUS_PROTO)
	@$(PROTOC) $(PROTO_PATH) $(PLUGIN) $(STATUS_PROTO)

	## Generate google/rpc/code.pb.go
	@curl -sS $(GOOGLEAPIS_URL)/$(CODE_PROTO) -o googleapis/$(CODE_PROTO)
	@gogoreplace 'option go_package = "google.golang.org/genproto/googleapis/rpc/code;code";' '' googleapis/$(CODE_PROTO)
	@$(PROTOC) $(PROTO_PATH) $(PLUGIN) $(CODE_PROTO)

	## Generate google/rpc/error_details.pb.go
	@curl -sS $(GOOGLEAPIS_URL)/$(ERR_PROTO) -o googleapis/$(ERR_PROTO)
	@gogoreplace 'option go_package = "google.golang.org/genproto/googleapis/rpc/errdetails;errdetails";' '' googleapis/$(ERR_PROTO)
	@$(PROTOC) $(PROTO_PATH) $(PLUGIN) $(ERR_PROTO)

	# Format code
	@gofmt -l -s -w .

gofmt:
	gofmt -l -s -w .

.PHONY: update gofmt
