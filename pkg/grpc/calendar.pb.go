// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

package calendarpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Event struct {
	UUID                 int64                `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	NotifyTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=notifyTime,proto3" json:"notifyTime,omitempty"`
	Description          string               `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	UserId               uint64               `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Event) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Event) GetNotifyTime() *timestamp.Timestamp {
	if m != nil {
		return m.NotifyTime
	}
	return nil
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CheckReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckReq) Reset()         { *m = CheckReq{} }
func (m *CheckReq) String() string { return proto.CompactTextString(m) }
func (*CheckReq) ProtoMessage()    {}
func (*CheckReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{1}
}

func (m *CheckReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckReq.Unmarshal(m, b)
}
func (m *CheckReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckReq.Marshal(b, m, deterministic)
}
func (m *CheckReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckReq.Merge(m, src)
}
func (m *CheckReq) XXX_Size() int {
	return xxx_messageInfo_CheckReq.Size(m)
}
func (m *CheckReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckReq proto.InternalMessageInfo

type CheckRes struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRes) Reset()         { *m = CheckRes{} }
func (m *CheckRes) String() string { return proto.CompactTextString(m) }
func (*CheckRes) ProtoMessage()    {}
func (*CheckRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{2}
}

func (m *CheckRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRes.Unmarshal(m, b)
}
func (m *CheckRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRes.Marshal(b, m, deterministic)
}
func (m *CheckRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRes.Merge(m, src)
}
func (m *CheckRes) XXX_Size() int {
	return xxx_messageInfo_CheckRes.Size(m)
}
func (m *CheckRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRes.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRes proto.InternalMessageInfo

func (m *CheckRes) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CreateEventReq struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventReq) Reset()         { *m = CreateEventReq{} }
func (m *CreateEventReq) String() string { return proto.CompactTextString(m) }
func (*CreateEventReq) ProtoMessage()    {}
func (*CreateEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{3}
}

func (m *CreateEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventReq.Unmarshal(m, b)
}
func (m *CreateEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventReq.Marshal(b, m, deterministic)
}
func (m *CreateEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventReq.Merge(m, src)
}
func (m *CreateEventReq) XXX_Size() int {
	return xxx_messageInfo_CreateEventReq.Size(m)
}
func (m *CreateEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventReq proto.InternalMessageInfo

func (m *CreateEventReq) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type CreateEventRes struct {
	UUID                 int64    `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventRes) Reset()         { *m = CreateEventRes{} }
func (m *CreateEventRes) String() string { return proto.CompactTextString(m) }
func (*CreateEventRes) ProtoMessage()    {}
func (*CreateEventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{4}
}

func (m *CreateEventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRes.Unmarshal(m, b)
}
func (m *CreateEventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRes.Marshal(b, m, deterministic)
}
func (m *CreateEventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRes.Merge(m, src)
}
func (m *CreateEventRes) XXX_Size() int {
	return xxx_messageInfo_CreateEventRes.Size(m)
}
func (m *CreateEventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRes proto.InternalMessageInfo

func (m *CreateEventRes) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

func (m *CreateEventRes) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetEventReq struct {
	UUID                 int64    `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventReq) Reset()         { *m = GetEventReq{} }
func (m *GetEventReq) String() string { return proto.CompactTextString(m) }
func (*GetEventReq) ProtoMessage()    {}
func (*GetEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{5}
}

func (m *GetEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventReq.Unmarshal(m, b)
}
func (m *GetEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventReq.Marshal(b, m, deterministic)
}
func (m *GetEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventReq.Merge(m, src)
}
func (m *GetEventReq) XXX_Size() int {
	return xxx_messageInfo_GetEventReq.Size(m)
}
func (m *GetEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventReq proto.InternalMessageInfo

func (m *GetEventReq) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

type GetEventRes struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventRes) Reset()         { *m = GetEventRes{} }
func (m *GetEventRes) String() string { return proto.CompactTextString(m) }
func (*GetEventRes) ProtoMessage()    {}
func (*GetEventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{6}
}

func (m *GetEventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventRes.Unmarshal(m, b)
}
func (m *GetEventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventRes.Marshal(b, m, deterministic)
}
func (m *GetEventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventRes.Merge(m, src)
}
func (m *GetEventRes) XXX_Size() int {
	return xxx_messageInfo_GetEventRes.Size(m)
}
func (m *GetEventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventRes proto.InternalMessageInfo

func (m *GetEventRes) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *GetEventRes) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UpdateEventReq struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventReq) Reset()         { *m = UpdateEventReq{} }
func (m *UpdateEventReq) String() string { return proto.CompactTextString(m) }
func (*UpdateEventReq) ProtoMessage()    {}
func (*UpdateEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{7}
}

