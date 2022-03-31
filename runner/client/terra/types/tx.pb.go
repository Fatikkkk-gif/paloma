// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: runner/client/terra/proto/tx.proto

package types

import (
	encoding_json "encoding/json"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// MsgExecuteContract represents a message to
// submits the given message data to a smart contract.
type MsgExecuteContract struct {
	// Sender is the that actor that signed the messages
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	// Contract is the address of the smart contract
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty" yaml:"contract"`
	// ExecuteMsg json encoded message to be passed to the contract
	ExecuteMsg encoding_json.RawMessage `protobuf:"bytes,3,opt,name=execute_msg,json=executeMsg,proto3,casttype=encoding/json.RawMessage" json:"execute_msg,omitempty" yaml:"execute_msg"`
	// Coins that are transferred to the contract on execution
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins" yaml:"coins"`
}

func (m *MsgExecuteContract) Reset()         { *m = MsgExecuteContract{} }
func (m *MsgExecuteContract) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteContract) ProtoMessage()    {}
func (*MsgExecuteContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18a0ea03d7a50ef, []int{0}
}
func (m *MsgExecuteContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteContract.Merge(m, src)
}
func (m *MsgExecuteContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteContract proto.InternalMessageInfo

// MsgExecuteContractResponse defines the Msg/ExecuteContract response type.
type MsgExecuteContractResponse struct {
	// Data contains base64-encoded bytes to returned from the contract
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty" yaml:"data"`
}

func (m *MsgExecuteContractResponse) Reset()         { *m = MsgExecuteContractResponse{} }
func (m *MsgExecuteContractResponse) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteContractResponse) ProtoMessage()    {}
func (*MsgExecuteContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18a0ea03d7a50ef, []int{1}
}
func (m *MsgExecuteContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteContractResponse.Merge(m, src)
}
func (m *MsgExecuteContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteContractResponse proto.InternalMessageInfo

func (m *MsgExecuteContractResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgExecuteContract)(nil), "terra.wasm.v1beta1.MsgExecuteContract")
	proto.RegisterType((*MsgExecuteContractResponse)(nil), "terra.wasm.v1beta1.MsgExecuteContractResponse")
}

func init() {
	proto.RegisterFile("runner/client/terra/proto/tx.proto", fileDescriptor_a18a0ea03d7a50ef)
}

