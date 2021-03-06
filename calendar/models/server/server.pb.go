// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/server/server.proto

package calendar_server

import (
	context "context"
	fmt "fmt"
	math "math"
	event "otus_home/calendar/models/event"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddEventRequest struct {
	Event                *event.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AddEventRequest) Reset()         { *m = AddEventRequest{} }
func (m *AddEventRequest) String() string { return proto.CompactTextString(m) }
func (*AddEventRequest) ProtoMessage()    {}
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4346842d3ce3a1, []int{0}
}

func (m *AddEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEventRequest.Unmarshal(m, b)
}
func (m *AddEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEventRequest.Marshal(b, m, deterministic)
}
func (m *AddEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEventRequest.Merge(m, src)
}
func (m *AddEventRequest) XXX_Size() int {
	return xxx_messageInfo_AddEventRequest.Size(m)
}
func (m *AddEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddEventRequest proto.InternalMessageInfo

func (m *AddEventRequest) GetEvent() *event.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type AddEventResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddEventResponse) Reset()         { *m = AddEventResponse{} }
func (m *AddEventResponse) String() string { return proto.CompactTextString(m) }
func (*AddEventResponse) ProtoMessage()    {}
func (*AddEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4346842d3ce3a1, []int{1}
}

func (m *AddEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEventResponse.Unmarshal(m, b)
}
func (m *AddEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEventResponse.Marshal(b, m, deterministic)
}
func (m *AddEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEventResponse.Merge(m, src)
}
func (m *AddEventResponse) XXX_Size() int {
	return xxx_messageInfo_AddEventResponse.Size(m)
}
func (m *AddEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddEventResponse proto.InternalMessageInfo

func (m *AddEventResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type UpdateEventRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventRequest) Reset()         { *m = UpdateEventRequest{} }
func (m *UpdateEventRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEventRequest) ProtoMessage()    {}
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4346842d3ce3a1, []int{2}
}

func (m *UpdateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventRequest.Unmarshal(m, b)
}
func (m *UpdateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventRequest.Merge(m, src)
}
func (m *UpdateEventRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEventRequest.Size(m)
}
func (m *UpdateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventRequest proto.InternalMessageInfo

func (m *UpdateEventRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateEventResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventResponse) Reset()         { *m = UpdateEventResponse{} }
func (m *UpdateEventResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateEventResponse) ProtoMessage()    {}
func (*UpdateEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4346842d3ce3a1, []int{3}
}

func (m *UpdateEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventResponse.Unmarshal(m, b)
}
func (m *UpdateEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventResponse.Marshal(b, m, deterministic)
}
func (m *UpdateEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventResponse.Merge(m, src)
}
func (m *UpdateEventResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateEventResponse.Size(m)
}
func (m *UpdateEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventResponse proto.InternalMessageInfo

func (m *UpdateEventResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type DeleteEventRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventRequest) Reset()         { *m = DeleteEventRequest{} }
func (m *DeleteEventRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEventRequest) ProtoMessage()    {}
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4346842d3ce3a1, []int{4}
}

func (m *DeleteEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventRequest.Unmarshal(m, b)
}
func (m *DeleteEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventRequest.Merge(m, src)
}
func (m *DeleteEventRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEventRequest.Size(m)
}
func (m *DeleteEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventRequest proto.InternalMessageInfo

func (m *DeleteEventRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteEventResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventResponse) Reset()         { *m = DeleteEventResponse{} }
func (m *DeleteEventResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEventResponse) ProtoMessage()    {}
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4346842d3ce3a1, []int{5}
}

func (m *DeleteEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventResponse.Unmarshal(m, b)
}
func (m *DeleteEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventResponse.Marshal(b, m, deterministic)
}
func (m *DeleteEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventResponse.Merge(m, src)
}
func (m *DeleteEventResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEventResponse.Size(m)
}
func (m *DeleteEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventResponse proto.InternalMessageInfo

func (m *DeleteEventResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*AddEventRequest)(nil), "calendar.AddEventRequest")
	proto.RegisterType((*AddEventResponse)(nil), "calendar.AddEventResponse")
	proto.RegisterType((*UpdateEventRequest)(nil), "calendar.UpdateEventRequest")
	proto.RegisterType((*UpdateEventResponse)(nil), "calendar.UpdateEventResponse")
	proto.RegisterType((*DeleteEventRequest)(nil), "calendar.DeleteEventRequest")
	proto.RegisterType((*DeleteEventResponse)(nil), "calendar.DeleteEventResponse")
}