func (m *UpdateEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventReq.Unmarshal(m, b)
}
func (m *UpdateEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventReq.Marshal(b, m, deterministic)
}
func (m *UpdateEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventReq.Merge(m, src)
}
func (m *UpdateEventReq) XXX_Size() int {
	return xxx_messageInfo_UpdateEventReq.Size(m)
}
func (m *UpdateEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventReq proto.InternalMessageInfo

func (m *UpdateEventReq) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateEventRes struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventRes) Reset()         { *m = UpdateEventRes{} }
func (m *UpdateEventRes) String() string { return proto.CompactTextString(m) }
func (*UpdateEventRes) ProtoMessage()    {}
func (*UpdateEventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{8}
}

func (m *UpdateEventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventRes.Unmarshal(m, b)
}
func (m *UpdateEventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventRes.Marshal(b, m, deterministic)
}
func (m *UpdateEventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventRes.Merge(m, src)
}
func (m *UpdateEventRes) XXX_Size() int {
	return xxx_messageInfo_UpdateEventRes.Size(m)
}
func (m *UpdateEventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventRes proto.InternalMessageInfo

func (m *UpdateEventRes) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *UpdateEventRes) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "calendarpb.Event")
	proto.RegisterType((*CheckReq)(nil), "calendarpb.CheckReq")
	proto.RegisterType((*CheckRes)(nil), "calendarpb.CheckRes")
	proto.RegisterType((*CreateEventReq)(nil), "calendarpb.CreateEventReq")
	proto.RegisterType((*CreateEventRes)(nil), "calendarpb.CreateEventRes")
	proto.RegisterType((*GetEventReq)(nil), "calendarpb.GetEventReq")
	proto.RegisterType((*GetEventRes)(nil), "calendarpb.GetEventRes")
	proto.RegisterType((*UpdateEventReq)(nil), "calendarpb.UpdateEventReq")
	proto.RegisterType((*UpdateEventRes)(nil), "calendarpb.UpdateEventRes")
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor_e3d25d49f056cdb2) }

var fileDescriptor_e3d25d49f056cdb2 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0xab, 0xda, 0x40,
	0x14, 0x25, 0xd1, 0x58, 0xbd, 0x29, 0x42, 0x2f, 0xd2, 0x86, 0x6c, 0x9a, 0x66, 0xd3, 0x2c, 0x4a,
	0x2c, 0x76, 0xa5, 0x74, 0x67, 0xa5, 0x08, 0x85, 0x42, 0x5a, 0x7f, 0x40, 0x4c, 0xae, 0x36, 0x54,
	0x93, 0x38, 0x33, 0x0a, 0xfd, 0x05, 0xdd, 0xf7, 0x17, 0x97, 0xcc, 0x18, 0xdf, 0x44, 0x0c, 0x3e,
	0xde, 0xdb, 0xdd, 0x8f, 0x73, 0x73, 0x4e, 0xce, 0x61, 0x60, 0x98, 0xc4, 0x3b, 0xca, 0xd3, 0x98,
	0x85, 0x25, 0x2b, 0x44, 0x81, 0x50, 0xf7, 0xe5, 0xda, 0x7d, 0xbb, 0x2d, 0x8a, 0xed, 0x8e, 0xc6,
	0x72, 0xb3, 0x3e, 0x6e, 0xc6, 0x22, 0xdb, 0x13, 0x17, 0xf1, 0xbe, 0x54, 0x60, 0xff, 0xaf, 0x09,
	0xd6, 0xe2, 0x44, 0xb9, 0x40, 0x84, 0xee, 0x6a, 0xb5, 0xfc, 0xe2, 0x18, 0x9e, 0x11, 0x74, 0x22,
	0x59, 0xe3, 0x08, 0x2c, 0x91, 0x89, 0x1d, 0x39, 0xa6, 0x67, 0x04, 0x83, 0x48, 0x35, 0xf8, 0x11,
	0x2c, 0x2e, 0x62, 0x26, 0x9c, 0x8e, 0x67, 0x04, 0xf6, 0xc4, 0x0d, 0x15, 0x49, 0x58, 0x93, 0x84,
	0x3f, 0x6b, 0x92, 0x48, 0x01, 0xf1, 0x03, 0x74, 0x28, 0x4f, 0x9d, 0xee, 0x5d, 0x7c, 0x05, 0xc3,
	0x19, 0x40, 0x5e, 0x88, 0x6c, 0xf3, 0xa7, 0x9a, 0x3b, 0xd6, 0xdd, 0x23, 0x0d, 0x8d, 0x1e, 0xd8,
	0x29, 0xf1, 0x84, 0x65, 0xa5, 0xc8, 0x8a, 0xdc, 0xe9, 0x49, 0xdd, 0xfa, 0x08, 0x5f, 0x43, 0xef,
	0xc8, 0x89, 0x2d, 0x53, 0xe7, 0x85, 0x67, 0x04, 0xdd, 0xe8, 0xdc, 0xf9, 0x00, 0xfd, 0xf9, 0x2f,
	0x4a, 0x7e, 0x47, 0x74, 0xf0, 0xbd, 0x4b, 0xcd, 0x2b, 0x0f, 0x88, 0xb1, 0x82, 0x49, 0x63, 0x06,
	0x91, 0x6a, 0xfc, 0x29, 0x0c, 0xe7, 0x8c, 0x62, 0x41, 0xd2, 0xbc, 0x88, 0x0e, 0xf8, 0x1e, 0x2c,
	0xaa, 0x6a, 0x89, 0xb3, 0x27, 0xaf, 0xc2, 0x87, 0x18, 0x42, 0x05, 0x52, 0x7b, 0x7f, 0x76, 0x75,
	0xca, 0xdb, 0xac, 0x57, 0xb4, 0xa6, 0x4e, 0xfb, 0x0e, 0xec, 0xaf, 0x24, 0x2e, 0x9c, 0x37, 0x0e,
	0xfd, 0x6f, 0x3a, 0x84, 0x3f, 0x5a, 0x56, 0x0b, 0xe1, 0x14, 0x86, 0xab, 0x32, 0x7d, 0xd2, 0x7f,
	0x7e, 0xbf, 0x3a, 0x7d, 0xae, 0x96, 0xc9, 0x3f, 0x13, 0xfa, 0xf3, 0xf3, 0x05, 0xce, 0xe0, 0xa5,
	0x8c, 0xe8, 0x07, 0xb1, 0x53, 0x96, 0x10, 0x8e, 0xf4, 0x8f, 0xd5, 0x41, 0xba, 0xb7, 0xa6, 0x1c,
	0x17, 0x60, 0x6b, 0x09, 0xa0, 0xdb, 0x00, 0x35, 0x52, 0x75, 0xdb, 0x77, 0x1c, 0x3f, 0x43, 0xbf,
	0x76, 0x1a, 0xdf, 0xe8, 0x38, 0x2d, 0x22, 0xb7, 0x65, 0x21, 0x45, 0x68, 0xf6, 0x34, 0x45, 0x34,
	0x2d, 0x77, 0xdb, 0x77, 0x7c, 0xdd, 0x93, 0x0f, 0xe2, 0xd3, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x40, 0x99, 0xbe, 0x0d, 0x06, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	CheckService(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckRes, error)
	CreateEvent(ctx context.Context, in *CreateEventReq, opts ...grpc.CallOption) (*CreateEventRes, error)
	GetEvent(ctx context.Context, in *GetEventReq, opts ...grpc.CallOption) (*GetEventRes, error)
	UpdateEvent(ctx context.Context, in *UpdateEventReq, opts ...grpc.CallOption) (*UpdateEventRes, error)
}

