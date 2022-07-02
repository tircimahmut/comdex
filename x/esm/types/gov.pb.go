// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/esm/v1beta1/gov.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ESMTriggerParamsProposal struct {
	Title            string           `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description      string           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	EsmTriggerParams ESMTriggerParams `protobuf:"bytes,3,opt,name=esmTriggerParams,proto3" json:"esmTriggerParams"`
}

func (m *ESMTriggerParamsProposal) Reset()         { *m = ESMTriggerParamsProposal{} }
func (m *ESMTriggerParamsProposal) String() string { return proto.CompactTextString(m) }
func (*ESMTriggerParamsProposal) ProtoMessage()    {}
func (*ESMTriggerParamsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6824f910863b2597, []int{0}
}
func (m *ESMTriggerParamsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ESMTriggerParamsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ESMTriggerParamsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ESMTriggerParamsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ESMTriggerParamsProposal.Merge(m, src)
}
func (m *ESMTriggerParamsProposal) XXX_Size() int {
	return m.Size()
}
func (m *ESMTriggerParamsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ESMTriggerParamsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ESMTriggerParamsProposal proto.InternalMessageInfo

func (m *ESMTriggerParamsProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ESMTriggerParamsProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ESMTriggerParamsProposal) GetEsmTriggerParams() ESMTriggerParams {
	if m != nil {
		return m.EsmTriggerParams
	}
	return ESMTriggerParams{}
}

func init() {
	proto.RegisterType((*ESMTriggerParamsProposal)(nil), "comdex.esm.v1beta1.ESMTriggerParamsProposal")
}

func init() { proto.RegisterFile("comdex/esm/v1beta1/gov.proto", fileDescriptor_6824f910863b2597) }

var fileDescriptor_6824f910863b2597 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0x9b, 0xf7, 0x55, 0xc1, 0xce, 0xc3, 0x08, 0x22, 0x65, 0x48, 0x36, 0x8a, 0xc8, 0x2e,
	0x6b, 0x98, 0x5e, 0xc4, 0xe3, 0x40, 0x3c, 0x09, 0x63, 0x8a, 0x07, 0x6f, 0x69, 0x97, 0xc5, 0x40,
	0xb3, 0x7f, 0x49, 0xe2, 0x70, 0xdf, 0xc2, 0x8f, 0xb5, 0xe3, 0x0e, 0x1e, 0x3c, 0x15, 0x69, 0xbf,
	0xc1, 0x3e, 0x81, 0xb4, 0x29, 0x32, 0x9d, 0xb7, 0x24, 0xbf, 0x87, 0xdf, 0x13, 0x1e, 0xff, 0x34,
	0x01, 0x35, 0xe5, 0xaf, 0x94, 0x1b, 0x45, 0x17, 0xc3, 0x98, 0x5b, 0x36, 0xa4, 0x02, 0x16, 0x51,
	0xa6, 0xc1, 0x02, 0xc6, 0x8e, 0x46, 0xdc, 0xa8, 0xa8, 0xa1, 0x9d, 0x63, 0x01, 0x02, 0x6a, 0x4c,
	0xab, 0x93, 0x4b, 0x76, 0x48, 0x02, 0x46, 0x81, 0xa1, 0x31, 0x33, 0xfc, 0x5b, 0x94, 0x80, 0x9c,
	0x37, 0xfc, 0xaf, 0x9e, 0xca, 0x5a, 0xd3, 0xf0, 0x1d, 0xf9, 0xc1, 0xcd, 0xfd, 0xdd, 0x83, 0x96,
	0x42, 0x70, 0x3d, 0x66, 0x9a, 0x29, 0x33, 0xd6, 0x90, 0x81, 0x61, 0x29, 0x3e, 0xf7, 0xf7, 0xad,
	0xb4, 0x29, 0x0f, 0x50, 0x0f, 0xf5, 0x0f, 0x47, 0xed, 0x4d, 0xde, 0x3d, 0x5a, 0x32, 0x95, 0x5e,
	0x87, 0xf5, 0x73, 0x38, 0x71, 0x18, 0x5f, 0xf9, 0xad, 0x29, 0x37, 0x89, 0x96, 0x99, 0x95, 0x30,
	0x0f, 0xfe, 0xd5, 0xe9, 0x93, 0x4d, 0xde, 0xc5, 0x2e, 0xbd, 0x05, 0xc3, 0xc9, 0x76, 0x14, 0x3f,
	0xfa, 0x6d, 0x6e, 0xd4, 0x8f, 0xf6, 0xe0, 0x7f, 0x0f, 0xf5, 0x5b, 0x17, 0x67, 0xd1, 0xee, 0x02,
	0xd1, 0xef, 0x9f, 0x8e, 0xf6, 0x56, 0x79, 0xd7, 0x9b, 0xec, 0x38, 0x46, 0xb7, 0xab, 0x82, 0xa0,
	0x75, 0x41, 0xd0, 0x67, 0x41, 0xd0, 0x5b, 0x49, 0xbc, 0x75, 0x49, 0xbc, 0x8f, 0x92, 0x78, 0x4f,
	0x03, 0x21, 0xed, 0xf3, 0x4b, 0x5c, 0xd9, 0xa9, 0x6b, 0x18, 0xc0, 0x6c, 0x26, 0x13, 0xc9, 0xd2,
	0xe6, 0x4e, 0xdd, 0x56, 0x76, 0x99, 0x71, 0x13, 0x1f, 0xd4, 0x33, 0x5d, 0x7e, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xcf, 0xd7, 0xb4, 0xc8, 0xae, 0x01, 0x00, 0x00,
}

func (m *ESMTriggerParamsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ESMTriggerParamsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ESMTriggerParamsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.EsmTriggerParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ESMTriggerParamsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.EsmTriggerParams.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ESMTriggerParamsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: ESMTriggerParamsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ESMTriggerParamsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EsmTriggerParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EsmTriggerParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