func init() { proto.RegisterFile("models/server/server.proto", fileDescriptor_0b4346842d3ce3a1) }

var fileDescriptor_0b4346842d3ce3a1 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xcd, 0x4f, 0x49,
	0xcd, 0x29, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x82, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x1c, 0xc9, 0x89, 0x39, 0xa9, 0x79, 0x29, 0x89, 0x45, 0x52, 0x12, 0x50, 0x55, 0xa9,
	0x65, 0xa9, 0x79, 0x25, 0x10, 0x12, 0xa2, 0x46, 0xc9, 0x94, 0x8b, 0xdf, 0x31, 0x25, 0xc5, 0x15,
	0x24, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xc4, 0xc5, 0x0a, 0x56, 0x21, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0xa3, 0x07, 0x51, 0x0f, 0x51, 0x03, 0x91, 0x52, 0xd2, 0xe2,
	0x12, 0x40, 0x68, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe3, 0x62, 0x2b, 0x2e, 0x49,
	0x2c, 0x29, 0x2d, 0x06, 0x6b, 0xe4, 0x0c, 0x82, 0xf2, 0x94, 0x54, 0xb8, 0x84, 0x42, 0x0b, 0x52,
	0x12, 0x4b, 0x52, 0x51, 0x6c, 0xe1, 0xe3, 0x62, 0xca, 0x4c, 0x01, 0xab, 0x64, 0x0d, 0x62, 0xca,
	0x4c, 0x51, 0xd2, 0xe5, 0x12, 0x46, 0x51, 0x45, 0xd8, 0x50, 0x97, 0xd4, 0x9c, 0x54, 0xc2, 0x86,
	0xa2, 0xa8, 0xc2, 0x6f, 0xa8, 0xd1, 0x2b, 0x46, 0x2e, 0x5e, 0x67, 0x68, 0x98, 0x81, 0x75, 0x08,
	0x39, 0x72, 0x71, 0xc0, 0xfc, 0x29, 0x24, 0xa9, 0x07, 0x0b, 0x4f, 0x3d, 0xb4, 0x20, 0x93, 0x92,
	0xc2, 0x26, 0x05, 0xb5, 0xcc, 0x8b, 0x8b, 0x1b, 0xc9, 0x63, 0x42, 0x32, 0x08, 0xa5, 0x98, 0xa1,
	0x22, 0x25, 0x8b, 0x43, 0x16, 0x61, 0x16, 0x92, 0x7f, 0x90, 0xcd, 0xc2, 0x0c, 0x0c, 0x64, 0xb3,
	0xb0, 0x04, 0x82, 0x93, 0x60, 0x14, 0x3f, 0x4c, 0x3e, 0x1e, 0x92, 0x6c, 0x92, 0xd8, 0xc0, 0x69,
	0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xa2, 0xc1, 0x14, 0x55, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarEventClient is the client API for CalendarEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarEventClient interface {
	AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error)
}

type calendarEventClient struct {
	cc *grpc.ClientConn
}

func NewCalendarEventClient(cc *grpc.ClientConn) CalendarEventClient {
	return &calendarEventClient{cc}
}

func (c *calendarEventClient) AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error) {
	out := new(AddEventResponse)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvent/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error) {
	out := new(UpdateEventResponse)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvent/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarEventClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error) {
	out := new(DeleteEventResponse)
	err := c.cc.Invoke(ctx, "/calendar.CalendarEvent/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarEventServer is the server API for CalendarEvent service.
type CalendarEventServer interface {
	AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error)
}

// UnimplementedCalendarEventServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarEventServer struct {
}

func (*UnimplementedCalendarEventServer) AddEvent(ctx context.Context, req *AddEventRequest) (*AddEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (*UnimplementedCalendarEventServer) UpdateEvent(ctx context.Context, req *UpdateEventRequest) (*UpdateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (*UnimplementedCalendarEventServer) DeleteEvent(ctx context.Context, req *DeleteEventRequest) (*DeleteEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}

func RegisterCalendarEventServer(s *grpc.Server, srv CalendarEventServer) {
	s.RegisterService(&_CalendarEvent_serviceDesc, srv)
}

func _CalendarEvent_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvent/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventServer).AddEvent(ctx, req.(*AddEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvent_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvent/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarEvent_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarEventServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.CalendarEvent/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarEventServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalendarEvent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.CalendarEvent",
	HandlerType: (*CalendarEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _CalendarEvent_AddEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _CalendarEvent_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _CalendarEvent_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "models/server/server.proto",
}
