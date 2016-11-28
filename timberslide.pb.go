// Code generated by protoc-gen-go.
// source: timberslide.proto
// DO NOT EDIT!

/*
Package ts is a generated protocol buffer package.

It is generated from these files:
	timberslide.proto

It has these top-level messages:
	Event
	EventReply
	Topic
	Register
	RegisterReply
	TopicsReq
	TopicsReply
*/
package ts

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

type Event struct {
	Topic   string `protobuf:"bytes,1,opt,name=Topic" json:"Topic,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Done    bool   `protobuf:"varint,3,opt,name=Done" json:"Done,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EventReply struct {
}

func (m *EventReply) Reset()                    { *m = EventReply{} }
func (m *EventReply) String() string            { return proto.CompactTextString(m) }
func (*EventReply) ProtoMessage()               {}
func (*EventReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Topic struct {
	ID       int64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Position int64  `protobuf:"varint,3,opt,name=Position" json:"Position,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Register struct {
	Topic string `protobuf:"bytes,1,opt,name=Topic" json:"Topic,omitempty"`
	User  string `protobuf:"bytes,2,opt,name=User" json:"User,omitempty"`
}

func (m *Register) Reset()                    { *m = Register{} }
func (m *Register) String() string            { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()               {}
func (*Register) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RegisterReply struct {
}

func (m *RegisterReply) Reset()                    { *m = RegisterReply{} }
func (m *RegisterReply) String() string            { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()               {}
func (*RegisterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type TopicsReq struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
}

func (m *TopicsReq) Reset()                    { *m = TopicsReq{} }
func (m *TopicsReq) String() string            { return proto.CompactTextString(m) }
func (*TopicsReq) ProtoMessage()               {}
func (*TopicsReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type TopicsReply struct {
	Topics []*Topic `protobuf:"bytes,1,rep,name=topics" json:"topics,omitempty"`
}

func (m *TopicsReply) Reset()                    { *m = TopicsReply{} }
func (m *TopicsReply) String() string            { return proto.CompactTextString(m) }
func (*TopicsReply) ProtoMessage()               {}
func (*TopicsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TopicsReply) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "ts.Event")
	proto.RegisterType((*EventReply)(nil), "ts.EventReply")
	proto.RegisterType((*Topic)(nil), "ts.Topic")
	proto.RegisterType((*Register)(nil), "ts.Register")
	proto.RegisterType((*RegisterReply)(nil), "ts.RegisterReply")
	proto.RegisterType((*TopicsReq)(nil), "ts.TopicsReq")
	proto.RegisterType((*TopicsReply)(nil), "ts.TopicsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ingest service

type IngestClient interface {
	StreamEvents(ctx context.Context, opts ...grpc.CallOption) (Ingest_StreamEventsClient, error)
}

type ingestClient struct {
	cc *grpc.ClientConn
}

func NewIngestClient(cc *grpc.ClientConn) IngestClient {
	return &ingestClient{cc}
}

func (c *ingestClient) StreamEvents(ctx context.Context, opts ...grpc.CallOption) (Ingest_StreamEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ingest_serviceDesc.Streams[0], c.cc, "/ts.Ingest/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingestStreamEventsClient{stream}
	return x, nil
}

type Ingest_StreamEventsClient interface {
	Send(*Event) error
	Recv() (*EventReply, error)
	grpc.ClientStream
}

type ingestStreamEventsClient struct {
	grpc.ClientStream
}

func (x *ingestStreamEventsClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingestStreamEventsClient) Recv() (*EventReply, error) {
	m := new(EventReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Ingest service

type IngestServer interface {
	StreamEvents(Ingest_StreamEventsServer) error
}

func RegisterIngestServer(s *grpc.Server, srv IngestServer) {
	s.RegisterService(&_Ingest_serviceDesc, srv)
}

func _Ingest_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngestServer).StreamEvents(&ingestStreamEventsServer{stream})
}

type Ingest_StreamEventsServer interface {
	Send(*EventReply) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type ingestStreamEventsServer struct {
	grpc.ServerStream
}

func (x *ingestStreamEventsServer) Send(m *EventReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingestStreamEventsServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Ingest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ts.Ingest",
	HandlerType: (*IngestServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _Ingest_StreamEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "timberslide.proto",
}

// Client API for Streamer service

type StreamerClient interface {
	GetStream(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Streamer_GetStreamClient, error)
}

type streamerClient struct {
	cc *grpc.ClientConn
}

func NewStreamerClient(cc *grpc.ClientConn) StreamerClient {
	return &streamerClient{cc}
}

func (c *streamerClient) GetStream(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Streamer_GetStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Streamer_serviceDesc.Streams[0], c.cc, "/ts.Streamer/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Streamer_GetStreamClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type streamerGetStreamClient struct {
	grpc.ClientStream
}

func (x *streamerGetStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Streamer service

type StreamerServer interface {
	GetStream(*Topic, Streamer_GetStreamServer) error
}

func RegisterStreamerServer(s *grpc.Server, srv StreamerServer) {
	s.RegisterService(&_Streamer_serviceDesc, srv)
}

func _Streamer_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamerServer).GetStream(m, &streamerGetStreamServer{stream})
}

type Streamer_GetStreamServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type streamerGetStreamServer struct {
	grpc.ServerStream
}

func (x *streamerGetStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Streamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ts.Streamer",
	HandlerType: (*StreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Streamer_GetStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "timberslide.proto",
}

// Client API for Topics service

type TopicsClient interface {
	RegisterTopic(ctx context.Context, in *Register, opts ...grpc.CallOption) (*RegisterReply, error)
	GetTopics(ctx context.Context, in *TopicsReq, opts ...grpc.CallOption) (*TopicsReply, error)
}

type topicsClient struct {
	cc *grpc.ClientConn
}

func NewTopicsClient(cc *grpc.ClientConn) TopicsClient {
	return &topicsClient{cc}
}

func (c *topicsClient) RegisterTopic(ctx context.Context, in *Register, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := grpc.Invoke(ctx, "/ts.Topics/registerTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicsClient) GetTopics(ctx context.Context, in *TopicsReq, opts ...grpc.CallOption) (*TopicsReply, error) {
	out := new(TopicsReply)
	err := grpc.Invoke(ctx, "/ts.Topics/getTopics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Topics service

type TopicsServer interface {
	RegisterTopic(context.Context, *Register) (*RegisterReply, error)
	GetTopics(context.Context, *TopicsReq) (*TopicsReply, error)
}

func RegisterTopicsServer(s *grpc.Server, srv TopicsServer) {
	s.RegisterService(&_Topics_serviceDesc, srv)
}

func _Topics_RegisterTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Register)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicsServer).RegisterTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ts.Topics/RegisterTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicsServer).RegisterTopic(ctx, req.(*Register))
	}
	return interceptor(ctx, in, info, handler)
}

func _Topics_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicsServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ts.Topics/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicsServer).GetTopics(ctx, req.(*TopicsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Topics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ts.Topics",
	HandlerType: (*TopicsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerTopic",
			Handler:    _Topics_RegisterTopic_Handler,
		},
		{
			MethodName: "getTopics",
			Handler:    _Topics_GetTopics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timberslide.proto",
}

func init() { proto.RegisterFile("timberslide.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0x69, 0x0b, 0x7d, 0xed, 0xe5, 0x5f, 0xb8, 0x79, 0x8b, 0xa6, 0x2b, 0x5e, 0x93, 0x17,
	0xbb, 0x11, 0xb1, 0xba, 0x71, 0x8f, 0x21, 0xc4, 0x68, 0xcc, 0xa8, 0x1f, 0x00, 0xf4, 0xa6, 0x69,
	0x84, 0xb6, 0xce, 0x4c, 0x4c, 0xf8, 0xf6, 0x66, 0xee, 0xb4, 0x85, 0x8d, 0xbb, 0x73, 0xe0, 0xdc,
	0xdf, 0x3d, 0x73, 0x53, 0x98, 0xe9, 0xe2, 0xb0, 0x23, 0xa9, 0xf6, 0xc5, 0x07, 0x2d, 0x6a, 0x59,
	0xe9, 0x0a, 0x5d, 0xad, 0x92, 0x07, 0x18, 0xdc, 0x7f, 0x53, 0xa9, 0xf1, 0x2f, 0x0c, 0x5e, 0xab,
	0xba, 0x78, 0x8f, 0x9c, 0xb9, 0x93, 0x86, 0xc2, 0x1a, 0x8c, 0xe0, 0xcf, 0x23, 0x29, 0xb5, 0xcd,
	0x29, 0x72, 0xf9, 0xf7, 0xd6, 0x22, 0x42, 0x7f, 0x55, 0x95, 0x14, 0x79, 0x73, 0x27, 0x0d, 0x04,
	0xeb, 0x64, 0x04, 0xc0, 0x30, 0x41, 0xf5, 0xfe, 0x98, 0xac, 0x1b, 0x22, 0x4e, 0xc0, 0xdd, 0xac,
	0x98, 0xeb, 0x09, 0x77, 0xb3, 0x32, 0xa3, 0x4f, 0xdb, 0x43, 0x4b, 0x64, 0x8d, 0x31, 0x04, 0xcf,
	0x95, 0x2a, 0x74, 0x51, 0x95, 0x8c, 0xf4, 0x44, 0xe7, 0x93, 0x5b, 0x08, 0x04, 0xe5, 0x85, 0xd2,
	0x24, 0x7f, 0xa9, 0x89, 0xd0, 0x7f, 0x53, 0x24, 0x5b, 0xa2, 0xd1, 0xc9, 0x14, 0xc6, 0xed, 0x94,
	0xed, 0x73, 0x01, 0x21, 0xa7, 0x95, 0xa0, 0x2f, 0xb3, 0xcf, 0xa4, 0x4a, 0xd3, 0xc3, 0xa2, 0x3a,
	0x9f, 0x2c, 0x61, 0xd8, 0x06, 0xeb, 0xfd, 0x11, 0xff, 0x81, 0xaf, 0xd9, 0x46, 0xce, 0xdc, 0x4b,
	0x87, 0x59, 0xb8, 0xd0, 0x6a, 0xc1, 0x01, 0xd1, 0xfc, 0x91, 0xdd, 0x81, 0xbf, 0x29, 0x73, 0x52,
	0x1a, 0xaf, 0x60, 0xf4, 0xa2, 0x25, 0x6d, 0x0f, 0x7c, 0x08, 0x85, 0x1c, 0x66, 0x1d, 0x4f, 0x3a,
	0x69, 0xfb, 0xf4, 0x52, 0x67, 0xe9, 0x64, 0xd7, 0x10, 0xd8, 0x01, 0x92, 0xf8, 0x1f, 0xc2, 0x35,
	0x69, 0x6b, 0xf1, 0xb4, 0x26, 0x3e, 0x41, 0x92, 0xde, 0xd2, 0xc9, 0x3e, 0xc1, 0xb7, 0xfd, 0x30,
	0x83, 0xb1, 0x6c, 0xde, 0x68, 0x0f, 0x31, 0x32, 0xc9, 0xf6, 0xd9, 0xf1, 0xec, 0xdc, 0x35, 0x4b,
	0xf1, 0x12, 0xc2, 0x9c, 0x74, 0x03, 0x18, 0x77, 0x4b, 0xcc, 0x55, 0xe2, 0xe9, 0xb9, 0xe5, 0xf8,
	0xce, 0xe7, 0x6f, 0xe5, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x24, 0xde, 0xef, 0x40, 0x02,
	0x00, 0x00,
}
