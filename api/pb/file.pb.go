// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type InApplyFid struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	ExName               string   `protobuf:"bytes,3,opt,name=ExName,proto3" json:"ExName,omitempty"`
	Md5                  string   `protobuf:"bytes,4,opt,name=Md5,proto3" json:"Md5,omitempty"`
	SliceTotal           int32    `protobuf:"varint,5,opt,name=SliceTotal,proto3" json:"SliceTotal,omitempty"`
	ExpiredTime          int64    `protobuf:"varint,6,opt,name=ExpiredTime,proto3" json:"ExpiredTime,omitempty"`
	Width                int32    `protobuf:"varint,7,opt,name=Width,proto3" json:"Width,omitempty"`
	Height               int32    `protobuf:"varint,8,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InApplyFid) Reset()         { *m = InApplyFid{} }
func (m *InApplyFid) String() string { return proto.CompactTextString(m) }
func (*InApplyFid) ProtoMessage()    {}
func (*InApplyFid) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *InApplyFid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InApplyFid.Unmarshal(m, b)
}
func (m *InApplyFid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InApplyFid.Marshal(b, m, deterministic)
}
func (m *InApplyFid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InApplyFid.Merge(m, src)
}
func (m *InApplyFid) XXX_Size() int {
	return xxx_messageInfo_InApplyFid.Size(m)
}
func (m *InApplyFid) XXX_DiscardUnknown() {
	xxx_messageInfo_InApplyFid.DiscardUnknown(m)
}

var xxx_messageInfo_InApplyFid proto.InternalMessageInfo

func (m *InApplyFid) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InApplyFid) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *InApplyFid) GetExName() string {
	if m != nil {
		return m.ExName
	}
	return ""
}

func (m *InApplyFid) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *InApplyFid) GetSliceTotal() int32 {
	if m != nil {
		return m.SliceTotal
	}
	return 0
}

func (m *InApplyFid) GetExpiredTime() int64 {
	if m != nil {
		return m.ExpiredTime
	}
	return 0
}

func (m *InApplyFid) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *InApplyFid) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type InFid struct {
	Fid                  int64    `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InFid) Reset()         { *m = InFid{} }
func (m *InFid) String() string { return proto.CompactTextString(m) }
func (*InFid) ProtoMessage()    {}
func (*InFid) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{1}
}

func (m *InFid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InFid.Unmarshal(m, b)
}
func (m *InFid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InFid.Marshal(b, m, deterministic)
}
func (m *InFid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InFid.Merge(m, src)
}
func (m *InFid) XXX_Size() int {
	return xxx_messageInfo_InFid.Size(m)
}
func (m *InFid) XXX_DiscardUnknown() {
	xxx_messageInfo_InFid.DiscardUnknown(m)
}

var xxx_messageInfo_InFid proto.InternalMessageInfo

func (m *InFid) GetFid() int64 {
	if m != nil {
		return m.Fid
	}
	return 0
}

type InMd5 struct {
	Md5                  string   `protobuf:"bytes,1,opt,name=Md5,proto3" json:"Md5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InMd5) Reset()         { *m = InMd5{} }
func (m *InMd5) String() string { return proto.CompactTextString(m) }
func (*InMd5) ProtoMessage()    {}
func (*InMd5) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{2}
}

func (m *InMd5) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InMd5.Unmarshal(m, b)
}
func (m *InMd5) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InMd5.Marshal(b, m, deterministic)
}
func (m *InMd5) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InMd5.Merge(m, src)
}
func (m *InMd5) XXX_Size() int {
	return xxx_messageInfo_InMd5.Size(m)
}
func (m *InMd5) XXX_DiscardUnknown() {
	xxx_messageInfo_InMd5.DiscardUnknown(m)
}

var xxx_messageInfo_InMd5 proto.InternalMessageInfo

func (m *InMd5) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

type OutApplyFid struct {
	Fid                  int64    `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutApplyFid) Reset()         { *m = OutApplyFid{} }
func (m *OutApplyFid) String() string { return proto.CompactTextString(m) }
func (*OutApplyFid) ProtoMessage()    {}
func (*OutApplyFid) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{3}
}

