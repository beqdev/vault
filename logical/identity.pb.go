// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logical/identity.proto

package logical

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Entity struct {
	// ID is the unique identifier for the entity
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name is the human-friendly unique identifier for the entity
	Name string `sentinel:"" protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Aliases contains thhe alias mappings for the given entity
	Aliases []*Alias `sentinel:"" protobuf:"bytes,3,rep,name=aliases,proto3" json:"aliases,omitempty"`
	// Metadata represents the custom data tied to this entity
	Metadata             map[string]string `sentinel:"" protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_04442ca37d5e30be, []int{0}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetAliases() []*Alias {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Entity) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Alias struct {
	// MountType is the backend mount's type to which this identity belongs
	MountType string `sentinel:"" protobuf:"bytes,1,opt,name=mount_type,json=mountType,proto3" json:"mount_type,omitempty"`
	// MountAccessor is the identifier of the mount entry to which this
	// identity belongs
	MountAccessor string `sentinel:"" protobuf:"bytes,2,opt,name=mount_accessor,json=mountAccessor,proto3" json:"mount_accessor,omitempty"`
	// Name is the identifier of this identity in its authentication source
	Name string `sentinel:"" protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Metadata represents the custom data tied to this alias
	Metadata             map[string]string `sentinel:"" protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_04442ca37d5e30be, []int{1}
}

func (m *Alias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alias.Unmarshal(m, b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
}
func (m *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(m, src)
}
func (m *Alias) XXX_Size() int {
	return xxx_messageInfo_Alias.Size(m)
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetMountType() string {
	if m != nil {
		return m.MountType
	}
	return ""
}

func (m *Alias) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *Alias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Alias) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Entity)(nil), "logical.Entity")
	proto.RegisterMapType((map[string]string)(nil), "logical.Entity.MetadataEntry")
	proto.RegisterType((*Alias)(nil), "logical.Alias")
	proto.RegisterMapType((map[string]string)(nil), "logical.Alias.MetadataEntry")
}

func init() { proto.RegisterFile("logical/identity.proto", fileDescriptor_04442ca37d5e30be) }

var fileDescriptor_04442ca37d5e30be = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0xd2, 0x3f, 0x76, 0xa4, 0x45, 0x06, 0x91, 0x20, 0x16, 0x4a, 0x50, 0xc8, 0x29,
	0x01, 0xbd, 0x54, 0x3d, 0x55, 0xda, 0x43, 0x0f, 0x5e, 0x82, 0x27, 0x2f, 0x32, 0x4d, 0x97, 0x66,
	0x31, 0xc9, 0x86, 0x64, 0x52, 0xc8, 0x97, 0xf4, 0xec, 0xc7, 0x91, 0x6e, 0xb6, 0xc1, 0xe2, 0xd9,
	0xdb, 0xec, 0xef, 0xcd, 0xce, 0xbe, 0x79, 0x0b, 0x57, 0xa9, 0xda, 0xc9, 0x98, 0xd2, 0x50, 0x6e,
	0x45, 0xce, 0x92, 0x9b, 0xa0, 0x28, 0x15, 0x2b, 0x1c, 0x1a, 0xee, 0x7d, 0x59, 0x30, 0x58, 0x69,
	0x05, 0x27, 0x60, 0xaf, 0x97, 0xae, 0x35, 0xb3, 0xfc, 0x51, 0x64, 0xaf, 0x97, 0x88, 0xd0, 0xcb,
	0x29, 0x13, 0xae, 0xad, 0x89, 0xae, 0xd1, 0x87, 0x21, 0xa5, 0x92, 0x2a, 0x51, 0xb9, 0xce, 0xcc,
	0xf1, 0xcf, 0xef, 0x27, 0x81, 0x99, 0x14, 0x2c, 0x0e, 0x3c, 0x3a, 0xca, 0xf8, 0x08, 0x67, 0x99,
	0x60, 0xda, 0x12, 0x93, 0xdb, 0xd3, 0xad, 0xd3, 0xae, 0xb5, 0x7d, 0x30, 0x78, 0x35, 0xfa, 0x2a,
	0xe7, 0xb2, 0x89, 0xba, 0xf6, 0xeb, 0x67, 0x18, 0x9f, 0x48, 0x78, 0x01, 0xce, 0xa7, 0x68, 0x8c,
	0xb5, 0x43, 0x89, 0x97, 0xd0, 0xdf, 0x53, 0x5a, 0x1f, 0xcd, 0xb5, 0x87, 0x27, 0x7b, 0x6e, 0x79,
	0xdf, 0x16, 0xf4, 0xb5, 0x15, 0x9c, 0x02, 0x64, 0xaa, 0xce, 0xf9, 0x83, 0x9b, 0x42, 0x98, 0xcb,
	0x23, 0x4d, 0xde, 0x9a, 0x42, 0xe0, 0x1d, 0x4c, 0x5a, 0x99, 0xe2, 0x58, 0x54, 0x95, 0x2a, 0xcd,
	0xac, 0xb1, 0xa6, 0x0b, 0x03, 0xbb, 0x14, 0x9c, 0x5f, 0x29, 0xcc, 0xff, 0xec, 0x76, 0x73, 0x1a,
	0xc3, 0xbf, 0xac, 0xf6, 0x72, 0xfb, 0xee, 0xed, 0x24, 0x27, 0xf5, 0x26, 0x88, 0x55, 0x16, 0x26,
	0x54, 0x25, 0x32, 0x56, 0x65, 0x11, 0xee, 0xa9, 0x4e, 0x39, 0x34, 0x06, 0x36, 0x03, 0xfd, 0xc3,
	0x0f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xfb, 0x6f, 0x8c, 0xfb, 0x01, 0x00, 0x00,
}
