# Copyright 2019 Istio Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

PROTOC = protoc

GOOGLEPROTOBUF_SHA = 63e4a3ecc956cbab6714b25e8b868765ea7e6fe5
GOOGLEPROTOBUF_URL = https://raw.githubusercontent.com/protocolbuffers/protobuf/$(GOOGLEPROTOBUF_SHA)/src

GOOGLEAPIS_SHA = 7ebb7a62ed598d4e7e6cc41403cbf191f71c079d
GOOGLEAPIS_URL = https://raw.githubusercontent.com/googleapis/googleapis/$(GOOGLEAPIS_SHA)

CENSUS_SHA = e2601ef16f8a085a69d94ace5133f97438f8945f
CENSUS_URL = https://raw.githubusercontent.com/census-instrumentation/opencensus-proto/$(CENSUS_SHA)/src

PROMETHEUS_SHA = 6f3806018612930941127f2a7c6c453ba2c527d2
PROMETHEUS_URL = https://raw.githubusercontent.com/prometheus/client_model/$(PROMETHEUS_SHA)

K8SAPI_SHA = d651a1528133f00fab57f5bf42fb26500e8aa406
K8SAPI_URL = https://raw.githubusercontent.com/kubernetes/api/$(K8SAPI_SHA)

K8SAPIMACHINERY_SHA = ac02f8882ef690d3030ad4230a947f01d83fafe9
K8SAPIMACHINERY_URL = https://raw.githubusercontent.com/kubernetes/apimachinery/$(K8SAPIMACHINERY_SHA)

GOGO_PROTO_PKG := github.com/gogo/protobuf/gogoproto
GOGO_TYPES := github.com/gogo/protobuf/types
GOGO_DESCRIPTOR := github.com/gogo/protobuf/protoc-gen-gogo/descriptor