func (m *OutApplyFid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutApplyFid.Unmarshal(m, b)
}
func (m *OutApplyFid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutApplyFid.Marshal(b, m, deterministic)
}
func (m *OutApplyFid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutApplyFid.Merge(m, src)
}
func (m *OutApplyFid) XXX_Size() int {
	return xxx_messageInfo_OutApplyFid.Size(m)
}
func (m *OutApplyFid) XXX_DiscardUnknown() {
	xxx_messageInfo_OutApplyFid.DiscardUnknown(m)
}

var xxx_messageInfo_OutApplyFid proto.InternalMessageInfo

func (m *OutApplyFid) GetFid() int64 {
	if m != nil {
		return m.Fid
	}
	return 0
}

func (m *OutApplyFid) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type InUpSliceFileItem struct {
	Fid                  int64    `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	Part                 int32    `protobuf:"varint,2,opt,name=Part,proto3" json:"Part,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Md5                  string   `protobuf:"bytes,4,opt,name=Md5,proto3" json:"Md5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InUpSliceFileItem) Reset()         { *m = InUpSliceFileItem{} }
func (m *InUpSliceFileItem) String() string { return proto.CompactTextString(m) }
func (*InUpSliceFileItem) ProtoMessage()    {}
func (*InUpSliceFileItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{4}
}

func (m *InUpSliceFileItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InUpSliceFileItem.Unmarshal(m, b)
}
func (m *InUpSliceFileItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InUpSliceFileItem.Marshal(b, m, deterministic)
}
func (m *InUpSliceFileItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InUpSliceFileItem.Merge(m, src)
}
func (m *InUpSliceFileItem) XXX_Size() int {
	return xxx_messageInfo_InUpSliceFileItem.Size(m)
}
func (m *InUpSliceFileItem) XXX_DiscardUnknown() {
	xxx_messageInfo_InUpSliceFileItem.DiscardUnknown(m)
}

var xxx_messageInfo_InUpSliceFileItem proto.InternalMessageInfo

func (m *InUpSliceFileItem) GetFid() int64 {
	if m != nil {
		return m.Fid
	}
	return 0
}

func (m *InUpSliceFileItem) GetPart() int32 {
	if m != nil {
		return m.Part
	}
	return 0
}

func (m *InUpSliceFileItem) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InUpSliceFileItem) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

