// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

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

type Profile struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	About                string               `protobuf:"bytes,3,opt,name=about,proto3" json:"about,omitempty"`
	Level                string               `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	Rings                int32                `protobuf:"varint,5,opt,name=rings,proto3" json:"rings,omitempty"`
	ProfileImg           *Image               `protobuf:"bytes,6,opt,name=profileImg,proto3" json:"profileImg,omitempty"`
	BackgroundImg        *Image               `protobuf:"bytes,7,opt,name=backgroundImg,proto3" json:"backgroundImg,omitempty"`
	DateLastUpdated      *timestamp.Timestamp `protobuf:"bytes,8,opt,name=dateLastUpdated,proto3" json:"dateLastUpdated,omitempty"`
	Followers            []string             `protobuf:"bytes,9,rep,name=followers,proto3" json:"followers,omitempty"`
	Following            []string             `protobuf:"bytes,10,rep,name=following,proto3" json:"following,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{0}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Profile) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *Profile) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *Profile) GetRings() int32 {
	if m != nil {
		return m.Rings
	}
	return 0
}

func (m *Profile) GetProfileImg() *Image {
	if m != nil {
		return m.ProfileImg
	}
	return nil
}

func (m *Profile) GetBackgroundImg() *Image {
	if m != nil {
		return m.BackgroundImg
	}
	return nil
}

func (m *Profile) GetDateLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.DateLastUpdated
	}
	return nil
}

func (m *Profile) GetFollowers() []string {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *Profile) GetFollowing() []string {
	if m != nil {
		return m.Following
	}
	return nil
}

type Image struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Size                 int64    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Width                int32    `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{1}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Image) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type PostRequest struct {
	Username             string               `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	About                string               `protobuf:"bytes,2,opt,name=about,proto3" json:"about,omitempty"`
	Level                string               `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	Rings                int32                `protobuf:"varint,4,opt,name=rings,proto3" json:"rings,omitempty"`
	ProfileImg           *Image               `protobuf:"bytes,5,opt,name=profileImg,proto3" json:"profileImg,omitempty"`
	BackgroundImg        *Image               `protobuf:"bytes,6,opt,name=backgroundImg,proto3" json:"backgroundImg,omitempty"`
	DateLastUpdated      *timestamp.Timestamp `protobuf:"bytes,7,opt,name=dateLastUpdated,proto3" json:"dateLastUpdated,omitempty"`
	Followers            []string             `protobuf:"bytes,8,rep,name=followers,proto3" json:"followers,omitempty"`
	Following            []string             `protobuf:"bytes,9,rep,name=following,proto3" json:"following,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PostRequest) Reset()         { *m = PostRequest{} }
func (m *PostRequest) String() string { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()    {}
func (*PostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{2}
}

func (m *PostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostRequest.Unmarshal(m, b)
}
func (m *PostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostRequest.Marshal(b, m, deterministic)
}
func (m *PostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostRequest.Merge(m, src)
}
func (m *PostRequest) XXX_Size() int {
	return xxx_messageInfo_PostRequest.Size(m)
}
func (m *PostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostRequest proto.InternalMessageInfo

func (m *PostRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PostRequest) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *PostRequest) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *PostRequest) GetRings() int32 {
	if m != nil {
		return m.Rings
	}
	return 0
}

func (m *PostRequest) GetProfileImg() *Image {
	if m != nil {
		return m.ProfileImg
	}
	return nil
}

func (m *PostRequest) GetBackgroundImg() *Image {
	if m != nil {
		return m.BackgroundImg
	}
	return nil
}

func (m *PostRequest) GetDateLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.DateLastUpdated
	}
	return nil
}

func (m *PostRequest) GetFollowers() []string {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *PostRequest) GetFollowing() []string {
	if m != nil {
		return m.Following
	}
	return nil
}

