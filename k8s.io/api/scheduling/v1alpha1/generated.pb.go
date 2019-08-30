// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/api/scheduling/v1alpha1/generated.proto

package k8s_io_api_scheduling_v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	_ "istio.io/gogo-genproto/k8s.io/api/core/v1"
	v1 "istio.io/gogo-genproto/k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "istio.io/gogo-genproto/k8s.io/apimachinery/pkg/runtime"
	_ "istio.io/gogo-genproto/k8s.io/apimachinery/pkg/runtime/schema"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// DEPRECATED - This group version of PriorityClass is deprecated by scheduling.k8s.io/v1/PriorityClass.
// PriorityClass defines mapping from a priority class name to the priority
// integer value. The value can be any valid integer.
type PriorityClass struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// The value of this priority class. This is the actual priority that pods
	// receive when they have the name of this class in their pod spec.
	Value int32 `protobuf:"varint,2,opt,name=value" json:"value"`
	// globalDefault specifies whether this PriorityClass should be considered as
	// the default priority for pods that do not have any priority class.
	// Only one PriorityClass can be marked as `globalDefault`. However, if more than
	// one PriorityClasses exists with their `globalDefault` field set to true,
	// the smallest value of such global default PriorityClasses will be used as the default priority.
	// +optional
	GlobalDefault bool `protobuf:"varint,3,opt,name=globalDefault" json:"globalDefault"`
	// description is an arbitrary string that usually provides guidelines on
	// when this priority class should be used.
	// +optional
	Description string `protobuf:"bytes,4,opt,name=description" json:"description"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority.
	// One of Never, PreemptLowerPriority.
	// Defaults to PreemptLowerPriority if unset.
	// This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
	// +optional
	PreemptionPolicy string `protobuf:"bytes,5,opt,name=preemptionPolicy" json:"preemptionPolicy"`
}

func (m *PriorityClass) Reset()      { *m = PriorityClass{} }
func (*PriorityClass) ProtoMessage() {}
func (*PriorityClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_260442fbb28d876a, []int{0}
}
func (m *PriorityClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriorityClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriorityClass.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriorityClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriorityClass.Merge(m, src)
}
func (m *PriorityClass) XXX_Size() int {
	return m.Size()
}
func (m *PriorityClass) XXX_DiscardUnknown() {
	xxx_messageInfo_PriorityClass.DiscardUnknown(m)
}

var xxx_messageInfo_PriorityClass proto.InternalMessageInfo

func (m *PriorityClass) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PriorityClass) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PriorityClass) GetGlobalDefault() bool {
	if m != nil {
		return m.GlobalDefault
	}
	return false
}

func (m *PriorityClass) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PriorityClass) GetPreemptionPolicy() string {
	if m != nil {
		return m.PreemptionPolicy
	}
	return ""
}

// PriorityClassList is a collection of priority classes.
type PriorityClassList struct {
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// items is the list of PriorityClasses
	Items []*PriorityClass `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (m *PriorityClassList) Reset()      { *m = PriorityClassList{} }
func (*PriorityClassList) ProtoMessage() {}
func (*PriorityClassList) Descriptor() ([]byte, []int) {
	return fileDescriptor_260442fbb28d876a, []int{1}
}
func (m *PriorityClassList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriorityClassList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriorityClassList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriorityClassList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriorityClassList.Merge(m, src)
}
func (m *PriorityClassList) XXX_Size() int {
	return m.Size()
}
func (m *PriorityClassList) XXX_DiscardUnknown() {
	xxx_messageInfo_PriorityClassList.DiscardUnknown(m)
}

var xxx_messageInfo_PriorityClassList proto.InternalMessageInfo

func (m *PriorityClassList) GetMetadata() *v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PriorityClassList) GetItems() []*PriorityClass {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*PriorityClass)(nil), "k8s.io.api.scheduling.v1alpha1.PriorityClass")
	proto.RegisterType((*PriorityClassList)(nil), "k8s.io.api.scheduling.v1alpha1.PriorityClassList")
}

func init() {
	proto.RegisterFile("k8s.io/api/scheduling/v1alpha1/generated.proto", fileDescriptor_260442fbb28d876a)
}