type InDownSliceFileItem struct {
	Fid                  int64    `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=Offset,proto3" json:"Offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InDownSliceFileItem) Reset()         { *m = InDownSliceFileItem{} }
func (m *InDownSliceFileItem) String() string { return proto.CompactTextString(m) }
func (*InDownSliceFileItem) ProtoMessage()    {}
func (*InDownSliceFileItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{5}
}

func (m *InDownSliceFileItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InDownSliceFileItem.Unmarshal(m, b)
}
func (m *InDownSliceFileItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InDownSliceFileItem.Marshal(b, m, deterministic)
}
func (m *InDownSliceFileItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InDownSliceFileItem.Merge(m, src)
}
func (m *InDownSliceFileItem) XXX_Size() int {
	return xxx_messageInfo_InDownSliceFileItem.Size(m)
}
func (m *InDownSliceFileItem) XXX_DiscardUnknown() {
	xxx_messageInfo_InDownSliceFileItem.DiscardUnknown(m)
}

var xxx_messageInfo_InDownSliceFileItem proto.InternalMessageInfo

func (m *InDownSliceFileItem) GetFid() int64 {
	if m != nil {
		return m.Fid
	}
	return 0
}

func (m *InDownSliceFileItem) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *InDownSliceFileItem) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type OutDownSliceFileItem struct {
	Fid                  int64    `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Md5                  string   `protobuf:"bytes,3,opt,name=Md5,proto3" json:"Md5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutDownSliceFileItem) Reset()         { *m = OutDownSliceFileItem{} }
func (m *OutDownSliceFileItem) String() string { return proto.CompactTextString(m) }
func (*OutDownSliceFileItem) ProtoMessage()    {}
func (*OutDownSliceFileItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{6}
}

func (m *OutDownSliceFileItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutDownSliceFileItem.Unmarshal(m, b)
}
func (m *OutDownSliceFileItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutDownSliceFileItem.Marshal(b, m, deterministic)
}
func (m *OutDownSliceFileItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutDownSliceFileItem.Merge(m, src)
}
func (m *OutDownSliceFileItem) XXX_Size() int {
	return xxx_messageInfo_OutDownSliceFileItem.Size(m)
}
func (m *OutDownSliceFileItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OutDownSliceFileItem.DiscardUnknown(m)
}

var xxx_messageInfo_OutDownSliceFileItem proto.InternalMessageInfo

func (m *OutDownSliceFileItem) GetFid() int64 {
	if m != nil {
		return m.Fid
	}
	return 0
}

func (m *OutDownSliceFileItem) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *OutDownSliceFileItem) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

type OutDownFile struct {
	Fid                  int64    `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Md5                  string   `protobuf:"bytes,3,opt,name=Md5,proto3" json:"Md5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutDownFile) Reset()         { *m = OutDownFile{} }
func (m *OutDownFile) String() string { return proto.CompactTextString(m) }
func (*OutDownFile) ProtoMessage()    {}
func (*OutDownFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{7}
}

func (m *OutDownFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutDownFile.Unmarshal(m, b)
}
func (m *OutDownFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutDownFile.Marshal(b, m, deterministic)
}
func (m *OutDownFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutDownFile.Merge(m, src)
}
func (m *OutDownFile) XXX_Size() int {
	return xxx_messageInfo_OutDownFile.Size(m)
}
func (m *OutDownFile) XXX_DiscardUnknown() {
	xxx_messageInfo_OutDownFile.DiscardUnknown(m)
}

var xxx_messageInfo_OutDownFile proto.InternalMessageInfo

func (m *OutDownFile) GetFid() int64 {
	if m != nil {
		return m.Fid
	}
	return 0
}

func (m *OutDownFile) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *OutDownFile) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

type InUpFile struct {
	Fid                  int64    `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Md5                  string   `protobuf:"bytes,4,opt,name=Md5,proto3" json:"Md5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InUpFile) Reset()         { *m = InUpFile{} }
func (m *InUpFile) String() string { return proto.CompactTextString(m) }
func (*InUpFile) ProtoMessage()    {}
func (*InUpFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{8}
}

func (m *InUpFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InUpFile.Unmarshal(m, b)
}
func (m *InUpFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InUpFile.Marshal(b, m, deterministic)
}
func (m *InUpFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InUpFile.Merge(m, src)
}
func (m *InUpFile) XXX_Size() int {
	return xxx_messageInfo_InUpFile.Size(m)
}
func (m *InUpFile) XXX_DiscardUnknown() {
	xxx_messageInfo_InUpFile.DiscardUnknown(m)
}

var xxx_messageInfo_InUpFile proto.InternalMessageInfo

func (m *InUpFile) GetFid() int64 {
	if m != nil {
		return m.Fid
	}
	return 0
}

func (m *InUpFile) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InUpFile) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

type FileInfo struct {
	Fid                  int64    `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	BucketName           string   `protobuf:"bytes,3,opt,name=BucketName,proto3" json:"BucketName,omitempty"`
	Size                 int64    `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	ContentType          string   `protobuf:"bytes,5,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	Md5                  string   `protobuf:"bytes,6,opt,name=Md5,proto3" json:"Md5,omitempty"`
	ExName               string   `protobuf:"bytes,7,opt,name=ExName,proto3" json:"ExName,omitempty"`
	IsImage              bool     `protobuf:"varint,8,opt,name=IsImage,proto3" json:"IsImage,omitempty"`
	ExImage              *ImageEx `protobuf:"bytes,9,opt,name=ExImage,proto3" json:"ExImage,omitempty"`
	SliceTotal           int32    `protobuf:"varint,10,opt,name=SliceTotal,proto3" json:"SliceTotal,omitempty"`
	ExpiredTime          int64    `protobuf:"varint,11,opt,name=ExpiredTime,proto3" json:"ExpiredTime,omitempty"`
	Status               int32    `protobuf:"varint,12,opt,name=Status,proto3" json:"Status,omitempty"`
	CreateTime           int64    `protobuf:"varint,13,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime           int64    `protobuf:"varint,14,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{9}
}

