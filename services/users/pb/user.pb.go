// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package users

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Phone                string   `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type FindAllRequest struct {
	Page                 string   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              string   `protobuf:"bytes,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllRequest) Reset()         { *m = FindAllRequest{} }
func (m *FindAllRequest) String() string { return proto.CompactTextString(m) }
func (*FindAllRequest) ProtoMessage()    {}
func (*FindAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *FindAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllRequest.Unmarshal(m, b)
}
func (m *FindAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllRequest.Marshal(b, m, deterministic)
}
func (m *FindAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllRequest.Merge(m, src)
}
func (m *FindAllRequest) XXX_Size() int {
	return xxx_messageInfo_FindAllRequest.Size(m)
}
func (m *FindAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllRequest proto.InternalMessageInfo

func (m *FindAllRequest) GetPage() string {
	if m != nil {
		return m.Page
	}
	return ""
}

func (m *FindAllRequest) GetPerPage() string {
	if m != nil {
		return m.PerPage
	}
	return ""
}

type FindAllResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 []*User  `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllResponse) Reset()         { *m = FindAllResponse{} }
func (m *FindAllResponse) String() string { return proto.CompactTextString(m) }
func (*FindAllResponse) ProtoMessage()    {}
func (*FindAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *FindAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllResponse.Unmarshal(m, b)
}
func (m *FindAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllResponse.Marshal(b, m, deterministic)
}
func (m *FindAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllResponse.Merge(m, src)
}
func (m *FindAllResponse) XXX_Size() int {
	return xxx_messageInfo_FindAllResponse.Size(m)
}
func (m *FindAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllResponse proto.InternalMessageInfo

func (m *FindAllResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *FindAllResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FindAllResponse) GetData() []*User {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "users.User")
	proto.RegisterType((*FindAllRequest)(nil), "users.FindAllRequest")
	proto.RegisterType((*FindAllResponse)(nil), "users.FindAllResponse")
	proto.RegisterType((*CreateRequest)(nil), "users.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "users.CreateResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x31, 0x4f, 0x32, 0x41,
	0x10, 0x0d, 0x70, 0x07, 0x7c, 0x43, 0x3e, 0x4c, 0x26, 0x40, 0x56, 0x1a, 0xc8, 0x55, 0x54, 0xc4,
	0x60, 0xa3, 0x95, 0x51, 0x13, 0x6b, 0x73, 0x89, 0xb5, 0x59, 0xb3, 0x13, 0xbc, 0x04, 0xee, 0xd6,
	0x9d, 0x25, 0x36, 0xfe, 0x0c, 0x7f, 0xb0, 0xd9, 0xd9, 0x3d, 0x22, 0x94, 0x76, 0xf3, 0xde, 0xbc,
	0x79, 0xf3, 0x76, 0xee, 0x00, 0x0e, 0x4c, 0x6e, 0x6d, 0x5d, 0xe3, 0x1b, 0xcc, 0x43, 0xcd, 0xc5,
	0x77, 0x07, 0xb2, 0x17, 0x26, 0x87, 0x63, 0xe8, 0x56, 0x46, 0x75, 0x96, 0x9d, 0xd5, 0xbf, 0xb2,
	0x5b, 0x19, 0x9c, 0xc3, 0x30, 0x28, 0x6a, 0xbd, 0x27, 0xd5, 0x15, 0xf6, 0x88, 0x43, 0xcf, 0x6a,
	0xe6, 0xcf, 0xc6, 0x19, 0xd5, 0x8b, 0xbd, 0x16, 0xe3, 0x04, 0x72, 0xda, 0xeb, 0x6a, 0xa7, 0x32,
	0x69, 0x44, 0x80, 0x0a, 0x06, 0xda, 0x18, 0x47, 0xcc, 0x2a, 0x17, 0xbe, 0x85, 0x41, 0x6f, 0xdf,
	0x9b, 0x9a, 0x54, 0x3f, 0xea, 0x05, 0x14, 0x77, 0x30, 0x7e, 0xaa, 0x6a, 0x73, 0xbf, 0xdb, 0x95,
	0xf4, 0x71, 0x20, 0xf6, 0x88, 0x90, 0x59, 0xbd, 0xa5, 0x94, 0x50, 0x6a, 0xbc, 0x84, 0xa1, 0x25,
	0xf7, 0x2a, 0x7c, 0xcc, 0x38, 0xb0, 0xe4, 0x9e, 0xf5, 0x96, 0x0a, 0x03, 0x17, 0x47, 0x03, 0xb6,
	0x4d, 0xcd, 0x84, 0x33, 0xe8, 0xb3, 0xd7, 0xfe, 0xc0, 0xe2, 0x31, 0x2c, 0x13, 0x0a, 0xd9, 0xf6,
	0xc4, 0xfc, 0xcb, 0x24, 0x41, 0x5c, 0x40, 0x66, 0xb4, 0xd7, 0xaa, 0xb7, 0xec, 0xad, 0x46, 0x9b,
	0xd1, 0x5a, 0x4e, 0xb6, 0x0e, 0xe7, 0x2a, 0xa5, 0x51, 0x5c, 0xc1, 0xff, 0x47, 0x47, 0xda, 0x53,
	0x9b, 0x72, 0x01, 0x59, 0x10, 0xc9, 0x86, 0xf3, 0x89, 0x50, 0x17, 0x0f, 0x30, 0x6e, 0x27, 0xfe,
	0x1a, 0x6b, 0xf3, 0x05, 0x79, 0x70, 0x64, 0xbc, 0x81, 0x41, 0x7a, 0x24, 0x4e, 0xd3, 0xaa, 0xd3,
	0xab, 0xcd, 0x67, 0xe7, 0x74, 0x5a, 0x7a, 0x0b, 0x10, 0x63, 0xc8, 0xb7, 0x9f, 0x24, 0xd5, 0xc9,
	0x5b, 0xe6, 0xd3, 0x33, 0x36, 0x8e, 0xbe, 0xf5, 0xe5, 0xff, 0xb9, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0xbf, 0x7e, 0xc8, 0x4d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	FindAll(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error)
	// rpc FineOne(GetOneRequest) returns (GetOneResponse);
	CreateUser(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) FindAll(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error) {
	out := new(FindAllResponse)
	err := c.cc.Invoke(ctx, "/users.Users/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CreateUser(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/users.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	FindAll(context.Context, *FindAllRequest) (*FindAllResponse, error)
	// rpc FineOne(GetOneRequest) returns (GetOneResponse);
	CreateUser(context.Context, *CreateRequest) (*CreateResponse, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) FindAll(ctx context.Context, req *FindAllRequest) (*FindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (*UnimplementedUsersServer) CreateUser(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).FindAll(ctx, req.(*FindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _Users_FindAll_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
