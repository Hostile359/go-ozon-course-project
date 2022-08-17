// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: comment.proto

package api

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

type CommentAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	UserId  uint64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CommentAddRequest) Reset() {
	*x = CommentAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentAddRequest) ProtoMessage() {}

func (x *CommentAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentAddRequest.ProtoReflect.Descriptor instead.
func (*CommentAddRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentAddRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CommentAddRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CommentAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommentAddResponse) Reset() {
	*x = CommentAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentAddResponse) ProtoMessage() {}

func (x *CommentAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentAddResponse.ProtoReflect.Descriptor instead.
func (*CommentAddResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{1}
}

type CommentGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CommentGetRequest) Reset() {
	*x = CommentGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentGetRequest) ProtoMessage() {}

func (x *CommentGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentGetRequest.ProtoReflect.Descriptor instead.
func (*CommentGetRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentGetRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CommentGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *CommentGetResponse_Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentGetResponse) Reset() {
	*x = CommentGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentGetResponse) ProtoMessage() {}

func (x *CommentGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentGetResponse.ProtoReflect.Descriptor instead.
func (*CommentGetResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CommentGetResponse) GetComment() *CommentGetResponse_Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage uint64 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *CommentListRequest) Reset() {
	*x = CommentListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListRequest) ProtoMessage() {}

func (x *CommentListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListRequest.ProtoReflect.Descriptor instead.
func (*CommentListRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{4}
}

func (x *CommentListRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CommentListRequest) GetPerPage() uint64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type CommentListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*CommentListResponse_Comment `protobuf:"bytes,1,rep,name=Comments,proto3" json:"Comments,omitempty"`
}

func (x *CommentListResponse) Reset() {
	*x = CommentListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListResponse) ProtoMessage() {}

func (x *CommentListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListResponse.ProtoReflect.Descriptor instead.
func (*CommentListResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{5}
}

func (x *CommentListResponse) GetComments() []*CommentListResponse_Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type CommentUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	UserId  uint64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CommentUpdateRequest) Reset() {
	*x = CommentUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentUpdateRequest) ProtoMessage() {}

func (x *CommentUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentUpdateRequest.ProtoReflect.Descriptor instead.
func (*CommentUpdateRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{6}
}

func (x *CommentUpdateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentUpdateRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CommentUpdateRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CommentUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommentUpdateResponse) Reset() {
	*x = CommentUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentUpdateResponse) ProtoMessage() {}

func (x *CommentUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentUpdateResponse.ProtoReflect.Descriptor instead.
func (*CommentUpdateResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{7}
}

type CommentDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CommentDeleteRequest) Reset() {
	*x = CommentDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentDeleteRequest) ProtoMessage() {}

func (x *CommentDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentDeleteRequest.ProtoReflect.Descriptor instead.
func (*CommentDeleteRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{8}
}

func (x *CommentDeleteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CommentDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommentDeleteResponse) Reset() {
	*x = CommentDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentDeleteResponse) ProtoMessage() {}

func (x *CommentDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentDeleteResponse.ProtoReflect.Descriptor instead.
func (*CommentDeleteResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{9}
}

type CommentGetResponse_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	UserId  uint64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CommentGetResponse_Comment) Reset() {
	*x = CommentGetResponse_Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentGetResponse_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentGetResponse_Comment) ProtoMessage() {}

func (x *CommentGetResponse_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentGetResponse_Comment.ProtoReflect.Descriptor instead.
func (*CommentGetResponse_Comment) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CommentGetResponse_Comment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentGetResponse_Comment) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CommentGetResponse_Comment) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CommentListResponse_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	UserId  uint64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CommentListResponse_Comment) Reset() {
	*x = CommentListResponse_Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListResponse_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListResponse_Comment) ProtoMessage() {}

func (x *CommentListResponse_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListResponse_Comment.ProtoReflect.Descriptor instead.
func (*CommentListResponse_Comment) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{5, 0}
}

func (x *CommentListResponse_Comment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentListResponse_Comment) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CommentListResponse_Comment) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_comment_proto protoreflect.FileDescriptor

var file_comment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a,
	0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4c, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x4c, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x12, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65,
	0x22, 0xb4, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x31, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x4c, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x91, 0x05, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x7b, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x12, 0x29, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x7d, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x47, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7c, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x1a, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x84, 0x01, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x31,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x31, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x2a, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x48, 0x6f, 0x73, 0x74, 0x69, 0x6c, 0x65, 0x33, 0x35, 0x39, 0x2f, 0x68,
	0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x31, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_proto_rawDescOnce sync.Once
	file_comment_proto_rawDescData = file_comment_proto_rawDesc
)

