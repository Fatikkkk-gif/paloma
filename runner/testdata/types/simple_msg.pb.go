// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: runner/testdata/proto/simple_msg.proto

package types

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

type SimpleMessage struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Hello  string `protobuf:"bytes,2,opt,name=hello,proto3" json:"hello,omitempty"`
	World  string `protobuf:"bytes,3,opt,name=world,proto3" json:"world,omitempty"`
}

func (m *SimpleMessage) Reset()         { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()    {}
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f53203c7ce7a8710, []int{0}
}
func (m *SimpleMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimpleMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimpleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMessage.Merge(m, src)
}
func (m *SimpleMessage) XXX_Size() int {
	return m.Size()
}
func (m *SimpleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMessage proto.InternalMessageInfo

func (m *SimpleMessage) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SimpleMessage) GetHello() string {
	if m != nil {
		return m.Hello
	}
	return ""
}

func (m *SimpleMessage) GetWorld() string {
	if m != nil {
		return m.World
	}
	return ""
}

func init() {
	proto.RegisterType((*SimpleMessage)(nil), "runner.SimpleMessage")
}

func init() {
	proto.RegisterFile("runner/testdata/proto/simple_msg.proto", fileDescriptor_f53203c7ce7a8710)
}

var fileDescriptor_f53203c7ce7a8710 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x2a, 0xcd, 0xcb,
	0x4b, 0x2d, 0xd2, 0x2f, 0x49, 0x2d, 0x2e, 0x49, 0x49, 0x2c, 0x49, 0xd4, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0xcc, 0x2d, 0xc8, 0x49, 0x8d, 0xcf, 0x2d, 0x4e, 0xd7, 0x03, 0x0b, 0x08,
	0xb1, 0x41, 0xd4, 0x29, 0x05, 0x73, 0xf1, 0x06, 0x83, 0xe5, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3,
	0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x53, 0xf3, 0x52, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xa0, 0x3c, 0x21, 0x11, 0x2e, 0xd6, 0x8c, 0xd4, 0x9c, 0x9c, 0x7c, 0x09, 0x26, 0xb0,
	0x30, 0x84, 0x03, 0x12, 0x2d, 0xcf, 0x2f, 0xca, 0x49, 0x91, 0x60, 0x86, 0x88, 0x82, 0x39, 0x4e,
	0xbe, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9c, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x96, 0x9f, 0x53, 0x9a, 0x9b, 0x9a, 0x96, 0xa9,
	0x9f, 0x5c, 0x94, 0x9f, 0x97, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x8f, 0xee, 0xf8, 0x92, 0xca, 0x82,
	0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x93, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xb0, 0xd2,
	0x50, 0xdc, 0x00, 0x00, 0x00,
}

func (m *SimpleMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimpleMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.World) > 0 {
		i -= len(m.World)
		copy(dAtA[i:], m.World)
		i = encodeVarintSimpleMsg(dAtA, i, uint64(len(m.World)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Hello) > 0 {
		i -= len(m.Hello)
		copy(dAtA[i:], m.Hello)
		i = encodeVarintSimpleMsg(dAtA, i, uint64(len(m.Hello)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintSimpleMsg(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSimpleMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovSimpleMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SimpleMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovSimpleMsg(uint64(l))
	}
	l = len(m.Hello)
	if l > 0 {
		n += 1 + l + sovSimpleMsg(uint64(l))
	}
	l = len(m.World)
	if l > 0 {
		n += 1 + l + sovSimpleMsg(uint64(l))
	}
	return n
}

func sovSimpleMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSimpleMsg(x uint64) (n int) {
	return sovSimpleMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SimpleMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSimpleMsg
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
			return fmt.Errorf("proto: SimpleMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimpleMsg
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
				return ErrInvalidLengthSimpleMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSimpleMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hello", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimpleMsg
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
				return ErrInvalidLengthSimpleMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSimpleMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hello = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field World", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimpleMsg
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
				return ErrInvalidLengthSimpleMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSimpleMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.World = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSimpleMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSimpleMsg
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
func skipSimpleMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSimpleMsg
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
					return 0, ErrIntOverflowSimpleMsg
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
					return 0, ErrIntOverflowSimpleMsg
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
				return 0, ErrInvalidLengthSimpleMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSimpleMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSimpleMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSimpleMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSimpleMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSimpleMsg = fmt.Errorf("proto: unexpected end of group")
)