var fileDescriptor_260442fbb28d876a = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0xe3, 0x99, 0x89, 0x34, 0xb8, 0x1a, 0x09, 0xb2, 0x8a, 0xba, 0x30, 0x51, 0x17, 0x28,
	0x42, 0xc2, 0x9e, 0x8e, 0x46, 0x88, 0xf5, 0x0c, 0x2b, 0x34, 0x88, 0x51, 0xdf, 0xc0, 0x4d, 0x0e,
	0xa9, 0xa9, 0x13, 0x5b, 0xb6, 0x13, 0xa9, 0x3b, 0x1e, 0x81, 0x87, 0x60, 0xc1, 0xa3, 0x74, 0xd9,
	0x65, 0x57, 0x88, 0xa6, 0x1b, 0x96, 0x7d, 0x03, 0x50, 0x68, 0xe9, 0x2d, 0x5c, 0x66, 0xe9, 0x73,
	0xbe, 0xef, 0xd7, 0xd1, 0x2f, 0x63, 0x3a, 0x7e, 0x65, 0xa9, 0x50, 0x8c, 0x6b, 0xc1, 0x6c, 0x32,
	0x82, 0xb4, 0x94, 0xa2, 0xc8, 0x58, 0xd5, 0xe7, 0x52, 0x8f, 0x78, 0x9f, 0x65, 0x50, 0x80, 0xe1,
	0x0e, 0x52, 0xaa, 0x8d, 0x72, 0x2a, 0x20, 0x6b, 0x9e, 0x72, 0x2d, 0xe8, 0x8e, 0xa7, 0xbf, 0xf9,
	0x6e, 0x6f, 0x2f, 0x2f, 0x51, 0x06, 0x58, 0xd5, 0xca, 0xe8, 0x5e, 0xef, 0x98, 0x9c, 0x27, 0x23,
	0x51, 0x80, 0x99, 0x30, 0x3d, 0xce, 0x9a, 0x81, 0x65, 0x39, 0x38, 0xfe, 0x27, 0x8b, 0xfd, 0xcd,
	0x32, 0x65, 0xe1, 0x44, 0x0e, 0x2d, 0xe1, 0xe5, 0xff, 0x84, 0xe6, 0xfe, 0x9c, 0x1f, 0x7b, 0xbd,
	0x1f, 0x08, 0x5f, 0xdc, 0x1b, 0xa1, 0x8c, 0x70, 0x93, 0x5b, 0xc9, 0xad, 0x0d, 0xee, 0xf0, 0x79,
	0x73, 0x55, 0xca, 0x1d, 0x0f, 0x51, 0x84, 0xe2, 0xce, 0xd5, 0x25, 0xdd, 0xf5, 0xb0, 0x0d, 0xa7,
	0x7a, 0x9c, 0x35, 0x03, 0x4b, 0x1b, 0x9a, 0x56, 0x7d, 0xfa, 0x6e, 0xf8, 0x01, 0x12, 0xf7, 0x16,
	0x1c, 0x1f, 0x6c, 0x13, 0x82, 0x2e, 0xf6, 0x2b, 0x2e, 0x4b, 0x08, 0x4f, 0x22, 0x14, 0xfb, 0x37,
	0x67, 0xd3, 0xaf, 0x4f, 0xbd, 0xc1, 0x7a, 0x14, 0x3c, 0xc7, 0x17, 0x99, 0x54, 0x43, 0x2e, 0x5f,
	0xc3, 0x7b, 0x5e, 0x4a, 0x17, 0x9e, 0x46, 0x28, 0x3e, 0xdf, 0x30, 0x87, 0xab, 0xe0, 0x19, 0xee,
	0xa4, 0x60, 0x13, 0x23, 0xb4, 0x13, 0xaa, 0x08, 0xcf, 0x22, 0x14, 0x3f, 0xda, 0x90, 0xfb, 0x8b,
	0xe0, 0x12, 0x3f, 0xd6, 0x06, 0x20, 0xff, 0xf5, 0xba, 0x57, 0x52, 0x24, 0x93, 0xd0, 0xdf, 0x83,
	0x5b, 0xdb, 0xde, 0x67, 0x84, 0x9f, 0x1c, 0x34, 0x70, 0x27, 0xac, 0x0b, 0xde, 0xb4, 0x5a, 0xa0,
	0x0f, 0x6b, 0xa1, 0xb1, 0x8f, 0x3a, 0xb8, 0xc5, 0xbe, 0x70, 0x90, 0xdb, 0xf0, 0x24, 0x3a, 0x8d,
	0x3b, 0x57, 0x2f, 0xe8, 0xbf, 0xbf, 0x15, 0x3d, 0xb8, 0x66, 0xb0, 0x76, 0x6f, 0xae, 0x67, 0x0b,
	0xe2, 0xcd, 0x17, 0xc4, 0x5b, 0x2d, 0x08, 0xfa, 0x58, 0x13, 0xf4, 0xa5, 0x26, 0x68, 0x5a, 0x13,
	0x34, 0xab, 0x09, 0xfa, 0x56, 0x13, 0xf4, 0xbd, 0x26, 0xde, 0xaa, 0x26, 0xe8, 0xd3, 0x92, 0x78,
	0xb3, 0x25, 0xf1, 0xe6, 0x4b, 0xe2, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x62, 0x64, 0x50, 0xa5,
	0xf2, 0x02, 0x00, 0x00,
}

