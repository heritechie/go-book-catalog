// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rpc_create_book.proto

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

type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Year        string `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	AuthorId    string `protobuf:"bytes,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	PublisherId string `protobuf:"bytes,5,opt,name=publisher_id,json=publisherId,proto3" json:"publisher_id,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_book_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBookRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *CreateBookRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *CreateBookRequest) GetPublisherId() string {
	if x != nil {
		return x.PublisherId
	}
	return ""
}

type CreateBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookResponse) Reset() {
	*x = CreateBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponse) ProtoMessage() {}

func (x *CreateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResponse.ProtoReflect.Descriptor instead.
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_book_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookResponse) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

var File_rpc_create_book_proto protoreflect.FileDescriptor

var file_rpc_create_book_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x42, 0x3c, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x72, 0x69,
	0x74, 0x65, 0x63, 0x68, 0x69, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x72, 0x65, 0x70, 0x6f,
	0x2d, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_create_book_proto_rawDescOnce sync.Once
	file_rpc_create_book_proto_rawDescData = file_rpc_create_book_proto_rawDesc
)

func file_rpc_create_book_proto_rawDescGZIP() []byte {
	file_rpc_create_book_proto_rawDescOnce.Do(func() {
		file_rpc_create_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_book_proto_rawDescData)
	})
	return file_rpc_create_book_proto_rawDescData
}

var file_rpc_create_book_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_book_proto_goTypes = []interface{}{
	(*CreateBookRequest)(nil),  // 0: pb.CreateBookRequest
	(*CreateBookResponse)(nil), // 1: pb.CreateBookResponse
	(*Book)(nil),               // 2: pb.Book
}
var file_rpc_create_book_proto_depIdxs = []int32{
	2, // 0: pb.CreateBookResponse.book:type_name -> pb.Book
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_create_book_proto_init() }
func file_rpc_create_book_proto_init() {
	if File_rpc_create_book_proto != nil {
		return
	}
	file_book_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
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
		file_rpc_create_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookResponse); i {
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
			RawDescriptor: file_rpc_create_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_book_proto_goTypes,
		DependencyIndexes: file_rpc_create_book_proto_depIdxs,
		MessageInfos:      file_rpc_create_book_proto_msgTypes,
	}.Build()
	File_rpc_create_book_proto = out.File
	file_rpc_create_book_proto_rawDesc = nil
	file_rpc_create_book_proto_goTypes = nil
	file_rpc_create_book_proto_depIdxs = nil
}