importmaps := \
	gogoproto/gogo.proto=$(GOGO_PROTO_PKG) \
	google/protobuf/any.proto=$(GOGO_TYPES) \
	google/protobuf/descriptor.proto=$(GOGO_DESCRIPTOR) \
	google/protobuf/duration.proto=$(GOGO_TYPES) \
	google/protobuf/empty.proto=$(GOGO_TYPES) \
	google/protobuf/timestamp.proto=$(GOGO_TYPES) \
	google/protobuf/wrappers.proto=$(GOGO_TYPES) \
	google/protobuf/struct.proto=$(GOGO_TYPES) \
	google/rpc/code.proto=istio.io/gogo-genproto/googleapis/google/rpc \
  google/rpc/error_details.proto=istio.io/gogo-genproto/googleapis/google/rpc \
  google/rpc/status.proto=istio.io/gogo-genproto/googleapis/google/rpc \
	opencensus/proto/resource/v1/resource.proto=istio.io/gogo-genproto/opencensus/proto/resource/v1 \
	k8s.io/api/admission/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/admission/v1 \
	k8s.io/api/admission/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/admission/v1beta1 \
	k8s.io/api/admissionregistration/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/admissionregistration/v1 \
	k8s.io/api/admissionregistration/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/admissionregistration/v1beta1 \
	k8s.io/api/apps/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/apps/v1 \
	k8s.io/api/apps/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/apps/v1beta1 \
	k8s.io/api/apps/v1beta2/generated.proto=istio.io/gogo-genproto/k8s.io/api/apps/v1beta2 \
	k8s.io/api/auditregistration/v1alpha1/generated.proto=istio.io/gogo-genproto/k8s.io/api/auditregistration/v1alpha1 \
	k8s.io/api/authentication/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/authentication/v1 \
	k8s.io/api/authentication/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/authentication/v1beta1 \
	k8s.io/api/authorization/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/authorization/v1 \
	k8s.io/api/authorization/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/authorization/v1beta1 \
	k8s.io/api/autoscaling/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/autoscaling/v1 \
	k8s.io/api/autoscaling/v2beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/autoscaling/v2beta1 \
	k8s.io/api/autoscaling/v2beta2/generated.proto=istio.io/gogo-genproto/k8s.io/api/autoscaling/v2beta2 \
	k8s.io/api/batch/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/batch/v1 \
	k8s.io/api/batch/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/batch/v1beta1 \
	k8s.io/api/batch/v2alpha1/generated.proto=istio.io/gogo-genproto/k8s.io/api/batch/v2alpha1 \
	k8s.io/api/certificates/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/certificates/v1beta1 \
	k8s.io/api/coordination/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/coordination/v1 \
	k8s.io/api/coordination/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/coordination/v1beta1 \
	k8s.io/api/core/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/core/v1 \
	k8s.io/api/events/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/events/v1beta1 \
	k8s.io/api/extensions/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/extensions/v1beta1 \
	k8s.io/api/imagepolicy/v1alpha1/generated.proto=istio.io/gogo-genproto/k8s.io/api/imagepolicy/v1alpha1 \
	k8s.io/api/networking/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/networking/v1 \
	k8s.io/api/networking/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/networking/v1beta1 \
	k8s.io/api/node/v1alpha1/generated.proto=istio.io/gogo-genproto/k8s.io/api/node/v1alpha1 \
	k8s.io/api/node/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/node/v1beta1 \
	k8s.io/api/policy/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/policy/v1beta1 \
	k8s.io/api/rbac/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/rbac/v1 \
	k8s.io/api/rbac/v1alpha1/generated.proto=istio.io/gogo-genproto/k8s.io/api/rbac/v1alpha1 \
	k8s.io/api/rbac/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/rbac/v1beta1 \
	k8s.io/api/scheduling/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/scheduling/v1 \
	k8s.io/api/scheduling/v1alpha1/generated.proto=istio.io/gogo-genproto/k8s.io/api/scheduling/v1alpha1 \
	k8s.io/api/scheduling/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/scheduling/v1beta1 \
	k8s.io/api/settings/v1alpha1/generated.proto=istio.io/gogo-genproto/k8s.io/api/settings/v1alpha1 \
	k8s.io/api/storage/v1/generated.proto=istio.io/gogo-genproto/k8s.io/api/storage/v1 \
	k8s.io/api/storage/v1alpha1/generated.proto=istio.io/gogo-genproto/k8s.io/api/storage/v1alpha1 \
	k8s.io/api/storage/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/api/storage/v1beta1 \
	k8s.io/apimachinery/pkg/runtime/generated.proto=istio.io/gogo-genproto/k8s.io/apimachinery/pkg/runtime \
	k8s.io/apimachinery/pkg/api/resource/generated.proto=istio.io/gogo-genproto/k8s.io/apimachinery/pkg/api/resource \
	k8s.io/apimachinery/pkg/runtime/schema/generated.proto=istio.io/gogo-genproto/k8s.io/apimachinery/pkg/runtime/schema \
	k8s.io/apimachinery/pkg/util/intstr/generated.proto=istio.io/gogo-genproto/k8s.io/apimachinery/pkg/util/intstr \
	k8s.io/apimachinery/pkg/apis/meta/v1/generated.proto=istio.io/gogo-genproto/k8s.io/apimachinery/pkg/apis/meta/v1 \
	k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/apimachinery/pkg/apis/meta/v1beta1 \
	k8s.io/apimachinery/pkg/apis/testapigroup/v1/generated.proto=istio.io/gogo-genproto/k8s.io/apimachinery/pkg/apis/testapigroup/v1/generated.proto

comma := ,
empty :=
space := $(empty) $(empty)
mapping_with_spaces := $(foreach map,$(importmaps),M$(map),)
MAPPING := $(subst $(space),$(empty),$(mapping_with_spaces))
GOGOSLICK_PLUGIN := --gogoslick_out=plugins=grpc,$(MAPPING)
GOGOFASTER_PLUGIN := --gogofaster_out=plugins=grpc,$(MAPPING)

googleprotobuf_protos = \
	google/protobuf/any.proto \
	google/protobuf/duration.proto \
	google/protobuf/descriptor.proto \
	google/protobuf/empty.proto \
	google/protobuf/struct.proto \
	google/protobuf/timestamp.proto \
	google/protobuf/wrappers.proto

googleprotobuf_packages = \
	google/protobuf

googleapis_protos = \
	google/api/http.proto \
	google/api/annotations.proto \
	google/rpc/status.proto \
	google/rpc/code.proto \
	google/rpc/error_details.proto \
	google/type/color.proto \
	google/type/date.proto \
	google/type/dayofweek.proto \
	google/type/latlng.proto \
	google/type/money.proto \
	google/type/postal_address.proto \
	google/type/timeofday.proto

googleapis_packages = \
	google/api \
	google/rpc \
	google/type

