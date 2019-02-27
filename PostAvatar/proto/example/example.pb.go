// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostAvatar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	//二进制图片流
	Avatar []byte `protobuf:"bytes,1,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	//文件大小
	Filesize int64 `protobuf:"varint,2,opt,name=Filesize,proto3" json:"Filesize,omitempty"`
	//文件后缀
	Fileext              string   `protobuf:"bytes,3,opt,name=Fileext,proto3" json:"Fileext,omitempty"`
	SessionId            string   `protobuf:"bytes,4,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAvatar() []byte {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *Request) GetFilesize() int64 {
	if m != nil {
		return m.Filesize
	}
	return 0
}

func (m *Request) GetFileext() string {
	if m != nil {
		return m.Fileext
	}
	return ""
}

func (m *Request) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type Response struct {
	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	//不完整的头像地址
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=Avatar_url,json=AvatarUrl,proto3" json:"Avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.PostAvatar.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostAvatar.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostAvatar.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.PostAvatar.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.PostAvatar.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.PostAvatar.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.PostAvatar.Pong")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0x7e, 0x69, 0x9b, 0x66, 0xf0, 0x50, 0x97, 0xa2, 0xa1, 0x55, 0xa9, 0x39, 0xd5,
	0x4b, 0x04, 0xfd, 0x05, 0x1e, 0x2a, 0x78, 0x10, 0xca, 0x16, 0xf1, 0xa6, 0xc4, 0x3a, 0x84, 0x60,
	0xb2, 0x5b, 0x67, 0x36, 0xa5, 0xfa, 0xeb, 0x25, 0xbb, 0x5b, 0xe3, 0xc1, 0x7a, 0xca, 0x3e, 0xf3,
	0x4e, 0x32, 0x4f, 0x98, 0x85, 0xc9, 0x9a, 0xb4, 0xd1, 0x97, 0xb8, 0xcd, 0xaa, 0x75, 0x89, 0xbb,
	0x67, 0x6a, 0xab, 0xe2, 0x38, 0xd7, 0x69, 0x55, 0xac, 0x48, 0xa7, 0x4c, 0x9b, 0x74, 0xa1, 0xd9,
	0xdc, 0x6c, 0x32, 0x93, 0x51, 0x32, 0x81, 0xf0, 0x1e, 0x99, 0xb3, 0x1c, 0xc5, 0x10, 0x02, 0xce,
	0x3e, 0xe2, 0xce, 0xb4, 0x33, 0x8b, 0x64, 0x73, 0x4c, 0x6a, 0x08, 0x25, 0xbe, 0xd7, 0xc8, 0x46,
	0x1c, 0x41, 0xdf, 0xbd, 0x61, 0xf3, 0x03, 0xe9, 0x49, 0x8c, 0x61, 0x70, 0x5b, 0x94, 0xc8, 0xc5,
	0x27, 0xc6, 0xff, 0xa7, 0x9d, 0x59, 0x20, 0xbf, 0x59, 0xc4, 0x10, 0x36, 0x67, 0xdc, 0x9a, 0x38,
	0xb0, 0x1f, 0xdd, 0xa1, 0x38, 0x81, 0x68, 0x89, 0xcc, 0x85, 0x56, 0x77, 0xaf, 0x71, 0xd7, 0x66,
	0x6d, 0x21, 0x79, 0x84, 0x81, 0x44, 0x5e, 0x6b, 0xc5, 0x28, 0x46, 0xd0, 0x9b, 0x13, 0x29, 0xed,
	0xb5, 0x1c, 0x34, 0x36, 0x73, 0xa2, 0x8a, 0x73, 0x3b, 0x33, 0x92, 0x9e, 0xc4, 0x29, 0x80, 0xf3,
	0x7a, 0xae, 0xa9, 0xf4, 0x43, 0x23, 0x57, 0x79, 0xa0, 0x32, 0x99, 0xc1, 0x70, 0x69, 0x08, 0xb3,
	0xaa, 0x50, 0xf9, 0xee, 0xc7, 0x46, 0xd0, 0x5b, 0xe9, 0x5a, 0x19, 0x3b, 0x20, 0x90, 0x0e, 0x92,
	0x0b, 0x38, 0xfc, 0xd1, 0xd9, 0xba, 0xfc, 0xd2, 0x7a, 0x06, 0xdd, 0x45, 0xa1, 0xf2, 0xc6, 0x89,
	0x0d, 0xe9, 0x37, 0xf4, 0xb1, 0x27, 0x9b, 0xeb, 0xfd, 0xf9, 0xd5, 0x13, 0x84, 0x73, 0xb7, 0x2b,
	0xb1, 0x04, 0x68, 0x57, 0x23, 0xa6, 0xe9, 0x9e, 0xa5, 0xa5, 0xde, 0x7d, 0x7c, 0xfe, 0x47, 0x87,
	0x73, 0x4e, 0xfe, 0xbd, 0xf4, 0xed, 0x0d, 0xb8, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x4a,
	0x46, 0x40, 0x20, 0x02, 0x00, 0x00,
}
