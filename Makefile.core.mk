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
	k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto=istio.io/gogo-genproto/k8s.io/apimachinery/pkg/apis/meta/v1beta1

comma := ,
empty :=
space := $(empty) $(empty)
mapping_with_spaces := $(foreach map,$(importmaps),M$(map),)
MAPPING := $(subst $(space),$(empty),$(mapping_with_spaces))
GOGOSLICK_PLUGIN := --gogoslick_out=plugins=grpc,$(MAPPING)
GOGOFASTER_PLUGIN := --gogofaster_out=plugins=grpc,$(MAPPING)

google_protos = \
	google/protobuf/any.proto \
	google/protobuf/duration.proto \
	google/protobuf/descriptor.proto \
	google/protobuf/empty.proto \
	google/protobuf/struct.proto \
	google/protobuf/timestamp.proto \
	google/protobuf/wrappers.proto \
	google/api/field_behavior.proto \
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

google_packages = \
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

k8s_protos = \
	api/admission/v1/generated.proto \
	api/admission/v1beta1/generated.proto \
	api/admissionregistration/v1/generated.proto \
	api/admissionregistration/v1beta1/generated.proto \
	api/apps/v1/generated.proto \
	api/apps/v1beta1/generated.proto \
	api/apps/v1beta2/generated.proto \
	api/auditregistration/v1alpha1/generated.proto \
	api/authentication/v1/generated.proto \
	api/authentication/v1beta1/generated.proto \
	api/authorization/v1/generated.proto \
	api/authorization/v1beta1/generated.proto \
	api/autoscaling/v1/generated.proto \
	api/autoscaling/v2beta1/generated.proto \
	api/autoscaling/v2beta2/generated.proto \
	api/batch/v1/generated.proto \
	api/batch/v1beta1/generated.proto \
	api/batch/v2alpha1/generated.proto \
	api/certificates/v1beta1/generated.proto \
	api/coordination/v1/generated.proto \
	api/coordination/v1beta1/generated.proto \
	api/core/v1/generated.proto \
	api/events/v1beta1/generated.proto \
	api/extensions/v1beta1/generated.proto \
	api/imagepolicy/v1alpha1/generated.proto \
	api/networking/v1/generated.proto \
	api/networking/v1beta1/generated.proto \
	api/node/v1alpha1/generated.proto \
	api/node/v1beta1/generated.proto \
	api/policy/v1beta1/generated.proto \
	api/rbac/v1/generated.proto \
	api/rbac/v1alpha1/generated.proto \
	api/rbac/v1beta1/generated.proto \
	api/scheduling/v1/generated.proto \
	api/scheduling/v1alpha1/generated.proto \
	api/scheduling/v1beta1/generated.proto \
	api/settings/v1alpha1/generated.proto \
	api/storage/v1/generated.proto \
	api/storage/v1alpha1/generated.proto \
	api/storage/v1beta1/generated.proto \
	apimachinery/pkg/runtime/generated.proto \
	apimachinery/pkg/api/resource/generated.proto \
	apimachinery/pkg/runtime/schema/generated.proto \
	apimachinery/pkg/util/intstr/generated.proto \
	apimachinery/pkg/apis/meta/v1/generated.proto \
	apimachinery/pkg/apis/meta/v1beta1/generated.proto

k8s_packages = \
	api/admission/v1 \
	api/admission/v1beta1 \
	api/admissionregistration/v1 \
	api/admissionregistration/v1beta1 \
	api/apps/v1 \
	api/apps/v1beta1 \
	api/apps/v1beta2 \
	api/auditregistration/v1alpha1 \
	api/authentication/v1 \
	api/authentication/v1beta1 \
	api/authorization/v1 \
	api/authorization/v1beta1 \
	api/autoscaling/v1 \
	api/autoscaling/v2beta1 \
	api/autoscaling/v2beta2 \
	api/batch/v1 \
	api/batch/v1beta1 \
	api/batch/v2alpha1 \
	api/certificates/v1beta1 \
	api/coordination/v1 \
	api/coordination/v1beta1 \
	api/core/v1 \
	api/events/v1beta1 \
	api/extensions/v1beta1 \
	api/imagepolicy/v1alpha1 \
	api/networking/v1 \
	api/networking/v1beta1 \
	api/node/v1alpha1 \
	api/node/v1beta1 \
	api/policy/v1beta1 \
	api/rbac/v1 \
	api/rbac/v1alpha1 \
	api/rbac/v1beta1 \
	api/scheduling/v1 \
	api/scheduling/v1alpha1 \
	api/scheduling/v1beta1 \
	api/settings/v1alpha1 \
	api/storage/v1 \
	api/storage/v1alpha1 \
	api/storage/v1beta1 \
	apimachinery/pkg/runtime \
	apimachinery/pkg/api/resource \
	apimachinery/pkg/runtime/schema \
	apimachinery/pkg/util/intstr \
	apimachinery/pkg/apis/meta/v1 \
	apimachinery/pkg/apis/meta/v1beta1

