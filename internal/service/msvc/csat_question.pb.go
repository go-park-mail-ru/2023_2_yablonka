// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: csat_question.proto

package microservice

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

type CSATQuestionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *CSATQuestionID) Reset() {
	*x = CSATQuestionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSATQuestionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSATQuestionID) ProtoMessage() {}

func (x *CSATQuestionID) ProtoReflect() protoreflect.Message {
	mi := &file_csat_question_proto_msgTypes[0]
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
	return file_csat_question_proto_rawDescGZIP(), []int{0}
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
		mi := &file_csat_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCSATQuestionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCSATQuestionInfo) ProtoMessage() {}

func (x *NewCSATQuestionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_csat_question_proto_msgTypes[1]
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
	return file_csat_question_proto_rawDescGZIP(), []int{1}
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
		mi := &file_csat_question_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCSATAnswerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCSATAnswerInfo) ProtoMessage() {}

func (x *NewCSATAnswerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_csat_question_proto_msgTypes[2]
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
	return file_csat_question_proto_rawDescGZIP(), []int{2}
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

type CSATQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *CSATQuestion) Reset() {
	*x = CSATQuestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_question_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSATQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSATQuestion) ProtoMessage() {}

func (x *CSATQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_csat_question_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSATQuestion.ProtoReflect.Descriptor instead.
func (*CSATQuestion) Descriptor() ([]byte, []int) {
	return file_csat_question_proto_rawDescGZIP(), []int{3}
}

func (x *CSATQuestion) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CSATQuestion) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CSATQuestion) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdatedCSATQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *UpdatedCSATQuestion) Reset() {
	*x = UpdatedCSATQuestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_question_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedCSATQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedCSATQuestion) ProtoMessage() {}

func (x *UpdatedCSATQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_csat_question_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedCSATQuestion.ProtoReflect.Descriptor instead.
func (*UpdatedCSATQuestion) Descriptor() ([]byte, []int) {
	return file_csat_question_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatedCSATQuestion) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdatedCSATQuestion) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CreateReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question *CSATQuestion `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	Error    *Error        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateReturn) Reset() {
	*x = CreateReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csat_question_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReturn) ProtoMessage() {}

func (x *CreateReturn) ProtoReflect() protoreflect.Message {
	mi := &file_csat_question_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReturn.ProtoReflect.Descriptor instead.
func (*CreateReturn) Descriptor() ([]byte, []int) {
	return file_csat_question_proto_rawDescGZIP(), []int{5}
}

func (x *CreateReturn) GetQuestion() *CSATQuestion {
	if x != nil {
		return x.Question
	}
	return nil
}

func (x *CreateReturn) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_csat_question_proto protoreflect.FileDescriptor

var file_csat_question_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x73, 0x61, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e,
	0x43, 0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x43, 0x53, 0x41, 0x54, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x11, 0x4e, 0x65, 0x77,
	0x43, 0x53, 0x41, 0x54, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e,
	0x0a, 0x0a, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x4c, 0x0a, 0x0c, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x43,
	0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5f, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x31, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xea, 0x01, 0x0a, 0x13, 0x43,
	0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x43,
	0x53, 0x41, 0x54, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x06, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x43,
	0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x43, 0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x53, 0x41, 0x54, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x06, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x3b, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_csat_question_proto_rawDescOnce sync.Once
	file_csat_question_proto_rawDescData = file_csat_question_proto_rawDesc
)

func file_csat_question_proto_rawDescGZIP() []byte {
	file_csat_question_proto_rawDescOnce.Do(func() {
		file_csat_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_csat_question_proto_rawDescData)
	})
	return file_csat_question_proto_rawDescData
}

var file_csat_question_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_csat_question_proto_goTypes = []interface{}{
	(*CSATQuestionID)(nil),      // 0: service.CSATQuestionID
	(*NewCSATQuestionInfo)(nil), // 1: service.NewCSATQuestionInfo
	(*NewCSATAnswerInfo)(nil),   // 2: service.NewCSATAnswerInfo
	(*CSATQuestion)(nil),        // 3: service.CSATQuestion
	(*UpdatedCSATQuestion)(nil), // 4: service.UpdatedCSATQuestion
	(*CreateReturn)(nil),        // 5: service.CreateReturn
	(*Error)(nil),               // 6: Error
}
var file_csat_question_proto_depIdxs = []int32{
	3, // 0: service.CreateReturn.question:type_name -> service.CSATQuestion
	6, // 1: service.CreateReturn.error:type_name -> Error
	2, // 2: service.CSATQuestionService.CheckRating:input_type -> service.NewCSATAnswerInfo
	1, // 3: service.CSATQuestionService.Create:input_type -> service.NewCSATQuestionInfo
	4, // 4: service.CSATQuestionService.Update:input_type -> service.UpdatedCSATQuestion
	0, // 5: service.CSATQuestionService.Delete:input_type -> service.CSATQuestionID
	6, // 6: service.CSATQuestionService.CheckRating:output_type -> Error
	5, // 7: service.CSATQuestionService.Create:output_type -> service.CreateReturn
	6, // 8: service.CSATQuestionService.Update:output_type -> Error
	6, // 9: service.CSATQuestionService.Delete:output_type -> Error
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_csat_question_proto_init() }
func file_csat_question_proto_init() {
	if File_csat_question_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_csat_question_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_csat_question_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_csat_question_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_csat_question_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSATQuestion); i {
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
		file_csat_question_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedCSATQuestion); i {
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
		file_csat_question_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReturn); i {
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
			RawDescriptor: file_csat_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_csat_question_proto_goTypes,
		DependencyIndexes: file_csat_question_proto_depIdxs,
		MessageInfos:      file_csat_question_proto_msgTypes,
	}.Build()
	File_csat_question_proto = out.File
	file_csat_question_proto_rawDesc = nil
	file_csat_question_proto_goTypes = nil
	file_csat_question_proto_depIdxs = nil
}
