// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageUserInfo struct {
	OpenId               string   `protobuf:"bytes,1,opt,name=open_id,json=openId,proto3" json:"open_id,omitempty"`
	ProfilePhoto         string   `protobuf:"bytes,2,opt,name=profile_photo,json=profilePhoto,proto3" json:"profile_photo,omitempty"`
	Sex                  int32    `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	UserName             string   `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageUserInfo) Reset()         { *m = MessageUserInfo{} }
func (m *MessageUserInfo) String() string { return proto.CompactTextString(m) }
func (*MessageUserInfo) ProtoMessage()    {}
func (*MessageUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *MessageUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageUserInfo.Unmarshal(m, b)
}
func (m *MessageUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageUserInfo.Marshal(b, m, deterministic)
}
func (m *MessageUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageUserInfo.Merge(m, src)
}
func (m *MessageUserInfo) XXX_Size() int {
	return xxx_messageInfo_MessageUserInfo.Size(m)
}
func (m *MessageUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MessageUserInfo proto.InternalMessageInfo

func (m *MessageUserInfo) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *MessageUserInfo) GetProfilePhoto() string {
	if m != nil {
		return m.ProfilePhoto
	}
	return ""
}

func (m *MessageUserInfo) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *MessageUserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type MessageInfo struct {
	MessageId            int64            `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Content              string           `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Urls                 []string         `protobuf:"bytes,3,rep,name=urls,proto3" json:"urls,omitempty"`
	CreateTime           int64            `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	MessageType          int32            `protobuf:"varint,5,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	UserInfo             *MessageUserInfo `protobuf:"bytes,6,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	DiggCount            int32            `protobuf:"varint,7,opt,name=digg_count,json=diggCount,proto3" json:"digg_count,omitempty"`
	CommentCount         int32            `protobuf:"varint,8,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	Digg                 bool             `protobuf:"varint,9,opt,name=digg,proto3" json:"digg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MessageInfo) Reset()         { *m = MessageInfo{} }
func (m *MessageInfo) String() string { return proto.CompactTextString(m) }
func (*MessageInfo) ProtoMessage()    {}
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *MessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageInfo.Unmarshal(m, b)
}
func (m *MessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageInfo.Marshal(b, m, deterministic)
}
func (m *MessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageInfo.Merge(m, src)
}
func (m *MessageInfo) XXX_Size() int {
	return xxx_messageInfo_MessageInfo.Size(m)
}
func (m *MessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MessageInfo proto.InternalMessageInfo

func (m *MessageInfo) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MessageInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MessageInfo) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *MessageInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *MessageInfo) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *MessageInfo) GetUserInfo() *MessageUserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *MessageInfo) GetDiggCount() int32 {
	if m != nil {
		return m.DiggCount
	}
	return 0
}

func (m *MessageInfo) GetCommentCount() int32 {
	if m != nil {
		return m.CommentCount
	}
	return 0
}

func (m *MessageInfo) GetDigg() bool {
	if m != nil {
		return m.Digg
	}
	return false
}

//创建帖子
// POST /create/message
type CreateMessageRequest struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Uris                 []string `protobuf:"bytes,2,rep,name=uris,proto3" json:"uris,omitempty"`
	MessageType          int32    `protobuf:"varint,3,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageRequest) Reset()         { *m = CreateMessageRequest{} }
func (m *CreateMessageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMessageRequest) ProtoMessage()    {}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *CreateMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageRequest.Unmarshal(m, b)
}
func (m *CreateMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageRequest.Marshal(b, m, deterministic)
}
func (m *CreateMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageRequest.Merge(m, src)
}
func (m *CreateMessageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMessageRequest.Size(m)
}
func (m *CreateMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageRequest proto.InternalMessageInfo

func (m *CreateMessageRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateMessageRequest) GetUris() []string {
	if m != nil {
		return m.Uris
	}
	return nil
}

func (m *CreateMessageRequest) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

type CreateMessageResponse struct {
	StatusCode           int32        `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message              string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MessageInfo          *MessageInfo `protobuf:"bytes,3,opt,name=message_info,json=messageInfo,proto3" json:"message_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateMessageResponse) Reset()         { *m = CreateMessageResponse{} }
func (m *CreateMessageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMessageResponse) ProtoMessage()    {}
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *CreateMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageResponse.Unmarshal(m, b)
}
func (m *CreateMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageResponse.Marshal(b, m, deterministic)
}
func (m *CreateMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageResponse.Merge(m, src)
}
func (m *CreateMessageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMessageResponse.Size(m)
}
func (m *CreateMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageResponse proto.InternalMessageInfo

func (m *CreateMessageResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *CreateMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateMessageResponse) GetMessageInfo() *MessageInfo {
	if m != nil {
		return m.MessageInfo
	}
	return nil
}

// GET /get/user_message
type GetUserMessageRequest struct {
	OpenId               string   `protobuf:"bytes,1,opt,name=open_id,json=openId,proto3" json:"open_id,omitempty"`
	Index                int64    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Count                int64    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserMessageRequest) Reset()         { *m = GetUserMessageRequest{} }
func (m *GetUserMessageRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserMessageRequest) ProtoMessage()    {}
func (*GetUserMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *GetUserMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserMessageRequest.Unmarshal(m, b)
}
func (m *GetUserMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserMessageRequest.Marshal(b, m, deterministic)
}
func (m *GetUserMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserMessageRequest.Merge(m, src)
}
func (m *GetUserMessageRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserMessageRequest.Size(m)
}
func (m *GetUserMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserMessageRequest proto.InternalMessageInfo

func (m *GetUserMessageRequest) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *GetUserMessageRequest) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *GetUserMessageRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetUserMessageResonse struct {
	StatusCode           int32          `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message              string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MessageInfos         []*MessageInfo `protobuf:"bytes,3,rep,name=message_infos,json=messageInfos,proto3" json:"message_infos,omitempty"`
	Count                int64          `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	HasMore              bool           `protobuf:"varint,5,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetUserMessageResonse) Reset()         { *m = GetUserMessageResonse{} }