func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetFid() int64 {
	if m != nil {
		return m.Fid
	}
	return 0
}

func (m *FileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileInfo) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *FileInfo) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileInfo) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *FileInfo) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *FileInfo) GetExName() string {
	if m != nil {
		return m.ExName
	}
	return ""
}

func (m *FileInfo) GetIsImage() bool {
	if m != nil {
		return m.IsImage
	}
	return false
}

func (m *FileInfo) GetExImage() *ImageEx {
	if m != nil {
		return m.ExImage
	}
	return nil
}

func (m *FileInfo) GetSliceTotal() int32 {
	if m != nil {
		return m.SliceTotal
	}
	return 0
}

func (m *FileInfo) GetExpiredTime() int64 {
	if m != nil {
		return m.ExpiredTime
	}
	return 0
}

func (m *FileInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *FileInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *FileInfo) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type ImageEx struct {
	Height               int32    `protobuf:"varint,1,opt,name=Height,proto3" json:"Height,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=Width,proto3" json:"Width,omitempty"`
	ThumbnailFid         int64    `protobuf:"varint,3,opt,name=ThumbnailFid,proto3" json:"ThumbnailFid,omitempty"`
	ThumbnailHeight      int32    `protobuf:"varint,4,opt,name=ThumbnailHeight,proto3" json:"ThumbnailHeight,omitempty"`
	ThumbnailWidth       int32    `protobuf:"varint,5,opt,name=ThumbnailWidth,proto3" json:"ThumbnailWidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageEx) Reset()         { *m = ImageEx{} }
func (m *ImageEx) String() string { return proto.CompactTextString(m) }
func (*ImageEx) ProtoMessage()    {}
func (*ImageEx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{10}
}

func (m *ImageEx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageEx.Unmarshal(m, b)
}
func (m *ImageEx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageEx.Marshal(b, m, deterministic)
}
func (m *ImageEx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageEx.Merge(m, src)
}
func (m *ImageEx) XXX_Size() int {
	return xxx_messageInfo_ImageEx.Size(m)
}
func (m *ImageEx) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageEx.DiscardUnknown(m)
}

var xxx_messageInfo_ImageEx proto.InternalMessageInfo

func (m *ImageEx) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ImageEx) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ImageEx) GetThumbnailFid() int64 {
	if m != nil {
		return m.ThumbnailFid
	}
	return 0
}

func (m *ImageEx) GetThumbnailHeight() int32 {
	if m != nil {
		return m.ThumbnailHeight
	}
	return 0
}

func (m *ImageEx) GetThumbnailWidth() int32 {
	if m != nil {
		return m.ThumbnailWidth
	}
	return 0
}

type InCancelUpload struct {
	Fid                  int64    `protobuf:"varint,1,opt,name=Fid,proto3" json:"Fid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InCancelUpload) Reset()         { *m = InCancelUpload{} }
func (m *InCancelUpload) String() string { return proto.CompactTextString(m) }
func (*InCancelUpload) ProtoMessage()    {}
func (*InCancelUpload) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{11}
}

func (m *InCancelUpload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InCancelUpload.Unmarshal(m, b)
}
func (m *InCancelUpload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InCancelUpload.Marshal(b, m, deterministic)
}
func (m *InCancelUpload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InCancelUpload.Merge(m, src)
}
func (m *InCancelUpload) XXX_Size() int {
	return xxx_messageInfo_InCancelUpload.Size(m)
}
func (m *InCancelUpload) XXX_DiscardUnknown() {
	xxx_messageInfo_InCancelUpload.DiscardUnknown(m)
}

var xxx_messageInfo_InCancelUpload proto.InternalMessageInfo

func (m *InCancelUpload) GetFid() int64 {
	if m != nil {
		return m.Fid
	}
	return 0
}

