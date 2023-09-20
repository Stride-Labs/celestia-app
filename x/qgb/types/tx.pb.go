// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: celestia/qgb/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// MsgRegisterEVMAddress registers an evm address to a validator.
type MsgRegisterEVMAddress struct {
	// The operating address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// The matching HEX encoded EVM address.
	EvmAddress string `protobuf:"bytes,2,opt,name=evm_address,json=evmAddress,proto3" json:"evm_address,omitempty"`
}

func (m *MsgRegisterEVMAddress) Reset()         { *m = MsgRegisterEVMAddress{} }
func (m *MsgRegisterEVMAddress) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterEVMAddress) ProtoMessage()    {}
func (*MsgRegisterEVMAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ed1095628e2204, []int{0}
}
func (m *MsgRegisterEVMAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterEVMAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterEVMAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterEVMAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterEVMAddress.Merge(m, src)
}
func (m *MsgRegisterEVMAddress) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterEVMAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterEVMAddress.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterEVMAddress proto.InternalMessageInfo

func (m *MsgRegisterEVMAddress) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *MsgRegisterEVMAddress) GetEvmAddress() string {
	if m != nil {
		return m.EvmAddress
	}
	return ""
}

// MsgRegisterEVMAddressResponse is the response to registering an EVM address.
type MsgRegisterEVMAddressResponse struct {
}

func (m *MsgRegisterEVMAddressResponse) Reset()         { *m = MsgRegisterEVMAddressResponse{} }
func (m *MsgRegisterEVMAddressResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterEVMAddressResponse) ProtoMessage()    {}
func (*MsgRegisterEVMAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ed1095628e2204, []int{1}
}
func (m *MsgRegisterEVMAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterEVMAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterEVMAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterEVMAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterEVMAddressResponse.Merge(m, src)
}
func (m *MsgRegisterEVMAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterEVMAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterEVMAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterEVMAddressResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterEVMAddress)(nil), "celestia.qgb.v1.MsgRegisterEVMAddress")
	proto.RegisterType((*MsgRegisterEVMAddressResponse)(nil), "celestia.qgb.v1.MsgRegisterEVMAddressResponse")
}

func init() { proto.RegisterFile("celestia/qgb/v1/tx.proto", fileDescriptor_85ed1095628e2204) }

