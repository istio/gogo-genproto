// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/type/date.proto

package google_type

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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

// Represents a whole or partial calendar date, e.g. a birthday. The time of day
// and time zone are either specified elsewhere or are not significant. The date
// is relative to the Proleptic Gregorian Calendar. This can represent:
//
// * A full date, with non-zero year, month and day values
// * A month and day value, with a zero year, e.g. an anniversary
// * A year on its own, with zero month and day values
// * A year and month value, with a zero day, e.g. a credit card expiration date
//
// Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and `google.protobuf.Timestamp`.
type Date struct {
	// Year of date. Must be from 1 to 9999, or 0 if specifying a date without
	// a year.
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Month of year. Must be from 1 to 12, or 0 if specifying a year without a
	// month and day.
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Day of month. Must be from 1 to 31 and valid for the year and month, or 0
	// if specifying a year by itself or a year and month where the day is not
	// significant.
	Day                  int32    `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_92c30699df886e3f, []int{0}
}
func (m *Date) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Date.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return m.Size()
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func init() {
	proto.RegisterType((*Date)(nil), "google.type.Date")
}

func init() { proto.RegisterFile("google/type/date.proto", fileDescriptor_92c30699df886e3f) }

var fileDescriptor_92c30699df886e3f = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x88, 0xeb, 0x81, 0xc4, 0x95, 0x9c, 0xb8, 0x58, 0x5c, 0x12, 0x4b,
	0x52, 0x85, 0x84, 0xb8, 0x58, 0x2a, 0x53, 0x13, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xdc, 0xfc, 0xbc, 0x92, 0x0c, 0x09, 0x26, 0xb0, 0x20, 0x84,
	0x23, 0x24, 0xc0, 0xc5, 0x9c, 0x92, 0x58, 0x29, 0xc1, 0x0c, 0x16, 0x03, 0x31, 0x9d, 0x8c, 0x4f,
	0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0x46, 0x2e, 0xfe, 0xe4, 0xfc,
	0x5c, 0x3d, 0x24, 0x2b, 0x9c, 0x38, 0x41, 0x16, 0x04, 0x80, 0xac, 0x0e, 0x60, 0xfc, 0xc1, 0xc8,
	0xb8, 0x88, 0x89, 0xd9, 0x3d, 0x24, 0x20, 0x89, 0x0d, 0xec, 0x18, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x26, 0x66, 0x69, 0x6f, 0xa6, 0x00, 0x00, 0x00,
}

func (m *Date) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Date) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Date) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Day != 0 {
		i = encodeVarintDate(dAtA, i, uint64(m.Day))
		i--
		dAtA[i] = 0x18
	}
	if m.Month != 0 {
		i = encodeVarintDate(dAtA, i, uint64(m.Month))
		i--
		dAtA[i] = 0x10
	}
	if m.Year != 0 {
		i = encodeVarintDate(dAtA, i, uint64(m.Year))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDate(dAtA []byte, offset int, v uint64) int {
	offset -= sovDate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Date) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Year != 0 {
		n += 1 + sovDate(uint64(m.Year))
	}
	if m.Month != 0 {
		n += 1 + sovDate(uint64(m.Month))
	}
	if m.Day != 0 {
		n += 1 + sovDate(uint64(m.Day))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDate(x uint64) (n int) {
	return sovDate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Date) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDate
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
			return fmt.Errorf("proto: Date: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Date: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Year", wireType)
			}
			m.Year = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Year |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Month", wireType)
			}
			m.Month = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Month |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Day", wireType)
			}
			m.Day = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Day |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDate
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
func skipDate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDate
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
					return 0, ErrIntOverflowDate
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
					return 0, ErrIntOverflowDate
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
				return 0, ErrInvalidLengthDate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDate = fmt.Errorf("proto: unexpected end of group")
)