type calendarClient struct {
	cc *grpc.ClientConn
}

func NewCalendarClient(cc *grpc.ClientConn) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) CheckService(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckRes, error) {
	out := new(CheckRes)
	err := c.cc.Invoke(ctx, "/calendarpb.Calendar/CheckService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) CreateEvent(ctx context.Context, in *CreateEventReq, opts ...grpc.CallOption) (*CreateEventRes, error) {
	out := new(CreateEventRes)
	err := c.cc.Invoke(ctx, "/calendarpb.Calendar/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetEvent(ctx context.Context, in *GetEventReq, opts ...grpc.CallOption) (*GetEventRes, error) {
	out := new(GetEventRes)
	err := c.cc.Invoke(ctx, "/calendarpb.Calendar/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) UpdateEvent(ctx context.Context, in *UpdateEventReq, opts ...grpc.CallOption) (*UpdateEventRes, error) {
	out := new(UpdateEventRes)
	err := c.cc.Invoke(ctx, "/calendarpb.Calendar/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	CheckService(context.Context, *CheckReq) (*CheckRes, error)
	CreateEvent(context.Context, *CreateEventReq) (*CreateEventRes, error)
	GetEvent(context.Context, *GetEventReq) (*GetEventRes, error)
	UpdateEvent(context.Context, *UpdateEventReq) (*UpdateEventRes, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) CheckService(ctx context.Context, req *CheckReq) (*CheckRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckService not implemented")
}
func (*UnimplementedCalendarServer) CreateEvent(ctx context.Context, req *CreateEventReq) (*CreateEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedCalendarServer) GetEvent(ctx context.Context, req *GetEventReq) (*GetEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServer) UpdateEvent(ctx context.Context, req *UpdateEventReq) (*UpdateEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_CheckService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).CheckService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendarpb.Calendar/CheckService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).CheckService(ctx, req.(*CheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendarpb.Calendar/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).CreateEvent(ctx, req.(*CreateEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendarpb.Calendar/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEvent(ctx, req.(*GetEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendarpb.Calendar/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).UpdateEvent(ctx, req.(*UpdateEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calendarpb.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckService",
			Handler:    _Calendar_CheckService_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Calendar_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Calendar_GetEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Calendar_UpdateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}