type PostResponse struct {
	Profile              string   `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResponse) Reset()         { *m = PostResponse{} }
func (m *PostResponse) String() string { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()    {}
func (*PostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{3}
}

func (m *PostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResponse.Unmarshal(m, b)
}
func (m *PostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResponse.Marshal(b, m, deterministic)
}
func (m *PostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResponse.Merge(m, src)
}
func (m *PostResponse) XXX_Size() int {
	return xxx_messageInfo_PostResponse.Size(m)
}
func (m *PostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostResponse proto.InternalMessageInfo

func (m *PostResponse) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID               uint64   `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetRequest) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type GetResponse struct {
	Profile              *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{5}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type GetMultipleRequest struct {
	Skip                 uint64   `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Take                 uint64   `protobuf:"varint,2,opt,name=take,proto3" json:"take,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMultipleRequest) Reset()         { *m = GetMultipleRequest{} }
func (m *GetMultipleRequest) String() string { return proto.CompactTextString(m) }
func (*GetMultipleRequest) ProtoMessage()    {}
func (*GetMultipleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{6}
}

func (m *GetMultipleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMultipleRequest.Unmarshal(m, b)
}
func (m *GetMultipleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMultipleRequest.Marshal(b, m, deterministic)
}
func (m *GetMultipleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMultipleRequest.Merge(m, src)
}
func (m *GetMultipleRequest) XXX_Size() int {
	return xxx_messageInfo_GetMultipleRequest.Size(m)
}
func (m *GetMultipleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMultipleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMultipleRequest proto.InternalMessageInfo

func (m *GetMultipleRequest) GetSkip() uint64 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *GetMultipleRequest) GetTake() uint64 {
	if m != nil {
		return m.Take
	}
	return 0
}

type GetMultipleResponse struct {
	Profiles             []*Profile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetMultipleResponse) Reset()         { *m = GetMultipleResponse{} }
func (m *GetMultipleResponse) String() string { return proto.CompactTextString(m) }
func (*GetMultipleResponse) ProtoMessage()    {}
func (*GetMultipleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{7}
}

func (m *GetMultipleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMultipleResponse.Unmarshal(m, b)
}
func (m *GetMultipleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMultipleResponse.Marshal(b, m, deterministic)
}
func (m *GetMultipleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMultipleResponse.Merge(m, src)
}
func (m *GetMultipleResponse) XXX_Size() int {
	return xxx_messageInfo_GetMultipleResponse.Size(m)
}
func (m *GetMultipleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMultipleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMultipleResponse proto.InternalMessageInfo

func (m *GetMultipleResponse) GetProfiles() []*Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

type PutRequest struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	About                string               `protobuf:"bytes,3,opt,name=about,proto3" json:"about,omitempty"`
	Level                string               `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	Rings                int32                `protobuf:"varint,5,opt,name=rings,proto3" json:"rings,omitempty"`
	ProfileImg           *Image               `protobuf:"bytes,6,opt,name=profileImg,proto3" json:"profileImg,omitempty"`
	BackgroundImg        *Image               `protobuf:"bytes,7,opt,name=backgroundImg,proto3" json:"backgroundImg,omitempty"`
	DateLastUpdated      *timestamp.Timestamp `protobuf:"bytes,8,opt,name=dateLastUpdated,proto3" json:"dateLastUpdated,omitempty"`
	Followers            []string             `protobuf:"bytes,9,rep,name=followers,proto3" json:"followers,omitempty"`
	Following            []string             `protobuf:"bytes,10,rep,name=following,proto3" json:"following,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{8}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PutRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PutRequest) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *PutRequest) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *PutRequest) GetRings() int32 {
	if m != nil {
		return m.Rings
	}
	return 0
}

func (m *PutRequest) GetProfileImg() *Image {
	if m != nil {
		return m.ProfileImg
	}
	return nil
}

func (m *PutRequest) GetBackgroundImg() *Image {
	if m != nil {
		return m.BackgroundImg
	}
	return nil
}

func (m *PutRequest) GetDateLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.DateLastUpdated
	}
	return nil
}

func (m *PutRequest) GetFollowers() []string {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *PutRequest) GetFollowing() []string {
	if m != nil {
		return m.Following
	}
	return nil
}

