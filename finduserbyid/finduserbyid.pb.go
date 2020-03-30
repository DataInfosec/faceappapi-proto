// Code generated by protoc-gen-go. DO NOT EDIT.
// source: finduserbyid/finduserbyid.proto

package finduserbyid

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

type UserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45f7443560d1cf8f, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CompanyRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompanyRequest) Reset()         { *m = CompanyRequest{} }
func (m *CompanyRequest) String() string { return proto.CompactTextString(m) }
func (*CompanyRequest) ProtoMessage()    {}
func (*CompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45f7443560d1cf8f, []int{1}
}

func (m *CompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompanyRequest.Unmarshal(m, b)
}
func (m *CompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompanyRequest.Marshal(b, m, deterministic)
}
func (m *CompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompanyRequest.Merge(m, src)
}
func (m *CompanyRequest) XXX_Size() int {
	return xxx_messageInfo_CompanyRequest.Size(m)
}
func (m *CompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompanyRequest proto.InternalMessageInfo

func (m *CompanyRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserResponse struct {
	Firstname            string   `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	Company              string   `protobuf:"bytes,6,opt,name=company,proto3" json:"company,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45f7443560d1cf8f, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *UserResponse) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserResponse) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

type CompanyResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	City                 string   `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompanyResponse) Reset()         { *m = CompanyResponse{} }
func (m *CompanyResponse) String() string { return proto.CompactTextString(m) }
func (*CompanyResponse) ProtoMessage()    {}
func (*CompanyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45f7443560d1cf8f, []int{3}
}

func (m *CompanyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompanyResponse.Unmarshal(m, b)
}
func (m *CompanyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompanyResponse.Marshal(b, m, deterministic)
}
func (m *CompanyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompanyResponse.Merge(m, src)
}
func (m *CompanyResponse) XXX_Size() int {
	return xxx_messageInfo_CompanyResponse.Size(m)
}
func (m *CompanyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompanyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompanyResponse proto.InternalMessageInfo

func (m *CompanyResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CompanyResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CompanyResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CompanyResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CompanyResponse) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CompanyResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CompanyResponse) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "finduserbyid.UserRequest")
	proto.RegisterType((*CompanyRequest)(nil), "finduserbyid.CompanyRequest")
	proto.RegisterType((*UserResponse)(nil), "finduserbyid.UserResponse")
	proto.RegisterType((*CompanyResponse)(nil), "finduserbyid.CompanyResponse")
}

func init() { proto.RegisterFile("finduserbyid/finduserbyid.proto", fileDescriptor_45f7443560d1cf8f) }

var fileDescriptor_45f7443560d1cf8f = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x5d, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x4d, 0x4c, 0x1b, 0x7b, 0x2d, 0x11, 0x06, 0x1f, 0xc6, 0xd0, 0x62, 0xc9, 0x93, 0x4f,
	0x15, 0x74, 0x09, 0x82, 0xaf, 0x42, 0x8b, 0x0b, 0x98, 0x66, 0x6e, 0x71, 0xa0, 0xf9, 0x71, 0x66,
	0x22, 0x64, 0x2f, 0x6e, 0xa0, 0xbb, 0x94, 0xcc, 0x4f, 0x4c, 0xc4, 0xf6, 0x6d, 0xce, 0x3d, 0x97,
	0x93, 0xef, 0x72, 0x02, 0xf7, 0x7b, 0x51, 0xf2, 0x46, 0xa1, 0xdc, 0xb5, 0x82, 0x3f, 0x0e, 0xc5,
	0xba, 0x96, 0x95, 0xae, 0xc8, 0x7c, 0x38, 0xcb, 0x96, 0x70, 0xfd, 0xae, 0x50, 0x6e, 0xf0, 0xb3,
	0x41, 0xa5, 0x49, 0x02, 0xa1, 0xe0, 0x34, 0x58, 0x05, 0x0f, 0xb3, 0x4d, 0x28, 0x78, 0xb6, 0x82,
	0xe4, 0xa5, 0x2a, 0x6a, 0x56, 0xb6, 0xa7, 0x36, 0xbe, 0x03, 0x98, 0xdb, 0x04, 0x55, 0x57, 0xa5,
	0x42, 0xb2, 0x80, 0xd9, 0x5e, 0x48, 0xa5, 0x4b, 0x56, 0xa0, 0xdb, 0xfb, 0x1d, 0x90, 0x14, 0xae,
	0x0e, 0xcc, 0x99, 0xa1, 0x31, 0x7b, 0x4d, 0x6e, 0x61, 0x82, 0x05, 0x13, 0x07, 0x7a, 0x69, 0x0c,
	0x2b, 0x08, 0x81, 0x48, 0xb7, 0x35, 0xd2, 0xc8, 0x0c, 0xcd, 0xdb, 0x41, 0x4c, 0x3c, 0x04, 0xa1,
	0x10, 0xe7, 0x16, 0x93, 0x4e, 0xcd, 0xd0, 0xcb, 0xec, 0x18, 0xc0, 0x4d, 0x7f, 0x81, 0x23, 0xfc,
	0x73, 0x42, 0xf7, 0x85, 0x01, 0x4f, 0x74, 0x86, 0x85, 0x42, 0xcc, 0x38, 0x97, 0xa8, 0x94, 0xc3,
	0xf1, 0xb2, 0xcb, 0xc8, 0x85, 0x6e, 0x1d, 0x93, 0x79, 0x77, 0x19, 0x4a, 0x33, 0x8d, 0x8e, 0xc9,
	0x0a, 0xcb, 0xda, 0x94, 0x5a, 0xb6, 0x34, 0xf6, 0xac, 0x46, 0x3e, 0x1d, 0x03, 0x48, 0xb6, 0x1f,
	0x4c, 0x22, 0xdf, 0xa2, 0xfc, 0x12, 0x39, 0x2a, 0xf2, 0x6a, 0xeb, 0x71, 0x9a, 0xdc, 0xad, 0x47,
	0x85, 0x0e, 0x9a, 0x4b, 0xd3, 0xff, 0x2c, 0x7b, 0x70, 0x76, 0x41, 0xde, 0xfa, 0x1e, 0x7d, 0xd4,
	0x62, 0xbc, 0x3f, 0x6e, 0x39, 0x5d, 0x9e, 0x70, 0x7d, 0xe0, 0x6e, 0x6a, 0x7e, 0xa6, 0xe7, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xde, 0x47, 0x61, 0x6f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SharedServicesClient is the client API for SharedServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SharedServicesClient interface {
	UserService(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	CompanyService(ctx context.Context, in *CompanyRequest, opts ...grpc.CallOption) (*CompanyResponse, error)
}

type sharedServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedServicesClient(cc grpc.ClientConnInterface) SharedServicesClient {
	return &sharedServicesClient{cc}
}

func (c *sharedServicesClient) UserService(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/finduserbyid.SharedServices/UserService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedServicesClient) CompanyService(ctx context.Context, in *CompanyRequest, opts ...grpc.CallOption) (*CompanyResponse, error) {
	out := new(CompanyResponse)
	err := c.cc.Invoke(ctx, "/finduserbyid.SharedServices/CompanyService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedServicesServer is the server API for SharedServices service.
type SharedServicesServer interface {
	UserService(context.Context, *UserRequest) (*UserResponse, error)
	CompanyService(context.Context, *CompanyRequest) (*CompanyResponse, error)
}

// UnimplementedSharedServicesServer can be embedded to have forward compatible implementations.
type UnimplementedSharedServicesServer struct {
}

func (*UnimplementedSharedServicesServer) UserService(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserService not implemented")
}
func (*UnimplementedSharedServicesServer) CompanyService(ctx context.Context, req *CompanyRequest) (*CompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompanyService not implemented")
}

func RegisterSharedServicesServer(s *grpc.Server, srv SharedServicesServer) {
	s.RegisterService(&_SharedServices_serviceDesc, srv)
}

func _SharedServices_UserService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedServicesServer).UserService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finduserbyid.SharedServices/UserService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedServicesServer).UserService(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedServices_CompanyService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedServicesServer).CompanyService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finduserbyid.SharedServices/CompanyService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedServicesServer).CompanyService(ctx, req.(*CompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SharedServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "finduserbyid.SharedServices",
	HandlerType: (*SharedServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserService",
			Handler:    _SharedServices_UserService_Handler,
		},
		{
			MethodName: "CompanyService",
			Handler:    _SharedServices_CompanyService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finduserbyid/finduserbyid.proto",
}