all: build

TMPDIR := $(shell mktemp -d)

google_prep:
	@mkdir -p ${TMPDIR}/google/ googleapis
	@cp -r common-protos/google/* ${TMPDIR}/google/
	@rm -fr ${TMPDIR}/google/api/*.proto
	@cp -r common-protos/google/api/http.proto ${TMPDIR}/google/api
	@cp -r common-protos/google/api/annotations.proto ${TMPDIR}/google/api
	@cp -r common-protos/google/api/field_behavior.proto ${TMPDIR}/google/api
	@rm -fr ${TMPDIR}/google/type/calendar_period.proto
	@rm -fr ${TMPDIR}/google/type/quaternion.proto
	@rm -fr ${TMPDIR}/google/type/expr.proto
	@rm -fr ${TMPDIR}/google/type/fraction.proto

$(google_protos): %: google_prep
	@sed -i -e '/^option go_package/d' ${TMPDIR}/$@

$(google_packages): %: $(google_protos)
	@$(PROTOC) $(GOGOSLICK_PLUGIN):googleapis -I${TMPDIR} ${TMPDIR}/$@/*.proto

$(googleexpr_protos): %: google_prep
	@sed -i -e '/^option go_package/d' ${TMPDIR}/$@

$(googleexpr_packages): %: $(googleexpr_protos)
	@$(PROTOC) $(GOGOFASTER_PLUGIN):googleapis -I${TMPDIR} ${TMPDIR}/$@/*.proto

census_prep:
	@mkdir -p ${TMPDIR}/opencensus opencensus
	@cp -r common-protos/github.com/census-instrumentation/opencensus-proto/src/opencensus/* ${TMPDIR}/opencensus

$(census_protos): %: google_prep census_prep
	@sed -i -e '/^option go_package/d' ${TMPDIR}/$@

$(census_packages): %: $(census_protos)
	@$(PROTOC) $(GOGOFASTER_PLUGIN):. -I${TMPDIR} ${TMPDIR}/$@/*.proto

prometheus_prep:
	@mkdir -p ${TMPDIR}/prometheus prometheus
	@cp common-protos/github.com/prometheus/client_model/metrics.proto ${TMPDIR}/prometheus/metrics.proto
	@sed -i -e '/^option go_package/d' ${TMPDIR}/prometheus/metrics.proto

prometheus/metrics.pb.go: prometheus_prep
	@$(PROTOC) $(GOGOFASTER_PLUGIN):prometheus -I${TMPDIR}/prometheus ${TMPDIR}/prometheus/metrics.proto

k8s_prep:
	@mkdir -p ${TMPDIR}/k8s.io/
	@cp -r common-protos/k8s.io/* ${TMPDIR}/k8s.io/

$(k8s_protos): %:
	@sed -i -e '/^option go_package/d' ${TMPDIR}/k8s.io/$@

$(k8s_packages): %: k8s_prep $(k8s_protos)
	@$(PROTOC) $(GOGOSLICK_PLUGIN):. -I${TMPDIR} ${TMPDIR}/k8s.io/$@/*.proto

gen: clean $(google_packages) $(googleexpr_packages) $(k8s_packages) $(census_packages) prometheus/metrics.pb.go tidy-go mirror-licenses

gen-check: gen check-clean-repo

build:
	@go build ./...

clean:
	@rm -fr googleapis opencensus prometheus k8s.io

lint: lint-all

.PHONY: all build clean

include common/Makefile.common.mk
