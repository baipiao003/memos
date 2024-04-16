// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v2/tag_service.proto

package apiv2

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

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The creator of tags.
	// Format: users/{id}
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

type UpsertTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpsertTagRequest) Reset() {
	*x = UpsertTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertTagRequest) ProtoMessage() {}

func (x *UpsertTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertTagRequest.ProtoReflect.Descriptor instead.
func (*UpsertTagRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpsertTagRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpsertTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *UpsertTagResponse) Reset() {
	*x = UpsertTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertTagResponse) ProtoMessage() {}

func (x *UpsertTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertTagResponse.ProtoReflect.Descriptor instead.
func (*UpsertTagResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertTagResponse) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type BatchUpsertTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*UpsertTagRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *BatchUpsertTagRequest) Reset() {
	*x = BatchUpsertTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchUpsertTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUpsertTagRequest) ProtoMessage() {}

func (x *BatchUpsertTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUpsertTagRequest.ProtoReflect.Descriptor instead.
func (*BatchUpsertTagRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{3}
}

func (x *BatchUpsertTagRequest) GetRequests() []*UpsertTagRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

type BatchUpsertTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BatchUpsertTagResponse) Reset() {
	*x = BatchUpsertTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchUpsertTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUpsertTagResponse) ProtoMessage() {}

func (x *BatchUpsertTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUpsertTagResponse.ProtoReflect.Descriptor instead.
func (*BatchUpsertTagResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{4}
}

type ListTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTagsRequest) Reset() {
	*x = ListTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsRequest) ProtoMessage() {}

func (x *ListTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsRequest.ProtoReflect.Descriptor instead.
func (*ListTagsRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{5}
}

type ListTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ListTagsResponse) Reset() {
	*x = ListTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsResponse) ProtoMessage() {}

func (x *ListTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsResponse.ProtoReflect.Descriptor instead.
func (*ListTagsResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListTagsResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type RenameTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The creator of tags.
	// Format: users/{id}
	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	OldName string `protobuf:"bytes,2,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	NewName string `protobuf:"bytes,3,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
}

func (x *RenameTagRequest) Reset() {
	*x = RenameTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameTagRequest) ProtoMessage() {}

func (x *RenameTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameTagRequest.ProtoReflect.Descriptor instead.
func (*RenameTagRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{7}
}

func (x *RenameTagRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RenameTagRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *RenameTagRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type RenameTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *RenameTagResponse) Reset() {
	*x = RenameTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameTagResponse) ProtoMessage() {}

func (x *RenameTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameTagResponse.ProtoReflect.Descriptor instead.
func (*RenameTagResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{8}
}

func (x *RenameTagResponse) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type DeleteTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *DeleteTagRequest) Reset() {
	*x = DeleteTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagRequest) ProtoMessage() {}

func (x *DeleteTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTagRequest) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type DeleteTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTagResponse) Reset() {
	*x = DeleteTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagResponse) ProtoMessage() {}

func (x *DeleteTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagResponse.ProtoReflect.Descriptor instead.
func (*DeleteTagResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{10}
}

type GetTagSuggestionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The creator of tags.
	// Format: users/{id}
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetTagSuggestionsRequest) Reset() {
	*x = GetTagSuggestionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagSuggestionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagSuggestionsRequest) ProtoMessage() {}

func (x *GetTagSuggestionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagSuggestionsRequest.ProtoReflect.Descriptor instead.
func (*GetTagSuggestionsRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{11}
}

func (x *GetTagSuggestionsRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type GetTagSuggestionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GetTagSuggestionsResponse) Reset() {
	*x = GetTagSuggestionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_tag_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagSuggestionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagSuggestionsResponse) ProtoMessage() {}

func (x *GetTagSuggestionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_tag_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagSuggestionsResponse.ProtoReflect.Descriptor instead.
func (*GetTagSuggestionsResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_tag_service_proto_rawDescGZIP(), []int{12}
}

func (x *GetTagSuggestionsResponse) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_api_v2_tag_service_proto protoreflect.FileDescriptor

var file_api_v2_tag_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x26, 0x0a, 0x10, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x11, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x53, 0x0a,
	0x15, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x39, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x5c, 0x0a, 0x10, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x11, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x22, 0x37, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x13, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x2f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x32, 0xb1, 0x05, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x65, 0x0a, 0x09, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1e, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x0e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x67, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a,
	0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12, 0x5f, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x6c, 0x0a, 0x09, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x3a, 0x01, 0x2a, 0x32, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61,
	0x67, 0x73, 0x3a, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x62, 0x0a, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a,
	0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x85, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x67, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xa7, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x0f, 0x54, 0x61, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x32, 0xa2,
	0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa, 0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x70,
	0x69, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x32, 0xe2, 0x02, 0x18, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0e, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v2_tag_service_proto_rawDescOnce sync.Once
	file_api_v2_tag_service_proto_rawDescData = file_api_v2_tag_service_proto_rawDesc
)

func file_api_v2_tag_service_proto_rawDescGZIP() []byte {
	file_api_v2_tag_service_proto_rawDescOnce.Do(func() {
		file_api_v2_tag_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_tag_service_proto_rawDescData)
	})
	return file_api_v2_tag_service_proto_rawDescData
}

var file_api_v2_tag_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_v2_tag_service_proto_goTypes = []interface{}{
	(*Tag)(nil),                       // 0: memos.api.v2.Tag
	(*UpsertTagRequest)(nil),          // 1: memos.api.v2.UpsertTagRequest
	(*UpsertTagResponse)(nil),         // 2: memos.api.v2.UpsertTagResponse
	(*BatchUpsertTagRequest)(nil),     // 3: memos.api.v2.BatchUpsertTagRequest
	(*BatchUpsertTagResponse)(nil),    // 4: memos.api.v2.BatchUpsertTagResponse
	(*ListTagsRequest)(nil),           // 5: memos.api.v2.ListTagsRequest
	(*ListTagsResponse)(nil),          // 6: memos.api.v2.ListTagsResponse
	(*RenameTagRequest)(nil),          // 7: memos.api.v2.RenameTagRequest
	(*RenameTagResponse)(nil),         // 8: memos.api.v2.RenameTagResponse
	(*DeleteTagRequest)(nil),          // 9: memos.api.v2.DeleteTagRequest
	(*DeleteTagResponse)(nil),         // 10: memos.api.v2.DeleteTagResponse
	(*GetTagSuggestionsRequest)(nil),  // 11: memos.api.v2.GetTagSuggestionsRequest
	(*GetTagSuggestionsResponse)(nil), // 12: memos.api.v2.GetTagSuggestionsResponse
}
var file_api_v2_tag_service_proto_depIdxs = []int32{
	0,  // 0: memos.api.v2.UpsertTagResponse.tag:type_name -> memos.api.v2.Tag
	1,  // 1: memos.api.v2.BatchUpsertTagRequest.requests:type_name -> memos.api.v2.UpsertTagRequest
	0,  // 2: memos.api.v2.ListTagsResponse.tags:type_name -> memos.api.v2.Tag
	0,  // 3: memos.api.v2.RenameTagResponse.tag:type_name -> memos.api.v2.Tag
	0,  // 4: memos.api.v2.DeleteTagRequest.tag:type_name -> memos.api.v2.Tag
	1,  // 5: memos.api.v2.TagService.UpsertTag:input_type -> memos.api.v2.UpsertTagRequest
	3,  // 6: memos.api.v2.TagService.BatchUpsertTag:input_type -> memos.api.v2.BatchUpsertTagRequest
	5,  // 7: memos.api.v2.TagService.ListTags:input_type -> memos.api.v2.ListTagsRequest
	7,  // 8: memos.api.v2.TagService.RenameTag:input_type -> memos.api.v2.RenameTagRequest
	9,  // 9: memos.api.v2.TagService.DeleteTag:input_type -> memos.api.v2.DeleteTagRequest
	11, // 10: memos.api.v2.TagService.GetTagSuggestions:input_type -> memos.api.v2.GetTagSuggestionsRequest
	2,  // 11: memos.api.v2.TagService.UpsertTag:output_type -> memos.api.v2.UpsertTagResponse
	4,  // 12: memos.api.v2.TagService.BatchUpsertTag:output_type -> memos.api.v2.BatchUpsertTagResponse
	6,  // 13: memos.api.v2.TagService.ListTags:output_type -> memos.api.v2.ListTagsResponse
	8,  // 14: memos.api.v2.TagService.RenameTag:output_type -> memos.api.v2.RenameTagResponse
	10, // 15: memos.api.v2.TagService.DeleteTag:output_type -> memos.api.v2.DeleteTagResponse
	12, // 16: memos.api.v2.TagService.GetTagSuggestions:output_type -> memos.api.v2.GetTagSuggestionsResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_v2_tag_service_proto_init() }
func file_api_v2_tag_service_proto_init() {
	if File_api_v2_tag_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v2_tag_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_api_v2_tag_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertTagRequest); i {
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
		file_api_v2_tag_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertTagResponse); i {
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
		file_api_v2_tag_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchUpsertTagRequest); i {
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
		file_api_v2_tag_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchUpsertTagResponse); i {
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
		file_api_v2_tag_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsRequest); i {
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
		file_api_v2_tag_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsResponse); i {
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
		file_api_v2_tag_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameTagRequest); i {
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
		file_api_v2_tag_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameTagResponse); i {
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
		file_api_v2_tag_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagRequest); i {
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
		file_api_v2_tag_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagResponse); i {
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
		file_api_v2_tag_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagSuggestionsRequest); i {
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
		file_api_v2_tag_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagSuggestionsResponse); i {
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
			RawDescriptor: file_api_v2_tag_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_tag_service_proto_goTypes,
		DependencyIndexes: file_api_v2_tag_service_proto_depIdxs,
		MessageInfos:      file_api_v2_tag_service_proto_msgTypes,
	}.Build()
	File_api_v2_tag_service_proto = out.File
	file_api_v2_tag_service_proto_rawDesc = nil
	file_api_v2_tag_service_proto_goTypes = nil
	file_api_v2_tag_service_proto_depIdxs = nil
}