var fileDescriptor_85ed1095628e2204 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0x4a, 0x2b, 0x41,
	0x14, 0xc7, 0x33, 0xb9, 0xdc, 0x0b, 0x77, 0x2c, 0x34, 0x4b, 0xc4, 0xb8, 0xc4, 0x89, 0x44, 0x11,
	0x09, 0x64, 0x87, 0x68, 0x67, 0x67, 0x20, 0x85, 0x45, 0x9a, 0x15, 0x2c, 0x6c, 0xc2, 0x6c, 0x32,
	0x8c, 0x03, 0xd9, 0x3d, 0x9b, 0x9d, 0x71, 0x89, 0x9d, 0x58, 0x5a, 0x89, 0x3e, 0x81, 0xe0, 0x03,
	0xa4, 0xf0, 0x21, 0x2c, 0x83, 0x36, 0x96, 0x92, 0x08, 0x79, 0x0d, 0xc9, 0x7e, 0x21, 0xb2, 0x85,
	0xcd, 0x72, 0xf6, 0xfc, 0xfe, 0xf3, 0x3f, 0x5f, 0xb8, 0xd2, 0xe7, 0x43, 0xae, 0xb4, 0x64, 0x74,
	0x24, 0x1c, 0x1a, 0xb6, 0xa8, 0x1e, 0x5b, 0x7e, 0x00, 0x1a, 0x8c, 0xd5, 0x94, 0x58, 0x23, 0xe1,
	0x58, 0x61, 0xcb, 0x2c, 0x0b, 0x10, 0x10, 0x31, 0xba, 0x8c, 0x62, 0x99, 0xb9, 0xd9, 0x07, 0xe5,
	0x82, 0xea, 0xc5, 0x20, 0xfe, 0x49, 0x50, 0x55, 0x00, 0x88, 0x21, 0xa7, 0xcc, 0x97, 0x94, 0x79,
	0x1e, 0x68, 0xa6, 0x25, 0x78, 0x29, 0xdd, 0x88, 0xb5, 0xd4, 0x55, 0x62, 0x59, 0xd7, 0x55, 0x22,
	0x01, 0x25, 0xe6, 0x4a, 0x0f, 0x68, 0xf4, 0x8d, 0x53, 0xf5, 0x27, 0x84, 0xd7, 0xbb, 0x4a, 0xd8,
	0x5c, 0x48, 0xa5, 0x79, 0xd0, 0x39, 0xeb, 0x1e, 0x0f, 0x06, 0x01, 0x57, 0xca, 0xe8, 0xe0, 0x52,
	0xc8, 0x86, 0x72, 0xc0, 0x34, 0x04, 0x3d, 0x16, 0x27, 0x2b, 0x68, 0x1b, 0xed, 0xff, 0x6f, 0x57,
	0x5e, 0x9f, 0x9b, 0xe5, 0xa4, 0xa1, 0x44, 0x7e, 0xaa, 0x03, 0xe9, 0x09, 0x7b, 0x2d, 0x7b, 0x92,
	0xda, 0xd4, 0xf0, 0x0a, 0x0f, 0xdd, 0xcc, 0xa0, 0xb8, 0x34, 0xb0, 0x31, 0x0f, 0xdd, 0x44, 0x70,
	0xb4, 0x73, 0xbb, 0x98, 0x34, 0x48, 0xb6, 0xac, 0xdc, 0x66, 0xea, 0x35, 0xbc, 0x95, 0x0b, 0x6c,
	0xae, 0x7c, 0xf0, 0x14, 0x3f, 0x78, 0x44, 0xf8, 0x4f, 0x57, 0x09, 0xe3, 0x1e, 0x61, 0x23, 0x67,
	0x98, 0x3d, 0xeb, 0xc7, 0xce, 0xad, 0x5c, 0x3b, 0xd3, 0xfa, 0x9d, 0x2e, 0x2d, 0x5b, 0xdf, 0xbd,
	0x79, 0xfb, 0x7c, 0x28, 0x12, 0xa3, 0x9a, 0x1e, 0x39, 0x48, 0xb4, 0xbd, 0x6f, 0x43, 0x9b, 0x7f,
	0xaf, 0x17, 0x93, 0x06, 0x6a, 0x9f, 0xbc, 0xcc, 0x08, 0x9a, 0xce, 0x08, 0xfa, 0x98, 0x11, 0x74,
	0x37, 0x27, 0x85, 0xe9, 0x9c, 0x14, 0xde, 0xe7, 0xa4, 0x70, 0x4e, 0x85, 0xd4, 0x17, 0x97, 0x8e,
	0xd5, 0x07, 0x97, 0xa6, 0x0d, 0x40, 0x20, 0xb2, 0xb8, 0xc9, 0x7c, 0x9f, 0x8e, 0xa3, 0x1a, 0xfa,
	0xca, 0xe7, 0xca, 0xf9, 0x17, 0x5d, 0xef, 0xf0, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xad, 0x0d, 0x20,
	0xf7, 0x65, 0x02, 0x00, 0x00,
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
	// RegisterEVMAddress records an evm address for the validator which is used
	// by the relayer to aggregate signatures. A validator can only register a
	// single EVM address. The EVM address can be overridden by a later message.
	// There are no validity checks of the EVM addresses existence on the Ethereum
	// state machine.
	RegisterEVMAddress(ctx context.Context, in *MsgRegisterEVMAddress, opts ...grpc.CallOption) (*MsgRegisterEVMAddressResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterEVMAddress(ctx context.Context, in *MsgRegisterEVMAddress, opts ...grpc.CallOption) (*MsgRegisterEVMAddressResponse, error) {
	out := new(MsgRegisterEVMAddressResponse)
	err := c.cc.Invoke(ctx, "/celestia.qgb.v1.Msg/RegisterEVMAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RegisterEVMAddress records an evm address for the validator which is used
	// by the relayer to aggregate signatures. A validator can only register a
	// single EVM address. The EVM address can be overridden by a later message.
	// There are no validity checks of the EVM addresses existence on the Ethereum
	// state machine.
	RegisterEVMAddress(context.Context, *MsgRegisterEVMAddress) (*MsgRegisterEVMAddressResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterEVMAddress(ctx context.Context, req *MsgRegisterEVMAddress) (*MsgRegisterEVMAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterEVMAddress not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterEVMAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterEVMAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterEVMAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/celestia.qgb.v1.Msg/RegisterEVMAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterEVMAddress(ctx, req.(*MsgRegisterEVMAddress))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "celestia.qgb.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterEVMAddress",
			Handler:    _Msg_RegisterEVMAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "celestia/qgb/v1/tx.proto",
}

func (m *MsgRegisterEVMAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterEVMAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterEVMAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EvmAddress) > 0 {
		i -= len(m.EvmAddress)
		copy(dAtA[i:], m.EvmAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EvmAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterEVMAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterEVMAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterEVMAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgRegisterEVMAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EvmAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterEVMAddressResponse) Size() (n int) {
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
func (m *MsgRegisterEVMAddress) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterEVMAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterEVMAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvmAddress", wireType)
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
			m.EvmAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MsgRegisterEVMAddressResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterEVMAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterEVMAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
