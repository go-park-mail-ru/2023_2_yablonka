// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: csat/api/csat_question.proto

package csat_microservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CSATQuestionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *CSATQuestionID) Reset() {
	*x = CSATQuestionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_api_csat_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSATQuestionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSATQuestionID) ProtoMessage() {}

func (x *CSATQuestionID) ProtoReflect() protoreflect.Message {
	mi := &file_csat_api_csat_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSATQuestionID.ProtoReflect.Descriptor instead.
func (*CSATQuestionID) Descriptor() ([]byte, []int) {
	return file_csat_api_csat_question_proto_rawDescGZIP(), []int{0}
}

func (x *CSATQuestionID) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type NewCSATQuestionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *NewCSATQuestionInfo) Reset() {
	*x = NewCSATQuestionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_api_csat_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCSATQuestionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCSATQuestionInfo) ProtoMessage() {}

func (x *NewCSATQuestionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_csat_api_csat_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCSATQuestionInfo.ProtoReflect.Descriptor instead.
func (*NewCSATQuestionInfo) Descriptor() ([]byte, []int) {
	return file_csat_api_csat_question_proto_rawDescGZIP(), []int{1}
}

func (x *NewCSATQuestionInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NewCSATQuestionInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type NewCSATAnswerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionID uint64 `protobuf:"varint,1,opt,name=QuestionID,proto3" json:"QuestionID,omitempty"`
	Rating     uint64 `protobuf:"varint,2,opt,name=Rating,proto3" json:"Rating,omitempty"`
}

func (x *NewCSATAnswerInfo) Reset() {
	*x = NewCSATAnswerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_api_csat_question_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCSATAnswerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCSATAnswerInfo) ProtoMessage() {}

func (x *NewCSATAnswerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_csat_api_csat_question_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCSATAnswerInfo.ProtoReflect.Descriptor instead.
func (*NewCSATAnswerInfo) Descriptor() ([]byte, []int) {
	return file_csat_api_csat_question_proto_rawDescGZIP(), []int{2}
}

func (x *NewCSATAnswerInfo) GetQuestionID() uint64 {
	if x != nil {
		return x.QuestionID
	}
	return 0
}

func (x *NewCSATAnswerInfo) GetRating() uint64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type CSATQuestionFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *CSATQuestionFull) Reset() {
	*x = CSATQuestionFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_api_csat_question_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSATQuestionFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSATQuestionFull) ProtoMessage() {}

func (x *CSATQuestionFull) ProtoReflect() protoreflect.Message {
	mi := &file_csat_api_csat_question_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSATQuestionFull.ProtoReflect.Descriptor instead.
func (*CSATQuestionFull) Descriptor() ([]byte, []int) {
	return file_csat_api_csat_question_proto_rawDescGZIP(), []int{3}
}

func (x *CSATQuestionFull) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CSATQuestionFull) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CSATQuestionFull) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdatedCSATQuestionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Type    string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *UpdatedCSATQuestionInfo) Reset() {
	*x = UpdatedCSATQuestionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_api_csat_question_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedCSATQuestionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedCSATQuestionInfo) ProtoMessage() {}

func (x *UpdatedCSATQuestionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_csat_api_csat_question_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedCSATQuestionInfo.ProtoReflect.Descriptor instead.
func (*UpdatedCSATQuestionInfo) Descriptor() ([]byte, []int) {
	return file_csat_api_csat_question_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatedCSATQuestionInfo) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdatedCSATQuestionInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdatedCSATQuestionInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AllQuestionStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*CSATQuestionFull `protobuf:"bytes,1,rep,name=Questions,proto3" json:"Questions,omitempty"`
}

func (x *AllQuestionStats) Reset() {
	*x = AllQuestionStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_api_csat_question_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllQuestionStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllQuestionStats) ProtoMessage() {}