type PutResponse struct {
	Profile              string   `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{9}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{10}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{11}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Profile)(nil), "pb.Profile")
	proto.RegisterType((*Image)(nil), "pb.Image")
	proto.RegisterType((*PostRequest)(nil), "pb.PostRequest")
	proto.RegisterType((*PostResponse)(nil), "pb.PostResponse")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "pb.GetResponse")
	proto.RegisterType((*GetMultipleRequest)(nil), "pb.GetMultipleRequest")
	proto.RegisterType((*GetMultipleResponse)(nil), "pb.GetMultipleResponse")
	proto.RegisterType((*PutRequest)(nil), "pb.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "pb.PutResponse")
	proto.RegisterType((*DeleteRequest)(nil), "pb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "pb.DeleteResponse")
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor_744bf7a47b381504) }

var fileDescriptor_744bf7a47b381504 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0x51, 0x6f, 0x94, 0x40,
	0x10, 0x2e, 0x1c, 0x70, 0xc7, 0x60, 0xaf, 0xba, 0x9a, 0x4a, 0x88, 0x49, 0x09, 0x89, 0x29, 0xc6,
	0x84, 0xc6, 0xda, 0x47, 0x63, 0x7c, 0x68, 0xd2, 0x34, 0xd1, 0xe4, 0x82, 0xfa, 0x03, 0xa0, 0x4c,
	0xe9, 0xa6, 0x1c, 0x20, 0x2c, 0x6d, 0xf4, 0xcd, 0x17, 0xff, 0x9b, 0x3f, 0xc0, 0xc4, 0x9f, 0x63,
	0x76, 0x59, 0x0e, 0x8e, 0xf6, 0x34, 0xd5, 0x57, 0xdf, 0x66, 0xbe, 0x99, 0xd9, 0x61, 0xbf, 0xef,
	0x63, 0x61, 0xbb, 0xac, 0x8a, 0x73, 0x9a, 0x61, 0x50, 0x56, 0x05, 0x2b, 0x88, 0x5a, 0xc6, 0xce,
	0x5e, 0x5a, 0x14, 0x69, 0x86, 0x07, 0x02, 0x89, 0x9b, 0xf3, 0x03, 0x46, 0x97, 0x58, 0xb3, 0x68,
	0x59, 0xb6, 0x4d, 0xde, 0x0f, 0x15, 0xa6, 0x8b, 0x76, 0x8c, 0xcc, 0x41, 0xa5, 0x89, 0xad, 0xb8,
	0x8a, 0x6f, 0x86, 0x2a, 0x4d, 0x88, 0x03, 0xb3, 0xa6, 0xc6, 0x2a, 0x8f, 0x96, 0x68, 0xab, 0x02,
	0x5d, 0xe5, 0xe4, 0x11, 0xe8, 0x51, 0x5c, 0x34, 0xcc, 0x9e, 0x88, 0x42, 0x9b, 0x70, 0x34, 0xc3,
	0x2b, 0xcc, 0x6c, 0xad, 0x45, 0x45, 0xc2, 0xd1, 0x8a, 0xe6, 0x69, 0x6d, 0xeb, 0xae, 0xe2, 0xeb,
	0x61, 0x9b, 0x90, 0x67, 0x00, 0xf2, 0x7b, 0x4f, 0x97, 0xa9, 0x6d, 0xb8, 0x8a, 0x6f, 0x1d, 0x9a,
	0x41, 0x19, 0x07, 0xa7, 0xcb, 0x28, 0xc5, 0x70, 0x50, 0x24, 0x07, 0xb0, 0x1d, 0x47, 0x67, 0x97,
	0x69, 0x55, 0x34, 0x79, 0xc2, 0xbb, 0xa7, 0xe3, 0xee, 0xf5, 0x3a, 0x39, 0x86, 0x9d, 0x24, 0x62,
	0xf8, 0x36, 0xaa, 0xd9, 0xc7, 0x92, 0x47, 0x89, 0x3d, 0x13, 0x23, 0x4e, 0xd0, 0x12, 0x12, 0x74,
	0x84, 0x04, 0x1f, 0x3a, 0x42, 0xc2, 0xf1, 0x08, 0x79, 0x02, 0xe6, 0x79, 0x91, 0x65, 0xc5, 0x35,
	0x56, 0xb5, 0x6d, 0xba, 0x13, 0xdf, 0x0c, 0x7b, 0xa0, 0xaf, 0xd2, 0x3c, 0xb5, 0x61, 0x58, 0xa5,
	0x79, 0xea, 0x7d, 0x55, 0x40, 0x17, 0x9f, 0x76, 0x83, 0x55, 0x02, 0xda, 0x80, 0x51, 0x11, 0x73,
	0x8c, 0x7d, 0x2e, 0x51, 0x92, 0x29, 0x62, 0x8e, 0xd5, 0xf4, 0x0b, 0x0a, 0x2a, 0x27, 0xa1, 0x88,
	0x39, 0x93, 0xd7, 0x34, 0x61, 0x17, 0x1d, 0x93, 0x22, 0x21, 0xbb, 0x60, 0x5c, 0x20, 0x4d, 0x2f,
	0x98, 0x60, 0x51, 0x0f, 0x65, 0xe6, 0x7d, 0x57, 0xc1, 0x5a, 0x14, 0x35, 0x0b, 0xf1, 0x53, 0x83,
	0x35, 0x5b, 0xd3, 0x53, 0xd9, 0xa4, 0xa7, 0x7a, 0xab, 0x9e, 0x93, 0x5b, 0xf5, 0xd4, 0x36, 0xeb,
	0xa9, 0xdf, 0x49, 0x4f, 0xe3, 0xee, 0x7a, 0x4e, 0xff, 0x51, 0xcf, 0xd9, 0x6f, 0xf5, 0x34, 0xc7,
	0x7a, 0xfa, 0x70, 0xaf, 0xa5, 0xb2, 0x2e, 0x8b, 0xbc, 0x46, 0x62, 0xc3, 0x54, 0x5e, 0x48, 0x52,
	0xd9, 0xa5, 0xde, 0x11, 0xc0, 0x09, 0xae, 0x38, 0x1f, 0xab, 0xbf, 0x0b, 0x06, 0xe7, 0xfc, 0xf4,
	0x58, 0x10, 0xad, 0x85, 0x32, 0xf3, 0x8e, 0xc0, 0x12, 0x53, 0xf2, 0xf8, 0xa7, 0xeb, 0xc7, 0x5b,
	0x87, 0x16, 0xe7, 0x46, 0xfe, 0xa8, 0xfd, 0xae, 0x57, 0x40, 0x4e, 0x90, 0xbd, 0x6b, 0x32, 0x46,
	0xcb, 0x0c, 0xbb, 0x9d, 0xdc, 0x39, 0x97, 0xb4, 0x14, 0x93, 0x5a, 0x28, 0x62, 0xe1, 0xb0, 0xe8,
	0x12, 0xe5, 0x56, 0x11, 0x7b, 0xaf, 0xe1, 0xe1, 0xda, 0xb4, 0xdc, 0xbd, 0x0f, 0x33, 0x79, 0x7e,
	0x6d, 0x2b, 0xee, 0x64, 0xbc, 0x7c, 0x55, 0xf4, 0x7e, 0xaa, 0x00, 0x8b, 0x66, 0xe3, 0x55, 0xff,
	0x3f, 0x1f, 0x7f, 0xff, 0x7c, 0xec, 0x83, 0x25, 0x98, 0xfd, 0xa3, 0xdb, 0xf6, 0x60, 0xfb, 0x18,
	0x33, 0x64, 0xb8, 0x41, 0x05, 0xcf, 0x85, 0x79, 0xd7, 0x20, 0x0f, 0x1b, 0x75, 0x1c, 0x7e, 0x53,
	0x61, 0x2e, 0xc5, 0x7d, 0x8f, 0xd5, 0x15, 0x3d, 0x43, 0xf2, 0x1c, 0x34, 0xee, 0x76, 0xb2, 0x23,
	0x84, 0xef, 0x9f, 0x10, 0xe7, 0x7e, 0x0f, 0xb4, 0xa7, 0x79, 0x5b, 0xc4, 0x87, 0xc9, 0x09, 0x32,
	0x32, 0xe7, 0xa5, 0xde, 0xf9, 0xce, 0xce, 0x2a, 0x5f, 0x75, 0xbe, 0x11, 0x26, 0xef, 0x0c, 0x47,
	0x76, 0x65, 0xc7, 0xc8, 0xbf, 0xce, 0xe3, 0x1b, 0xf8, 0x70, 0xd7, 0xa2, 0x91, 0xbb, 0x7a, 0xeb,
	0xb5, 0xbb, 0x06, 0x84, 0x79, 0x5b, 0xe4, 0x05, 0x18, 0xed, 0xbd, 0xc9, 0x03, 0x5e, 0x5c, 0x23,
	0xc9, 0x21, 0x43, 0xa8, 0x1b, 0x89, 0x0d, 0xa1, 0xea, 0xcb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x6c, 0xaa, 0xb5, 0x42, 0x48, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetMultiple(ctx context.Context, in *GetMultipleRequest, opts ...grpc.CallOption) (*GetMultipleResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type profileServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/pb.ProfileService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pb.ProfileService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetMultiple(ctx context.Context, in *GetMultipleRequest, opts ...grpc.CallOption) (*GetMultipleResponse, error) {
	out := new(GetMultipleResponse)
	err := c.cc.Invoke(ctx, "/pb.ProfileService/GetMultiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/pb.ProfileService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.ProfileService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	Post(context.Context, *PostRequest) (*PostResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetMultiple(context.Context, *GetMultipleRequest) (*GetMultipleResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (*UnimplementedProfileServiceServer) Post(ctx context.Context, req *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (*UnimplementedProfileServiceServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedProfileServiceServer) GetMultiple(ctx context.Context, req *GetMultipleRequest) (*GetMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiple not implemented")
}
func (*UnimplementedProfileServiceServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedProfileServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProfileService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProfileService/GetMultiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetMultiple(ctx, req.(*GetMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProfileService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _ProfileService_Post_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProfileService_Get_Handler,
		},
		{
			MethodName: "GetMultiple",
			Handler:    _ProfileService_GetMultiple_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _ProfileService_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProfileService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
