// Code generated by protoc-gen-go. DO NOT EDIT.
// source: space.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Space struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator              string               `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Managers             []string             `protobuf:"bytes,3,rep,name=managers,proto3" json:"managers,omitempty"`
	Topic                string               `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	Details              string               `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
	Type                 string               `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Description          string               `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Tags                 []string             `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	Followers            []string             `protobuf:"bytes,9,rep,name=followers,proto3" json:"followers,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,10,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Space) Reset()         { *m = Space{} }
func (m *Space) String() string { return proto.CompactTextString(m) }
func (*Space) ProtoMessage()    {}
func (*Space) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{0}
}

func (m *Space) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Space.Unmarshal(m, b)
}
func (m *Space) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Space.Marshal(b, m, deterministic)
}
func (m *Space) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Space.Merge(m, src)
}
func (m *Space) XXX_Size() int {
	return xxx_messageInfo_Space.Size(m)
}
func (m *Space) XXX_DiscardUnknown() {
	xxx_messageInfo_Space.DiscardUnknown(m)
}

var xxx_messageInfo_Space proto.InternalMessageInfo

func (m *Space) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Space) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Space) GetManagers() []string {
	if m != nil {
		return m.Managers
	}
	return nil
}

func (m *Space) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Space) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *Space) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Space) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Space) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Space) GetFollowers() []string {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *Space) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type PostSpaceRequest struct {
	Creator              string               `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Managers             []string             `protobuf:"bytes,2,rep,name=managers,proto3" json:"managers,omitempty"`
	Topic                string               `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Details              string               `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	Type                 string               `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Description          string               `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Tags                 []string             `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Followers            []string             `protobuf:"bytes,8,rep,name=followers,proto3" json:"followers,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,9,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PostSpaceRequest) Reset()         { *m = PostSpaceRequest{} }
func (m *PostSpaceRequest) String() string { return proto.CompactTextString(m) }
func (*PostSpaceRequest) ProtoMessage()    {}
func (*PostSpaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{1}
}

func (m *PostSpaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostSpaceRequest.Unmarshal(m, b)
}
func (m *PostSpaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostSpaceRequest.Marshal(b, m, deterministic)
}
func (m *PostSpaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostSpaceRequest.Merge(m, src)
}
func (m *PostSpaceRequest) XXX_Size() int {
	return xxx_messageInfo_PostSpaceRequest.Size(m)
}
func (m *PostSpaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostSpaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostSpaceRequest proto.InternalMessageInfo

func (m *PostSpaceRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PostSpaceRequest) GetManagers() []string {
	if m != nil {
		return m.Managers
	}
	return nil
}

func (m *PostSpaceRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PostSpaceRequest) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *PostSpaceRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PostSpaceRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PostSpaceRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PostSpaceRequest) GetFollowers() []string {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *PostSpaceRequest) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type PostSpaceResponse struct {
	Space                string   `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostSpaceResponse) Reset()         { *m = PostSpaceResponse{} }
func (m *PostSpaceResponse) String() string { return proto.CompactTextString(m) }
func (*PostSpaceResponse) ProtoMessage()    {}
func (*PostSpaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{2}
}

func (m *PostSpaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostSpaceResponse.Unmarshal(m, b)
}
func (m *PostSpaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostSpaceResponse.Marshal(b, m, deterministic)
}
func (m *PostSpaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostSpaceResponse.Merge(m, src)
}
func (m *PostSpaceResponse) XXX_Size() int {
	return xxx_messageInfo_PostSpaceResponse.Size(m)
}
func (m *PostSpaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostSpaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostSpaceResponse proto.InternalMessageInfo

func (m *PostSpaceResponse) GetSpace() string {
	if m != nil {
		return m.Space
	}
	return ""
}

type GetSpaceRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID               uint64   `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSpaceRequest) Reset()         { *m = GetSpaceRequest{} }
func (m *GetSpaceRequest) String() string { return proto.CompactTextString(m) }
func (*GetSpaceRequest) ProtoMessage()    {}
func (*GetSpaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{3}
}

func (m *GetSpaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSpaceRequest.Unmarshal(m, b)
}
func (m *GetSpaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSpaceRequest.Marshal(b, m, deterministic)
}
func (m *GetSpaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpaceRequest.Merge(m, src)
}
func (m *GetSpaceRequest) XXX_Size() int {
	return xxx_messageInfo_GetSpaceRequest.Size(m)
}
func (m *GetSpaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpaceRequest proto.InternalMessageInfo

func (m *GetSpaceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetSpaceRequest) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type GetSpaceResponse struct {
	Space                *Space   `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSpaceResponse) Reset()         { *m = GetSpaceResponse{} }
