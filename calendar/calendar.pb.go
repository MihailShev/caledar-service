// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

package calendar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	UUID                 string               `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	UserId               uint64               `protobuf:"varint,6,opt,name=userId,proto3" json:"userId,omitempty"`
	NoticeTime           uint32               `protobuf:"varint,7,opt,name=noticeTime,proto3" json:"noticeTime,omitempty"`
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

func (m *Event) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
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

func (m *Event) GetNoticeTime() uint32 {
	if m != nil {
		return m.NoticeTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "calendar.Event")
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor_e3d25d49f056cdb2) }

var fileDescriptor_e3d25d49f056cdb2 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8d, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0xdd, 0x5d, 0x75, 0x8a, 0x1e, 0x06, 0x91, 0xd0, 0x83, 0x06, 0x4f, 0x7b, 0x90,
	0x54, 0xf4, 0x15, 0xf4, 0xd0, 0x6b, 0xb0, 0x0f, 0x90, 0x6e, 0xc6, 0x12, 0xd8, 0x26, 0x4b, 0x32,
	0xf5, 0xa5, 0x7d, 0x09, 0x69, 0xe2, 0x42, 0x6f, 0xde, 0xe6, 0xff, 0xf8, 0x98, 0x0f, 0x6e, 0x07,
	0x3b, 0x52, 0x70, 0x36, 0xe9, 0x29, 0x45, 0x8e, 0x78, 0x35, 0xef, 0xd5, 0xe3, 0x3e, 0xc6, 0xfd,
	0x48, 0xeb, 0xc2, 0x77, 0xc7, 0xaf, 0x35, 0xfb, 0x03, 0x65, 0xb6, 0x87, 0xa9, 0xaa, 0x4f, 0x3f,
	0x02, 0xda, 0x8f, 0x6f, 0x0a, 0x8c, 0x08, 0xcd, 0x76, 0xbb, 0x79, 0x97, 0x42, 0x89, 0xfe, 0xda,
	0x94, 0x1b, 0xef, 0xa0, 0x65, 0xcf, 0x23, 0xc9, 0x8b, 0x02, 0xeb, 0xc0, 0x17, 0x68, 0x33, 0xdb,
	0xc4, 0x72, 0xa1, 0x44, 0xbf, 0x7c, 0x5d, 0xe9, 0x1a, 0xd1, 0x73, 0x44, 0x7f, 0xce, 0x11, 0x53,
	0x45, 0x7c, 0x86, 0x05, 0x05, 0x27, 0x9b, 0x7f, 0xfd, 0x93, 0x86, 0x0a, 0x96, 0x8e, 0xf2, 0x90,
	0xfc, 0xc4, 0x3e, 0x06, 0xd9, 0x96, 0xf6, 0x39, 0xc2, 0x7b, 0xe8, 0x8e, 0x99, 0xd2, 0xc6, 0xc9,
	0x4e, 0x89, 0xbe, 0x31, 0x7f, 0x0b, 0x1f, 0x00, 0x42, 0x64, 0x3f, 0xd0, 0xe9, 0xa3, 0xbc, 0x54,
	0xa2, 0xbf, 0x31, 0x67, 0x64, 0xd7, 0x95, 0xe4, 0xdb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5,
	0x57, 0xe4, 0x1e, 0x31, 0x01, 0x00, 0x00,
}
