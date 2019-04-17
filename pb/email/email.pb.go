// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email/email.proto

package email

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/lyft/protoc-gen-validate/validate"
	grpc "google.golang.org/grpc"
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

type SendEmailRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Subject              string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Emails               []string `protobuf:"bytes,5,rep,name=emails,proto3" json:"emails,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEmailRequest) Reset()         { *m = SendEmailRequest{} }
func (m *SendEmailRequest) String() string { return proto.CompactTextString(m) }
func (*SendEmailRequest) ProtoMessage()    {}
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74366e5b77a2734d, []int{0}
}

func (m *SendEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEmailRequest.Unmarshal(m, b)
}
func (m *SendEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEmailRequest.Marshal(b, m, deterministic)
}
func (m *SendEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEmailRequest.Merge(m, src)
}
func (m *SendEmailRequest) XXX_Size() int {
	return xxx_messageInfo_SendEmailRequest.Size(m)
}
func (m *SendEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendEmailRequest proto.InternalMessageInfo

func (m *SendEmailRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SendEmailRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SendEmailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendEmailRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendEmailRequest) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

func init() {
	proto.RegisterType((*SendEmailRequest)(nil), "go.emailer.com.SendEmailRequest")
}

func init() { proto.RegisterFile("email/email.proto", fileDescriptor_74366e5b77a2734d) }

var fileDescriptor_74366e5b77a2734d = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xb4, 0x40,
	0x10, 0xc7, 0xb3, 0xc0, 0x71, 0x1f, 0x93, 0x2f, 0x97, 0x73, 0x0b, 0x6f, 0xc5, 0x86, 0x50, 0x5d,
	0xb5, 0x24, 0xfa, 0x06, 0x24, 0x14, 0xb6, 0x5c, 0x6c, 0x6c, 0x0c, 0xb0, 0x23, 0x59, 0x03, 0x2e,
	0x2e, 0xcb, 0x25, 0xf7, 0x1a, 0x56, 0x3e, 0x8b, 0x95, 0xaf, 0xe3, 0x5b, 0x98, 0x5d, 0xc4, 0xe4,
	0x6c, 0x26, 0x33, 0xf3, 0x9f, 0x99, 0xfc, 0xe7, 0x07, 0x17, 0xd8, 0x57, 0xb2, 0xcb, 0x5c, 0xe4,
	0x83, 0x56, 0x46, 0xd1, 0x4d, 0xab, 0xb8, 0xab, 0x51, 0xf3, 0x46, 0xf5, 0xf1, 0x75, 0xab, 0x54,
	0xdb, 0x61, 0xe6, 0xd4, 0x7a, 0x7a, 0xca, 0xb0, 0x1f, 0xcc, 0x69, 0x1e, 0x8e, 0x77, 0xc7, 0xaa,
	0x93, 0xa2, 0x32, 0x98, 0x2d, 0xc9, 0x2c, 0xa4, 0xef, 0x04, 0xb6, 0x07, 0x7c, 0x11, 0x85, 0xbd,
	0x54, 0xe2, 0xeb, 0x84, 0xa3, 0xa1, 0x1b, 0xf0, 0xa4, 0x60, 0x24, 0x21, 0xfb, 0xa8, 0xf4, 0xa4,
	0xa0, 0x29, 0xac, 0xa7, 0x11, 0xf5, 0xa3, 0x14, 0xcc, 0xb3, 0xcd, 0x3c, 0xfa, 0xf8, 0xfa, 0xf4,
	0x03, 0xed, 0x6d, 0x49, 0x19, 0x5a, 0xe5, 0x4e, 0x50, 0x06, 0xeb, 0x71, 0xaa, 0x9f, 0xb1, 0x31,
	0xcc, 0x77, 0x8b, 0x4b, 0x49, 0x29, 0x04, 0xb5, 0x12, 0x27, 0x16, 0xb8, 0xb6, 0xcb, 0x69, 0x0a,
	0xa1, 0xf3, 0x3e, 0xb2, 0x55, 0xe2, 0xef, 0xa3, 0x1c, 0xec, 0xc1, 0xd5, 0x1b, 0xf1, 0xfe, 0x91,
	0xf2, 0x47, 0xb9, 0xb9, 0x87, 0xff, 0xce, 0xd5, 0x01, 0xf5, 0x51, 0x36, 0x48, 0x0b, 0x88, 0x7e,
	0x9d, 0xd2, 0x84, 0x9f, 0xbf, 0xcf, 0xff, 0x3e, 0x11, 0x5f, 0xf2, 0x19, 0x08, 0x5f, 0x80, 0xf0,
	0xc2, 0x02, 0xc9, 0xaf, 0x1e, 0x76, 0xe7, 0xab, 0xd9, 0x50, 0xcf, 0x60, 0xeb, 0xd0, 0x8d, 0xde,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xaf, 0x0f, 0x01, 0x6e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailServiceClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type emailServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmailServiceClient(cc *grpc.ClientConn) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/go.emailer.com.EmailService/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
type EmailServiceServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*empty.Empty, error)
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.emailer.com.EmailService/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.emailer.com.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _EmailService_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email/email.proto",
}
