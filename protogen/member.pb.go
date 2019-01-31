// Code generated by protoc-gen-go. DO NOT EDIT.
// source: member.proto

package member

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MemberDetailRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberDetailRequest) Reset()         { *m = MemberDetailRequest{} }
func (m *MemberDetailRequest) String() string { return proto.CompactTextString(m) }
func (*MemberDetailRequest) ProtoMessage()    {}
func (*MemberDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_member_c54c7c6bc0844618, []int{0}
}
func (m *MemberDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberDetailRequest.Unmarshal(m, b)
}
func (m *MemberDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberDetailRequest.Marshal(b, m, deterministic)
}
func (dst *MemberDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberDetailRequest.Merge(dst, src)
}
func (m *MemberDetailRequest) XXX_Size() int {
	return xxx_messageInfo_MemberDetailRequest.Size(m)
}
func (m *MemberDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberDetailRequest proto.InternalMessageInfo

func (m *MemberDetailRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MemberDetailResponse struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Gender               string   `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	IdentityNo           string   `protobuf:"bytes,6,opt,name=identity_no,json=identityNo,proto3" json:"identity_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberDetailResponse) Reset()         { *m = MemberDetailResponse{} }
func (m *MemberDetailResponse) String() string { return proto.CompactTextString(m) }
func (*MemberDetailResponse) ProtoMessage()    {}
func (*MemberDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_member_c54c7c6bc0844618, []int{1}
}
func (m *MemberDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberDetailResponse.Unmarshal(m, b)
}
func (m *MemberDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberDetailResponse.Marshal(b, m, deterministic)
}
func (dst *MemberDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberDetailResponse.Merge(dst, src)
}
func (m *MemberDetailResponse) XXX_Size() int {
	return xxx_messageInfo_MemberDetailResponse.Size(m)
}
func (m *MemberDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberDetailResponse proto.InternalMessageInfo

func (m *MemberDetailResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *MemberDetailResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *MemberDetailResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *MemberDetailResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *MemberDetailResponse) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *MemberDetailResponse) GetIdentityNo() string {
	if m != nil {
		return m.IdentityNo
	}
	return ""
}

func init() {
	proto.RegisterType((*MemberDetailRequest)(nil), "MemberDetailRequest")
	proto.RegisterType((*MemberDetailResponse)(nil), "MemberDetailResponse")
}

func init() { proto.RegisterFile("member.proto", fileDescriptor_member_c54c7c6bc0844618) }

var fileDescriptor_member_c54c7c6bc0844618 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x87, 0x69, 0x75, 0x83, 0x1d, 0xff, 0x1c, 0x62, 0x95, 0xa0, 0x88, 0xb2, 0x20, 0x78, 0xea,
	0x41, 0xf1, 0x0d, 0xc4, 0x9b, 0x7b, 0x58, 0x1f, 0xa0, 0xa4, 0x66, 0xd4, 0x81, 0x26, 0xa9, 0x49,
	0x14, 0x7c, 0x32, 0x5f, 0x4f, 0x3a, 0x69, 0x0f, 0x0b, 0x3d, 0x7e, 0xdf, 0x37, 0x87, 0xfc, 0x02,
	0x47, 0x16, 0x6d, 0x87, 0xa1, 0x19, 0x82, 0x4f, 0x7e, 0x7d, 0x0b, 0xa7, 0x2f, 0xcc, 0x4f, 0x98,
	0x34, 0xf5, 0x5b, 0xfc, 0xfa, 0xc6, 0x98, 0xe4, 0x09, 0x94, 0x64, 0x54, 0x71, 0x53, 0xdc, 0x55,
	0xdb, 0x92, 0xcc, 0xfa, 0xaf, 0x80, 0x7a, 0xf7, 0x2e, 0x0e, 0xde, 0x45, 0x94, 0x35, 0xac, 0xd0,
	0x6a, 0xea, 0xa7, 0xdb, 0x0c, 0xa3, 0x1d, 0x3e, 0xbd, 0x43, 0x55, 0x66, 0xcb, 0x20, 0xaf, 0x00,
	0xde, 0x29, 0xc4, 0xd4, 0x3a, 0x6d, 0x51, 0xed, 0x71, 0xaa, 0xd8, 0x6c, 0xb4, 0x45, 0x79, 0x09,
	0x55, 0xaf, 0xe7, 0xba, 0xcf, 0xf5, 0x60, 0x14, 0x1c, 0xcf, 0x41, 0x7c, 0xa0, 0x33, 0x18, 0xd4,
	0x8a, 0xcb, 0x44, 0xf2, 0x1a, 0x0e, 0xc9, 0xa0, 0x4b, 0x94, 0x7e, 0x5b, 0xe7, 0x95, 0xe0, 0x08,
	0xb3, 0xda, 0xf8, 0xfb, 0x67, 0x38, 0xce, 0x0f, 0x7f, 0xc5, 0xf0, 0x43, 0x6f, 0x28, 0x1f, 0x41,
	0xe4, 0x0d, 0xb2, 0x6e, 0x16, 0xa6, 0x5f, 0x9c, 0x35, 0x4b, 0x43, 0x3b, 0xc1, 0xff, 0xf5, 0xf0,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x03, 0x03, 0x5c, 0x3f, 0x01, 0x00, 0x00,
}
