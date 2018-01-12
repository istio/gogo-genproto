# BUILD UP PROTOC ARGs
PROTO_PATH := --proto_path=googleapis

MAPPING := Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,
MAPPING := $(MAPPING)Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,
MAPPING := $(MAPPING)Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,
PLUGIN := --plugin=protoc-gen-gogoslick=gogoslick --gogoslick_out=$(MAPPING):googleapis
PROTOC = protoc

SHA = c8c975543a134177cc41b64cbbf10b88fe66aa1d
GOOGLEAPIS_URL = https://raw.githubusercontent.com/googleapis/googleapis/$(SHA)

googleapis_protos = \
	google/rpc/status.proto \
	google/rpc/code.proto \
	google/rpc/error_details.proto

all: depend $(googleapis_protos) format

depend:
	@dep ensure

protoc.version:
	# Record protoc version
	@echo `protoc --version` > protoc.version

gogoslick: depend
	@go build -o gogoslick vendor/github.com/gogo/protobuf/protoc-gen-gogoslick/main.go

$(googleapis_protos): %: gogoslick protoc.version
	## Download $@
	@curl -sS $(GOOGLEAPIS_URL)/$@ -o googleapis/$@.tmp
	@sed -e '/^option go_package/d' googleapis/$@.tmp > googleapis/$@
	@rm googleapis/$@.tmp
	## Generate $@
	@$(PROTOC) $(PROTO_PATH) $(PLUGIN) $@

format: $(googleapis_protos)
	# Format code
	@gofmt -l -s -w .

clean:
	@rm gogoslick

.PHONY: all depend format $(googleapis_protos) protoc.version clean
