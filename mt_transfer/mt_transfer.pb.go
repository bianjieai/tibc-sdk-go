// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tibc/apps/mt_transfer/v1/mt_transfer.proto

package mt_transfer

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

type MultiTokenPacketData struct {
	// the class to which the Mt to be transferred belongs
	Class string `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	// the mt id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// the address defined by MT outside the chain
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// the mt sender
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// the mt receiver
	Receiver string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// identify whether it is far away from the source chain
	AwayFromOrigin bool `protobuf:"varint,6,opt,name=away_from_origin,json=awayFromOrigin,proto3" json:"away_from_origin,omitempty"`
	// the destination contract address to receive the nft
	DestContract string `protobuf:"bytes,7,opt,name=dest_contract,json=destContract,proto3" json:"dest_contract,omitempty"`
	// the amount defined by MT outside the chain
	Amount uint64 `protobuf:"varint,8,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MultiTokenPacketData) Reset()         { *m = MultiTokenPacketData{} }
func (m *MultiTokenPacketData) String() string { return proto.CompactTextString(m) }
func (*MultiTokenPacketData) ProtoMessage()    {}
func (*MultiTokenPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e1dfa992d03735, []int{0}
}
func (m *MultiTokenPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiTokenPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiTokenPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiTokenPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiTokenPacketData.Merge(m, src)
}
func (m *MultiTokenPacketData) XXX_Size() int {
	return m.Size()
}
func (m *MultiTokenPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiTokenPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_MultiTokenPacketData proto.InternalMessageInfo

func (m *MultiTokenPacketData) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *MultiTokenPacketData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MultiTokenPacketData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MultiTokenPacketData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MultiTokenPacketData) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MultiTokenPacketData) GetAwayFromOrigin() bool {
	if m != nil {
		return m.AwayFromOrigin
	}
	return false
}

func (m *MultiTokenPacketData) GetDestContract() string {
	if m != nil {
		return m.DestContract
	}
	return ""
}

func (m *MultiTokenPacketData) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// ClassTrace contains the base class for Multi Token and the
// source tracing information path.
type ClassTrace struct {
	// path defines the chain of sourceChain/destChain
	// identifiers used for tracing the source of the Non fungible token.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// base class of the relayed non fungible token.
	BaseClass string `protobuf:"bytes,2,opt,name=base_class,json=baseClass,proto3" json:"base_class,omitempty"`
}

func (m *ClassTrace) Reset()         { *m = ClassTrace{} }
func (m *ClassTrace) String() string { return proto.CompactTextString(m) }
func (*ClassTrace) ProtoMessage()    {}
func (*ClassTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e1dfa992d03735, []int{1}
}
func (m *ClassTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassTrace.Merge(m, src)
}
func (m *ClassTrace) XXX_Size() int {
	return m.Size()
}
func (m *ClassTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassTrace.DiscardUnknown(m)
}

var xxx_messageInfo_ClassTrace proto.InternalMessageInfo

func (m *ClassTrace) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ClassTrace) GetBaseClass() string {
	if m != nil {
		return m.BaseClass
	}
	return ""
}

func init() {
	proto.RegisterType((*MultiTokenPacketData)(nil), "tibc.apps.mt_transfer.v1.MultiTokenPacketData")
	proto.RegisterType((*ClassTrace)(nil), "tibc.apps.mt_transfer.v1.ClassTrace")
}

func init() {
	proto.RegisterFile("tibc/apps/mt_transfer/v1/mt_transfer.proto", fileDescriptor_c6e1dfa992d03735)
}

