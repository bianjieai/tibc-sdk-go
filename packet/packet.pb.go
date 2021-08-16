// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tibc/core/packet/v1/packet.proto

package packet

import (
	fmt "fmt"
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

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	// number corresponds to the order of sends and receives, where a Packet
	// with an earlier sequence number must be sent and received before a Packet
	// with a later sequence number.
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// identifies the port on the sending chain and destination chain.
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	// identifies the chain id of the sending chain.
	SourceChain string `protobuf:"bytes,3,opt,name=source_chain,json=sourceChain,proto3" json:"source_chain,omitempty" yaml:"source_chain"`
	// identifies the chain id of the receiving chain.
	DestinationChain string `protobuf:"bytes,4,opt,name=destination_chain,json=destinationChain,proto3" json:"destination_chain,omitempty" yaml:"destination_port"`
	// identifies the chain id of the relay chain.
	RelayChain string `protobuf:"bytes,5,opt,name=relay_chain,json=relayChain,proto3" json:"relay_chain,omitempty" yaml:"relay_chain"`
	// actual opaque bytes transferred directly to the application module
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_59da8d378a3b7468, []int{0}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return m.Size()
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

// PacketState defines the generic type necessary to retrieve and store
// packet commitments, acknowledgements, and receipts.
// Caller is responsible for knowing the context necessary to interpret this
// state as a commitment, acknowledgement, or a receipt.
type PacketState struct {
	// the sending chain identifier.
	SourceChain string `protobuf:"bytes,1,opt,name=source_chain,json=sourceChain,proto3" json:"source_chain,omitempty" yaml:"source_chain"`
	// the receiving chain identifier.
	DestinationChain string `protobuf:"bytes,2,opt,name=destination_chain,json=destinationChain,proto3" json:"destination_chain,omitempty" yaml:"source_chain"`
	// packet sequence.
	Sequence uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// embedded data that represents packet state.
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PacketState) Reset()         { *m = PacketState{} }
func (m *PacketState) String() string { return proto.CompactTextString(m) }
func (*PacketState) ProtoMessage()    {}
func (*PacketState) Descriptor() ([]byte, []int) {
	return fileDescriptor_59da8d378a3b7468, []int{1}
}
func (m *PacketState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketState.Merge(m, src)
}
func (m *PacketState) XXX_Size() int {
	return m.Size()
}
func (m *PacketState) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketState.DiscardUnknown(m)
}

var xxx_messageInfo_PacketState proto.InternalMessageInfo

// Acknowledgement is the recommended acknowledgement format to be used by
// app-specific protocols.
// NOTE: The field numbers 21 and 22 were explicitly chosen to avoid accidental
// conflicts with other protobuf message formats used for acknowledgements.
// The first byte of any message with this format will be the non-ASCII values
// `0xaa` (result) or `0xb2` (error). Implemented as defined by ICS:
// https://github.com/cosmos/ics/tree/master/spec/ics-004-channel-and-packet-semantics#acknowledgement-envelope
type Acknowledgement struct {
	// response contains either a result or an error and must be non-empty
	//
	// Types that are valid to be assigned to Response:
	//	*Acknowledgement_Result
	//	*Acknowledgement_Error
	Response isAcknowledgement_Response `protobuf_oneof:"response"`
}

