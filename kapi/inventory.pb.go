// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: inventory.proto

package kapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PlayRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PlayRequests) Reset() {
	*x = PlayRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequests) ProtoMessage() {}

func (x *PlayRequests) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequests.ProtoReflect.Descriptor instead.
func (*PlayRequests) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *PlayRequests) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PlayReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *PlayReply) Reset() {
	*x = PlayReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayReply) ProtoMessage() {}

func (x *PlayReply) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayReply.ProtoReflect.Descriptor instead.
func (*PlayReply) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *PlayReply) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

type InventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*Node `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *InventoryRequest) Reset() {
	*x = InventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryRequest) ProtoMessage() {}

func (x *InventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryRequest.ProtoReflect.Descriptor instead.
func (*InventoryRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *InventoryRequest) GetItem() []*Node {
	if x != nil {
		return x.Item
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip       string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port     string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Role     string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *Node) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Node) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Node) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type InventoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InventoryReply) Reset() {
	*x = InventoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryReply) ProtoMessage() {}

func (x *InventoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryReply.ProtoReflect.Descriptor instead.
func (*InventoryReply) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *InventoryReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_inventory_proto protoreflect.FileDescriptor

var file_inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6b, 0x61, 0x70, 0x69, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x1d, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73,
	0x22, 0x32, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x22, 0x6e, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xc3, 0x01, 0x0a, 0x0d, 0x41, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x59, 0x61,
	0x6d, 0x6c, 0x12, 0x16, 0x2e, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6b, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x6f, 0x6f,
	0x6b, 0x12, 0x12, 0x2e, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x1a, 0x0f, 0x2e, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x75, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x2e,
	0x6b, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x1a, 0x0f, 0x2e, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6b, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData = file_inventory_proto_rawDesc
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_proto_rawDescData)
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_inventory_proto_goTypes = []interface{}{
	(*PlayRequests)(nil),     // 0: kapi.PlayRequests
	(*PlayReply)(nil),        // 1: kapi.PlayReply
	(*InventoryRequest)(nil), // 2: kapi.InventoryRequest
	(*Node)(nil),             // 3: kapi.Node
	(*InventoryReply)(nil),   // 4: kapi.InventoryReply
}
var file_inventory_proto_depIdxs = []int32{
	3, // 0: kapi.InventoryRequest.item:type_name -> kapi.Node
	2, // 1: kapi.AnsibleServer.GenerateYaml:input_type -> kapi.InventoryRequest
	0, // 2: kapi.AnsibleServer.RunPlaybook:input_type -> kapi.PlayRequests
	0, // 3: kapi.AnsibleServer.StreamRunPlaybook:input_type -> kapi.PlayRequests
	4, // 4: kapi.AnsibleServer.GenerateYaml:output_type -> kapi.InventoryReply
	1, // 5: kapi.AnsibleServer.RunPlaybook:output_type -> kapi.PlayReply
	1, // 6: kapi.AnsibleServer.StreamRunPlaybook:output_type -> kapi.PlayReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequests); i {
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
		file_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayReply); i {
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
		file_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryRequest); i {
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
		file_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_inventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryReply); i {
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
			RawDescriptor: file_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_rawDesc = nil
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnsibleServerClient is the client API for AnsibleServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnsibleServerClient interface {
	GenerateYaml(ctx context.Context, in *InventoryRequest, opts ...grpc.CallOption) (*InventoryReply, error)
	RunPlaybook(ctx context.Context, in *PlayRequests, opts ...grpc.CallOption) (*PlayReply, error)
	StreamRunPlaybook(ctx context.Context, in *PlayRequests, opts ...grpc.CallOption) (AnsibleServer_StreamRunPlaybookClient, error)
}

type ansibleServerClient struct {
	cc grpc.ClientConnInterface
}

func NewAnsibleServerClient(cc grpc.ClientConnInterface) AnsibleServerClient {
	return &ansibleServerClient{cc}
}

func (c *ansibleServerClient) GenerateYaml(ctx context.Context, in *InventoryRequest, opts ...grpc.CallOption) (*InventoryReply, error) {
	out := new(InventoryReply)
	err := c.cc.Invoke(ctx, "/kapi.AnsibleServer/GenerateYaml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ansibleServerClient) RunPlaybook(ctx context.Context, in *PlayRequests, opts ...grpc.CallOption) (*PlayReply, error) {
	out := new(PlayReply)
	err := c.cc.Invoke(ctx, "/kapi.AnsibleServer/RunPlaybook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ansibleServerClient) StreamRunPlaybook(ctx context.Context, in *PlayRequests, opts ...grpc.CallOption) (AnsibleServer_StreamRunPlaybookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AnsibleServer_serviceDesc.Streams[0], "/kapi.AnsibleServer/StreamRunPlaybook", opts...)
	if err != nil {
		return nil, err
	}
	x := &ansibleServerStreamRunPlaybookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AnsibleServer_StreamRunPlaybookClient interface {
	Recv() (*PlayReply, error)
	grpc.ClientStream
}

type ansibleServerStreamRunPlaybookClient struct {
	grpc.ClientStream
}

func (x *ansibleServerStreamRunPlaybookClient) Recv() (*PlayReply, error) {
	m := new(PlayReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnsibleServerServer is the server API for AnsibleServer service.
type AnsibleServerServer interface {
	GenerateYaml(context.Context, *InventoryRequest) (*InventoryReply, error)
	RunPlaybook(context.Context, *PlayRequests) (*PlayReply, error)
	StreamRunPlaybook(*PlayRequests, AnsibleServer_StreamRunPlaybookServer) error
}

// UnimplementedAnsibleServerServer can be embedded to have forward compatible implementations.
type UnimplementedAnsibleServerServer struct {
}

func (*UnimplementedAnsibleServerServer) GenerateYaml(context.Context, *InventoryRequest) (*InventoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateYaml not implemented")
}
func (*UnimplementedAnsibleServerServer) RunPlaybook(context.Context, *PlayRequests) (*PlayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunPlaybook not implemented")
}
func (*UnimplementedAnsibleServerServer) StreamRunPlaybook(*PlayRequests, AnsibleServer_StreamRunPlaybookServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRunPlaybook not implemented")
}

func RegisterAnsibleServerServer(s *grpc.Server, srv AnsibleServerServer) {
	s.RegisterService(&_AnsibleServer_serviceDesc, srv)
}

func _AnsibleServer_GenerateYaml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnsibleServerServer).GenerateYaml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kapi.AnsibleServer/GenerateYaml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnsibleServerServer).GenerateYaml(ctx, req.(*InventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnsibleServer_RunPlaybook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnsibleServerServer).RunPlaybook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kapi.AnsibleServer/RunPlaybook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnsibleServerServer).RunPlaybook(ctx, req.(*PlayRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnsibleServer_StreamRunPlaybook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlayRequests)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnsibleServerServer).StreamRunPlaybook(m, &ansibleServerStreamRunPlaybookServer{stream})
}

type AnsibleServer_StreamRunPlaybookServer interface {
	Send(*PlayReply) error
	grpc.ServerStream
}

type ansibleServerStreamRunPlaybookServer struct {
	grpc.ServerStream
}

func (x *ansibleServerStreamRunPlaybookServer) Send(m *PlayReply) error {
	return x.ServerStream.SendMsg(m)
}

var _AnsibleServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kapi.AnsibleServer",
	HandlerType: (*AnsibleServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateYaml",
			Handler:    _AnsibleServer_GenerateYaml_Handler,
		},
		{
			MethodName: "RunPlaybook",
			Handler:    _AnsibleServer_RunPlaybook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRunPlaybook",
			Handler:       _AnsibleServer_StreamRunPlaybook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "inventory.proto",
}