googleexpr_protos = \
	google/api/expr/v1alpha1/cel_service.proto \
	google/api/expr/v1alpha1/checked.proto \
	google/api/expr/v1alpha1/conformance_service.proto \
	google/api/expr/v1alpha1/eval.proto \
	google/api/expr/v1alpha1/explain.proto \
	google/api/expr/v1alpha1/syntax.proto \
	google/api/expr/v1alpha1/value.proto

googleexpr_packages = \
	google/api/expr/v1alpha1

census_protos = \
	opencensus/proto/stats/v1/stats.proto \
	opencensus/proto/trace/v1/trace.proto \
	opencensus/proto/trace/v1/trace_config.proto \
	opencensus/proto/resource/v1/resource.proto

census_packages = \
	opencensus/proto/stats/v1 \
	opencensus/proto/trace/v1 \
	opencensus/proto/resource/v1

k8sapi_protos = \
	admission/v1/generated.proto \
	admission/v1beta1/generated.proto \
	admissionregistration/v1/generated.proto \
	admissionregistration/v1beta1/generated.proto \
	apps/v1/generated.proto \
	apps/v1beta1/generated.proto \
	apps/v1beta2/generated.proto \
	auditregistration/v1alpha1/generated.proto \
	authentication/v1/generated.proto \
	authentication/v1beta1/generated.proto \
	authorization/v1/generated.proto \
	authorization/v1beta1/generated.proto \
	autoscaling/v1/generated.proto \
	autoscaling/v2beta1/generated.proto \
	autoscaling/v2beta2/generated.proto \
	batch/v1/generated.proto \
	batch/v1beta1/generated.proto \
	batch/v2alpha1/generated.proto \
	certificates/v1beta1/generated.proto \
	coordination/v1/generated.proto \
	coordination/v1beta1/generated.proto \
	core/v1/generated.proto \
	events/v1beta1/generated.proto \
	extensions/v1beta1/generated.proto \
	imagepolicy/v1alpha1/generated.proto \
	networking/v1/generated.proto \
	networking/v1beta1/generated.proto \
	node/v1alpha1/generated.proto \
	node/v1beta1/generated.proto \
	policy/v1beta1/generated.proto \
	rbac/v1/generated.proto \
	rbac/v1alpha1/generated.proto \
	rbac/v1beta1/generated.proto \
	scheduling/v1/generated.proto \
	scheduling/v1alpha1/generated.proto \
	scheduling/v1beta1/generated.proto \
	settings/v1alpha1/generated.proto \
	storage/v1/generated.proto \
	storage/v1alpha1/generated.proto \
	storage/v1beta1/generated.proto

k8sapi_packages = \
	admission/v1 \
	admission/v1beta1 \
	admissionregistration/v1 \
	admissionregistration/v1beta1 \
	apps/v1 \
	apps/v1beta1 \
	apps/v1beta2 \
	auditregistration/v1alpha1 \
	authentication/v1 \
	authentication/v1beta1 \
	authorization/v1 \
	authorization/v1beta1 \
	autoscaling/v1 \
	autoscaling/v2beta1 \
	autoscaling/v2beta2 \
	batch/v1 \
	batch/v1beta1 \
	batch/v2alpha1 \
	certificates/v1beta1 \
	coordination/v1 \
	coordination/v1beta1 \
	core/v1 \
	events/v1beta1 \
	extensions/v1beta1 \
	imagepolicy/v1alpha1 \
	networking/v1 \
	networking/v1beta1 \
	node/v1alpha1 \
	node/v1beta1 \
	policy/v1beta1 \
	rbac/v1 \
	rbac/v1alpha1 \
	rbac/v1beta1 \
	scheduling/v1 \
	scheduling/v1alpha1 \
	scheduling/v1beta1 \
	settings/v1alpha1 \
	storage/v1 \
	storage/v1alpha1 \
	storage/v1beta1

k8sapimachinery_protos = \
	pkg/runtime/generated.proto \
	pkg/api/resource/generated.proto \
	pkg/runtime/schema/generated.proto \
	pkg/util/intstr/generated.proto \
	pkg/apis/meta/v1/generated.proto \
	pkg/apis/meta/v1beta1/generated.proto \
	pkg/apis/testapigroup/v1/generated.proto

