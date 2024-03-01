// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: api/protos/console.proto

package console

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

type GetConsoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetConsoleRequest) Reset() {
	*x = GetConsoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_console_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsoleRequest) ProtoMessage() {}

func (x *GetConsoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_console_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsoleRequest.ProtoReflect.Descriptor instead.
func (*GetConsoleRequest) Descriptor() ([]byte, []int) {
	return file_api_protos_console_proto_rawDescGZIP(), []int{0}
}

func (x *GetConsoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetConsoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Console *Console `protobuf:"bytes,1,opt,name=console,proto3" json:"console,omitempty"`
}

func (x *GetConsoleResponse) Reset() {
	*x = GetConsoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_console_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsoleResponse) ProtoMessage() {}

func (x *GetConsoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_console_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsoleResponse.ProtoReflect.Descriptor instead.
func (*GetConsoleResponse) Descriptor() ([]byte, []int) {
	return file_api_protos_console_proto_rawDescGZIP(), []int{1}
}

func (x *GetConsoleResponse) GetConsole() *Console {
	if x != nil {
		return x.Console
	}
	return nil
}

type CreateConsoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Manufacturer string `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ReleaseDate  string `protobuf:"bytes,3,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
}

func (x *CreateConsoleRequest) Reset() {
	*x = CreateConsoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_console_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsoleRequest) ProtoMessage() {}

func (x *CreateConsoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_console_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsoleRequest.ProtoReflect.Descriptor instead.
func (*CreateConsoleRequest) Descriptor() ([]byte, []int) {
	return file_api_protos_console_proto_rawDescGZIP(), []int{2}
}

func (x *CreateConsoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateConsoleRequest) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *CreateConsoleRequest) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

type CreateConsoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Console *Console `protobuf:"bytes,1,opt,name=console,proto3" json:"console,omitempty"`
}

func (x *CreateConsoleResponse) Reset() {
	*x = CreateConsoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_console_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsoleResponse) ProtoMessage() {}

func (x *CreateConsoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_console_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsoleResponse.ProtoReflect.Descriptor instead.
func (*CreateConsoleResponse) Descriptor() ([]byte, []int) {
	return file_api_protos_console_proto_rawDescGZIP(), []int{3}
}

func (x *CreateConsoleResponse) GetConsole() *Console {
	if x != nil {
		return x.Console
	}
	return nil
}

type Console struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Manufacturer string `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ReleaseDate  string `protobuf:"bytes,4,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
}

func (x *Console) Reset() {
	*x = Console{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_console_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Console) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Console) ProtoMessage() {}

func (x *Console) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_console_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Console.ProtoReflect.Descriptor instead.
func (*Console) Descriptor() ([]byte, []int) {
	return file_api_protos_console_proto_rawDescGZIP(), []int{4}
}

func (x *Console) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Console) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Console) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *Console) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

var File_api_protos_console_proto protoreflect.FileDescriptor

var file_api_protos_console_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x65, 0x7a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x58, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x61, 0x72,
	0x65, 0x7a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x22, 0x71, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x5b, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x61, 0x72,
	0x65, 0x7a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x22, 0x74, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e,
	0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x32, 0x87, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x12, 0x32, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x65, 0x7a, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x65, 0x7a,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7e, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x12, 0x35, 0x2e, 0x67, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x65, 0x7a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x65, 0x7a, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x65,
	0x7a, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protos_console_proto_rawDescOnce sync.Once
	file_api_protos_console_proto_rawDescData = file_api_protos_console_proto_rawDesc
)

func file_api_protos_console_proto_rawDescGZIP() []byte {
	file_api_protos_console_proto_rawDescOnce.Do(func() {
		file_api_protos_console_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protos_console_proto_rawDescData)
	})
	return file_api_protos_console_proto_rawDescData
}

var file_api_protos_console_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_protos_console_proto_goTypes = []interface{}{
	(*GetConsoleRequest)(nil),     // 0: gandarez.video_game_api.console.GetConsoleRequest
	(*GetConsoleResponse)(nil),    // 1: gandarez.video_game_api.console.GetConsoleResponse
	(*CreateConsoleRequest)(nil),  // 2: gandarez.video_game_api.console.CreateConsoleRequest
	(*CreateConsoleResponse)(nil), // 3: gandarez.video_game_api.console.CreateConsoleResponse
	(*Console)(nil),               // 4: gandarez.video_game_api.console.Console
}
var file_api_protos_console_proto_depIdxs = []int32{
	4, // 0: gandarez.video_game_api.console.GetConsoleResponse.console:type_name -> gandarez.video_game_api.console.Console
	4, // 1: gandarez.video_game_api.console.CreateConsoleResponse.console:type_name -> gandarez.video_game_api.console.Console
	0, // 2: gandarez.video_game_api.console.ConsoleService.GetConsole:input_type -> gandarez.video_game_api.console.GetConsoleRequest
	2, // 3: gandarez.video_game_api.console.ConsoleService.CreateConsole:input_type -> gandarez.video_game_api.console.CreateConsoleRequest
	1, // 4: gandarez.video_game_api.console.ConsoleService.GetConsole:output_type -> gandarez.video_game_api.console.GetConsoleResponse
	3, // 5: gandarez.video_game_api.console.ConsoleService.CreateConsole:output_type -> gandarez.video_game_api.console.CreateConsoleResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_protos_console_proto_init() }
func file_api_protos_console_proto_init() {
	if File_api_protos_console_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protos_console_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsoleRequest); i {
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
		file_api_protos_console_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsoleResponse); i {
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
		file_api_protos_console_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsoleRequest); i {
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
		file_api_protos_console_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsoleResponse); i {
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
		file_api_protos_console_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Console); i {
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
			RawDescriptor: file_api_protos_console_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_protos_console_proto_goTypes,
		DependencyIndexes: file_api_protos_console_proto_depIdxs,
		MessageInfos:      file_api_protos_console_proto_msgTypes,
	}.Build()
	File_api_protos_console_proto = out.File
	file_api_protos_console_proto_rawDesc = nil
	file_api_protos_console_proto_goTypes = nil
	file_api_protos_console_proto_depIdxs = nil
}