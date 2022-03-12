// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetAccountReq struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountReq) Reset()         { *m = GetAccountReq{} }
func (m *GetAccountReq) String() string { return proto.CompactTextString(m) }
func (*GetAccountReq) ProtoMessage()    {}
func (*GetAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_de28e89bd9adcdc1, []int{0}
}
func (m *GetAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountReq.Unmarshal(m, b)
}
func (m *GetAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountReq.Marshal(b, m, deterministic)
}
func (dst *GetAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountReq.Merge(dst, src)
}
func (m *GetAccountReq) XXX_Size() int {
	return xxx_messageInfo_GetAccountReq.Size(m)
}
func (m *GetAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountReq proto.InternalMessageInfo

func (m *GetAccountReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type GetAccountReply struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	UserInfo             *User    `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountReply) Reset()         { *m = GetAccountReply{} }
func (m *GetAccountReply) String() string { return proto.CompactTextString(m) }
func (*GetAccountReply) ProtoMessage()    {}
func (*GetAccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_de28e89bd9adcdc1, []int{1}
}
func (m *GetAccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountReply.Unmarshal(m, b)
}
func (m *GetAccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountReply.Marshal(b, m, deterministic)
}
func (dst *GetAccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountReply.Merge(dst, src)
}
func (m *GetAccountReply) XXX_Size() int {
	return xxx_messageInfo_GetAccountReply.Size(m)
}
func (m *GetAccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountReply proto.InternalMessageInfo

func (m *GetAccountReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *GetAccountReply) GetUserInfo() *User {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_de28e89bd9adcdc1, []int{2}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*GetAccountReq)(nil), "account.GetAccountReq")
	proto.RegisterType((*GetAccountReply)(nil), "account.GetAccountReply")
	proto.RegisterType((*User)(nil), "account.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	GetAccount(ctx context.Context, in *GetAccountReq, opts ...grpc.CallOption) (*GetAccountReply, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) GetAccount(ctx context.Context, in *GetAccountReq, opts ...grpc.CallOption) (*GetAccountReply, error) {
	out := new(GetAccountReply)
	err := c.cc.Invoke(ctx, "/account.AccountService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	GetAccount(context.Context, *GetAccountReq) (*GetAccountReply, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccount(ctx, req.(*GetAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _AccountService_GetAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_account_de28e89bd9adcdc1) }

var fileDescriptor_account_de28e89bd9adcdc1 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x14, 0xb9,
	0x78, 0xdd, 0x53, 0x4b, 0x1c, 0x21, 0xbc, 0xa0, 0xd4, 0x42, 0x21, 0x01, 0x2e, 0xe6, 0xd2, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x10, 0x53, 0x29, 0x86, 0x8b, 0x1f, 0x59, 0x49,
	0x41, 0x4e, 0xa5, 0x90, 0x2c, 0x17, 0x57, 0x6a, 0x51, 0x51, 0x7e, 0x51, 0x7c, 0x72, 0x7e, 0x4a,
	0x2a, 0x58, 0x2d, 0x6b, 0x10, 0x27, 0x58, 0xc4, 0x39, 0x3f, 0x25, 0x55, 0x48, 0x8b, 0x8b, 0xb3,
	0xb4, 0x38, 0xb5, 0x28, 0x3e, 0x33, 0x2f, 0x2d, 0x5f, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x88,
	0x57, 0x0f, 0xe6, 0x80, 0xd0, 0xe2, 0xd4, 0xa2, 0x20, 0x0e, 0x90, 0xbc, 0x67, 0x5e, 0x5a, 0xbe,
	0x92, 0x0e, 0x17, 0x0b, 0x48, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x17, 0x62, 0x18, 0x67,
	0x10, 0x98, 0x0d, 0x72, 0x4b, 0x62, 0x7a, 0x2a, 0xd8, 0x04, 0xd6, 0x20, 0x10, 0xd3, 0x28, 0x88,
	0x8b, 0x0f, 0xea, 0x90, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x07, 0x2e, 0x2e, 0x84,
	0xeb, 0x84, 0xc4, 0xe0, 0xd6, 0xa0, 0xf8, 0x4a, 0x4a, 0x02, 0xab, 0x78, 0x41, 0x4e, 0xa5, 0x12,
	0x43, 0x12, 0x1b, 0x38, 0x48, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x78, 0x63, 0xd3,
	0x23, 0x01, 0x00, 0x00,
}