func (x *AllQuestionStats) ProtoReflect() protoreflect.Message {
	mi := &file_csat_api_csat_question_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllQuestionStats.ProtoReflect.Descriptor instead.
func (*AllQuestionStats) Descriptor() ([]byte, []int) {
	return file_csat_api_csat_question_proto_rawDescGZIP(), []int{5}
}

func (x *AllQuestionStats) GetQuestions() []*CSATQuestionFull {
	if x != nil {
		return x.Questions
	}
	return nil
}

type RatingStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating  uint64 `protobuf:"varint,1,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Count   uint64 `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	Average uint64 `protobuf:"varint,3,opt,name=Average,proto3" json:"Average,omitempty"`
}

func (x *RatingStats) Reset() {
	*x = RatingStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_api_csat_question_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingStats) ProtoMessage() {}

func (x *RatingStats) ProtoReflect() protoreflect.Message {
	mi := &file_csat_api_csat_question_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingStats.ProtoReflect.Descriptor instead.
func (*RatingStats) Descriptor() ([]byte, []int) {
	return file_csat_api_csat_question_proto_rawDescGZIP(), []int{6}
}

func (x *RatingStats) GetRating() uint64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *RatingStats) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RatingStats) GetAverage() uint64 {
	if x != nil {
		return x.Average
	}
	return 0
}

type QuestionWithStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint64         `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Content string         `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Type    string         `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Stats   []*RatingStats `protobuf:"bytes,4,rep,name=Stats,proto3" json:"Stats,omitempty"`
}

func (x *QuestionWithStats) Reset() {
	*x = QuestionWithStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_api_csat_question_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionWithStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionWithStats) ProtoMessage() {}

func (x *QuestionWithStats) ProtoReflect() protoreflect.Message {
	mi := &file_csat_api_csat_question_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionWithStats.ProtoReflect.Descriptor instead.
func (*QuestionWithStats) Descriptor() ([]byte, []int) {
	return file_csat_api_csat_question_proto_rawDescGZIP(), []int{7}
}

func (x *QuestionWithStats) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *QuestionWithStats) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *QuestionWithStats) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *QuestionWithStats) GetStats() []*RatingStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type AllQuestionsWithStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*QuestionWithStats `protobuf:"bytes,1,rep,name=Questions,proto3" json:"Questions,omitempty"`
}

func (x *AllQuestionsWithStats) Reset() {
	*x = AllQuestionsWithStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_api_csat_question_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllQuestionsWithStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllQuestionsWithStats) ProtoMessage() {}

func (x *AllQuestionsWithStats) ProtoReflect() protoreflect.Message {
	mi := &file_csat_api_csat_question_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllQuestionsWithStats.ProtoReflect.Descriptor instead.
func (*AllQuestionsWithStats) Descriptor() ([]byte, []int) {
	return file_csat_api_csat_question_proto_rawDescGZIP(), []int{8}
}

func (x *AllQuestionsWithStats) GetQuestions() []*QuestionWithStats {
	if x != nil {
		return x.Questions
	}
	return nil
}

var File_csat_api_csat_question_proto protoreflect.FileDescriptor

var file_csat_api_csat_question_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x73, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x73, 0x61, 0x74, 0x5f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e, 0x43,
	0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x43,
	0x53, 0x41, 0x54, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a,
	0x0a, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x50, 0x0a, 0x10, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x43, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x09, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x55, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x75, 0x0a, 0x11,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x22, 0x49, 0x0a, 0x15, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x09,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x09, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xef,
	0x02, 0x0a, 0x13, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x53, 0x41, 0x54, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x69, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x43, 0x53, 0x41, 0x54, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x63, 0x73, 0x61, 0x74, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_csat_api_csat_question_proto_rawDescOnce sync.Once
	file_csat_api_csat_question_proto_rawDescData = file_csat_api_csat_question_proto_rawDesc
)

func file_csat_api_csat_question_proto_rawDescGZIP() []byte {
	file_csat_api_csat_question_proto_rawDescOnce.Do(func() {
		file_csat_api_csat_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_csat_api_csat_question_proto_rawDescData)
	})
	return file_csat_api_csat_question_proto_rawDescData
}

var file_csat_api_csat_question_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_csat_api_csat_question_proto_goTypes = []interface{}{
	(*CSATQuestionID)(nil),          // 0: CSATQuestionID
	(*NewCSATQuestionInfo)(nil),     // 1: NewCSATQuestionInfo
	(*NewCSATAnswerInfo)(nil),       // 2: NewCSATAnswerInfo
	(*CSATQuestionFull)(nil),        // 3: CSATQuestionFull
	(*UpdatedCSATQuestionInfo)(nil), // 4: UpdatedCSATQuestionInfo
	(*AllQuestionStats)(nil),        // 5: AllQuestionStats
	(*RatingStats)(nil),             // 6: RatingStats
	(*QuestionWithStats)(nil),       // 7: QuestionWithStats
	(*AllQuestionsWithStats)(nil),   // 8: AllQuestionsWithStats
	(*emptypb.Empty)(nil),           // 9: google.protobuf.Empty
}
var file_csat_api_csat_question_proto_depIdxs = []int32{
	3, // 0: AllQuestionStats.Questions:type_name -> CSATQuestionFull
	6, // 1: QuestionWithStats.Stats:type_name -> RatingStats
	7, // 2: AllQuestionsWithStats.Questions:type_name -> QuestionWithStats
	2, // 3: CSATQuestionService.CheckRating:input_type -> NewCSATAnswerInfo
	9, // 4: CSATQuestionService.GetAll:input_type -> google.protobuf.Empty
	9, // 5: CSATQuestionService.GetStats:input_type -> google.protobuf.Empty
	1, // 6: CSATQuestionService.Create:input_type -> NewCSATQuestionInfo
	4, // 7: CSATQuestionService.Update:input_type -> UpdatedCSATQuestionInfo
	0, // 8: CSATQuestionService.Delete:input_type -> CSATQuestionID
	9, // 9: CSATQuestionService.CheckRating:output_type -> google.protobuf.Empty
	5, // 10: CSATQuestionService.GetAll:output_type -> AllQuestionStats
	8, // 11: CSATQuestionService.GetStats:output_type -> AllQuestionsWithStats
	3, // 12: CSATQuestionService.Create:output_type -> CSATQuestionFull
	9, // 13: CSATQuestionService.Update:output_type -> google.protobuf.Empty
	9, // 14: CSATQuestionService.Delete:output_type -> google.protobuf.Empty
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_csat_api_csat_question_proto_init() }
func file_csat_api_csat_question_proto_init() {
	if File_csat_api_csat_question_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_csat_api_csat_question_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSATQuestionID); i {
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
		file_csat_api_csat_question_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCSATQuestionInfo); i {
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
		file_csat_api_csat_question_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCSATAnswerInfo); i {
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
		file_csat_api_csat_question_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSATQuestionFull); i {
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
		file_csat_api_csat_question_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedCSATQuestionInfo); i {
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
		file_csat_api_csat_question_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllQuestionStats); i {
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
		file_csat_api_csat_question_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingStats); i {
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
		file_csat_api_csat_question_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionWithStats); i {
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
		file_csat_api_csat_question_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllQuestionsWithStats); i {
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
			RawDescriptor: file_csat_api_csat_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_csat_api_csat_question_proto_goTypes,
		DependencyIndexes: file_csat_api_csat_question_proto_depIdxs,
		MessageInfos:      file_csat_api_csat_question_proto_msgTypes,
	}.Build()
	File_csat_api_csat_question_proto = out.File
	file_csat_api_csat_question_proto_rawDesc = nil
	file_csat_api_csat_question_proto_goTypes = nil
	file_csat_api_csat_question_proto_depIdxs = nil
}