func (m *GetSpaceResponse) String() string { return proto.CompactTextString(m) }
func (*GetSpaceResponse) ProtoMessage()    {}
func (*GetSpaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{4}
}

func (m *GetSpaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSpaceResponse.Unmarshal(m, b)
}
func (m *GetSpaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSpaceResponse.Marshal(b, m, deterministic)
}
func (m *GetSpaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpaceResponse.Merge(m, src)
}
func (m *GetSpaceResponse) XXX_Size() int {
	return xxx_messageInfo_GetSpaceResponse.Size(m)
}
func (m *GetSpaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpaceResponse proto.InternalMessageInfo

func (m *GetSpaceResponse) GetSpace() *Space {
	if m != nil {
		return m.Space
	}
	return nil
}

type GetMultipleSpacesRequest struct {
	Skip                 uint64   `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Take                 uint64   `protobuf:"varint,2,opt,name=take,proto3" json:"take,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMultipleSpacesRequest) Reset()         { *m = GetMultipleSpacesRequest{} }
func (m *GetMultipleSpacesRequest) String() string { return proto.CompactTextString(m) }
func (*GetMultipleSpacesRequest) ProtoMessage()    {}
func (*GetMultipleSpacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{5}
}

func (m *GetMultipleSpacesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMultipleSpacesRequest.Unmarshal(m, b)
}
func (m *GetMultipleSpacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMultipleSpacesRequest.Marshal(b, m, deterministic)
}
func (m *GetMultipleSpacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMultipleSpacesRequest.Merge(m, src)
}
func (m *GetMultipleSpacesRequest) XXX_Size() int {
	return xxx_messageInfo_GetMultipleSpacesRequest.Size(m)
}
func (m *GetMultipleSpacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMultipleSpacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMultipleSpacesRequest proto.InternalMessageInfo

func (m *GetMultipleSpacesRequest) GetSkip() uint64 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *GetMultipleSpacesRequest) GetTake() uint64 {
	if m != nil {
		return m.Take
	}
	return 0
}

type GetMultipleSpacesResponse struct {
	Spaces               []*Space `protobuf:"bytes,1,rep,name=spaces,proto3" json:"spaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMultipleSpacesResponse) Reset()         { *m = GetMultipleSpacesResponse{} }
func (m *GetMultipleSpacesResponse) String() string { return proto.CompactTextString(m) }
func (*GetMultipleSpacesResponse) ProtoMessage()    {}
func (*GetMultipleSpacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{6}
}

func (m *GetMultipleSpacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMultipleSpacesResponse.Unmarshal(m, b)
}
func (m *GetMultipleSpacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMultipleSpacesResponse.Marshal(b, m, deterministic)
}
func (m *GetMultipleSpacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMultipleSpacesResponse.Merge(m, src)
}
func (m *GetMultipleSpacesResponse) XXX_Size() int {
	return xxx_messageInfo_GetMultipleSpacesResponse.Size(m)
}
func (m *GetMultipleSpacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMultipleSpacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMultipleSpacesResponse proto.InternalMessageInfo

func (m *GetMultipleSpacesResponse) GetSpaces() []*Space {
	if m != nil {
		return m.Spaces
	}
	return nil
}

type PutSpaceRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator              string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Managers             []string `protobuf:"bytes,3,rep,name=managers,proto3" json:"managers,omitempty"`
	Topic                string   `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	Details              string   `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Tags                 []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	Followers            []string `protobuf:"bytes,9,rep,name=followers,proto3" json:"followers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutSpaceRequest) Reset()         { *m = PutSpaceRequest{} }
func (m *PutSpaceRequest) String() string { return proto.CompactTextString(m) }
func (*PutSpaceRequest) ProtoMessage()    {}
func (*PutSpaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{7}
}

func (m *PutSpaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutSpaceRequest.Unmarshal(m, b)
}
func (m *PutSpaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutSpaceRequest.Marshal(b, m, deterministic)
}
func (m *PutSpaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutSpaceRequest.Merge(m, src)
}
func (m *PutSpaceRequest) XXX_Size() int {
	return xxx_messageInfo_PutSpaceRequest.Size(m)
}
func (m *PutSpaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutSpaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutSpaceRequest proto.InternalMessageInfo

func (m *PutSpaceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PutSpaceRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PutSpaceRequest) GetManagers() []string {
	if m != nil {
		return m.Managers
	}
	return nil
}

func (m *PutSpaceRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PutSpaceRequest) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *PutSpaceRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PutSpaceRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PutSpaceRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PutSpaceRequest) GetFollowers() []string {
	if m != nil {
		return m.Followers
	}
	return nil
}

type PutSpaceResponse struct {
	Space                string   `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutSpaceResponse) Reset()         { *m = PutSpaceResponse{} }
func (m *PutSpaceResponse) String() string { return proto.CompactTextString(m) }
func (*PutSpaceResponse) ProtoMessage()    {}
func (*PutSpaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{8}
}

func (m *PutSpaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutSpaceResponse.Unmarshal(m, b)
}
func (m *PutSpaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutSpaceResponse.Marshal(b, m, deterministic)
}
func (m *PutSpaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutSpaceResponse.Merge(m, src)
}
func (m *PutSpaceResponse) XXX_Size() int {
	return xxx_messageInfo_PutSpaceResponse.Size(m)
}
func (m *PutSpaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutSpaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutSpaceResponse proto.InternalMessageInfo

func (m *PutSpaceResponse) GetSpace() string {
	if m != nil {
		return m.Space
	}
	return ""
}

type DeleteSpaceRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSpaceRequest) Reset()         { *m = DeleteSpaceRequest{} }
func (m *DeleteSpaceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSpaceRequest) ProtoMessage()    {}
func (*DeleteSpaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{9}
}

func (m *DeleteSpaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSpaceRequest.Unmarshal(m, b)
}
func (m *DeleteSpaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSpaceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSpaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSpaceRequest.Merge(m, src)
}
func (m *DeleteSpaceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSpaceRequest.Size(m)
}
func (m *DeleteSpaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSpaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSpaceRequest proto.InternalMessageInfo

func (m *DeleteSpaceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteSpaceResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSpaceResponse) Reset()         { *m = DeleteSpaceResponse{} }
func (m *DeleteSpaceResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteSpaceResponse) ProtoMessage()    {}
func (*DeleteSpaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8a3f24abfdc04ca, []int{10}
}

func (m *DeleteSpaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSpaceResponse.Unmarshal(m, b)
}
func (m *DeleteSpaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSpaceResponse.Marshal(b, m, deterministic)
}
func (m *DeleteSpaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSpaceResponse.Merge(m, src)
}
func (m *DeleteSpaceResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteSpaceResponse.Size(m)
}
func (m *DeleteSpaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSpaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSpaceResponse proto.InternalMessageInfo

func (m *DeleteSpaceResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Space)(nil), "pb.Space")
	proto.RegisterType((*PostSpaceRequest)(nil), "pb.PostSpaceRequest")
	proto.RegisterType((*PostSpaceResponse)(nil), "pb.PostSpaceResponse")
	proto.RegisterType((*GetSpaceRequest)(nil), "pb.GetSpaceRequest")
	proto.RegisterType((*GetSpaceResponse)(nil), "pb.GetSpaceResponse")
	proto.RegisterType((*GetMultipleSpacesRequest)(nil), "pb.GetMultipleSpacesRequest")
	proto.RegisterType((*GetMultipleSpacesResponse)(nil), "pb.GetMultipleSpacesResponse")
	proto.RegisterType((*PutSpaceRequest)(nil), "pb.PutSpaceRequest")
	proto.RegisterType((*PutSpaceResponse)(nil), "pb.PutSpaceResponse")
	proto.RegisterType((*DeleteSpaceRequest)(nil), "pb.DeleteSpaceRequest")
	proto.RegisterType((*DeleteSpaceResponse)(nil), "pb.DeleteSpaceResponse")
}

func init() { proto.RegisterFile("space.proto", fileDescriptor_b8a3f24abfdc04ca) }

var fileDescriptor_b8a3f24abfdc04ca = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x6b, 0xc7, 0xf9, 0xf0, 0x04, 0xd1, 0x74, 0x1b, 0xca, 0x62, 0x15, 0x35, 0x58, 0x20,
	0x85, 0x8b, 0x2b, 0xa5, 0x07, 0x04, 0x07, 0x84, 0x50, 0xa5, 0x8a, 0x03, 0x52, 0xe5, 0xf2, 0x02,
	0x4e, 0x32, 0x8d, 0xac, 0x3a, 0xd9, 0xc5, 0xbb, 0x06, 0xf1, 0x16, 0x5c, 0x78, 0x31, 0x5e, 0x84,
	0x57, 0x40, 0x1e, 0xaf, 0x13, 0xc7, 0x89, 0x23, 0xce, 0xdc, 0x76, 0x66, 0xff, 0xf3, 0xb1, 0x3f,
	0xff, 0x0d, 0x7d, 0x25, 0xa3, 0x19, 0x06, 0x32, 0x15, 0x5a, 0x30, 0x5b, 0x4e, 0xbd, 0x8b, 0x85,
	0x10, 0x8b, 0x04, 0x2f, 0x29, 0x33, 0xcd, 0xee, 0x2f, 0x75, 0xbc, 0x44, 0xa5, 0xa3, 0xa5, 0x2c,
	0x44, 0xfe, 0x2f, 0x1b, 0xda, 0x77, 0x79, 0x11, 0x7b, 0x0c, 0x76, 0x3c, 0xe7, 0xd6, 0xc8, 0x1a,
	0xbb, 0xa1, 0x1d, 0xcf, 0x19, 0x87, 0xee, 0x2c, 0xc5, 0x48, 0x8b, 0x94, 0xdb, 0x94, 0x2c, 0x43,
	0xe6, 0x41, 0x6f, 0x19, 0xad, 0xa2, 0x05, 0xa6, 0x8a, 0xb7, 0x46, 0xad, 0xb1, 0x1b, 0xae, 0x63,
	0x36, 0x84, 0xb6, 0x16, 0x32, 0x9e, 0x71, 0x87, 0x6a, 0x8a, 0x20, 0xef, 0x35, 0x47, 0x1d, 0xc5,
	0x89, 0xe2, 0xed, 0xa2, 0x97, 0x09, 0x19, 0x03, 0x47, 0xff, 0x90, 0xc8, 0x3b, 0x94, 0xa6, 0x33,
	0x1b, 0x41, 0x7f, 0x8e, 0x6a, 0x96, 0xc6, 0x52, 0xc7, 0x62, 0xc5, 0xbb, 0x74, 0x55, 0x4d, 0x51,
	0x55, 0xb4, 0x50, 0xbc, 0x47, 0xd3, 0xe9, 0xcc, 0xce, 0xc1, 0xbd, 0x17, 0x49, 0x22, 0xbe, 0xe7,
	0x6b, 0xb9, 0x74, 0xb1, 0x49, 0xb0, 0x00, 0x9c, 0xfc, 0xe9, 0x1c, 0x46, 0xd6, 0xb8, 0x3f, 0xf1,
	0x82, 0x82, 0x4b, 0x50, 0x72, 0x09, 0xbe, 0x94, 0x5c, 0x42, 0xd2, 0xf9, 0x3f, 0x6d, 0x18, 0xdc,
	0x0a, 0xa5, 0x89, 0x4d, 0x88, 0x5f, 0x33, 0x54, 0xba, 0x8a, 0xc4, 0x6a, 0x46, 0x62, 0x37, 0x21,
	0x69, 0x35, 0x20, 0x71, 0xf6, 0x23, 0x69, 0x37, 0x23, 0xe9, 0x34, 0x23, 0xe9, 0x36, 0x21, 0xe9,
	0x35, 0x21, 0x71, 0xff, 0x11, 0xc9, 0x6b, 0x38, 0xa9, 0x10, 0x51, 0x52, 0xac, 0x14, 0xe6, 0x8f,
	0x23, 0xcf, 0x19, 0x20, 0x45, 0xe0, 0xbf, 0x85, 0xe3, 0x1b, 0xdc, 0x66, 0x57, 0xb7, 0xd7, 0x19,
	0x74, 0x32, 0x85, 0xe9, 0xa7, 0x6b, 0x72, 0x97, 0x13, 0x9a, 0xc8, 0xbf, 0x82, 0xc1, 0xa6, 0xd4,
	0x0c, 0xb9, 0xa8, 0x0e, 0xe9, 0x4f, 0xdc, 0x40, 0x4e, 0x83, 0x42, 0x61, 0xe6, 0x7d, 0x04, 0x7e,
	0x83, 0xfa, 0x73, 0x96, 0xe8, 0x58, 0x26, 0x48, 0x57, 0xaa, 0x1c, 0xcc, 0xc0, 0x51, 0x0f, 0xb1,
	0xa4, 0x5a, 0x27, 0xa4, 0x73, 0x01, 0xeb, 0x01, 0xcd, 0x68, 0x3a, 0xfb, 0xef, 0xe1, 0xd9, 0x9e,
	0x1e, 0x66, 0x83, 0x17, 0xd0, 0xa1, 0x49, 0x8a, 0x5b, 0xa3, 0xd6, 0xf6, 0x0a, 0xe6, 0xc2, 0xff,
	0x63, 0xc1, 0xf1, 0x6d, 0x76, 0xf8, 0xd1, 0xff, 0xd5, 0x3f, 0xe5, 0x8f, 0x61, 0xb0, 0x79, 0xf0,
	0x41, 0x3f, 0xbc, 0x04, 0x76, 0x8d, 0x09, 0x6a, 0x3c, 0x44, 0xc7, 0x7f, 0x05, 0xa7, 0x5b, 0x2a,
	0xd3, 0xb2, 0x26, 0x9b, 0xfc, 0xb6, 0xe1, 0x11, 0x29, 0xee, 0x30, 0xfd, 0x16, 0xcf, 0x90, 0xbd,
	0x03, 0x77, 0x6d, 0x4c, 0x36, 0xcc, 0xbf, 0x4c, 0xfd, 0xcf, 0xf5, 0x9e, 0xd4, 0xb2, 0x45, 0x6b,
	0xff, 0x88, 0xbd, 0x81, 0x5e, 0x69, 0x37, 0x76, 0x9a, 0x8b, 0x6a, 0xbe, 0xf5, 0x86, 0xdb, 0xc9,
	0x75, 0x61, 0x08, 0x27, 0x3b, 0x76, 0x61, 0xe7, 0x46, 0xbc, 0xd7, 0x89, 0xde, 0xf3, 0x86, 0xdb,
	0xea, 0x32, 0x25, 0xd0, 0x62, 0x99, 0x9a, 0x9f, 0x8a, 0x65, 0xea, 0xcc, 0xfd, 0x23, 0xf6, 0x01,
	0xfa, 0x15, 0x72, 0xec, 0x2c, 0x97, 0xed, 0x02, 0xf7, 0x9e, 0xee, 0xe4, 0xcb, 0x0e, 0xd3, 0x0e,
	0xfd, 0xf6, 0x57, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x52, 0xb4, 0xe7, 0x53, 0x42, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SpaceServiceClient is the client API for SpaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpaceServiceClient interface {
	PostSpace(ctx context.Context, in *PostSpaceRequest, opts ...grpc.CallOption) (*PostSpaceResponse, error)
	GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error)
	GetMultipleSpaces(ctx context.Context, in *GetMultipleSpacesRequest, opts ...grpc.CallOption) (*GetMultipleSpacesResponse, error)
	PutSpace(ctx context.Context, in *PutSpaceRequest, opts ...grpc.CallOption) (*PutSpaceResponse, error)
	DeleteSpace(ctx context.Context, in *DeleteSpaceRequest, opts ...grpc.CallOption) (*DeleteSpaceResponse, error)
}

type spaceServiceClient struct {
	cc *grpc.ClientConn
}

func NewSpaceServiceClient(cc *grpc.ClientConn) SpaceServiceClient {
	return &spaceServiceClient{cc}
}

func (c *spaceServiceClient) PostSpace(ctx context.Context, in *PostSpaceRequest, opts ...grpc.CallOption) (*PostSpaceResponse, error) {
	out := new(PostSpaceResponse)
	err := c.cc.Invoke(ctx, "/pb.SpaceService/PostSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error) {
	out := new(GetSpaceResponse)
	err := c.cc.Invoke(ctx, "/pb.SpaceService/GetSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetMultipleSpaces(ctx context.Context, in *GetMultipleSpacesRequest, opts ...grpc.CallOption) (*GetMultipleSpacesResponse, error) {
	out := new(GetMultipleSpacesResponse)
	err := c.cc.Invoke(ctx, "/pb.SpaceService/GetMultipleSpaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) PutSpace(ctx context.Context, in *PutSpaceRequest, opts ...grpc.CallOption) (*PutSpaceResponse, error) {
	out := new(PutSpaceResponse)
	err := c.cc.Invoke(ctx, "/pb.SpaceService/PutSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) DeleteSpace(ctx context.Context, in *DeleteSpaceRequest, opts ...grpc.CallOption) (*DeleteSpaceResponse, error) {
	out := new(DeleteSpaceResponse)
	err := c.cc.Invoke(ctx, "/pb.SpaceService/DeleteSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServiceServer is the server API for SpaceService service.
type SpaceServiceServer interface {
	PostSpace(context.Context, *PostSpaceRequest) (*PostSpaceResponse, error)
	GetSpace(context.Context, *GetSpaceRequest) (*GetSpaceResponse, error)
	GetMultipleSpaces(context.Context, *GetMultipleSpacesRequest) (*GetMultipleSpacesResponse, error)
	PutSpace(context.Context, *PutSpaceRequest) (*PutSpaceResponse, error)
	DeleteSpace(context.Context, *DeleteSpaceRequest) (*DeleteSpaceResponse, error)
}

// UnimplementedSpaceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSpaceServiceServer struct {
}

func (*UnimplementedSpaceServiceServer) PostSpace(ctx context.Context, req *PostSpaceRequest) (*PostSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSpace not implemented")
}
func (*UnimplementedSpaceServiceServer) GetSpace(ctx context.Context, req *GetSpaceRequest) (*GetSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpace not implemented")
}
func (*UnimplementedSpaceServiceServer) GetMultipleSpaces(ctx context.Context, req *GetMultipleSpacesRequest) (*GetMultipleSpacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultipleSpaces not implemented")
}
func (*UnimplementedSpaceServiceServer) PutSpace(ctx context.Context, req *PutSpaceRequest) (*PutSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSpace not implemented")
}
func (*UnimplementedSpaceServiceServer) DeleteSpace(ctx context.Context, req *DeleteSpaceRequest) (*DeleteSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSpace not implemented")
}

func RegisterSpaceServiceServer(s *grpc.Server, srv SpaceServiceServer) {
	s.RegisterService(&_SpaceService_serviceDesc, srv)
}

func _SpaceService_PostSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).PostSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpaceService/PostSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).PostSpace(ctx, req.(*PostSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_GetSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpaceService/GetSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetSpace(ctx, req.(*GetSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_GetMultipleSpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultipleSpacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetMultipleSpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpaceService/GetMultipleSpaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetMultipleSpaces(ctx, req.(*GetMultipleSpacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_PutSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).PutSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpaceService/PutSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).PutSpace(ctx, req.(*PutSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_DeleteSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).DeleteSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpaceService/DeleteSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).DeleteSpace(ctx, req.(*DeleteSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SpaceService",
	HandlerType: (*SpaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostSpace",
			Handler:    _SpaceService_PostSpace_Handler,
		},
		{
			MethodName: "GetSpace",
			Handler:    _SpaceService_GetSpace_Handler,
		},
		{
			MethodName: "GetMultipleSpaces",
			Handler:    _SpaceService_GetMultipleSpaces_Handler,
		},
		{
			MethodName: "PutSpace",
			Handler:    _SpaceService_PutSpace_Handler,
		},
		{
			MethodName: "DeleteSpace",
			Handler:    _SpaceService_DeleteSpace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "space.proto",
}
