// Code generated by protoc-gen-go.
// source: todo/todo.proto
// DO NOT EDIT!

/*
Package todo is a generated protocol buffer package.

It is generated from these files:
	todo/todo.proto

It has these top-level messages:
	TodoItem
	TodoList
*/
package todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type TodoItem struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TodoItem) Reset()                    { *m = TodoItem{} }
func (m *TodoItem) String() string            { return proto.CompactTextString(m) }
func (*TodoItem) ProtoMessage()               {}
func (*TodoItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TodoItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TodoList struct {
	Name  string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Todos []*TodoItem `protobuf:"bytes,2,rep,name=todos" json:"todos,omitempty"`
}

func (m *TodoList) Reset()                    { *m = TodoList{} }
func (m *TodoList) String() string            { return proto.CompactTextString(m) }
func (*TodoList) ProtoMessage()               {}
func (*TodoList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TodoList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TodoList) GetTodos() []*TodoItem {
	if m != nil {
		return m.Todos
	}
	return nil
}

func init() {
	proto.RegisterType((*TodoItem)(nil), "TodoItem")
	proto.RegisterType((*TodoList)(nil), "TodoList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DoSomething service

type DoSomethingClient interface {
	// Sends a greeting
	AddTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	ListTodos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*TodoList, error)
}

type doSomethingClient struct {
	cc *grpc.ClientConn
}

func NewDoSomethingClient(cc *grpc.ClientConn) DoSomethingClient {
	return &doSomethingClient{cc}
}

func (c *doSomethingClient) AddTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/DoSomething/AddTodoItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doSomethingClient) ListTodos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*TodoList, error) {
	out := new(TodoList)
	err := grpc.Invoke(ctx, "/DoSomething/ListTodos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DoSomething service

type DoSomethingServer interface {
	// Sends a greeting
	AddTodoItem(context.Context, *TodoItem) (*google_protobuf.Empty, error)
	ListTodos(context.Context, *google_protobuf.Empty) (*TodoList, error)
}

func RegisterDoSomethingServer(s *grpc.Server, srv DoSomethingServer) {
	s.RegisterService(&_DoSomething_serviceDesc, srv)
}

func _DoSomething_AddTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoSomethingServer).AddTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DoSomething/AddTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoSomethingServer).AddTodoItem(ctx, req.(*TodoItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoSomething_ListTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoSomethingServer).ListTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DoSomething/ListTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoSomethingServer).ListTodos(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DoSomething_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DoSomething",
	HandlerType: (*DoSomethingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTodoItem",
			Handler:    _DoSomething_AddTodoItem_Handler,
		},
		{
			MethodName: "ListTodos",
			Handler:    _DoSomething_ListTodos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo/todo.proto",
}

func init() { proto.RegisterFile("todo/todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc9, 0x4f, 0xc9,
	0xd7, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x52, 0xd2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39,
	0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x6a, 0x6e, 0x41, 0x49, 0x25, 0x44, 0x52, 0x49,
	0x8e, 0x8b, 0x23, 0x24, 0x3f, 0x25, 0xdf, 0xb3, 0x24, 0x35, 0x57, 0x48, 0x88, 0x8b, 0x25, 0x2f,
	0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0xb2, 0x87, 0xc8, 0xfb,
	0x64, 0x16, 0x97, 0x60, 0x93, 0x17, 0x92, 0xe7, 0x62, 0x05, 0x59, 0x55, 0x2c, 0xc1, 0xa4, 0xc0,
	0xac, 0xc1, 0x6d, 0xc4, 0xa9, 0x07, 0x33, 0x2d, 0x08, 0x22, 0x6e, 0x54, 0xcc, 0xc5, 0xed, 0x92,
	0x1f, 0x9c, 0x9f, 0x9b, 0x5a, 0x92, 0x91, 0x99, 0x97, 0x2e, 0x64, 0xc4, 0xc5, 0xed, 0x98, 0x92,
	0x02, 0xb7, 0x12, 0xa1, 0x5e, 0x4a, 0x4c, 0x0f, 0xe2, 0x4e, 0x3d, 0x98, 0x3b, 0xf5, 0x5c, 0x41,
	0xee, 0x54, 0x62, 0x10, 0x32, 0xe0, 0xe2, 0x04, 0xd9, 0x0f, 0x52, 0x59, 0x2c, 0x84, 0x43, 0x99,
	0x14, 0xc4, 0x24, 0x90, 0x3a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xa4, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x4c, 0x32, 0x4f, 0x8e, 0x0c, 0x01, 0x00, 0x00,
}