func (m *GetUserMessageResonse) String() string { return proto.CompactTextString(m) }
func (*GetUserMessageResonse) ProtoMessage()    {}
func (*GetUserMessageResonse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *GetUserMessageResonse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserMessageResonse.Unmarshal(m, b)
}
func (m *GetUserMessageResonse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserMessageResonse.Marshal(b, m, deterministic)
}
func (m *GetUserMessageResonse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserMessageResonse.Merge(m, src)
}
func (m *GetUserMessageResonse) XXX_Size() int {
	return xxx_messageInfo_GetUserMessageResonse.Size(m)
}
func (m *GetUserMessageResonse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserMessageResonse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserMessageResonse proto.InternalMessageInfo

func (m *GetUserMessageResonse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *GetUserMessageResonse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetUserMessageResonse) GetMessageInfos() []*MessageInfo {
	if m != nil {
		return m.MessageInfos
	}
	return nil
}

func (m *GetUserMessageResonse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetUserMessageResonse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

//获取全部帖子
//GET /get/message
type GetMessageRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageRequest) Reset()         { *m = GetMessageRequest{} }
func (m *GetMessageRequest) String() string { return proto.CompactTextString(m) }
func (*GetMessageRequest) ProtoMessage()    {}
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *GetMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageRequest.Unmarshal(m, b)
}
func (m *GetMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageRequest.Marshal(b, m, deterministic)
}
func (m *GetMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageRequest.Merge(m, src)
}
func (m *GetMessageRequest) XXX_Size() int {
	return xxx_messageInfo_GetMessageRequest.Size(m)
}
func (m *GetMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageRequest proto.InternalMessageInfo

type GetMessageResponse struct {
	StatusCode           int32          `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message              string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MessageInfos         []*MessageInfo `protobuf:"bytes,3,rep,name=message_infos,json=messageInfos,proto3" json:"message_infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetMessageResponse) Reset()         { *m = GetMessageResponse{} }
func (m *GetMessageResponse) String() string { return proto.CompactTextString(m) }
func (*GetMessageResponse) ProtoMessage()    {}
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{7}
}

func (m *GetMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageResponse.Unmarshal(m, b)
}
func (m *GetMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageResponse.Marshal(b, m, deterministic)
}
func (m *GetMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageResponse.Merge(m, src)
}
func (m *GetMessageResponse) XXX_Size() int {
	return xxx_messageInfo_GetMessageResponse.Size(m)
}
func (m *GetMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageResponse proto.InternalMessageInfo

func (m *GetMessageResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *GetMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetMessageResponse) GetMessageInfos() []*MessageInfo {
	if m != nil {
		return m.MessageInfos
	}
	return nil
}

//获取热门帖子
//GET /get/hot_message
type GetHotMessageRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHotMessageRequest) Reset()         { *m = GetHotMessageRequest{} }
func (m *GetHotMessageRequest) String() string { return proto.CompactTextString(m) }
func (*GetHotMessageRequest) ProtoMessage()    {}
func (*GetHotMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{8}
}

func (m *GetHotMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHotMessageRequest.Unmarshal(m, b)
}
func (m *GetHotMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHotMessageRequest.Marshal(b, m, deterministic)
}
func (m *GetHotMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHotMessageRequest.Merge(m, src)
}
func (m *GetHotMessageRequest) XXX_Size() int {
	return xxx_messageInfo_GetHotMessageRequest.Size(m)
}
func (m *GetHotMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHotMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHotMessageRequest proto.InternalMessageInfo

type GetHotMessageResponse struct {
	StatusCode           int32          `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message              string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MessageInfos         []*MessageInfo `protobuf:"bytes,3,rep,name=message_infos,json=messageInfos,proto3" json:"message_infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetHotMessageResponse) Reset()         { *m = GetHotMessageResponse{} }
func (m *GetHotMessageResponse) String() string { return proto.CompactTextString(m) }
func (*GetHotMessageResponse) ProtoMessage()    {}
func (*GetHotMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{9}
}

func (m *GetHotMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHotMessageResponse.Unmarshal(m, b)
}
func (m *GetHotMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHotMessageResponse.Marshal(b, m, deterministic)
}
func (m *GetHotMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHotMessageResponse.Merge(m, src)
}
func (m *GetHotMessageResponse) XXX_Size() int {
	return xxx_messageInfo_GetHotMessageResponse.Size(m)
}
func (m *GetHotMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHotMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHotMessageResponse proto.InternalMessageInfo

func (m *GetHotMessageResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *GetHotMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetHotMessageResponse) GetMessageInfos() []*MessageInfo {
	if m != nil {
		return m.MessageInfos
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageUserInfo)(nil), "MessageUserInfo")
	proto.RegisterType((*MessageInfo)(nil), "MessageInfo")
	proto.RegisterType((*CreateMessageRequest)(nil), "CreateMessageRequest")
	proto.RegisterType((*CreateMessageResponse)(nil), "CreateMessageResponse")
	proto.RegisterType((*GetUserMessageRequest)(nil), "GetUserMessageRequest")
	proto.RegisterType((*GetUserMessageResonse)(nil), "GetUserMessageResonse")
	proto.RegisterType((*GetMessageRequest)(nil), "GetMessageRequest")
	proto.RegisterType((*GetMessageResponse)(nil), "GetMessageResponse")
	proto.RegisterType((*GetHotMessageRequest)(nil), "GetHotMessageRequest")
	proto.RegisterType((*GetHotMessageResponse)(nil), "GetHotMessageResponse")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcb, 0x8e, 0xd3, 0x3c,
	0x18, 0x55, 0xea, 0xde, 0xf2, 0xb5, 0xa3, 0x7f, 0x7e, 0xd3, 0x01, 0x23, 0x84, 0x28, 0x61, 0x93,
	0x0d, 0x45, 0xc0, 0x23, 0x74, 0x51, 0xba, 0x18, 0x84, 0xac, 0x61, 0x87, 0x14, 0x85, 0xe4, 0x6b,
	0x1b, 0x69, 0x62, 0x87, 0xd8, 0x91, 0x66, 0x36, 0x2c, 0x40, 0x3c, 0x0a, 0xaf, 0xc0, 0xf3, 0x21,
	0x5f, 0x32, 0x13, 0x52, 0x58, 0xa0, 0x59, 0xb0, 0xf3, 0x39, 0xbe, 0x9c, 0x8b, 0x9d, 0xc0, 0x49,
	0x89, 0x4a, 0xa5, 0x7b, 0x5c, 0x55, 0xb5, 0xd4, 0x32, 0xfa, 0x0c, 0xff, 0x9d, 0x3b, 0xe2, 0xbd,
	0xc2, 0x7a, 0x2b, 0x76, 0x92, 0x3e, 0x80, 0x89, 0xac, 0x50, 0x24, 0x45, 0xce, 0x82, 0x65, 0x10,
	0x87, 0x7c, 0x6c, 0xe0, 0x36, 0xa7, 0xcf, 0xe0, 0xa4, 0xaa, 0xe5, 0xae, 0xb8, 0xc4, 0xa4, 0x3a,
	0x48, 0x2d, 0xd9, 0xc0, 0x4e, 0xcf, 0x3d, 0xf9, 0xce, 0x70, 0xf4, 0x14, 0x88, 0xc2, 0x2b, 0x46,
	0x96, 0x41, 0x3c, 0xe2, 0x66, 0x48, 0x1f, 0x41, 0xd8, 0x28, 0xac, 0x13, 0x91, 0x96, 0xc8, 0x86,
	0x76, 0xcb, 0xd4, 0x10, 0x6f, 0xd3, 0x12, 0xa3, 0xef, 0x03, 0x98, 0x79, 0x03, 0x56, 0xfc, 0x31,
	0x80, 0x37, 0xd8, 0xea, 0x13, 0x1e, 0x7a, 0x66, 0x9b, 0x53, 0x06, 0x93, 0x4c, 0x0a, 0x8d, 0x42,
	0x7b, 0xf1, 0x16, 0x52, 0x0a, 0xc3, 0xa6, 0xbe, 0x54, 0x8c, 0x2c, 0x49, 0x1c, 0x72, 0x3b, 0xa6,
	0x4f, 0x60, 0x96, 0xd5, 0x98, 0x6a, 0x4c, 0x74, 0xe1, 0xb5, 0x09, 0x07, 0x47, 0x5d, 0x14, 0x25,
	0xd2, 0xa7, 0x30, 0x6f, 0xd5, 0xf4, 0x75, 0x85, 0x6c, 0x64, 0x5d, 0xcf, 0x3c, 0x77, 0x71, 0x5d,
	0x21, 0x7d, 0xee, 0xdd, 0x17, 0x62, 0x27, 0xd9, 0x78, 0x19, 0xc4, 0xb3, 0x57, 0xa7, 0xab, 0x5e,
	0x65, 0x2e, 0x4f, 0xeb, 0x3f, 0x2f, 0xf6, 0xfb, 0x24, 0x93, 0x8d, 0xd0, 0x6c, 0x62, 0xcf, 0x0b,
	0x0d, 0xb3, 0x36, 0x84, 0xa9, 0x30, 0x93, 0x65, 0x89, 0x42, 0xfb, 0x15, 0x53, 0xbb, 0x62, 0xee,
	0x49, 0xb7, 0x88, 0xc2, 0xd0, 0xec, 0x60, 0xe1, 0x32, 0x88, 0xa7, 0xdc, 0x8e, 0xa3, 0x3d, 0x2c,
	0xd6, 0xd6, 0xb7, 0x97, 0xe6, 0xf8, 0xa9, 0x41, 0xa5, 0xbb, 0x85, 0x04, 0xbf, 0x29, 0xa4, 0x50,
	0x6c, 0xd0, 0x16, 0x52, 0xa8, 0xa3, 0xbc, 0xe4, 0x28, 0x6f, 0xf4, 0x35, 0x80, 0xb3, 0x9e, 0x92,
	0xaa, 0xa4, 0x50, 0x68, 0xda, 0x54, 0x3a, 0xd5, 0x8d, 0x4a, 0x32, 0x99, 0xa3, 0x95, 0x1b, 0x71,
	0x70, 0xd4, 0x5a, 0xe6, 0x68, 0xbc, 0xf8, 0x93, 0xda, 0xcb, 0xf1, 0x90, 0xbe, 0xb8, 0xd5, 0xb5,
	0x3d, 0x12, 0xdb, 0xe3, 0x7c, 0xd5, 0xb9, 0xf9, 0x1b, 0x17, 0x06, 0x44, 0x1f, 0xe0, 0x6c, 0x83,
	0xda, 0xf4, 0xdb, 0xcb, 0xfb, 0xc7, 0xc7, 0xb9, 0x80, 0x51, 0x21, 0x72, 0xbc, 0xb2, 0xd2, 0x84,
	0x3b, 0x60, 0x58, 0xd7, 0x33, 0x71, 0xac, 0x05, 0xd1, 0x8f, 0xe0, 0xf8, 0x78, 0x75, 0xd7, 0x8c,
	0x2f, 0x6f, 0x3e, 0x2d, 0x9b, 0xd1, 0xbd, 0xc4, 0x7e, 0xc8, 0x79, 0x27, 0xa4, 0xba, 0x75, 0x37,
	0xec, 0xb8, 0xa3, 0x0f, 0x61, 0x7a, 0x48, 0x55, 0x52, 0xca, 0xda, 0x3d, 0xc8, 0x29, 0x9f, 0x1c,
	0x52, 0x75, 0x2e, 0x6b, 0x8c, 0xee, 0xc1, 0xff, 0x1b, 0xd4, 0xbf, 0x56, 0x12, 0x7d, 0x09, 0x80,
	0x76, 0xd9, 0xbb, 0x5f, 0xd7, 0xdf, 0x47, 0x89, 0xee, 0xc3, 0x62, 0x83, 0xfa, 0x8d, 0xec, 0x9b,
	0xfb, 0xe6, 0xaa, 0xee, 0x4e, 0xfc, 0x0b, 0x7f, 0x1f, 0xc7, 0xf6, 0x77, 0xf7, 0xfa, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1e, 0x53, 0xc3, 0x92, 0xff, 0x04, 0x00, 0x00,
}
