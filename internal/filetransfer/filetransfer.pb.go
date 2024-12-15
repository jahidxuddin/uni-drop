// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: filetransfer.proto

package filetransfer

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

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Data       []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ChunkIndex int64  `protobuf:"varint,3,opt,name=chunk_index,json=chunkIndex,proto3" json:"chunk_index,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	mi := &file_filetransfer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_filetransfer_proto_rawDescGZIP(), []int{0}
}

func (x *FileChunk) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FileChunk) GetChunkIndex() int64 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	mi := &file_filetransfer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_filetransfer_proto_rawDescGZIP(), []int{1}
}

func (x *FileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type DirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryPath string `protobuf:"bytes,1,opt,name=directory_path,json=directoryPath,proto3" json:"directory_path,omitempty"`
}

func (x *DirectoryRequest) Reset() {
	*x = DirectoryRequest{}
	mi := &file_filetransfer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryRequest) ProtoMessage() {}

func (x *DirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryRequest.ProtoReflect.Descriptor instead.
func (*DirectoryRequest) Descriptor() ([]byte, []int) {
	return file_filetransfer_proto_rawDescGZIP(), []int{2}
}

func (x *DirectoryRequest) GetDirectoryPath() string {
	if x != nil {
		return x.DirectoryPath
	}
	return ""
}

type FileList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *FileList) Reset() {
	*x = FileList{}
	mi := &file_filetransfer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileList) ProtoMessage() {}

func (x *FileList) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileList.ProtoReflect.Descriptor instead.
func (*FileList) Descriptor() ([]byte, []int) {
	return file_filetransfer_proto_rawDescGZIP(), []int{3}
}

func (x *FileList) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type UploadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	mi := &file_filetransfer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_filetransfer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_filetransfer_proto_rawDescGZIP(), []int{4}
}

func (x *UploadStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UploadStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_filetransfer_proto protoreflect.FileDescriptor

var file_filetransfer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x22, 0x5c, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x29, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x10, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x22, 0x20, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xde, 0x01, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x43, 0x0a,
	0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x28, 0x01, 0x12, 0x44, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x19, 0x5a,
	0x17, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filetransfer_proto_rawDescOnce sync.Once
	file_filetransfer_proto_rawDescData = file_filetransfer_proto_rawDesc
)

func file_filetransfer_proto_rawDescGZIP() []byte {
	file_filetransfer_proto_rawDescOnce.Do(func() {
		file_filetransfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_filetransfer_proto_rawDescData)
	})
	return file_filetransfer_proto_rawDescData
}

var file_filetransfer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_filetransfer_proto_goTypes = []any{
	(*FileChunk)(nil),        // 0: filetransfer.FileChunk
	(*FileRequest)(nil),      // 1: filetransfer.FileRequest
	(*DirectoryRequest)(nil), // 2: filetransfer.DirectoryRequest
	(*FileList)(nil),         // 3: filetransfer.FileList
	(*UploadStatus)(nil),     // 4: filetransfer.UploadStatus
}
var file_filetransfer_proto_depIdxs = []int32{
	0, // 0: filetransfer.FileTransfer.UploadFile:input_type -> filetransfer.FileChunk
	1, // 1: filetransfer.FileTransfer.DownloadFile:input_type -> filetransfer.FileRequest
	2, // 2: filetransfer.FileTransfer.ListFiles:input_type -> filetransfer.DirectoryRequest
	4, // 3: filetransfer.FileTransfer.UploadFile:output_type -> filetransfer.UploadStatus
	0, // 4: filetransfer.FileTransfer.DownloadFile:output_type -> filetransfer.FileChunk
	3, // 5: filetransfer.FileTransfer.ListFiles:output_type -> filetransfer.FileList
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_filetransfer_proto_init() }
func file_filetransfer_proto_init() {
	if File_filetransfer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filetransfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filetransfer_proto_goTypes,
		DependencyIndexes: file_filetransfer_proto_depIdxs,
		MessageInfos:      file_filetransfer_proto_msgTypes,
	}.Build()
	File_filetransfer_proto = out.File
	file_filetransfer_proto_rawDesc = nil
	file_filetransfer_proto_goTypes = nil
	file_filetransfer_proto_depIdxs = nil
}
