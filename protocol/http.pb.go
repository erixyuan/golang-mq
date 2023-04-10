// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: http.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ConsumeUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitLogOffset int64  `protobuf:"varint,1,opt,name=commitLogOffset,proto3" json:"commitLogOffset,omitempty"`
	Size            int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	TagHashCode     int64  `protobuf:"varint,3,opt,name=tagHashCode,proto3" json:"tagHashCode,omitempty"`
	MessageId       string `protobuf:"bytes,4,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Offset          int64  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ConsumeUnit) Reset() {
	*x = ConsumeUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeUnit) ProtoMessage() {}

func (x *ConsumeUnit) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeUnit.ProtoReflect.Descriptor instead.
func (*ConsumeUnit) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumeUnit) GetCommitLogOffset() int64 {
	if x != nil {
		return x.CommitLogOffset
	}
	return 0
}

func (x *ConsumeUnit) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ConsumeUnit) GetTagHashCode() int64 {
	if x != nil {
		return x.TagHashCode
	}
	return 0
}

func (x *ConsumeUnit) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *ConsumeUnit) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FetchMessageListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int32  `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Topic    string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	QueueId  int32  `protobuf:"varint,4,opt,name=queueId,proto3" json:"queueId,omitempty"`
}

func (x *FetchMessageListReq) Reset() {
	*x = FetchMessageListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMessageListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMessageListReq) ProtoMessage() {}

func (x *FetchMessageListReq) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMessageListReq.ProtoReflect.Descriptor instead.
func (*FetchMessageListReq) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{2}
}

func (x *FetchMessageListReq) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *FetchMessageListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FetchMessageListReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *FetchMessageListReq) GetQueueId() int32 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

type FetchMessageListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *FetchMessageListResp_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FetchMessageListResp) Reset() {
	*x = FetchMessageListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMessageListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMessageListResp) ProtoMessage() {}

func (x *FetchMessageListResp) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMessageListResp.ProtoReflect.Descriptor instead.
func (*FetchMessageListResp) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{3}
}

func (x *FetchMessageListResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FetchMessageListResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FetchMessageListResp) GetData() *FetchMessageListResp_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type TopicRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TopicRecord) Reset() {
	*x = TopicRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicRecord) ProtoMessage() {}

func (x *TopicRecord) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicRecord.ProtoReflect.Descriptor instead.
func (*TopicRecord) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{4}
}

func (x *TopicRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FetchTopicListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *FetchTopicListResp_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FetchTopicListResp) Reset() {
	*x = FetchTopicListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTopicListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTopicListResp) ProtoMessage() {}

func (x *FetchTopicListResp) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTopicListResp.ProtoReflect.Descriptor instead.
func (*FetchTopicListResp) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{5}
}

func (x *FetchTopicListResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FetchTopicListResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FetchTopicListResp) GetData() *FetchTopicListResp_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type FetchMessageListResp_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*ConsumeUnit `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *FetchMessageListResp_Data) Reset() {
	*x = FetchMessageListResp_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchMessageListResp_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchMessageListResp_Data) ProtoMessage() {}

func (x *FetchMessageListResp_Data) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchMessageListResp_Data.ProtoReflect.Descriptor instead.
func (*FetchMessageListResp_Data) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{3, 0}
}

func (x *FetchMessageListResp_Data) GetList() []*ConsumeUnit {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *FetchMessageListResp_Data) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type FetchTopicListResp_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*TopicRecord `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FetchTopicListResp_Data) Reset() {
	*x = FetchTopicListResp_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTopicListResp_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTopicListResp_Data) ProtoMessage() {}

func (x *FetchTopicListResp_Data) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTopicListResp_Data.ProtoReflect.Descriptor instead.
func (*FetchTopicListResp_Data) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{5, 0}
}

func (x *FetchTopicListResp_Data) GetList() []*TopicRecord {
	if x != nil {
		return x.List
	}
	return nil
}

var File_http_proto protoreflect.FileDescriptor

var file_http_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xa3, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4c,
	0x6f, 0x67, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x61, 0x67, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x61, 0x67, 0x48, 0x61, 0x73,
	0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x7b, 0x0a, 0x13, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18,
	0x0a, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x14, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x44, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x21, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_proto_rawDescOnce sync.Once
	file_http_proto_rawDescData = file_http_proto_rawDesc
)

func file_http_proto_rawDescGZIP() []byte {
	file_http_proto_rawDescOnce.Do(func() {
		file_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_proto_rawDescData)
	})
	return file_http_proto_rawDescData
}

var file_http_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_http_proto_goTypes = []interface{}{
	(*Response)(nil),                  // 0: proto.Response
	(*ConsumeUnit)(nil),               // 1: proto.ConsumeUnit
	(*FetchMessageListReq)(nil),       // 2: proto.FetchMessageListReq
	(*FetchMessageListResp)(nil),      // 3: proto.FetchMessageListResp
	(*TopicRecord)(nil),               // 4: proto.TopicRecord
	(*FetchTopicListResp)(nil),        // 5: proto.FetchTopicListResp
	(*FetchMessageListResp_Data)(nil), // 6: proto.FetchMessageListResp.Data
	(*FetchTopicListResp_Data)(nil),   // 7: proto.FetchTopicListResp.Data
}
var file_http_proto_depIdxs = []int32{
	6, // 0: proto.FetchMessageListResp.data:type_name -> proto.FetchMessageListResp.Data
	7, // 1: proto.FetchTopicListResp.data:type_name -> proto.FetchTopicListResp.Data
	1, // 2: proto.FetchMessageListResp.Data.list:type_name -> proto.ConsumeUnit
	4, // 3: proto.FetchTopicListResp.Data.list:type_name -> proto.TopicRecord
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_http_proto_init() }
func file_http_proto_init() {
	if File_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeUnit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchMessageListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_http_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchMessageListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_http_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_http_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTopicListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_http_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchMessageListResp_Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_http_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTopicListResp_Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_proto_goTypes,
		DependencyIndexes: file_http_proto_depIdxs,
		MessageInfos:      file_http_proto_msgTypes,
	}.Build()
	File_http_proto = out.File
	file_http_proto_rawDesc = nil
	file_http_proto_goTypes = nil
	file_http_proto_depIdxs = nil
}