var fileDescriptor_a18a0ea03d7a50ef = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0x86, 0xed, 0x84, 0x44, 0x61, 0x73, 0x28, 0x62, 0xa1, 0x30, 0x57, 0x78, 0x4f, 0xa6, 0x31,
	0x45, 0xbc, 0x4a, 0xe8, 0x52, 0x81, 0x23, 0xca, 0x6b, 0x4c, 0x47, 0x83, 0xf6, 0xf6, 0x46, 0x8e,
	0x21, 0xde, 0x39, 0x76, 0x36, 0xe4, 0xee, 0x0d, 0x28, 0x79, 0x84, 0xd4, 0x3c, 0x06, 0x55, 0xca,
	0x94, 0x54, 0x06, 0xdd, 0x35, 0xd4, 0x2e, 0xa9, 0x90, 0xbd, 0x3e, 0x14, 0x89, 0xca, 0xa3, 0xf9,
	0xbf, 0xf1, 0xcc, 0xfc, 0xb3, 0x2c, 0x45, 0xab, 0x2f, 0x80, 0x9c, 0x55, 0x0e, 0xad, 0xd4, 0x97,
	0x15, 0x18, 0x27, 0x1d, 0x58, 0xab, 0xe4, 0xc2, 0xa2, 0x43, 0xe9, 0x96, 0x59, 0x1f, 0x70, 0xde,
	0x67, 0xb3, 0x6b, 0x45, 0x75, 0xf6, 0xf9, 0x64, 0x06, 0x4e, 0x9d, 0x8c, 0x9f, 0x96, 0x58, 0xa2,
	0xe7, 0xba, 0xc8, 0x93, 0xe3, 0x58, 0x23, 0xd5, 0x48, 0x72, 0xa6, 0x08, 0xe4, 0x80, 0x4a, 0x8d,
	0x95, 0xf1, 0x7a, 0xf2, 0x7d, 0x87, 0xf1, 0x29, 0x95, 0x6f, 0x96, 0xa0, 0xaf, 0x1c, 0x9c, 0xa3,
	0x71, 0x56, 0x69, 0xc7, 0x5f, 0xb0, 0x7d, 0x02, 0x33, 0x07, 0x1b, 0x85, 0x93, 0x30, 0x7d, 0x98,
	0x3f, 0x6e, 0x1b, 0xf1, 0x68, 0xa5, 0xea, 0xcb, 0xb3, 0xc4, 0xe7, 0x93, 0x62, 0x00, 0xb8, 0x64,
	0x07, 0x7a, 0x28, 0x8b, 0x76, 0x7a, 0xf8, 0x49, 0xdb, 0x88, 0x23, 0x0f, 0x6f, 0x95, 0xa4, 0xf8,
	0x07, 0xf1, 0xb7, 0xec, 0x10, 0x7c, 0xbb, 0xf7, 0x35, 0x95, 0xd1, 0xee, 0x24, 0x4c, 0x47, 0xf9,
	0x69, 0xdb, 0x08, 0xee, 0x6b, 0xee, 0x89, 0xc9, 0x9f, 0x46, 0x44, 0x60, 0x34, 0xce, 0x2b, 0x53,
	0xca, 0x0f, 0x84, 0x26, 0x2b, 0xd4, 0xf5, 0x14, 0x88, 0x54, 0x09, 0x05, 0x1b, 0xc8, 0x29, 0x95,
	0xfc, 0x13, 0xdb, 0xeb, 0xb6, 0xa2, 0x68, 0x6f, 0xb2, 0x9b, 0x1e, 0x9e, 0x3e, 0xcb, 0xfc, 0xde,
	0x59, 0xb7, 0xf7, 0xd6, 0xa2, 0xec, 0x1c, 0x2b, 0x93, 0xbf, 0xba, 0x6d, 0x44, 0xd0, 0x36, 0x62,
	0xb4, 0x9d, 0xb0, 0x32, 0x94, 0x7c, 0xfb, 0x29, 0xd2, 0xb2, 0x72, 0x17, 0x57, 0xb3, 0x4c, 0x63,
	0x2d, 0x07, 0xd3, 0xfc, 0xe7, 0x98, 0xe6, 0x1f, 0xa5, 0x5b, 0x2d, 0x80, 0xfa, 0x1f, 0x50, 0xe1,
	0x3b, 0x9d, 0x1d, 0x7c, 0xb9, 0x11, 0xc1, 0xef, 0x1b, 0x11, 0x24, 0xaf, 0xd9, 0xf8, 0x7f, 0x0f,
	0x0b, 0xa0, 0x05, 0x1a, 0x02, 0xfe, 0x9c, 0x3d, 0x98, 0x2b, 0xa7, 0x7a, 0x27, 0x47, 0xf9, 0x51,
	0xdb, 0x88, 0x43, 0xdf, 0xba, 0xcb, 0x26, 0x45, 0x2f, 0xe6, 0xf9, 0xed, 0x3a, 0x0e, 0xef, 0xd6,
	0x71, 0xf8, 0x6b, 0x1d, 0x87, 0x5f, 0x37, 0x71, 0x70, 0xb7, 0x89, 0x83, 0x1f, 0x9b, 0x38, 0x78,
	0x77, 0x7f, 0xae, 0xfe, 0xec, 0xc7, 0x35, 0x1a, 0x58, 0x49, 0x8d, 0x16, 0xe4, 0x52, 0x76, 0x6f,
	0xc0, 0x4f, 0x37, 0xdb, 0xef, 0x4f, 0xfa, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x8d,
	0xb8, 0x9d, 0x48, 0x02, 0x00, 0x00,
}

func (m *MsgExecuteContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ExecuteMsg) > 0 {
		i -= len(m.ExecuteMsg)
		copy(dAtA[i:], m.ExecuteMsg)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ExecuteMsg)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecuteContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgExecuteContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ExecuteMsg)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgExecuteContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgExecuteContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgExecuteContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecuteMsg = append(m.ExecuteMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.ExecuteMsg == nil {
				m.ExecuteMsg = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgExecuteContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgExecuteContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
