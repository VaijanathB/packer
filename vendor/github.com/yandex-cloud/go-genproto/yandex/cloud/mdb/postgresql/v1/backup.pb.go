// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/postgresql/v1/backup.proto

package postgresql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A PostgreSQL Backup resource. For more information, see
// the [Developer's Guide](/docs/managed-postgresql/concepts/backup).
type Backup struct {
	// ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the backup belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format
	// (i.e. when the backup operation was completed).
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the PostgreSQL cluster that the backup was created for.
	SourceClusterId string `protobuf:"bytes,4,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Time when the backup operation was started.
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Backup) Reset()         { *m = Backup{} }
func (m *Backup) String() string { return proto.CompactTextString(m) }
func (*Backup) ProtoMessage()    {}
func (*Backup) Descriptor() ([]byte, []int) {
	return fileDescriptor_12b8aee21e233390, []int{0}
}

func (m *Backup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backup.Unmarshal(m, b)
}
func (m *Backup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backup.Marshal(b, m, deterministic)
}
func (m *Backup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backup.Merge(m, src)
}
func (m *Backup) XXX_Size() int {
	return xxx_messageInfo_Backup.Size(m)
}
func (m *Backup) XXX_DiscardUnknown() {
	xxx_messageInfo_Backup.DiscardUnknown(m)
}

var xxx_messageInfo_Backup proto.InternalMessageInfo

func (m *Backup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Backup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Backup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Backup) GetSourceClusterId() string {
	if m != nil {
		return m.SourceClusterId
	}
	return ""
}

func (m *Backup) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Backup)(nil), "yandex.cloud.mdb.postgresql.v1.Backup")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/postgresql/v1/backup.proto", fileDescriptor_12b8aee21e233390)
}

var fileDescriptor_12b8aee21e233390 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x99, 0xf9, 0x7f, 0x8b, 0x13, 0x41, 0x71, 0x56, 0x43, 0x05, 0x2d, 0xae, 0x8a, 0xd2,
	0x84, 0xea, 0x4a, 0x5c, 0xb5, 0xae, 0x5c, 0x88, 0x50, 0x5c, 0xb9, 0x19, 0x92, 0xdc, 0x34, 0x0e,
	0xce, 0xf4, 0x8e, 0xc9, 0x4d, 0xd1, 0x27, 0xf5, 0x75, 0x84, 0x64, 0x4a, 0x77, 0xba, 0xcc, 0xc9,
	0x77, 0xcf, 0x07, 0x87, 0x5d, 0x7f, 0xc9, 0x0d, 0x98, 0x4f, 0xa1, 0x5b, 0x0c, 0x20, 0x3a, 0x50,
	0xa2, 0x47, 0x4f, 0xd6, 0x19, 0xff, 0xd1, 0x8a, 0xed, 0x5c, 0x28, 0xa9, 0xdf, 0x43, 0xcf, 0x7b,
	0x87, 0x84, 0xe5, 0x79, 0x82, 0x79, 0x84, 0x79, 0x07, 0x8a, 0xef, 0x61, 0xbe, 0x9d, 0x8f, 0x2f,
	0x2c, 0xa2, 0x6d, 0x8d, 0x88, 0xb4, 0x0a, 0x6b, 0x41, 0x4d, 0x67, 0x3c, 0xc9, 0x6e, 0x28, 0xb8,
	0xfc, 0xce, 0xd8, 0x68, 0x19, 0x1b, 0xcb, 0x63, 0x96, 0x37, 0x50, 0x65, 0x93, 0x6c, 0x5a, 0xac,
	0xf2, 0x06, 0xca, 0x33, 0x56, 0xac, 0xb1, 0x05, 0xe3, 0xea, 0x06, 0xaa, 0x3c, 0xc6, 0x87, 0x29,
	0x78, 0x84, 0xf2, 0x8e, 0x31, 0xed, 0x8c, 0x24, 0x03, 0xb5, 0xa4, 0xea, 0xdf, 0x24, 0x9b, 0x1e,
	0xdd, 0x8c, 0x79, 0xb2, 0xf1, 0x9d, 0x8d, 0xbf, 0xec, 0x6c, 0xab, 0x62, 0xa0, 0x17, 0x54, 0x5e,
	0xb1, 0x53, 0x8f, 0xc1, 0x69, 0x53, 0xeb, 0x36, 0x78, 0x4a, 0xfd, 0xff, 0x63, 0xff, 0x49, 0xfa,
	0x78, 0x48, 0x79, 0xd2, 0x78, 0x92, 0x6e, 0xd0, 0x1c, 0xfc, 0xad, 0x19, 0xe8, 0x05, 0x2d, 0x9f,
	0x5f, 0x9f, 0x6c, 0x43, 0x6f, 0x41, 0x71, 0x8d, 0x9d, 0x48, 0x3b, 0xcd, 0xd2, 0xa8, 0x16, 0x67,
	0xd6, 0x6c, 0xe2, 0xb9, 0xf8, 0x7d, 0xed, 0xfb, 0xfd, 0x4b, 0x8d, 0xe2, 0xc1, 0xed, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xed, 0x24, 0x15, 0xfb, 0xa1, 0x01, 0x00, 0x00,
}
