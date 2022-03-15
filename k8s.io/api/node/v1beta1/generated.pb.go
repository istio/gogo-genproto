// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/api/node/v1beta1/generated.proto

package k8s_io_api_node_v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	v11 "istio.io/gogo-genproto/k8s.io/api/core/v1"
	resource "istio.io/gogo-genproto/k8s.io/apimachinery/pkg/api/resource"
	v1 "istio.io/gogo-genproto/k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "istio.io/gogo-genproto/k8s.io/apimachinery/pkg/runtime"
	_ "istio.io/gogo-genproto/k8s.io/apimachinery/pkg/runtime/schema"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Overhead structure represents the resource overhead associated with running a pod.
type Overhead struct {
	// PodFixed represents the fixed resource overhead associated with running a pod.
	// +optional
	PodFixed             map[string]*resource.Quantity `protobuf:"bytes,1,rep,name=podFixed" json:"podFixed,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Overhead) Reset()         { *m = Overhead{} }
func (m *Overhead) String() string { return proto.CompactTextString(m) }
func (*Overhead) ProtoMessage()    {}
func (*Overhead) Descriptor() ([]byte, []int) {
	return fileDescriptor_73bb62abe8438af4, []int{0}
}
func (m *Overhead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Overhead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Overhead.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Overhead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Overhead.Merge(m, src)
}
func (m *Overhead) XXX_Size() int {
	return m.Size()
}
func (m *Overhead) XXX_DiscardUnknown() {
	xxx_messageInfo_Overhead.DiscardUnknown(m)
}

var xxx_messageInfo_Overhead proto.InternalMessageInfo

func (m *Overhead) GetPodFixed() map[string]*resource.Quantity {
	if m != nil {
		return m.PodFixed
	}
	return nil
}

// RuntimeClass defines a class of container runtime supported in the cluster.
// The RuntimeClass is used to determine which container runtime is used to run
// all containers in a pod. RuntimeClasses are (currently) manually defined by a
// user or cluster provisioner, and referenced in the PodSpec. The Kubelet is
// responsible for resolving the RuntimeClassName reference before running the
// pod.  For more details, see
// https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
type RuntimeClass struct {
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Handler specifies the underlying runtime and configuration that the CRI
	// implementation will use to handle pods of this class. The possible values
	// are specific to the node & CRI configuration.  It is assumed that all
	// handlers are available on every node, and handlers of the same name are
	// equivalent on every node.
	// For example, a handler called "runc" might specify that the runc OCI
	// runtime (using native Linux containers) will be used to run the containers
	// in a pod.
	// The Handler must conform to the DNS Label (RFC 1123) requirements, and is
	// immutable.
	Handler *string `protobuf:"bytes,2,opt,name=handler" json:"handler,omitempty"`
	// Overhead represents the resource overhead associated with running a pod for a
	// given RuntimeClass. For more details, see
	// https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md
	// This field is alpha-level as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
	// +optional
	Overhead *Overhead `protobuf:"bytes,3,opt,name=overhead" json:"overhead,omitempty"`
	// Scheduling holds the scheduling constraints to ensure that pods running
	// with this RuntimeClass are scheduled to nodes that support it.
	// If scheduling is nil, this RuntimeClass is assumed to be supported by all
	// nodes.
	// +optional
	Scheduling           *Scheduling `protobuf:"bytes,4,opt,name=scheduling" json:"scheduling,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RuntimeClass) Reset()         { *m = RuntimeClass{} }
func (m *RuntimeClass) String() string { return proto.CompactTextString(m) }
func (*RuntimeClass) ProtoMessage()    {}
func (*RuntimeClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_73bb62abe8438af4, []int{1}
}
func (m *RuntimeClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuntimeClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RuntimeClass.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RuntimeClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeClass.Merge(m, src)
}
func (m *RuntimeClass) XXX_Size() int {
	return m.Size()
}
func (m *RuntimeClass) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeClass.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeClass proto.InternalMessageInfo

func (m *RuntimeClass) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *RuntimeClass) GetHandler() string {
	if m != nil && m.Handler != nil {
		return *m.Handler
	}
	return ""
}

func (m *RuntimeClass) GetOverhead() *Overhead {
	if m != nil {
		return m.Overhead
	}
	return nil
}

