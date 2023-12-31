// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: gokitimp/aggsvc/pb/aggsvc.proto

package pb

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

type AggregateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObuID int32   `protobuf:"varint,1,opt,name=ObuID,proto3" json:"ObuID,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Unix  int64   `protobuf:"varint,3,opt,name=Unix,proto3" json:"Unix,omitempty"`
}

func (x *AggregateRequest) Reset() {
	*x = AggregateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateRequest) ProtoMessage() {}

func (x *AggregateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateRequest.ProtoReflect.Descriptor instead.
func (*AggregateRequest) Descriptor() ([]byte, []int) {
	return file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescGZIP(), []int{0}
}

func (x *AggregateRequest) GetObuID() int32 {
	if x != nil {
		return x.ObuID
	}
	return 0
}

func (x *AggregateRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *AggregateRequest) GetUnix() int64 {
	if x != nil {
		return x.Unix
	}
	return 0
}

type AggregateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AggregateReply) Reset() {
	*x = AggregateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateReply) ProtoMessage() {}

func (x *AggregateReply) ProtoReflect() protoreflect.Message {
	mi := &file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateReply.ProtoReflect.Descriptor instead.
func (*AggregateReply) Descriptor() ([]byte, []int) {
	return file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescGZIP(), []int{1}
}

type GetInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObuID int32 `protobuf:"varint,1,opt,name=ObuID,proto3" json:"ObuID,omitempty"`
}

func (x *GetInvoiceRequest) Reset() {
	*x = GetInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceRequest) ProtoMessage() {}

func (x *GetInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GetInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescGZIP(), []int{2}
}

func (x *GetInvoiceRequest) GetObuID() int32 {
	if x != nil {
		return x.ObuID
	}
	return 0
}

type GetInvoiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObuID         int32   `protobuf:"varint,1,opt,name=ObuID,proto3" json:"ObuID,omitempty"`
	TotalDistance float64 `protobuf:"fixed64,2,opt,name=TotalDistance,proto3" json:"TotalDistance,omitempty"`
	TotalAmount   float64 `protobuf:"fixed64,3,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
}

func (x *GetInvoiceReply) Reset() {
	*x = GetInvoiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvoiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceReply) ProtoMessage() {}

func (x *GetInvoiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceReply.ProtoReflect.Descriptor instead.
func (*GetInvoiceReply) Descriptor() ([]byte, []int) {
	return file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescGZIP(), []int{3}
}

func (x *GetInvoiceReply) GetObuID() int32 {
	if x != nil {
		return x.ObuID
	}
	return 0
}

func (x *GetInvoiceReply) GetTotalDistance() float64 {
	if x != nil {
		return x.TotalDistance
	}
	return 0
}

func (x *GetInvoiceReply) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

var File_gokitimp_aggsvc_pb_aggsvc_proto protoreflect.FileDescriptor

var file_gokitimp_aggsvc_pb_aggsvc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x6f, 0x6b, 0x69, 0x74, 0x69, 0x6d, 0x70, 0x2f, 0x61, 0x67, 0x67, 0x73, 0x76,
	0x63, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x67, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x52, 0x0a, 0x10, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x62, 0x75,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f, 0x62, 0x75, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x78, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4f, 0x62, 0x75, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x4f, 0x62, 0x75, 0x49, 0x44, 0x22, 0x6f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x62, 0x75,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f, 0x62, 0x75, 0x49, 0x44, 0x12,
	0x24, 0x0a, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x81, 0x01, 0x0a, 0x0a, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6d, 0x6b, 0x71, 0x77, 0x65,
	0x72, 0x74, 0x79, 0x2f, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x66, 0x61, 0x72, 0x65,
	0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x67, 0x6f, 0x6b, 0x69, 0x74, 0x69, 0x6d, 0x70,
	0x2f, 0x61, 0x67, 0x67, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescOnce sync.Once
	file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescData = file_gokitimp_aggsvc_pb_aggsvc_proto_rawDesc
)

func file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescGZIP() []byte {
	file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescOnce.Do(func() {
		file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescData)
	})
	return file_gokitimp_aggsvc_pb_aggsvc_proto_rawDescData
}

var file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gokitimp_aggsvc_pb_aggsvc_proto_goTypes = []interface{}{
	(*AggregateRequest)(nil),  // 0: pb.AggregateRequest
	(*AggregateReply)(nil),    // 1: pb.AggregateReply
	(*GetInvoiceRequest)(nil), // 2: pb.GetInvoiceRequest
	(*GetInvoiceReply)(nil),   // 3: pb.GetInvoiceReply
}
var file_gokitimp_aggsvc_pb_aggsvc_proto_depIdxs = []int32{
	0, // 0: pb.Aggregator.Aggregate:input_type -> pb.AggregateRequest
	2, // 1: pb.Aggregator.GetInvoice:input_type -> pb.GetInvoiceRequest
	1, // 2: pb.Aggregator.Aggregate:output_type -> pb.AggregateReply
	3, // 3: pb.Aggregator.GetInvoice:output_type -> pb.GetInvoiceReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gokitimp_aggsvc_pb_aggsvc_proto_init() }
func file_gokitimp_aggsvc_pb_aggsvc_proto_init() {
	if File_gokitimp_aggsvc_pb_aggsvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateRequest); i {
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
		file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateReply); i {
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
		file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvoiceRequest); i {
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
		file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvoiceReply); i {
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
			RawDescriptor: file_gokitimp_aggsvc_pb_aggsvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gokitimp_aggsvc_pb_aggsvc_proto_goTypes,
		DependencyIndexes: file_gokitimp_aggsvc_pb_aggsvc_proto_depIdxs,
		MessageInfos:      file_gokitimp_aggsvc_pb_aggsvc_proto_msgTypes,
	}.Build()
	File_gokitimp_aggsvc_pb_aggsvc_proto = out.File
	file_gokitimp_aggsvc_pb_aggsvc_proto_rawDesc = nil
	file_gokitimp_aggsvc_pb_aggsvc_proto_goTypes = nil
	file_gokitimp_aggsvc_pb_aggsvc_proto_depIdxs = nil
}
