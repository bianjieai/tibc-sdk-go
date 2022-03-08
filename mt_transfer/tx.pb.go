// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tibc/apps/mt_transfer/v1/tx.proto

package mt_transfer

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgMtTransfer struct {
	// the class to which the mt to be transferred belongs
	Class string `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	// the mt id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// the mt sender
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// the mt receiver
	Receiver string `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// target chain of transmission
	DestChain string `protobuf:"bytes,5,opt,name=dest_chain,json=destChain,proto3" json:"dest_chain,omitempty"`
	// relay chain during transmission
	RealayChain string `protobuf:"bytes,6,opt,name=realay_chain,json=realayChain,proto3" json:"realay_chain,omitempty"`
	// the destination contract address to receive the nft
	DestContract string `protobuf:"bytes,7,opt,name=dest_contract,json=destContract,proto3" json:"dest_contract,omitempty"`
	// the amount defined by MT outside the chain
	Amount uint64 `protobuf:"varint,8,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MsgMtTransfer) Reset()         { *m = MsgMtTransfer{} }
func (m *MsgMtTransfer) String() string { return proto.CompactTextString(m) }
func (*MsgMtTransfer) ProtoMessage()    {}
func (*MsgMtTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2978e16e5405f501, []int{0}
}
func (m *MsgMtTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMtTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMtTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMtTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMtTransfer.Merge(m, src)
}
func (m *MsgMtTransfer) XXX_Size() int {
	return m.Size()
}
func (m *MsgMtTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMtTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMtTransfer proto.InternalMessageInfo

// MsgMtTransferResponse defines the Msg/MtTransfer response type.
type MsgMtTransferResponse struct {
}

func (m *MsgMtTransferResponse) Reset()         { *m = MsgMtTransferResponse{} }
func (m *MsgMtTransferResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMtTransferResponse) ProtoMessage()    {}
func (*MsgMtTransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2978e16e5405f501, []int{1}
}
func (m *MsgMtTransferResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMtTransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMtTransferResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMtTransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMtTransferResponse.Merge(m, src)
}
func (m *MsgMtTransferResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMtTransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMtTransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMtTransferResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgMtTransfer)(nil), "tibc.apps.mt_transfer.v1.MsgMtTransfer")
	proto.RegisterType((*MsgMtTransferResponse)(nil), "tibc.apps.mt_transfer.v1.MsgMtTransferResponse")
}

func init() { proto.RegisterFile("tibc/apps/mt_transfer/v1/tx.proto", fileDescriptor_2978e16e5405f501) }

var fileDescriptor_2978e16e5405f501 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbd, 0x4e, 0xeb, 0x30,
	0x18, 0x86, 0x93, 0xfe, 0x9d, 0xf6, 0x3b, 0xed, 0x19, 0xac, 0x1e, 0xb0, 0x2a, 0x91, 0xfe, 0x30,
	0xd0, 0x85, 0x58, 0x85, 0x8d, 0x09, 0xc1, 0xdc, 0xa5, 0x62, 0xea, 0x52, 0x39, 0x89, 0x9b, 0x1a,
	0x35, 0x71, 0x64, 0xbb, 0x15, 0xbd, 0x03, 0x46, 0x2e, 0xa1, 0x97, 0xc3, 0xd8, 0x91, 0x11, 0xb5,
	0x0b, 0x33, 0x57, 0x80, 0xe2, 0x04, 0xd4, 0x4a, 0x20, 0xb1, 0xe5, 0x7d, 0xbe, 0x47, 0x6f, 0x92,
	0xcf, 0x86, 0xae, 0xe6, 0x9e, 0x4f, 0x68, 0x92, 0x28, 0x12, 0xe9, 0x89, 0x96, 0x34, 0x56, 0x53,
	0x26, 0xc9, 0x72, 0x40, 0xf4, 0x83, 0x9b, 0x48, 0xa1, 0x05, 0xc2, 0xa9, 0xe2, 0xa6, 0x8a, 0xbb,
	0xa7, 0xb8, 0xcb, 0x41, 0xab, 0x19, 0x8a, 0x50, 0x18, 0x89, 0xa4, 0x4f, 0x99, 0xdf, 0x7b, 0xb7,
	0xa1, 0x31, 0x54, 0xe1, 0x50, 0xdf, 0xe5, 0x2a, 0x6a, 0x42, 0xd9, 0x9f, 0x53, 0xa5, 0xb0, 0xdd,
	0xb1, 0xfb, 0xb5, 0x51, 0x16, 0xd0, 0x3f, 0x28, 0xf0, 0x00, 0x17, 0x0c, 0x2a, 0xf0, 0x00, 0x1d,
	0x41, 0x45, 0xb1, 0x38, 0x60, 0x12, 0x17, 0x0d, 0xcb, 0x13, 0x6a, 0x41, 0x55, 0x32, 0x9f, 0xf1,
	0x25, 0x93, 0xb8, 0x64, 0x26, 0x5f, 0x19, 0x9d, 0x00, 0x04, 0x4c, 0xe9, 0x89, 0x3f, 0xa3, 0x3c,
	0xc6, 0x65, 0x33, 0xad, 0xa5, 0xe4, 0x36, 0x05, 0xa8, 0x0b, 0x75, 0xc9, 0xe8, 0x9c, 0xae, 0x72,
	0xa1, 0x62, 0x84, 0xbf, 0x19, 0xcb, 0x94, 0x53, 0x68, 0x64, 0x0d, 0x22, 0xd6, 0x92, 0xfa, 0x1a,
	0xff, 0x31, 0x4e, 0xdd, 0x94, 0xe4, 0x2c, 0xfd, 0x34, 0x1a, 0x89, 0x45, 0xac, 0x71, 0xb5, 0x63,
	0xf7, 0x4b, 0xa3, 0x3c, 0x5d, 0x55, 0x1f, 0xd7, 0x6d, 0xeb, 0x6d, 0xdd, 0xb6, 0x7a, 0xc7, 0xf0,
	0xff, 0xe0, 0x9f, 0x47, 0x4c, 0x25, 0x22, 0x56, 0xec, 0x22, 0x82, 0xe2, 0x50, 0x85, 0x68, 0x0a,
	0xb0, 0xb7, 0x90, 0x33, 0xf7, 0xa7, 0x9d, 0xba, 0x07, 0x2d, 0x2d, 0xf2, 0x4b, 0xf1, 0xf3, 0x75,
	0x37, 0xe3, 0xe7, 0xad, 0x63, 0x6f, 0xb6, 0x8e, 0xfd, 0xba, 0x75, 0xec, 0xa7, 0x9d, 0x63, 0x6d,
	0x76, 0x8e, 0xf5, 0xb2, 0x73, 0xac, 0xf1, 0x75, 0xc8, 0xf5, 0x6c, 0xe1, 0xb9, 0xbe, 0x88, 0x88,
	0xc7, 0x69, 0x7c, 0xcf, 0x19, 0xe5, 0x24, 0xad, 0x3f, 0x0f, 0x05, 0x89, 0x44, 0xb0, 0x98, 0x33,
	0x45, 0xbe, 0xbf, 0x0e, 0x7a, 0x95, 0x30, 0xe5, 0x55, 0xcc, 0xf9, 0x5e, 0x7e, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x1d, 0x0a, 0xa5, 0x3c, 0x34, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// MtTransfer defines a rpc handler method for MsgMtTransfer.
	MtTransfer(ctx context.Context, in *MsgMtTransfer, opts ...grpc.CallOption) (*MsgMtTransferResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MtTransfer(ctx context.Context, in *MsgMtTransfer, opts ...grpc.CallOption) (*MsgMtTransferResponse, error) {
	out := new(MsgMtTransferResponse)
	err := c.cc.Invoke(ctx, "/tibc.apps.mt_transfer.v1.Msg/MtTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// MtTransfer defines a rpc handler method for MsgMtTransfer.
	MtTransfer(context.Context, *MsgMtTransfer) (*MsgMtTransferResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MtTransfer(ctx context.Context, req *MsgMtTransfer) (*MsgMtTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MtTransfer not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MtTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMtTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MtTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tibc.apps.mt_transfer.v1.Msg/MtTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MtTransfer(ctx, req.(*MsgMtTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tibc.apps.mt_transfer.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MtTransfer",
			Handler:    _Msg_MtTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tibc/apps/mt_transfer/v1/tx.proto",
}

func (m *MsgMtTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMtTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMtTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.DestContract) > 0 {
		i -= len(m.DestContract)
		copy(dAtA[i:], m.DestContract)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DestContract)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.RealayChain) > 0 {
		i -= len(m.RealayChain)
		copy(dAtA[i:], m.RealayChain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RealayChain)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DestChain) > 0 {
		i -= len(m.DestChain)
		copy(dAtA[i:], m.DestChain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DestChain)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Class) > 0 {
		i -= len(m.Class)
		copy(dAtA[i:], m.Class)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Class)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMtTransferResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMtTransferResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMtTransferResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgMtTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Class)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DestChain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RealayChain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DestContract)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTx(uint64(m.Amount))
	}
	return n
}

func (m *MsgMtTransferResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgMtTransfer) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMtTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMtTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Class", wireType)
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
			m.Class = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
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
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestChain", wireType)
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
			m.DestChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RealayChain", wireType)
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
			m.RealayChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestContract", wireType)
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
			m.DestContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
func (m *MsgMtTransferResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMtTransferResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMtTransferResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