func (this *PriorityClass) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PriorityClass)
	if !ok {
		that2, ok := that.(PriorityClass)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.GlobalDefault != that1.GlobalDefault {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.PreemptionPolicy != that1.PreemptionPolicy {
		return false
	}
	return true
}
func (this *PriorityClassList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PriorityClassList)
	if !ok {
		that2, ok := that.(PriorityClassList)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if len(this.Items) != len(that1.Items) {
		return false
	}
	for i := range this.Items {
		if !this.Items[i].Equal(that1.Items[i]) {
			return false
		}
	}
	return true
}
func (this *PriorityClass) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&k8s_io_api_scheduling_v1alpha1.PriorityClass{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "GlobalDefault: "+fmt.Sprintf("%#v", this.GlobalDefault)+",\n")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	s = append(s, "PreemptionPolicy: "+fmt.Sprintf("%#v", this.PreemptionPolicy)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PriorityClassList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&k8s_io_api_scheduling_v1alpha1.PriorityClassList{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.Items != nil {
		s = append(s, "Items: "+fmt.Sprintf("%#v", this.Items)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGenerated(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PriorityClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriorityClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.PreemptionPolicy)
	copy(dAtA[i:], m.PreemptionPolicy)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.PreemptionPolicy)))
	i--
	dAtA[i] = 0x2a
	i -= len(m.Description)
	copy(dAtA[i:], m.Description)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Description)))
	i--
	dAtA[i] = 0x22
	i--
	if m.GlobalDefault {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x18
	i = encodeVarintGenerated(dAtA, i, uint64(m.Value))
	i--
	dAtA[i] = 0x10
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

func (m *PriorityClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriorityClassList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriorityClassList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *PriorityClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	n += 1 + sovGenerated(uint64(m.Value))
	n += 2
	l = len(m.Description)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.PreemptionPolicy)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PriorityClassList) Size() (n int) {
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
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PriorityClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PriorityClass{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ObjectMeta", "v1.ObjectMeta", 1) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`GlobalDefault:` + fmt.Sprintf("%v", this.GlobalDefault) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`PreemptionPolicy:` + fmt.Sprintf("%v", this.PreemptionPolicy) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PriorityClassList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]*PriorityClass{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(f.String(), "PriorityClass", "PriorityClass", 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&PriorityClassList{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ListMeta", "v1.ListMeta", 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PriorityClass) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PriorityClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClass: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalDefault", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.GlobalDefault = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreemptionPolicy", wireType)
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
			m.PreemptionPolicy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PriorityClassList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PriorityClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriorityClassList: illegal tag %d (wire type %d)", fieldNum, wire)
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
			m.Items = append(m.Items, &PriorityClass{})
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
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
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
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
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
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGenerated
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)