func (m *Acknowledgement) Reset()         { *m = Acknowledgement{} }
func (m *Acknowledgement) String() string { return proto.CompactTextString(m) }
func (*Acknowledgement) ProtoMessage()    {}
func (*Acknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_59da8d378a3b7468, []int{2}
}
func (m *Acknowledgement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Acknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Acknowledgement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Acknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Acknowledgement.Merge(m, src)
}
func (m *Acknowledgement) XXX_Size() int {
	return m.Size()
}
func (m *Acknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_Acknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_Acknowledgement proto.InternalMessageInfo

type isAcknowledgement_Response interface {
	isAcknowledgement_Response()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Acknowledgement_Result struct {
	Result []byte `protobuf:"bytes,21,opt,name=result,proto3,oneof" json:"result,omitempty"`
}
type Acknowledgement_Error struct {
	Error string `protobuf:"bytes,22,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (*Acknowledgement_Result) isAcknowledgement_Response() {}
func (*Acknowledgement_Error) isAcknowledgement_Response()  {}

func (m *Acknowledgement) GetResponse() isAcknowledgement_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Acknowledgement) GetResult() []byte {
	if x, ok := m.GetResponse().(*Acknowledgement_Result); ok {
		return x.Result
	}
	return nil
}

func (m *Acknowledgement) GetError() string {
	if x, ok := m.GetResponse().(*Acknowledgement_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Acknowledgement) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Acknowledgement_Result)(nil),
		(*Acknowledgement_Error)(nil),
	}
}

func init() {
	proto.RegisterType((*Packet)(nil), "tibc.core.packet.v1.Packet")
	proto.RegisterType((*PacketState)(nil), "tibc.core.packet.v1.PacketState")
	proto.RegisterType((*Acknowledgement)(nil), "tibc.core.packet.v1.Acknowledgement")
}

func init() { proto.RegisterFile("tibc/core/packet/v1/packet.proto", fileDescriptor_59da8d378a3b7468) }

var fileDescriptor_59da8d378a3b7468 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0xe3, 0x2e, 0xab, 0x86, 0x3b, 0x09, 0xf0, 0xa0, 0x8b, 0x86, 0x94, 0x56, 0xb9, 0xd0,
	0xcb, 0x1a, 0x4d, 0x1c, 0x90, 0x7a, 0x5b, 0xe1, 0xd0, 0x1b, 0xc8, 0xdc, 0xb8, 0x20, 0xd7, 0xf9,
	0x2b, 0x33, 0x4d, 0xed, 0xe0, 0x38, 0x43, 0x7b, 0x03, 0x8e, 0xf0, 0x06, 0xbc, 0x0c, 0x12, 0xc7,
	0x1e, 0x39, 0x55, 0xa8, 0x7d, 0x83, 0x3e, 0x01, 0xb2, 0x1d, 0xa1, 0x48, 0x41, 0x48, 0xbb, 0xfd,
	0x3f, 0x7f, 0x5f, 0x7e, 0xc9, 0x17, 0xff, 0xf1, 0xd8, 0x88, 0x25, 0x4f, 0xb9, 0xd2, 0x90, 0x96,
	0x8c, 0xaf, 0xc0, 0xa4, 0xb7, 0x57, 0xcd, 0x34, 0x2d, 0xb5, 0x32, 0x8a, 0x9c, 0xd9, 0xc4, 0xd4,
	0x26, 0xa6, 0xcd, 0xf9, 0xed, 0xd5, 0xc5, 0x93, 0x5c, 0xe5, 0xca, 0xf9, 0xa9, 0x9d, 0x7c, 0x34,
	0xf9, 0xd6, 0xc3, 0xfd, 0xb7, 0x2e, 0x43, 0x2e, 0xf0, 0x49, 0x05, 0x9f, 0x6a, 0x90, 0x1c, 0x22,
	0x34, 0x46, 0x93, 0x90, 0xfe, 0xd5, 0x84, 0xe0, 0xb0, 0x54, 0xda, 0x44, 0xbd, 0x31, 0x9a, 0x3c,
	0xa0, 0x6e, 0x26, 0x33, 0x7c, 0x5a, 0xa9, 0x5a, 0x73, 0xf8, 0xc0, 0x6f, 0x98, 0x90, 0xd1, 0x91,
	0xf5, 0xe6, 0xe7, 0x87, 0xed, 0xe8, 0xec, 0x8e, 0xad, 0x8b, 0x59, 0xd2, 0x76, 0x13, 0x3a, 0xf0,
	0xf2, 0x95, 0x55, 0x64, 0x81, 0x1f, 0x67, 0x50, 0x19, 0x21, 0x99, 0x11, 0x4a, 0x36, 0x80, 0xd0,
	0x01, 0x9e, 0x1d, 0xb6, 0xa3, 0x73, 0x0f, 0x68, 0x47, 0xec, 0x2b, 0x13, 0xfa, 0xa8, 0x75, 0xe4,
	0x49, 0x2f, 0xf1, 0x40, 0x43, 0xc1, 0xee, 0x1a, 0xc6, 0xb1, 0x63, 0x0c, 0x0f, 0xdb, 0x11, 0xf1,
	0x8c, 0x96, 0x99, 0x50, 0xec, 0x94, 0x7f, 0x90, 0xe0, 0x30, 0x63, 0x86, 0x45, 0xfd, 0x31, 0x9a,
	0x9c, 0x52, 0x37, 0xcf, 0xc2, 0x2f, 0xdf, 0x47, 0x41, 0xf2, 0x03, 0xe1, 0x81, 0xff, 0x27, 0xef,
	0x0c, 0x33, 0xd0, 0x29, 0x8a, 0xee, 0x51, 0xf4, 0xf5, 0xbf, 0x8a, 0xf6, 0xfe, 0x0f, 0xe8, 0x96,
	0x6c, 0x5f, 0xcd, 0x51, 0xf7, 0x6a, 0x5c, 0x8f, 0xb0, 0xd3, 0xe3, 0x0d, 0x7e, 0x78, 0xcd, 0x57,
	0x52, 0x7d, 0x2e, 0x20, 0xcb, 0x61, 0x0d, 0xd2, 0x90, 0x08, 0xf7, 0x35, 0x54, 0x75, 0x61, 0xa2,
	0xa7, 0x36, 0xbe, 0x08, 0x68, 0xa3, 0xc9, 0x10, 0x1f, 0x83, 0xd6, 0x4a, 0x47, 0x43, 0xfb, 0x71,
	0x8b, 0x80, 0x7a, 0x39, 0xc7, 0xf8, 0x44, 0x43, 0x55, 0x2a, 0x59, 0xc1, 0xfc, 0xfa, 0xe7, 0x2e,
	0x46, 0x9b, 0x5d, 0x8c, 0x7e, 0xef, 0x62, 0xf4, 0x75, 0x1f, 0x07, 0x9b, 0x7d, 0x1c, 0xfc, 0xda,
	0xc7, 0xc1, 0xfb, 0xe7, 0xb9, 0x30, 0x37, 0xf5, 0x72, 0xca, 0xd5, 0x3a, 0x5d, 0x0a, 0x26, 0x3f,
	0x0a, 0x60, 0x22, 0xb5, 0x6b, 0x78, 0x59, 0x65, 0xab, 0xcb, 0x5c, 0x35, 0x0b, 0xba, 0xec, 0xbb,
	0xb5, 0x7b, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x30, 0x78, 0x98, 0xab, 0xc5, 0x02, 0x00, 0x00,
}

func (m *Packet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Packet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.RelayChain) > 0 {
		i -= len(m.RelayChain)
		copy(dAtA[i:], m.RelayChain)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.RelayChain)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DestinationChain) > 0 {
		i -= len(m.DestinationChain)
		copy(dAtA[i:], m.DestinationChain)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.DestinationChain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceChain) > 0 {
		i -= len(m.SourceChain)
		copy(dAtA[i:], m.SourceChain)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.SourceChain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sequence != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PacketState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if m.Sequence != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DestinationChain) > 0 {
		i -= len(m.DestinationChain)
		copy(dAtA[i:], m.DestinationChain)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.DestinationChain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourceChain) > 0 {
		i -= len(m.SourceChain)
		copy(dAtA[i:], m.SourceChain)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.SourceChain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Acknowledgement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Acknowledgement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Acknowledgement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Response != nil {
		{
			size := m.Response.Size()
			i -= size
			if _, err := m.Response.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Acknowledgement_Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Acknowledgement_Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Result != nil {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	return len(dAtA) - i, nil
}
func (m *Acknowledgement_Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Acknowledgement_Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Error)
	copy(dAtA[i:], m.Error)
	i = encodeVarintPacket(dAtA, i, uint64(len(m.Error)))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0xb2
	return len(dAtA) - i, nil
}
func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Packet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovPacket(uint64(m.Sequence))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.SourceChain)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.DestinationChain)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.RelayChain)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *PacketState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceChain)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.DestinationChain)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovPacket(uint64(m.Sequence))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *Acknowledgement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Response != nil {
		n += m.Response.Size()
	}
	return n
}

func (m *Acknowledgement_Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		l = len(m.Result)
		n += 2 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *Acknowledgement_Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	n += 2 + l + sovPacket(uint64(l))
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Packet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: Packet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Packet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
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
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *PacketState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: PacketState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
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
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *Acknowledgement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: Acknowledgement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Acknowledgement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Response = &Acknowledgement_Result{v}
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = &Acknowledgement_Error{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