func (m *RuntimeClass) GetScheduling() *Scheduling {
	if m != nil {
		return m.Scheduling
	}
	return nil
}

// RuntimeClassList is a list of RuntimeClass objects.
type RuntimeClassList struct {
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Items is a list of schema objects.
	Items                []*RuntimeClass `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RuntimeClassList) Reset()         { *m = RuntimeClassList{} }
func (m *RuntimeClassList) String() string { return proto.CompactTextString(m) }
func (*RuntimeClassList) ProtoMessage()    {}
func (*RuntimeClassList) Descriptor() ([]byte, []int) {
	return fileDescriptor_73bb62abe8438af4, []int{2}
}
func (m *RuntimeClassList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuntimeClassList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RuntimeClassList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RuntimeClassList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeClassList.Merge(m, src)
}
func (m *RuntimeClassList) XXX_Size() int {
	return m.Size()
}
func (m *RuntimeClassList) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeClassList.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeClassList proto.InternalMessageInfo

func (m *RuntimeClassList) GetMetadata() *v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *RuntimeClassList) GetItems() []*RuntimeClass {
	if m != nil {
		return m.Items
	}
	return nil
}

// Scheduling specifies the scheduling constraints for nodes supporting a
// RuntimeClass.
type Scheduling struct {
	// nodeSelector lists labels that must be present on nodes that support this
	// RuntimeClass. Pods using this RuntimeClass can only be scheduled to a
	// node matched by this selector. The RuntimeClass nodeSelector is merged
	// with a pod's existing nodeSelector. Any conflicts will cause the pod to
	// be rejected in admission.
	// +optional
	NodeSelector map[string]string `protobuf:"bytes,1,rep,name=nodeSelector" json:"nodeSelector,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// tolerations are appended (excluding duplicates) to pods running with this
	// RuntimeClass during admission, effectively unioning the set of nodes
	// tolerated by the pod and the RuntimeClass.
	// +optional
	// +listType=atomic
	Tolerations          []*v11.Toleration `protobuf:"bytes,2,rep,name=tolerations" json:"tolerations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Scheduling) Reset()         { *m = Scheduling{} }
func (m *Scheduling) String() string { return proto.CompactTextString(m) }
func (*Scheduling) ProtoMessage()    {}
func (*Scheduling) Descriptor() ([]byte, []int) {
	return fileDescriptor_73bb62abe8438af4, []int{3}
}
func (m *Scheduling) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scheduling) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Scheduling.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Scheduling) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scheduling.Merge(m, src)
}
func (m *Scheduling) XXX_Size() int {
	return m.Size()
}
func (m *Scheduling) XXX_DiscardUnknown() {
	xxx_messageInfo_Scheduling.DiscardUnknown(m)
}

var xxx_messageInfo_Scheduling proto.InternalMessageInfo

func (m *Scheduling) GetNodeSelector() map[string]string {
	if m != nil {
		return m.NodeSelector
	}
	return nil
}

func (m *Scheduling) GetTolerations() []*v11.Toleration {
	if m != nil {
		return m.Tolerations
	}
	return nil
}

func init() {
	proto.RegisterType((*Overhead)(nil), "k8s.io.api.node.v1beta1.Overhead")
	proto.RegisterMapType((map[string]*resource.Quantity)(nil), "k8s.io.api.node.v1beta1.Overhead.PodFixedEntry")
	proto.RegisterType((*RuntimeClass)(nil), "k8s.io.api.node.v1beta1.RuntimeClass")
	proto.RegisterType((*RuntimeClassList)(nil), "k8s.io.api.node.v1beta1.RuntimeClassList")
	proto.RegisterType((*Scheduling)(nil), "k8s.io.api.node.v1beta1.Scheduling")
	proto.RegisterMapType((map[string]string)(nil), "k8s.io.api.node.v1beta1.Scheduling.NodeSelectorEntry")
}

func init() {
	proto.RegisterFile("k8s.io/api/node/v1beta1/generated.proto", fileDescriptor_73bb62abe8438af4)
}

var fileDescriptor_73bb62abe8438af4 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x35, 0xc9, 0x57, 0x7d, 0xc9, 0x4d, 0x90, 0xca, 0x08, 0x09, 0x2b, 0x8b, 0x28, 0x04,
	0x21, 0xb2, 0x1a, 0x93, 0x08, 0x50, 0x05, 0x42, 0x20, 0x0a, 0x2c, 0xa0, 0x50, 0x98, 0xb2, 0x61,
	0x39, 0xb5, 0xaf, 0x92, 0xc1, 0x7f, 0xc6, 0x1a, 0x8f, 0x23, 0xf2, 0x2c, 0xbc, 0x0e, 0x0b, 0x96,
	0x3c, 0x02, 0xca, 0x92, 0x47, 0x60, 0x85, 0xc6, 0x4e, 0xdc, 0x69, 0xd2, 0xd0, 0xb2, 0x4b, 0xe4,
	0xfb, 0x3b, 0xe7, 0xde, 0xe3, 0x63, 0xb8, 0x1b, 0x1d, 0xe4, 0x4c, 0x2a, 0x5f, 0x64, 0xd2, 0x4f,
	0x55, 0x88, 0xfe, 0x7c, 0x7c, 0x8a, 0x46, 0x8c, 0xfd, 0x29, 0xa6, 0xa8, 0x85, 0xc1, 0x90, 0x65,
	0x5a, 0x19, 0x45, 0x6f, 0x56, 0x83, 0x4c, 0x64, 0x92, 0xd9, 0x41, 0xb6, 0x1a, 0xec, 0x0d, 0x1d,
	0x85, 0x40, 0x69, 0xab, 0xb0, 0x09, 0xf7, 0xee, 0x9f, 0xcd, 0x24, 0x22, 0x98, 0xc9, 0x14, 0xf5,
	0xc2, 0xcf, 0xa2, 0x69, 0x09, 0x69, 0xcc, 0x55, 0xa1, 0x03, 0xfc, 0x27, 0x2a, 0xf7, 0x13, 0x34,
	0xe2, 0x22, 0x2f, 0x7f, 0x17, 0xa5, 0x8b, 0xd4, 0xc8, 0x64, 0xdb, 0xe6, 0xe1, 0x65, 0x40, 0x1e,
	0xcc, 0x30, 0x11, 0x9b, 0xdc, 0xf0, 0x1b, 0x81, 0xd6, 0xf1, 0x1c, 0xf5, 0x0c, 0x45, 0x48, 0xdf,
	0x40, 0x2b, 0x53, 0xe1, 0x2b, 0xf9, 0x05, 0x43, 0x8f, 0x0c, 0x9a, 0xa3, 0xce, 0xc4, 0x67, 0x3b,
	0x12, 0x63, 0x6b, 0x88, 0xbd, 0x5f, 0x11, 0x2f, 0x53, 0xa3, 0x17, 0xbc, 0x16, 0xe8, 0x45, 0x70,
	0xed, 0xdc, 0x23, 0xba, 0x0f, 0xcd, 0x08, 0x17, 0x1e, 0x19, 0x90, 0x51, 0x9b, 0xdb, 0x9f, 0xf4,
	0x05, 0xec, 0xcd, 0x45, 0x5c, 0xa0, 0xd7, 0x18, 0x90, 0x51, 0x67, 0xc2, 0x1c, 0xb3, 0xfa, 0x08,
	0x96, 0x45, 0xd3, 0xd2, 0x7d, 0x9d, 0x30, 0xfb, 0x50, 0x88, 0xd4, 0x48, 0xb3, 0xe0, 0x15, 0xfc,
	0xa8, 0x71, 0x40, 0x86, 0xbf, 0x09, 0x74, 0x79, 0x75, 0xe9, 0x61, 0x2c, 0xf2, 0x9c, 0x1e, 0x41,
	0xcb, 0x66, 0x1b, 0x0a, 0x23, 0x4a, 0xc7, 0xce, 0xe4, 0xde, 0xdf, 0xd4, 0x73, 0x66, 0xa7, 0xd9,
	0x7c, 0xcc, 0x8e, 0x4f, 0x3f, 0x63, 0x60, 0xde, 0xa2, 0x11, 0xbc, 0x56, 0xa0, 0x1e, 0xfc, 0x3f,
	0x13, 0x69, 0x18, 0xa3, 0x2e, 0x57, 0x6d, 0xf3, 0xf5, 0x5f, 0xfa, 0x04, 0x5a, 0x6a, 0x95, 0x84,
	0xd7, 0x2c, 0x7d, 0x6e, 0x5d, 0x1a, 0x19, 0xaf, 0x11, 0x7a, 0x08, 0x60, 0x5f, 0x4c, 0x58, 0xc4,
	0x32, 0x9d, 0x7a, 0xff, 0x95, 0x02, 0xb7, 0x77, 0x0a, 0x9c, 0xd4, 0xa3, 0xdc, 0xc1, 0x86, 0x5f,
	0x09, 0xec, 0xbb, 0xc7, 0x1f, 0xc9, 0xdc, 0xd0, 0xd7, 0x5b, 0x01, 0xb0, 0xab, 0x05, 0x60, 0xe9,
	0x8d, 0xf3, 0x1f, 0xc3, 0x9e, 0x34, 0x98, 0xe4, 0x5e, 0xa3, 0x2c, 0xc5, 0x9d, 0x9d, 0x0b, 0xba,
	0x5b, 0xf0, 0x8a, 0x19, 0xfe, 0x22, 0x00, 0x67, 0x8b, 0xd3, 0x4f, 0xd0, 0xb5, 0xc8, 0x09, 0xc6,
	0x18, 0x18, 0xa5, 0x57, 0x3d, 0x7b, 0x70, 0x85, 0x9b, 0xd9, 0x3b, 0x87, 0xab, 0xda, 0x76, 0x4e,
	0x8a, 0x3e, 0x83, 0x8e, 0x51, 0xb1, 0xad, 0xb7, 0x54, 0xe9, 0x7a, 0xd9, 0xbe, 0xab, 0x6c, 0x3f,
	0x6d, 0x7b, 0xe3, 0xc7, 0x7a, 0x8c, 0xbb, 0x48, 0xef, 0x29, 0x5c, 0xdf, 0x32, 0xb9, 0xa0, 0xb7,
	0x37, 0xdc, 0xde, 0xb6, 0x9d, 0x1e, 0x3e, 0xef, 0x7e, 0x5f, 0xf6, 0xc9, 0x8f, 0x65, 0x9f, 0xfc,
	0x5c, 0xf6, 0xc9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xf5, 0x63, 0xa2, 0x98, 0x04, 0x00,
	0x00,
}

func (m *Overhead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Overhead) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Overhead) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PodFixed) > 0 {
		for k := range m.PodFixed {
			v := m.PodFixed[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintGenerated(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenerated(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RuntimeClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuntimeClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Scheduling != nil {
		{
			size, err := m.Scheduling.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Overhead != nil {
		{
			size, err := m.Overhead.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Handler != nil {
		i -= len(*m.Handler)
		copy(dAtA[i:], *m.Handler)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Handler)))
		i--
		dAtA[i] = 0x12
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RuntimeClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeClassList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuntimeClassList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Scheduling) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scheduling) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Scheduling) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tolerations) > 0 {
		for iNdEx := len(m.Tolerations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tolerations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.NodeSelector) > 0 {
		for k := range m.NodeSelector {
			v := m.NodeSelector[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenerated(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Overhead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PodFixed) > 0 {
		for k, v := range m.PodFixed {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovGenerated(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RuntimeClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Handler != nil {
		l = len(*m.Handler)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Overhead != nil {
		l = m.Overhead.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Scheduling != nil {
		l = m.Scheduling.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RuntimeClassList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Scheduling) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.NodeSelector) > 0 {
		for k, v := range m.NodeSelector {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if len(m.Tolerations) > 0 {
		for _, e := range m.Tolerations {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Overhead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Overhead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Overhead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodFixed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PodFixed == nil {
				m.PodFixed = make(map[string]*resource.Quantity)
			}
			var mapkey string
			var mapvalue *resource.Quantity
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenerated
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenerated
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &resource.Quantity{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.PodFixed[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RuntimeClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RuntimeClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Handler = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overhead", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Overhead == nil {
				m.Overhead = &Overhead{}
			}
			if err := m.Overhead.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scheduling", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Scheduling == nil {
				m.Scheduling = &Scheduling{}
			}
			if err := m.Scheduling.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RuntimeClassList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RuntimeClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &v1.ListMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &RuntimeClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Scheduling) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Scheduling: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scheduling: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeSelector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NodeSelector == nil {
				m.NodeSelector = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.NodeSelector[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tolerations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tolerations = append(m.Tolerations, &v11.Toleration{})
			if err := m.Tolerations[len(m.Tolerations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
