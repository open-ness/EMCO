// Code generated by protoc-gen-go. DO NOT EDIT.
// source: readynotify.proto

/*
Package readynotify is a generated protocol buffer package.

It is generated from these files:
	readynotify.proto

It has these top-level messages:
	Alert
	Topic
	Notification
*/
package readynotify

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

type Alert struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Alert) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Topic struct {
	ClientName string `protobuf:"bytes,1,opt,name=clientName" json:"clientName,omitempty"`
	AppContext string `protobuf:"bytes,2,opt,name=appContext" json:"appContext,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Topic) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Topic) GetAppContext() string {
	if m != nil {
		return m.AppContext
	}
	return ""
}

type Notification struct {
	ClientName string `protobuf:"bytes,1,opt,name=clientName" json:"clientName,omitempty"`
	ServerName string `protobuf:"bytes,2,opt,name=serverName" json:"serverName,omitempty"`
	AppContext string `protobuf:"bytes,3,opt,name=appContext" json:"appContext,omitempty"`
}

func (m *Notification) Reset()                    { *m = Notification{} }
func (m *Notification) String() string            { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()               {}
func (*Notification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Notification) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Notification) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *Notification) GetAppContext() string {
	if m != nil {
		return m.AppContext
	}
	return ""
}

func init() {
	proto.RegisterType((*Alert)(nil), "Alert")
	proto.RegisterType((*Topic)(nil), "Topic")
	proto.RegisterType((*Notification)(nil), "Notification")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ReadyNotify service

type ReadyNotifyClient interface {
	Alert(ctx context.Context, in *Topic, opts ...grpc.CallOption) (ReadyNotify_AlertClient, error)
}

type readyNotifyClient struct {
	cc *grpc.ClientConn
}

func NewReadyNotifyClient(cc *grpc.ClientConn) ReadyNotifyClient {
	return &readyNotifyClient{cc}
}

func (c *readyNotifyClient) Alert(ctx context.Context, in *Topic, opts ...grpc.CallOption) (ReadyNotify_AlertClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ReadyNotify_serviceDesc.Streams[0], c.cc, "/readyNotify/Alert", opts...)
	if err != nil {
		return nil, err
	}
	x := &readyNotifyAlertClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReadyNotify_AlertClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type readyNotifyAlertClient struct {
	grpc.ClientStream
}

func (x *readyNotifyAlertClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ReadyNotify service

type ReadyNotifyServer interface {
	Alert(*Topic, ReadyNotify_AlertServer) error
}

func RegisterReadyNotifyServer(s *grpc.Server, srv ReadyNotifyServer) {
	s.RegisterService(&_ReadyNotify_serviceDesc, srv)
}

func _ReadyNotify_Alert_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReadyNotifyServer).Alert(m, &readyNotifyAlertServer{stream})
}

type ReadyNotify_AlertServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type readyNotifyAlertServer struct {
	grpc.ServerStream
}

func (x *readyNotifyAlertServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

var _ReadyNotify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "readyNotify",
	HandlerType: (*ReadyNotifyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Alert",
			Handler:       _ReadyNotify_Alert_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "readynotify.proto",
}

func init() { proto.RegisterFile("readynotify.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4a, 0x4d, 0x4c,
	0xa9, 0xcc, 0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe4,
	0x62, 0x75, 0xcc, 0x49, 0x2d, 0x2a, 0x11, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0xdc, 0xb9, 0x58, 0x43, 0xf2,
	0x0b, 0x32, 0x93, 0x85, 0xe4, 0xb8, 0xb8, 0x92, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0xfc, 0x12, 0x73,
	0x61, 0xaa, 0x90, 0x44, 0x40, 0xf2, 0x89, 0x05, 0x05, 0xce, 0xf9, 0x79, 0x25, 0xa9, 0x15, 0x25,
	0x12, 0x4c, 0x10, 0x79, 0x84, 0x88, 0x52, 0x1e, 0x17, 0x8f, 0x1f, 0xc8, 0xee, 0xcc, 0xe4, 0xc4,
	0x92, 0xcc, 0xfc, 0x3c, 0x62, 0xcc, 0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x02, 0xcb, 0x43, 0xcd,
	0x43, 0x88, 0xa0, 0xd9, 0xc7, 0x8c, 0x6e, 0x9f, 0x91, 0x3e, 0x17, 0x37, 0xd8, 0xc3, 0x60, 0x4b,
	0x2b, 0x85, 0x14, 0x60, 0x5e, 0x65, 0xd3, 0x03, 0xfb, 0x47, 0x8a, 0x57, 0x0f, 0xd9, 0x39, 0x06,
	0x8c, 0x49, 0x6c, 0xe0, 0x30, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xc1, 0xd0, 0x25,
	0x28, 0x01, 0x00, 0x00,
}