var fileDescriptor_c6e1dfa992d03735 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6a, 0x2a, 0x31,
	0x14, 0x86, 0x8d, 0x57, 0xbd, 0x1a, 0xbc, 0x72, 0x09, 0x72, 0x09, 0x17, 0x3a, 0x0c, 0x76, 0x33,
	0x14, 0x3a, 0x83, 0xf4, 0x01, 0x5a, 0x6a, 0xe9, 0xae, 0xb4, 0x0c, 0xae, 0xdc, 0x0c, 0x67, 0x32,
	0x51, 0x53, 0x9d, 0xc9, 0x90, 0x64, 0x2c, 0xbe, 0x45, 0x1f, 0xab, 0x4b, 0x97, 0x5d, 0x16, 0x5d,
	0xf6, 0x25, 0x4a, 0xe2, 0x50, 0x5c, 0x74, 0x77, 0xfe, 0xef, 0xfc, 0x09, 0xe7, 0x3f, 0x07, 0x5f,
	0x18, 0x91, 0xb2, 0x08, 0xca, 0x52, 0x47, 0xb9, 0x49, 0x8c, 0x82, 0x42, 0xcf, 0xb9, 0x8a, 0x36,
	0xe3, 0x53, 0x19, 0x96, 0x4a, 0x1a, 0x49, 0xa8, 0xf5, 0x86, 0xd6, 0x1b, 0x9e, 0x36, 0x37, 0xe3,
	0xd1, 0x27, 0xc2, 0xc3, 0x87, 0x6a, 0x6d, 0xc4, 0x54, 0xae, 0x78, 0xf1, 0x04, 0x6c, 0xc5, 0xcd,
	0x1d, 0x18, 0x20, 0x43, 0xdc, 0x66, 0x6b, 0xd0, 0x9a, 0x22, 0x1f, 0x05, 0xbd, 0xf8, 0x28, 0xc8,
	0x00, 0x37, 0x45, 0x46, 0x9b, 0x0e, 0x35, 0x45, 0x46, 0x08, 0x6e, 0x65, 0x60, 0x80, 0xfe, 0xf2,
	0x51, 0xd0, 0x8f, 0x5d, 0x4d, 0xfe, 0xe1, 0x8e, 0xe6, 0x45, 0xc6, 0x15, 0x6d, 0x39, 0x5f, 0xad,
	0xc8, 0x7f, 0xdc, 0x55, 0x9c, 0x71, 0xb1, 0xe1, 0x8a, 0xb6, 0x5d, 0xe7, 0x5b, 0x93, 0x00, 0xff,
	0x85, 0x17, 0xd8, 0x26, 0x73, 0x25, 0xf3, 0x44, 0x2a, 0xb1, 0x10, 0x05, 0xed, 0xf8, 0x28, 0xe8,
	0xc6, 0x03, 0xcb, 0xef, 0x95, 0xcc, 0x1f, 0x1d, 0x25, 0xe7, 0xf8, 0x4f, 0xc6, 0xb5, 0x49, 0x98,
	0x2c, 0x8c, 0x02, 0x66, 0xe8, 0x6f, 0xf7, 0x55, 0xdf, 0xc2, 0x49, 0xcd, 0xec, 0x08, 0x90, 0xcb,
	0xaa, 0x30, 0xb4, 0xeb, 0xa3, 0xa0, 0x15, 0xd7, 0x6a, 0x74, 0x8d, 0xf1, 0xc4, 0xe6, 0x98, 0x2a,
	0x60, 0xdc, 0x0e, 0x5f, 0x82, 0x59, 0xd6, 0x09, 0x5d, 0x4d, 0xce, 0x30, 0x4e, 0x41, 0xf3, 0xe4,
	0x98, 0xfd, 0x18, 0xb4, 0x67, 0x89, 0x7b, 0x77, 0x3b, 0x7b, 0xdb, 0x7b, 0x68, 0xb7, 0xf7, 0xd0,
	0xc7, 0xde, 0x43, 0xaf, 0x07, 0xaf, 0xb1, 0x3b, 0x78, 0x8d, 0xf7, 0x83, 0xd7, 0x98, 0xdd, 0x2c,
	0x84, 0x59, 0x56, 0x69, 0xc8, 0x64, 0x1e, 0xa5, 0x02, 0x8a, 0x67, 0xc1, 0x41, 0x44, 0x76, 0xef,
	0x97, 0x0b, 0x19, 0xe5, 0x32, 0xab, 0xd6, 0x5c, 0x47, 0x3f, 0xdf, 0xcc, 0x6c, 0x4b, 0xae, 0xd3,
	0x8e, 0xbb, 0xd5, 0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x39, 0x53, 0x35, 0xd9, 0x01,
	0x00, 0x00,
}

func (m *MultiTokenPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiTokenPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiTokenPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintMtTransfer(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.DestContract) > 0 {
		i -= len(m.DestContract)
		copy(dAtA[i:], m.DestContract)
		i = encodeVarintMtTransfer(dAtA, i, uint64(len(m.DestContract)))
		i--
		dAtA[i] = 0x3a
	}
	if m.AwayFromOrigin {
		i--
		if m.AwayFromOrigin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintMtTransfer(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMtTransfer(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMtTransfer(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintMtTransfer(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Class) > 0 {
		i -= len(m.Class)
		copy(dAtA[i:], m.Class)
		i = encodeVarintMtTransfer(dAtA, i, uint64(len(m.Class)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClassTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BaseClass) > 0 {
		i -= len(m.BaseClass)
		copy(dAtA[i:], m.BaseClass)
		i = encodeVarintMtTransfer(dAtA, i, uint64(len(m.BaseClass)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintMtTransfer(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMtTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovMtTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MultiTokenPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Class)
	if l > 0 {
		n += 1 + l + sovMtTransfer(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMtTransfer(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMtTransfer(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMtTransfer(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovMtTransfer(uint64(l))
	}
	if m.AwayFromOrigin {
		n += 2
	}
	l = len(m.DestContract)
	if l > 0 {
		n += 1 + l + sovMtTransfer(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovMtTransfer(uint64(m.Amount))
	}
	return n
}

func (m *ClassTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovMtTransfer(uint64(l))
	}
	l = len(m.BaseClass)
	if l > 0 {
		n += 1 + l + sovMtTransfer(uint64(l))
	}
	return n
}

func sovMtTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMtTransfer(x uint64) (n int) {
	return sovMtTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MultiTokenPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMtTransfer
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
			return fmt.Errorf("proto: MultiTokenPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiTokenPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Class", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
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
				return ErrInvalidLengthMtTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMtTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Class = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
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
				return ErrInvalidLengthMtTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMtTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
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
				return ErrInvalidLengthMtTransfer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMtTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
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
				return ErrInvalidLengthMtTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMtTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
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
				return ErrInvalidLengthMtTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMtTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AwayFromOrigin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
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
			m.AwayFromOrigin = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
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
				return ErrInvalidLengthMtTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMtTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMtTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMtTransfer
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
func (m *ClassTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMtTransfer
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
			return fmt.Errorf("proto: ClassTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
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
				return ErrInvalidLengthMtTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMtTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseClass", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMtTransfer
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
				return ErrInvalidLengthMtTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMtTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseClass = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMtTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMtTransfer
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
func skipMtTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMtTransfer
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
					return 0, ErrIntOverflowMtTransfer
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
					return 0, ErrIntOverflowMtTransfer
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
				return 0, ErrInvalidLengthMtTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMtTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMtTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMtTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMtTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMtTransfer = fmt.Errorf("proto: unexpected end of group")
)
