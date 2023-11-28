// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: csrf/api/csrf.proto

package csrf_microservice

import (
	common "./common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CSRFToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *CSRFToken) Reset() {
	*x = CSRFToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csrf_api_csrf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSRFToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSRFToken) ProtoMessage() {}

func (x *CSRFToken) ProtoReflect() protoreflect.Message {
	mi := &file_csrf_api_csrf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSRFToken.ProtoReflect.Descriptor instead.
func (*CSRFToken) Descriptor() ([]byte, []int) {
	return file_csrf_api_csrf_proto_rawDescGZIP(), []int{0}
}

func (x *CSRFToken) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CSRFData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ExpirationDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=ExpirationDate,proto3" json:"ExpirationDate,omitempty"`
}

func (x *CSRFData) Reset() {
	*x = CSRFData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csrf_api_csrf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSRFData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSRFData) ProtoMessage() {}

func (x *CSRFData) ProtoReflect() protoreflect.Message {
	mi := &file_csrf_api_csrf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSRFData.ProtoReflect.Descriptor instead.
func (*CSRFData) Descriptor() ([]byte, []int) {
	return file_csrf_api_csrf_proto_rawDescGZIP(), []int{1}
}

func (x *CSRFData) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CSRFData) GetExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDate
	}
	return nil
}

var File_csrf_api_csrf_proto protoreflect.FileDescriptor

var file_csrf_api_csrf_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x73, 0x72, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x73, 0x72, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x09, 0x43, 0x53,
	0x52, 0x46, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5e, 0x0a,
	0x08, 0x43, 0x53, 0x52, 0x46, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x0e, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x32, 0xdc, 0x01,
	0x0a, 0x0b, 0x43, 0x53, 0x52, 0x46, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a,
	0x09, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43, 0x53, 0x52, 0x46, 0x12, 0x07, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x43, 0x53, 0x52, 0x46, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x53, 0x52, 0x46, 0x12, 0x0a,
	0x2e, 0x43, 0x53, 0x52, 0x46, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x53,
	0x52, 0x46, 0x12, 0x0a, 0x2e, 0x43, 0x53, 0x52, 0x46, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13,
	0x2e, 0x2f, 0x63, 0x73, 0x72, 0x66, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_csrf_api_csrf_proto_rawDescOnce sync.Once
	file_csrf_api_csrf_proto_rawDescData = file_csrf_api_csrf_proto_rawDesc
)

func file_csrf_api_csrf_proto_rawDescGZIP() []byte {
	file_csrf_api_csrf_proto_rawDescOnce.Do(func() {
		file_csrf_api_csrf_proto_rawDescData = protoimpl.X.CompressGZIP(file_csrf_api_csrf_proto_rawDescData)
	})
	return file_csrf_api_csrf_proto_rawDescData
}

var file_csrf_api_csrf_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_csrf_api_csrf_proto_goTypes = []interface{}{
	(*CSRFToken)(nil),             // 0: CSRFToken
	(*CSRFData)(nil),              // 1: CSRFData
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*common.UserID)(nil),         // 3: UserID
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
	(*durationpb.Duration)(nil),   // 5: google.protobuf.Duration
}
var file_csrf_api_csrf_proto_depIdxs = []int32{
	2, // 0: CSRFData.ExpirationDate:type_name -> google.protobuf.Timestamp
	3, // 1: CSRFService.SetupCSRF:input_type -> UserID
	0, // 2: CSRFService.VerifyCSRF:input_type -> CSRFToken
	0, // 3: CSRFService.DeleteCSRF:input_type -> CSRFToken
	4, // 4: CSRFService.GetLifetime:input_type -> google.protobuf.Empty
	1, // 5: CSRFService.SetupCSRF:output_type -> CSRFData
	4, // 6: CSRFService.VerifyCSRF:output_type -> google.protobuf.Empty
	4, // 7: CSRFService.DeleteCSRF:output_type -> google.protobuf.Empty
	5, // 8: CSRFService.GetLifetime:output_type -> google.protobuf.Duration
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_csrf_api_csrf_proto_init() }
func file_csrf_api_csrf_proto_init() {
	if File_csrf_api_csrf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_csrf_api_csrf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSRFToken); i {
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
		file_csrf_api_csrf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSRFData); i {
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
			RawDescriptor: file_csrf_api_csrf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_csrf_api_csrf_proto_goTypes,
		DependencyIndexes: file_csrf_api_csrf_proto_depIdxs,
		MessageInfos:      file_csrf_api_csrf_proto_msgTypes,
	}.Build()
	File_csrf_api_csrf_proto = out.File
	file_csrf_api_csrf_proto_rawDesc = nil
	file_csrf_api_csrf_proto_goTypes = nil
	file_csrf_api_csrf_proto_depIdxs = nil
}