func init() {
	proto.RegisterType((*InApplyFid)(nil), "storage.InApplyFid")
	proto.RegisterType((*InFid)(nil), "storage.InFid")
	proto.RegisterType((*InMd5)(nil), "storage.InMd5")
	proto.RegisterType((*OutApplyFid)(nil), "storage.OutApplyFid")
	proto.RegisterType((*InUpSliceFileItem)(nil), "storage.InUpSliceFileItem")
	proto.RegisterType((*InDownSliceFileItem)(nil), "storage.InDownSliceFileItem")
	proto.RegisterType((*OutDownSliceFileItem)(nil), "storage.OutDownSliceFileItem")
	proto.RegisterType((*OutDownFile)(nil), "storage.OutDownFile")
	proto.RegisterType((*InUpFile)(nil), "storage.InUpFile")
	proto.RegisterType((*FileInfo)(nil), "storage.FileInfo")
	proto.RegisterType((*ImageEx)(nil), "storage.ImageEx")
	proto.RegisterType((*InCancelUpload)(nil), "storage.InCancelUpload")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162) }

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x51, 0x8f, 0xd2, 0x40,
	0x10, 0x4e, 0x29, 0xa5, 0x30, 0x20, 0xe2, 0x4a, 0x4c, 0x7d, 0xb9, 0x34, 0x7d, 0x30, 0xc4, 0x07,
	0x1e, 0x34, 0xc6, 0x67, 0x8f, 0x83, 0xd8, 0x44, 0x0f, 0x53, 0x20, 0x3e, 0x9a, 0x3d, 0xba, 0xc0,
	0xc6, 0x76, 0xdb, 0xc0, 0x12, 0x39, 0x7f, 0x85, 0x7f, 0xc4, 0xbf, 0xe2, 0x6f, 0x32, 0x3b, 0x5d,
	0xe8, 0x1e, 0x10, 0x31, 0xf7, 0x36, 0xf3, 0xcd, 0xce, 0xcc, 0xb7, 0xb3, 0xf3, 0x2d, 0xc0, 0x82,
	0x27, 0xac, 0x9f, 0xaf, 0x33, 0x99, 0x11, 0x77, 0x23, 0xb3, 0x35, 0x5d, 0xb2, 0xe0, 0x8f, 0x05,
	0x10, 0x8a, 0x0f, 0x79, 0x9e, 0xdc, 0x8f, 0x78, 0x4c, 0x08, 0x54, 0x6f, 0x69, 0xca, 0x3c, 0xcb,
	0xb7, 0x7a, 0x8d, 0x08, 0x6d, 0x85, 0x4d, 0xf8, 0x4f, 0xe6, 0x55, 0x7c, 0xab, 0x67, 0x47, 0x68,
	0x93, 0x17, 0x50, 0x1b, 0xee, 0xf0, 0xa4, 0x8d, 0x27, 0xb5, 0x47, 0x3a, 0x60, 0x7f, 0x8e, 0xdf,
	0x79, 0x55, 0x04, 0x95, 0x49, 0xae, 0x00, 0x26, 0x09, 0x9f, 0xb3, 0x69, 0x26, 0x69, 0xe2, 0x39,
	0xbe, 0xd5, 0x73, 0x22, 0x03, 0x21, 0x3e, 0x34, 0x87, 0xbb, 0x9c, 0xaf, 0x59, 0x3c, 0xe5, 0x29,
	0xf3, 0x6a, 0xd8, 0xc4, 0x84, 0x48, 0x17, 0x9c, 0xaf, 0x3c, 0x96, 0x2b, 0xcf, 0xc5, 0xe4, 0xc2,
	0x51, 0x0c, 0x3e, 0x32, 0xbe, 0x5c, 0x49, 0xaf, 0x8e, 0xb0, 0xf6, 0x82, 0x97, 0xe0, 0x84, 0x42,
	0x5d, 0xa5, 0x03, 0xf6, 0x88, 0xc7, 0x78, 0x13, 0x3b, 0x52, 0x66, 0x11, 0x52, 0x9c, 0x34, 0x4b,
	0xeb, 0xc0, 0x32, 0x78, 0x0f, 0xcd, 0xf1, 0x56, 0x1e, 0xc6, 0x70, 0x92, 0xab, 0xda, 0x4d, 0x24,
	0x95, 0xdb, 0x0d, 0x8e, 0xc1, 0x89, 0xb4, 0x17, 0x7c, 0x83, 0x67, 0xa1, 0x98, 0xe5, 0x78, 0xa1,
	0x11, 0x4f, 0x58, 0x28, 0x59, 0x7a, 0x26, 0x9d, 0x40, 0xf5, 0x0b, 0x5d, 0x4b, 0x9d, 0x8c, 0xb6,
	0xc2, 0x6e, 0xa8, 0xa4, 0x38, 0xc1, 0x56, 0x84, 0xf6, 0xe9, 0xfc, 0x82, 0x19, 0x3c, 0x0f, 0xc5,
	0x4d, 0xf6, 0x43, 0x5c, 0x6a, 0xd1, 0x05, 0xe7, 0x13, 0x4f, 0xb9, 0xd4, 0xef, 0x54, 0x38, 0x8a,
	0xf7, 0x78, 0xb1, 0xd8, 0x30, 0x89, 0x6d, 0xec, 0x48, 0x7b, 0xc1, 0x2d, 0x74, 0xc7, 0x5b, 0xf9,
	0x3f, 0x75, 0xf7, 0x34, 0x2b, 0xa7, 0x34, 0xed, 0x92, 0xe6, 0x10, 0x07, 0xa8, 0xea, 0xa9, 0x52,
	0x8f, 0x2e, 0x73, 0x0d, 0x75, 0x35, 0xce, 0x0b, 0x35, 0xfe, 0x3d, 0xb1, 0x5f, 0x36, 0xd4, 0xf1,
	0x3e, 0x62, 0x91, 0x9d, 0x2f, 0x82, 0x8b, 0x5b, 0x31, 0x56, 0xfc, 0x0a, 0xe0, 0x7a, 0x3b, 0xff,
	0xce, 0xa4, 0xb1, 0xd2, 0x06, 0x72, 0x90, 0x40, 0xd5, 0x90, 0x80, 0x0f, 0xcd, 0x41, 0x26, 0x24,
	0x13, 0x72, 0x7a, 0x9f, 0x33, 0xdc, 0xec, 0x46, 0x64, 0x42, 0x7b, 0x6a, 0xb5, 0x52, 0x0c, 0xa5,
	0x6c, 0xdc, 0x07, 0xb2, 0xf1, 0xc0, 0x0d, 0x37, 0x61, 0x4a, 0x97, 0x0c, 0xb7, 0xb9, 0x1e, 0xed,
	0x5d, 0xf2, 0x1a, 0xdc, 0xe1, 0xae, 0x88, 0x34, 0x7c, 0xab, 0xd7, 0x7c, 0xd3, 0xe9, 0x6b, 0xe9,
	0xf6, 0x11, 0x1d, 0xee, 0xa2, 0xfd, 0x81, 0x23, 0xa9, 0xc1, 0x25, 0xa9, 0x35, 0x4f, 0xa5, 0x56,
	0x6e, 0x79, 0xcb, 0xdc, 0x72, 0x55, 0x79, 0xb0, 0x66, 0x54, 0x32, 0x4c, 0x7c, 0x82, 0x89, 0x06,
	0xa2, 0xe2, 0xb3, 0x3c, 0xde, 0xc7, 0xdb, 0x45, 0xbc, 0x44, 0x82, 0xdf, 0x16, 0xb8, 0x9a, 0xae,
	0x21, 0x5c, 0xcb, 0x14, 0x6e, 0x29, 0xf3, 0x8a, 0x29, 0xf3, 0x00, 0x5a, 0xd3, 0xd5, 0x36, 0xbd,
	0x13, 0x94, 0x27, 0xea, 0x21, 0x8b, 0x2d, 0x7e, 0x80, 0x91, 0x1e, 0x3c, 0x3d, 0xf8, 0xba, 0x74,
	0x15, 0x6b, 0x1c, 0xc3, 0xe4, 0x15, 0xb4, 0x0f, 0x50, 0xd1, 0xac, 0xf8, 0x90, 0x8e, 0xd0, 0x20,
	0x80, 0x76, 0x28, 0x06, 0x54, 0xcc, 0x59, 0x32, 0xcb, 0x93, 0x8c, 0x9e, 0xf9, 0x11, 0xee, 0x6a,
	0xf8, 0x93, 0xbe, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xb6, 0x14, 0x5c, 0x57, 0x05, 0x00,
	0x00,
}
