// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: memo.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateMemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Priority int32  `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *CreateMemoRequest) Reset() {
	*x = CreateMemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemoRequest) ProtoMessage() {}

func (x *CreateMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemoRequest.ProtoReflect.Descriptor instead.
func (*CreateMemoRequest) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMemoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateMemoRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateMemoRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type GetMemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMemoRequest) Reset() {
	*x = GetMemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoRequest) ProtoMessage() {}

func (x *GetMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoRequest.ProtoReflect.Descriptor instead.
func (*GetMemoRequest) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{1}
}

func (x *GetMemoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListMemosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids           []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Keyword       string  `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Priority      int32   `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	FromCreatedAt int64   `protobuf:"varint,4,opt,name=from_created_at,proto3" json:"from_created_at,omitempty"`
	ToCreatedAt   int64   `protobuf:"varint,5,opt,name=to_created_at,proto3" json:"to_created_at,omitempty"`
}

func (x *ListMemosRequest) Reset() {
	*x = ListMemosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMemosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemosRequest) ProtoMessage() {}

func (x *ListMemosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemosRequest.ProtoReflect.Descriptor instead.
func (*ListMemosRequest) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{2}
}

func (x *ListMemosRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ListMemosRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListMemosRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *ListMemosRequest) GetFromCreatedAt() int64 {
	if x != nil {
		return x.FromCreatedAt
	}
	return 0
}

func (x *ListMemosRequest) GetToCreatedAt() int64 {
	if x != nil {
		return x.ToCreatedAt
	}
	return 0
}

type UpdateMemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Priority int32  `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *UpdateMemoRequest) Reset() {
	*x = UpdateMemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMemoRequest) ProtoMessage() {}

func (x *UpdateMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMemoRequest.ProtoReflect.Descriptor instead.
func (*UpdateMemoRequest) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMemoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateMemoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateMemoRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateMemoRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type DeleteMemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMemoRequest) Reset() {
	*x = DeleteMemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemoRequest) ProtoMessage() {}

func (x *DeleteMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemoRequest.ProtoReflect.Descriptor instead.
func (*DeleteMemoRequest) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMemoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Priority  int32  `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	CreatedAt int64  `protobuf:"varint,5,opt,name=created_at,proto3" json:"created_at,omitempty"`
}

func (x *MemoResponse) Reset() {
	*x = MemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoResponse) ProtoMessage() {}

func (x *MemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoResponse.ProtoReflect.Descriptor instead.
func (*MemoResponse) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{5}
}

func (x *MemoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemoResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MemoResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MemoResponse) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *MemoResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type ListMemosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memos []*MemoResponse `protobuf:"bytes,1,rep,name=memos,proto3" json:"memos,omitempty"`
}

func (x *ListMemosResponse) Reset() {
	*x = ListMemosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMemosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemosResponse) ProtoMessage() {}

func (x *ListMemosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemosResponse.ProtoReflect.Descriptor instead.
func (*ListMemosResponse) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{6}
}

func (x *ListMemosResponse) GetMemos() []*MemoResponse {
	if x != nil {
		return x.Memos
	}
	return nil
}

// 추가된 메시지 정의
type DeleteMemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 삭제 작업의 성공 여부, 오류 메시지 등 필요한 필드를 추가할 수 있습니다.
	// 예시로, 성공 여부만 나타내는 simple boolean 필드를 추가합니다.
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteMemoResponse) Reset() {
	*x = DeleteMemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemoResponse) ProtoMessage() {}

func (x *DeleteMemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemoResponse.ProtoReflect.Descriptor instead.
func (*DeleteMemoResponse) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteMemoResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_memo_proto protoreflect.FileDescriptor

var file_memo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x28, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x74, 0x6f, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x22, 0x6f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x22, 0x3d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4d,
	0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0xa7, 0x03, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f,
	0x12, 0x17, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x6f,
	0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x6d, 0x6f, 0x12, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x14, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x53,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x17, 0x2e, 0x6d,
	0x65, 0x6d, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x65, 0x6d,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x3a, 0x01, 0x2a, 0x32, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x56, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4f, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x42, 0x10, 0x5a, 0x0e,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memo_proto_rawDescOnce sync.Once
	file_memo_proto_rawDescData = file_memo_proto_rawDesc
)

func file_memo_proto_rawDescGZIP() []byte {
	file_memo_proto_rawDescOnce.Do(func() {
		file_memo_proto_rawDescData = protoimpl.X.CompressGZIP(file_memo_proto_rawDescData)
	})
	return file_memo_proto_rawDescData
}

var file_memo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_memo_proto_goTypes = []interface{}{
	(*CreateMemoRequest)(nil),  // 0: memo.CreateMemoRequest
	(*GetMemoRequest)(nil),     // 1: memo.GetMemoRequest
	(*ListMemosRequest)(nil),   // 2: memo.ListMemosRequest
	(*UpdateMemoRequest)(nil),  // 3: memo.UpdateMemoRequest
	(*DeleteMemoRequest)(nil),  // 4: memo.DeleteMemoRequest
	(*MemoResponse)(nil),       // 5: memo.MemoResponse
	(*ListMemosResponse)(nil),  // 6: memo.ListMemosResponse
	(*DeleteMemoResponse)(nil), // 7: memo.DeleteMemoResponse
}
var file_memo_proto_depIdxs = []int32{
	5, // 0: memo.ListMemosResponse.memos:type_name -> memo.MemoResponse
	0, // 1: memo.MemoService.CreateMemo:input_type -> memo.CreateMemoRequest
	1, // 2: memo.MemoService.GetMemo:input_type -> memo.GetMemoRequest
	3, // 3: memo.MemoService.UpdateMemo:input_type -> memo.UpdateMemoRequest
	4, // 4: memo.MemoService.DeleteMemo:input_type -> memo.DeleteMemoRequest
	2, // 5: memo.MemoService.ListMemos:input_type -> memo.ListMemosRequest
	5, // 6: memo.MemoService.CreateMemo:output_type -> memo.MemoResponse
	5, // 7: memo.MemoService.GetMemo:output_type -> memo.MemoResponse
	5, // 8: memo.MemoService.UpdateMemo:output_type -> memo.MemoResponse
	7, // 9: memo.MemoService.DeleteMemo:output_type -> memo.DeleteMemoResponse
	6, // 10: memo.MemoService.ListMemos:output_type -> memo.ListMemosResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_memo_proto_init() }
func file_memo_proto_init() {
	if File_memo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_memo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMemoRequest); i {
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
		file_memo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemoRequest); i {
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
		file_memo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMemosRequest); i {
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
		file_memo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMemoRequest); i {
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
		file_memo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemoRequest); i {
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
		file_memo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoResponse); i {
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
		file_memo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMemosResponse); i {
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
		file_memo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemoResponse); i {
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
			RawDescriptor: file_memo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memo_proto_goTypes,
		DependencyIndexes: file_memo_proto_depIdxs,
		MessageInfos:      file_memo_proto_msgTypes,
	}.Build()
	File_memo_proto = out.File
	file_memo_proto_rawDesc = nil
	file_memo_proto_goTypes = nil
	file_memo_proto_depIdxs = nil
}
