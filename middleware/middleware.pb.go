// Code generated by protoc-gen-go. DO NOT EDIT.
// source: middleware/middleware.proto

package middleware

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type JwtRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwtRequest) Reset()         { *m = JwtRequest{} }
func (m *JwtRequest) String() string { return proto.CompactTextString(m) }
func (*JwtRequest) ProtoMessage()    {}
func (*JwtRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d89c3b3ad9147, []int{0}
}

func (m *JwtRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtRequest.Unmarshal(m, b)
}
func (m *JwtRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtRequest.Marshal(b, m, deterministic)
}
func (m *JwtRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtRequest.Merge(m, src)
}
func (m *JwtRequest) XXX_Size() int {
	return xxx_messageInfo_JwtRequest.Size(m)
}
func (m *JwtRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JwtRequest proto.InternalMessageInfo

func (m *JwtRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type JwtResponse struct {
	Firstname            string   `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Authorized           bool     `protobuf:"varint,4,opt,name=authorized,proto3" json:"authorized,omitempty"`
	Exp                  string   `protobuf:"bytes,5,opt,name=exp,proto3" json:"exp,omitempty"`
	Id                   string   `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	Company              string   `protobuf:"bytes,7,opt,name=company,proto3" json:"company,omitempty"`
	Type                 string   `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwtResponse) Reset()         { *m = JwtResponse{} }
func (m *JwtResponse) String() string { return proto.CompactTextString(m) }
func (*JwtResponse) ProtoMessage()    {}
func (*JwtResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d89c3b3ad9147, []int{1}
}

func (m *JwtResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtResponse.Unmarshal(m, b)
}
func (m *JwtResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtResponse.Marshal(b, m, deterministic)
}
func (m *JwtResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtResponse.Merge(m, src)
}
func (m *JwtResponse) XXX_Size() int {
	return xxx_messageInfo_JwtResponse.Size(m)
}
func (m *JwtResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JwtResponse proto.InternalMessageInfo

func (m *JwtResponse) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *JwtResponse) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *JwtResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *JwtResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func (m *JwtResponse) GetExp() string {
	if m != nil {
		return m.Exp
	}
	return ""
}

func (m *JwtResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JwtResponse) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *JwtResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*JwtRequest)(nil), "middleware.JwtRequest")
	proto.RegisterType((*JwtResponse)(nil), "middleware.JwtResponse")
}

func init() {
	proto.RegisterFile("middleware/middleware.proto", fileDescriptor_ce5d89c3b3ad9147)
}

var fileDescriptor_ce5d89c3b3ad9147 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xdd, 0xed, 0xbf, 0xed, 0x08, 0x22, 0x83, 0xe8, 0x50, 0x45, 0xca, 0x9e, 0x7a, 0xaa,
	0xa0, 0x0f, 0xe0, 0x5d, 0xf0, 0xb2, 0x3e, 0x41, 0x6c, 0x46, 0x0c, 0xee, 0x6e, 0x62, 0x92, 0x5a,
	0xeb, 0x3b, 0xfa, 0x4e, 0xd2, 0xd9, 0x6a, 0x56, 0x7a, 0xfb, 0x7e, 0xbf, 0x99, 0x0c, 0xe4, 0x83,
	0xcb, 0xc6, 0x68, 0x5d, 0xf3, 0x46, 0x79, 0xbe, 0x49, 0x71, 0xe9, 0xbc, 0x8d, 0x16, 0x21, 0x99,
	0xb2, 0x04, 0x78, 0xd8, 0xc4, 0x8a, 0xdf, 0xd7, 0x1c, 0x22, 0x9e, 0xc1, 0x28, 0xda, 0x37, 0x6e,
	0x29, 0x9b, 0x67, 0x8b, 0x69, 0xd5, 0x41, 0xf9, 0x9d, 0xc1, 0xb1, 0x2c, 0x05, 0x67, 0xdb, 0xc0,
	0x78, 0x05, 0xd3, 0x17, 0xe3, 0x43, 0x6c, 0x55, 0xc3, 0xfb, 0xcd, 0x24, 0x70, 0x06, 0x45, 0xad,
	0xf6, 0xc3, 0x5c, 0x86, 0x7f, 0xbc, 0xbb, 0xcf, 0x8d, 0x32, 0x35, 0x0d, 0xba, 0xfb, 0x02, 0x78,
	0x0d, 0xa0, 0xd6, 0xf1, 0xd5, 0x7a, 0xf3, 0xc5, 0x9a, 0x86, 0xf3, 0x6c, 0x51, 0x54, 0x3d, 0x83,
	0xa7, 0x30, 0xe0, 0x4f, 0x47, 0x23, 0x79, 0xb3, 0x8b, 0x78, 0x02, 0xb9, 0xd1, 0x34, 0x16, 0x91,
	0x1b, 0x8d, 0x04, 0x93, 0x95, 0x6d, 0x9c, 0x6a, 0xb7, 0x34, 0x11, 0xf9, 0x8b, 0x88, 0x30, 0x8c,
	0x5b, 0xc7, 0x54, 0x88, 0x96, 0x7c, 0xfb, 0x28, 0x7f, 0x7e, 0x62, 0xff, 0x61, 0x56, 0x8c, 0xf7,
	0xff, 0xe8, 0x7c, 0xd9, 0xab, 0x2b, 0x35, 0x33, 0xbb, 0x38, 0xf0, 0x5d, 0x19, 0xe5, 0xd1, 0xf3,
	0x58, 0x5a, 0xbd, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x40, 0x96, 0x11, 0xda, 0x74, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JwtServiceClient is the client API for JwtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JwtServiceClient interface {
	JwtService(ctx context.Context, in *JwtRequest, opts ...grpc.CallOption) (*JwtResponse, error)
}

type jwtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJwtServiceClient(cc grpc.ClientConnInterface) JwtServiceClient {
	return &jwtServiceClient{cc}
}

func (c *jwtServiceClient) JwtService(ctx context.Context, in *JwtRequest, opts ...grpc.CallOption) (*JwtResponse, error) {
	out := new(JwtResponse)
	err := c.cc.Invoke(ctx, "/middleware.JwtService/JwtService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JwtServiceServer is the server API for JwtService service.
type JwtServiceServer interface {
	JwtService(context.Context, *JwtRequest) (*JwtResponse, error)
}

// UnimplementedJwtServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJwtServiceServer struct {
}

func (*UnimplementedJwtServiceServer) JwtService(ctx context.Context, req *JwtRequest) (*JwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JwtService not implemented")
}

func RegisterJwtServiceServer(s *grpc.Server, srv JwtServiceServer) {
	s.RegisterService(&_JwtService_serviceDesc, srv)
}

func _JwtService_JwtService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).JwtService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/middleware.JwtService/JwtService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).JwtService(ctx, req.(*JwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JwtService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "middleware.JwtService",
	HandlerType: (*JwtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JwtService",
			Handler:    _JwtService_JwtService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "middleware/middleware.proto",
}