k8sapimachinery_packages = \
	pkg/runtime \
	pkg/api/resource \
	pkg/runtime/schema \
	pkg/util/intstr \
	pkg/apis/meta/v1 \
	pkg/apis/meta/v1beta1 \
	pkg/apis/testapigroup/v1

all: build

$(googleprotobuf_protos): %:
	@mkdir -p googleprotobuf/google/protobuf
	@curl -sS $(GOOGLEPROTOBUF_URL)/$@ -o googleprotobuf/$@

$(googleprotobuf_packages): %: $(googleprotobuf_protos)

$(googleapis_protos) $(googleexpr_protos): %:
	@mkdir -p googleapis/google/rpc googleapis/google/type googleapis/google/api/expr/v1alpha1
	@curl -sS $(GOOGLEAPIS_URL)/$@ -o googleapis/$@
	@sed -i -e '/^option go_package/d' googleapis/$@

$(googleapis_packages): %: $(googleapis_protos)
	@$(PROTOC) $(GOGOSLICK_PLUGIN):googleapis -Igoogleprotobuf -Igoogleapis googleapis/$@/*.proto

$(googleexpr_packages): %: $(googleexpr_protos)
	@$(PROTOC) $(GOGOFASTER_PLUGIN):googleapis -Igoogleprotobuf -Igoogleapis googleapis/$@/*.proto

$(census_protos): %:
	@mkdir -p oc/opencensus/proto/stats/v1 oc/opencensus/proto/trace/v1 oc/opencensus/proto/resource/v1
	@curl -sS $(CENSUS_URL)/$@ -o oc/$@
	@sed -i -e '/^option go_package/d' oc/$@

$(census_packages): %: $(census_protos)
	@$(PROTOC) $(GOGOFASTER_PLUGIN):oc -Igoogleprotobuf -Ioc oc/$@/*.proto

prometheus/metrics.proto:
	@mkdir -p prometheus
	@curl -sS $(PROMETHEUS_URL)/metrics.proto -o prometheus/metrics.proto

prometheus/metrics.pb.go: prometheus/metrics.proto
	@$(PROTOC) $(GOGOFASTER_PLUGIN):. prometheus/metrics.proto

k8sapi_prep:
	@mkdir -p k8sapi/k8s.io/api
	@cd k8sapi/k8s.io/api && mkdir -p $(k8sapi_packages)

$(k8sapi_protos): %:
	@curl -sS $(K8SAPI_URL)/$@ -o k8sapi/k8s.io/api/$@
	@sed -i -e '/^option go_package/d' k8sapi/k8s.io/api/$@

$(k8sapi_packages): %: k8sapi_prep $(k8sapi_protos)
	@$(PROTOC) $(GOGOSLICK_PLUGIN):k8sapi -Igoogleprotobuf -Ik8sapimachinery -Ik8sapi k8sapi/k8s.io/api/$@/*.proto

k8sapimachinery_prep:
	@mkdir -p k8sapimachinery/k8s.io/apimachinery
	@cd k8sapimachinery/k8s.io/apimachinery && mkdir -p $(k8sapimachinery_packages)

$(k8sapimachinery_protos): %:
	@curl -sS $(K8SAPIMACHINERY_URL)/$@ -o k8sapimachinery/k8s.io/apimachinery/$@
	@sed -i -e '/^option go_package/d' k8sapimachinery/k8s.io/apimachinery/$@

$(k8sapimachinery_packages): %: k8sapimachinery_prep $(k8sapimachinery_protos)
	@$(PROTOC) $(GOGOSLICK_PLUGIN):k8sapimachinery -Igoogleprotobuf -Ik8sapimachinery k8sapimachinery/k8s.io/apimachinery/$@/*.proto

generate: clean $(googleprotobuf_packages) $(googleapis_packages) $(googleexpr_packages) $(census_packages) prometheus/metrics.pb.go $(k8sapimachinery_packages) $(k8sapi_packages)
	@mv oc/opencensus ./opencensus
	@rm -fr oc
	@mv k8sapi/k8s.io ./k8s.io
	@rm -fr k8sapi
	@mv k8sapimachinery/k8s.io/apimachinery ./k8s.io/apimachinery
	@rm -fr k8sapimachinery

build: generate
	@go build ./...

clean:
	@rm -fr googleprotobuf googleapis opencensus prometheus k8sapimachinery k8sapi k8s.io tmp

.PHONY: all build $(googleapis_protos) $(googleapis_packages) clean

include Makefile.common.mk
