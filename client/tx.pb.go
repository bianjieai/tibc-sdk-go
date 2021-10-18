// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tibc/core/client/v1/tx.proto

package client

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/irisnet/core-sdk-go/codec/types"
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

// MsgUpdateClient defines an sdk.Msg to update a IBC client state using
// the given header.
type MsgUpdateClient struct {
	// client unique identifier
	ChainName string `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	// header to update the light client
	Header *types.Any `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	// signer address
	Signer string `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *MsgUpdateClient) Reset()         { *m = MsgUpdateClient{} }
func (m *MsgUpdateClient) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateClient) ProtoMessage()    {}
func (*MsgUpdateClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_d101b27e3bfc60c8, []int{0}
}
func (m *MsgUpdateClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateClient.Merge(m, src)
}
func (m *MsgUpdateClient) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateClient) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateClient.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateClient proto.InternalMessageInfo

// MsgUpdateClientResponse defines the Msg/UpdateClient response type.
type MsgUpdateClientResponse struct {
}

func (m *MsgUpdateClientResponse) Reset()         { *m = MsgUpdateClientResponse{} }
func (m *MsgUpdateClientResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateClientResponse) ProtoMessage()    {}
func (*MsgUpdateClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d101b27e3bfc60c8, []int{1}
}
func (m *MsgUpdateClientResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateClientResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateClientResponse.Merge(m, src)
}
func (m *MsgUpdateClientResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateClientResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateClient)(nil), "tibc.core.client.v1.MsgUpdateClient")
	proto.RegisterType((*MsgUpdateClientResponse)(nil), "tibc.core.client.v1.MsgUpdateClientResponse")
}

func init() { proto.RegisterFile("tibc/core/client/v1/tx.proto", fileDescriptor_d101b27e3bfc60c8) }

var fileDescriptor_d101b27e3bfc60c8 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4f, 0x32, 0x31,
	0x18, 0xc7, 0xaf, 0x2f, 0x09, 0x79, 0xa9, 0x26, 0x26, 0x27, 0x51, 0x20, 0x5a, 0x08, 0x31, 0x91,
	0x01, 0xda, 0x80, 0x9b, 0x1b, 0x3a, 0xe3, 0x40, 0xe2, 0xe2, 0x62, 0xda, 0xe3, 0xb1, 0x54, 0xa1,
	0x25, 0xd7, 0x42, 0x64, 0x73, 0x74, 0xf4, 0x23, 0xf0, 0x71, 0x1c, 0x19, 0x1d, 0x0d, 0xb7, 0xf8,
	0x31, 0xcc, 0xf5, 0x8e, 0x41, 0xe2, 0xe0, 0xd6, 0xa7, 0xbf, 0x7f, 0x7e, 0xfd, 0xb7, 0xc5, 0x27,
	0x4e, 0x89, 0x88, 0x45, 0x26, 0x06, 0x16, 0x4d, 0x14, 0x68, 0xc7, 0x16, 0x5d, 0xe6, 0x9e, 0xe9,
	0x2c, 0x36, 0xce, 0x84, 0x87, 0x29, 0xa5, 0x29, 0xa5, 0x19, 0xa5, 0x8b, 0x6e, 0xad, 0x2c, 0x8d,
	0x34, 0x9e, 0xb3, 0x74, 0x95, 0x45, 0x6b, 0x55, 0x69, 0x8c, 0x9c, 0x00, 0xf3, 0x93, 0x98, 0x3f,
	0x30, 0xae, 0x97, 0x19, 0x6a, 0xbe, 0x20, 0x7c, 0x30, 0xb0, 0xf2, 0x76, 0x36, 0xe2, 0x0e, 0xae,
	0xbd, 0x27, 0x3c, 0xc5, 0x38, 0x1a, 0x73, 0xa5, 0xef, 0x35, 0x9f, 0x42, 0x05, 0x35, 0x50, 0xab,
	0x34, 0x2c, 0xf9, 0x9d, 0x1b, 0x3e, 0x85, 0xb0, 0x8d, 0x8b, 0x63, 0xe0, 0x23, 0x88, 0x2b, 0xff,
	0x1a, 0xa8, 0xb5, 0xd7, 0x2b, 0xd3, 0x4c, 0x4f, 0xb7, 0x7a, 0xda, 0xd7, 0xcb, 0x61, 0x9e, 0x09,
	0x8f, 0x70, 0xd1, 0x2a, 0xa9, 0x21, 0xae, 0x14, 0xbc, 0x28, 0x9f, 0x2e, 0xff, 0xbf, 0xae, 0xea,
	0xc1, 0xd7, 0xaa, 0x1e, 0x34, 0xab, 0xf8, 0x78, 0xa7, 0xc1, 0x10, 0xec, 0xcc, 0x68, 0x0b, 0x3d,
	0x85, 0x0b, 0x03, 0x2b, 0x43, 0x81, 0xf7, 0x7f, 0x14, 0x3c, 0xa3, 0xbf, 0xdc, 0x9d, 0xee, 0x48,
	0x6a, 0xed, 0xbf, 0xa4, 0xb6, 0x47, 0x5d, 0xf5, 0xdf, 0x37, 0x04, 0xad, 0x37, 0x04, 0x7d, 0x6e,
	0x08, 0x7a, 0x4b, 0x48, 0xb0, 0x4e, 0x48, 0xf0, 0x91, 0x90, 0xe0, 0xee, 0x5c, 0x2a, 0x37, 0x9e,
	0x0b, 0x1a, 0x99, 0x29, 0x13, 0x8a, 0xeb, 0x47, 0x05, 0x5c, 0xb1, 0xd4, 0xdd, 0xb1, 0xa3, 0xa7,
	0x8e, 0x34, 0xf9, 0xef, 0x88, 0xa2, 0x7f, 0x80, 0x8b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5,
	0x36, 0x7a, 0x8d, 0xb8, 0x01, 0x00, 0x00,
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
	// UpdateClient defines a rpc handler method for MsgUpdateClient.
	UpdateClient(ctx context.Context, in *MsgUpdateClient, opts ...grpc.CallOption) (*MsgUpdateClientResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateClient(ctx context.Context, in *MsgUpdateClient, opts ...grpc.CallOption) (*MsgUpdateClientResponse, error) {
	out := new(MsgUpdateClientResponse)
	err := c.cc.Invoke(ctx, "/tibc.core.client.v1.Msg/UpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateClient defines a rpc handler method for MsgUpdateClient.
	UpdateClient(context.Context, *MsgUpdateClient) (*MsgUpdateClientResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateClient(ctx context.Context, req *MsgUpdateClient) (*MsgUpdateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tibc.core.client.v1.Msg/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateClient(ctx, req.(*MsgUpdateClient))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tibc.core.client.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateClient",
			Handler:    _Msg_UpdateClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tibc/core/client/v1/tx.proto",
}

func (m *MsgUpdateClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateClientResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateClientResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateClientResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgUpdateClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateClientResponse) Size() (n int) {
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
func (m *MsgUpdateClient) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateClient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
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
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
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
			if m.Header == nil {
				m.Header = &types.Any{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
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
			m.Signer = string(dAtA[iNdEx:postIndex])
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
func (m *MsgUpdateClientResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateClientResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateClientResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
