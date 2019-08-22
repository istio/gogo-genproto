// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto

package k8s_io_apimachinery_pkg_apis_meta_v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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

// PartialObjectMetadataList contains a list of objects containing only their metadata.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PartialObjectMetadataList struct {
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	// items contains each of the included items.
	Items []*v1.PartialObjectMetadata `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *PartialObjectMetadataList) Reset()      { *m = PartialObjectMetadataList{} }
func (*PartialObjectMetadataList) ProtoMessage() {}
func (*PartialObjectMetadataList) Descriptor() ([]byte, []int) {
	return fileDescriptor_39237a8d8061b52f, []int{0}
}
func (m *PartialObjectMetadataList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PartialObjectMetadataList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PartialObjectMetadataList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PartialObjectMetadataList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialObjectMetadataList.Merge(m, src)
}
func (m *PartialObjectMetadataList) XXX_Size() int {
	return m.Size()
}
func (m *PartialObjectMetadataList) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialObjectMetadataList.DiscardUnknown(m)
}

var xxx_messageInfo_PartialObjectMetadataList proto.InternalMessageInfo

func (m *PartialObjectMetadataList) GetMetadata() *v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PartialObjectMetadataList) GetItems() []*v1.PartialObjectMetadata {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*PartialObjectMetadataList)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadataList")
}

func init() {
	proto.RegisterFile("k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto", fileDescriptor_39237a8d8061b52f)
}

var fileDescriptor_39237a8d8061b52f = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcf, 0x31, 0x4a, 0x03, 0x41,
	0x14, 0x06, 0xe0, 0x19, 0x45, 0x90, 0x4d, 0xb7, 0x55, 0x4c, 0xf1, 0x08, 0x56, 0xb1, 0x99, 0x21,
	0x21, 0x88, 0x62, 0x67, 0x29, 0x8a, 0x9a, 0x1b, 0xbc, 0x6c, 0x1e, 0x9b, 0x71, 0x9d, 0xdd, 0x65,
	0xe7, 0x29, 0xd8, 0x79, 0x04, 0x8f, 0xe1, 0x19, 0x3c, 0x81, 0xe5, 0x96, 0x29, 0xdd, 0xd9, 0xc6,
	0x32, 0x47, 0x90, 0x25, 0x41, 0x24, 0x2a, 0xd9, 0x72, 0x7e, 0xf8, 0xfe, 0x7f, 0x5e, 0x70, 0x9a,
	0x9c, 0x38, 0x65, 0x32, 0x8d, 0xb9, 0xb1, 0x18, 0xcd, 0x4d, 0x4a, 0xc5, 0x93, 0xce, 0x93, 0xb8,
	0x09, 0x9c, 0xb6, 0xc4, 0xa8, 0x1f, 0x87, 0x53, 0x62, 0x1c, 0xea, 0x98, 0x52, 0x2a, 0x90, 0x69,
	0xa6, 0xf2, 0x22, 0xe3, 0x2c, 0x3c, 0x5a, 0x51, 0xf5, 0x93, 0xaa, 0x3c, 0x89, 0x9b, 0xc0, 0xa9,
	0x86, 0xaa, 0x35, 0xed, 0x8d, 0xdb, 0xac, 0x6c, 0x0e, 0xf4, 0xf4, 0x7f, 0xaa, 0x78, 0x48, 0xd9,
	0x58, 0xfa, 0x05, 0x8e, 0xb7, 0x01, 0x17, 0xcd, 0xc9, 0xe2, 0xa6, 0x3b, 0x7c, 0x93, 0xc1, 0xc1,
	0x0d, 0x16, 0x6c, 0xf0, 0xfe, 0x7a, 0x7a, 0x47, 0x11, 0x5f, 0x11, 0xe3, 0x0c, 0x19, 0x2f, 0x8d,
	0xe3, 0xf0, 0x22, 0xd8, 0xb7, 0xeb, 0x77, 0x77, 0xa7, 0x2f, 0x07, 0x9d, 0x91, 0x52, 0x6d, 0x4e,
	0x57, 0x8d, 0x6e, 0x9a, 0x26, 0xdf, 0x3e, 0xbc, 0x0d, 0xf6, 0x0c, 0x93, 0x75, 0x5d, 0xd9, 0xdf,
	0x1d, 0x74, 0x46, 0x67, 0xed, 0x8a, 0xfe, 0xfc, 0xdb, 0x64, 0xd5, 0x74, 0x3e, 0x2e, 0x2b, 0x10,
	0x8b, 0x0a, 0xc4, 0xb2, 0x02, 0xf9, 0xec, 0x41, 0xbe, 0x7a, 0x90, 0xef, 0x1e, 0x64, 0xe9, 0x41,
	0x7e, 0x78, 0x90, 0x9f, 0x1e, 0xc4, 0xd2, 0x83, 0x7c, 0xa9, 0x41, 0x94, 0x35, 0x88, 0x45, 0x0d,
	0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x51, 0xf9, 0x7a, 0xc5, 0xf8, 0x01, 0x00, 0x00,
}

func (this *PartialObjectMetadataList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PartialObjectMetadataList)
	if !ok {
		that2, ok := that.(PartialObjectMetadataList)
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
func (this *PartialObjectMetadataList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&k8s_io_apimachinery_pkg_apis_meta_v1beta1.PartialObjectMetadataList{")
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
func (m *PartialObjectMetadataList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartialObjectMetadataList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PartialObjectMetadataList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
		dAtA[i] = 0x12
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
func (m *PartialObjectMetadataList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PartialObjectMetadataList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]*PartialObjectMetadata{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(fmt.Sprintf("%v", f), "PartialObjectMetadata", "v1.PartialObjectMetadata", 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&PartialObjectMetadataList{`,
		`Items:` + repeatedStringForItems + `,`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ListMeta", "v1.ListMeta", 1) + `,`,
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
func (m *PartialObjectMetadataList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PartialObjectMetadataList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartialObjectMetadataList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
			m.Items = append(m.Items, &v1.PartialObjectMetadata{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
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