func file_comment_proto_rawDescGZIP() []byte {
	file_comment_proto_rawDescOnce.Do(func() {
		file_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_proto_rawDescData)
	})
	return file_comment_proto_rawDescData
}

var file_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_comment_proto_goTypes = []interface{}{
	(*CommentAddRequest)(nil),           // 0: ozon.dev.homework1.api.CommentAddRequest
	(*CommentAddResponse)(nil),          // 1: ozon.dev.homework1.api.CommentAddResponse
	(*CommentGetRequest)(nil),           // 2: ozon.dev.homework1.api.CommentGetRequest
	(*CommentGetResponse)(nil),          // 3: ozon.dev.homework1.api.CommentGetResponse
	(*CommentListRequest)(nil),          // 4: ozon.dev.homework1.api.CommentListRequest
	(*CommentListResponse)(nil),         // 5: ozon.dev.homework1.api.CommentListResponse
	(*CommentUpdateRequest)(nil),        // 6: ozon.dev.homework1.api.CommentUpdateRequest
	(*CommentUpdateResponse)(nil),       // 7: ozon.dev.homework1.api.CommentUpdateResponse
	(*CommentDeleteRequest)(nil),        // 8: ozon.dev.homework1.api.CommentDeleteRequest
	(*CommentDeleteResponse)(nil),       // 9: ozon.dev.homework1.api.CommentDeleteResponse
	(*CommentGetResponse_Comment)(nil),  // 10: ozon.dev.homework1.api.CommentGetResponse.Comment
	(*CommentListResponse_Comment)(nil), // 11: ozon.dev.homework1.api.CommentListResponse.Comment
}
var file_comment_proto_depIdxs = []int32{
	10, // 0: ozon.dev.homework1.api.CommentGetResponse.comment:type_name -> ozon.dev.homework1.api.CommentGetResponse.Comment
	11, // 1: ozon.dev.homework1.api.CommentListResponse.Comments:type_name -> ozon.dev.homework1.api.CommentListResponse.Comment
	0,  // 2: ozon.dev.homework1.api.Comment.CommentAdd:input_type -> ozon.dev.homework1.api.CommentAddRequest
	2,  // 3: ozon.dev.homework1.api.Comment.CommentGet:input_type -> ozon.dev.homework1.api.CommentGetRequest
	4,  // 4: ozon.dev.homework1.api.Comment.CommentList:input_type -> ozon.dev.homework1.api.CommentListRequest
	6,  // 5: ozon.dev.homework1.api.Comment.CommentUpdate:input_type -> ozon.dev.homework1.api.CommentUpdateRequest
	8,  // 6: ozon.dev.homework1.api.Comment.CommentDelete:input_type -> ozon.dev.homework1.api.CommentDeleteRequest
	1,  // 7: ozon.dev.homework1.api.Comment.CommentAdd:output_type -> ozon.dev.homework1.api.CommentAddResponse
	3,  // 8: ozon.dev.homework1.api.Comment.CommentGet:output_type -> ozon.dev.homework1.api.CommentGetResponse
	5,  // 9: ozon.dev.homework1.api.Comment.CommentList:output_type -> ozon.dev.homework1.api.CommentListResponse
	7,  // 10: ozon.dev.homework1.api.Comment.CommentUpdate:output_type -> ozon.dev.homework1.api.CommentUpdateResponse
	9,  // 11: ozon.dev.homework1.api.Comment.CommentDelete:output_type -> ozon.dev.homework1.api.CommentDeleteResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_comment_proto_init() }
func file_comment_proto_init() {
	if File_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentAddRequest); i {
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
		file_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentAddResponse); i {
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
		file_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentGetRequest); i {
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
		file_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentGetResponse); i {
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
		file_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListRequest); i {
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
		file_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListResponse); i {
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
		file_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentUpdateRequest); i {
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
		file_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentUpdateResponse); i {
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
		file_comment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentDeleteRequest); i {
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
		file_comment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentDeleteResponse); i {
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
		file_comment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentGetResponse_Comment); i {
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
		file_comment_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListResponse_Comment); i {
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
			RawDescriptor: file_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_proto_goTypes,
		DependencyIndexes: file_comment_proto_depIdxs,
		MessageInfos:      file_comment_proto_msgTypes,
	}.Build()
	File_comment_proto = out.File
	file_comment_proto_rawDesc = nil
	file_comment_proto_goTypes = nil
	file_comment_proto_depIdxs = nil
}