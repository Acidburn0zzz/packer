// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/dataproc/v1/subcluster.proto

package dataproc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type Role int32

const (
	Role_ROLE_UNSPECIFIED Role = 0
	// The subcluster fulfills the master role.
	//
	// Master can run the following services, depending on the requested components:
	// * HDFS: Namenode, Secondary Namenode
	// * YARN: ResourceManager, Timeline Server
	// * HBase Master
	// * Hive: Server, Metastore, HCatalog
	// * Spark History Server
	// * Zeppelin
	// * ZooKeeper
	Role_MASTERNODE Role = 1
	// The subcluster is a DATANODE in a Data Proc cluster.
	//
	// DATANODE can run the following services, depending on the requested components:
	// * HDFS DataNode
	// * YARN NodeManager
	// * HBase RegionServer
	// * Spark libraries
	Role_DATANODE Role = 2
	// The subcluster is a COMPUTENODE in a Data Proc cluster.
	//
	// COMPUTENODE can run the following services, depending on the requested components:
	// * YARN NodeManager
	// * Spark libraries
	Role_COMPUTENODE Role = 3
)

var Role_name = map[int32]string{
	0: "ROLE_UNSPECIFIED",
	1: "MASTERNODE",
	2: "DATANODE",
	3: "COMPUTENODE",
}

var Role_value = map[string]int32{
	"ROLE_UNSPECIFIED": 0,
	"MASTERNODE":       1,
	"DATANODE":         2,
	"COMPUTENODE":      3,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3761a92b6a1bd471, []int{0}
}

// A Data Proc subcluster. For details about the concept, see [documentation](/docs/data-proc/concepts/).
type Subcluster struct {
	// ID of the subcluster. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the Data Proc cluster that the subcluster belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the subcluster. The name is unique within the cluster.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Role that is fulfilled by hosts of the subcluster.
	Role Role `protobuf:"varint,5,opt,name=role,proto3,enum=yandex.cloud.dataproc.v1.Role" json:"role,omitempty"`
	// Resources allocated for each host in the subcluster.
	Resources *Resources `protobuf:"bytes,6,opt,name=resources,proto3" json:"resources,omitempty"`
	// ID of the VPC subnet used for hosts in the subcluster.
	SubnetId string `protobuf:"bytes,7,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Number of hosts in the subcluster.
	HostsCount           int64    `protobuf:"varint,8,opt,name=hosts_count,json=hostsCount,proto3" json:"hosts_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subcluster) Reset()         { *m = Subcluster{} }
func (m *Subcluster) String() string { return proto.CompactTextString(m) }
func (*Subcluster) ProtoMessage()    {}
func (*Subcluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_3761a92b6a1bd471, []int{0}
}

func (m *Subcluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subcluster.Unmarshal(m, b)
}
func (m *Subcluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subcluster.Marshal(b, m, deterministic)
}
func (m *Subcluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subcluster.Merge(m, src)
}
func (m *Subcluster) XXX_Size() int {
	return xxx_messageInfo_Subcluster.Size(m)
}
func (m *Subcluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Subcluster.DiscardUnknown(m)
}

var xxx_messageInfo_Subcluster proto.InternalMessageInfo

func (m *Subcluster) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subcluster) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *Subcluster) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Subcluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subcluster) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (m *Subcluster) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Subcluster) GetSubnetId() string {
	if m != nil {
		return m.SubnetId
	}
	return ""
}

func (m *Subcluster) GetHostsCount() int64 {
	if m != nil {
		return m.HostsCount
	}
	return 0
}

// A Data Proc host. For details about the concept, see [documentation](/docs/data-proc/concepts/).
type Host struct {
	// Name of the Data Proc host. The host name is assigned by Data Proc at creation time
	// and cannot be changed. The name is generated to be unique across all existing Data Proc
	// hosts in Yandex.Cloud, as it defines the FQDN of the host.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the Data Proc subcluster that the host belongs to.
	SubclusterId string `protobuf:"bytes,2,opt,name=subcluster_id,json=subclusterId,proto3" json:"subcluster_id,omitempty"`
	// Host status code.
	Health Health `protobuf:"varint,3,opt,name=health,proto3,enum=yandex.cloud.dataproc.v1.Health" json:"health,omitempty"`
	// ID of the Compute virtual machine that is used as the Data Proc host.
	ComputeInstanceId string `protobuf:"bytes,4,opt,name=compute_instance_id,json=computeInstanceId,proto3" json:"compute_instance_id,omitempty"`
	// Role of the host in the cluster.
	Role                 Role     `protobuf:"varint,5,opt,name=role,proto3,enum=yandex.cloud.dataproc.v1.Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_3761a92b6a1bd471, []int{1}
}

func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Host) GetSubclusterId() string {
	if m != nil {
		return m.SubclusterId
	}
	return ""
}

func (m *Host) GetHealth() Health {
	if m != nil {
		return m.Health
	}
	return Health_HEALTH_UNKNOWN
}

func (m *Host) GetComputeInstanceId() string {
	if m != nil {
		return m.ComputeInstanceId
	}
	return ""
}

func (m *Host) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("yandex.cloud.dataproc.v1.Role", Role_name, Role_value)
	proto.RegisterType((*Subcluster)(nil), "yandex.cloud.dataproc.v1.Subcluster")
	proto.RegisterType((*Host)(nil), "yandex.cloud.dataproc.v1.Host")
}

func init() {
	proto.RegisterFile("yandex/cloud/dataproc/v1/subcluster.proto", fileDescriptor_3761a92b6a1bd471)
}

var fileDescriptor_3761a92b6a1bd471 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xd1, 0x8e, 0xd2, 0x4e,
	0x14, 0xc6, 0xff, 0x85, 0xfe, 0x11, 0x0e, 0x2b, 0xe2, 0xe8, 0x45, 0x83, 0xbb, 0x2e, 0xd9, 0x8d,
	0x09, 0x9a, 0x30, 0x0d, 0x6c, 0x62, 0x34, 0x5e, 0xb1, 0x50, 0xb3, 0x8d, 0xee, 0xb2, 0x19, 0xd8,
	0x1b, 0x6f, 0xc8, 0x30, 0x33, 0x42, 0x93, 0xb6, 0xd3, 0x74, 0xa6, 0x44, 0x5f, 0xc1, 0x27, 0xd3,
	0x87, 0xf0, 0x5d, 0x4c, 0xa7, 0xad, 0xb8, 0x17, 0x78, 0xe1, 0x5d, 0xe7, 0x3b, 0xbf, 0xd3, 0xf3,
	0x9d, 0xaf, 0x53, 0x78, 0xf9, 0x95, 0xc6, 0x5c, 0x7c, 0x71, 0x59, 0x28, 0x33, 0xee, 0x72, 0xaa,
	0x69, 0x92, 0x4a, 0xe6, 0xee, 0x46, 0xae, 0xca, 0xd6, 0x2c, 0xcc, 0x94, 0x16, 0x29, 0x4e, 0x52,
	0xa9, 0x25, 0x72, 0x0a, 0x14, 0x1b, 0x14, 0x57, 0x28, 0xde, 0x8d, 0x7a, 0xa7, 0x1b, 0x29, 0x37,
	0xa1, 0x70, 0x0d, 0xb7, 0xce, 0x3e, 0xbb, 0x3a, 0x88, 0x84, 0xd2, 0x34, 0x4a, 0x8a, 0xd6, 0xde,
	0x8b, 0x83, 0x53, 0x98, 0x8c, 0x22, 0x19, 0x97, 0xd8, 0xc9, 0x3d, 0x6c, 0x47, 0xc3, 0x80, 0x53,
	0x1d, 0x54, 0xe5, 0xb3, 0x1f, 0x35, 0x80, 0xc5, 0x6f, 0x57, 0xa8, 0x03, 0xb5, 0x80, 0x3b, 0x56,
	0xdf, 0x1a, 0xb4, 0x48, 0x2d, 0xe0, 0xe8, 0x04, 0xa0, 0x2c, 0xad, 0x02, 0xee, 0xd4, 0x8c, 0xde,
	0x2a, 0x15, 0x9f, 0xa3, 0xb7, 0x00, 0x2c, 0x15, 0x54, 0x0b, 0xbe, 0xa2, 0xda, 0xa9, 0xf7, 0xad,
	0x41, 0x7b, 0xdc, 0xc3, 0x85, 0x73, 0x5c, 0x39, 0xc7, 0xcb, 0xca, 0x39, 0x69, 0x95, 0xf4, 0x44,
	0xa3, 0x63, 0xb0, 0x63, 0x1a, 0x09, 0xc7, 0xce, 0xdf, 0x79, 0xd9, 0xfc, 0xf6, 0x7d, 0x64, 0x8f,
	0x86, 0xaf, 0x2f, 0x88, 0x51, 0xd1, 0x18, 0xec, 0x54, 0x86, 0xc2, 0xf9, 0xbf, 0x6f, 0x0d, 0x3a,
	0xe3, 0xe7, 0xf8, 0x50, 0x4c, 0x98, 0xc8, 0x50, 0x10, 0xc3, 0xa2, 0x09, 0xb4, 0x52, 0xa1, 0x64,
	0x96, 0x32, 0xa1, 0x9c, 0x86, 0xf1, 0x72, 0xfe, 0x97, 0xc6, 0x0a, 0x25, 0xfb, 0x2e, 0xf4, 0x0c,
	0x5a, 0x2a, 0x5b, 0xc7, 0x42, 0xe7, 0xdb, 0x3e, 0x30, 0xdb, 0x36, 0x0b, 0xc1, 0xe7, 0xe8, 0x14,
	0xda, 0x5b, 0xa9, 0xb4, 0x5a, 0x31, 0x99, 0xc5, 0xda, 0x69, 0xf6, 0xad, 0x41, 0x9d, 0x80, 0x91,
	0xa6, 0xb9, 0x72, 0xf6, 0xd3, 0x02, 0xfb, 0x4a, 0x2a, 0x8d, 0x50, 0xb9, 0x5b, 0x91, 0x63, 0xb1,
	0xd1, 0x39, 0x3c, 0xdc, 0x7f, 0xfd, 0x7d, 0x98, 0x47, 0x7b, 0xd1, 0xe7, 0xe8, 0x0d, 0x34, 0xb6,
	0x82, 0x86, 0x7a, 0x6b, 0xb2, 0xec, 0x8c, 0xfb, 0x87, 0xfd, 0x5f, 0x19, 0x8e, 0x94, 0x3c, 0xc2,
	0xf0, 0x84, 0xc9, 0x28, 0xc9, 0xb4, 0x58, 0x05, 0xb1, 0xd2, 0x34, 0x66, 0x22, 0x1f, 0x62, 0xd2,
	0x25, 0x8f, 0xcb, 0x92, 0x5f, 0x56, 0x7c, 0xfe, 0x2f, 0x01, 0xbf, 0xfa, 0x00, 0x76, 0x7e, 0x42,
	0x4f, 0xa1, 0x4b, 0xe6, 0x1f, 0xbd, 0xd5, 0xdd, 0xcd, 0xe2, 0xd6, 0x9b, 0xfa, 0xef, 0x7d, 0x6f,
	0xd6, 0xfd, 0x0f, 0x75, 0x00, 0xae, 0x27, 0x8b, 0xa5, 0x47, 0x6e, 0xe6, 0x33, 0xaf, 0x6b, 0xa1,
	0x23, 0x68, 0xce, 0x26, 0xcb, 0x89, 0x39, 0xd5, 0xd0, 0x23, 0x68, 0x4f, 0xe7, 0xd7, 0xb7, 0x77,
	0x4b, 0xcf, 0x08, 0xf5, 0x4b, 0x01, 0xc7, 0xf7, 0x66, 0xd2, 0x24, 0xf8, 0x73, 0xee, 0x27, 0x6f,
	0x13, 0xe8, 0x6d, 0xb6, 0xc6, 0x4c, 0x46, 0x6e, 0x01, 0x0e, 0x8b, 0x2b, 0xbc, 0x91, 0xc3, 0x8d,
	0x88, 0xcd, 0xe5, 0x72, 0x0f, 0xfd, 0x02, 0xef, 0xaa, 0xe7, 0x75, 0xc3, 0x80, 0x17, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x71, 0xfc, 0x2b, 0xc1, 0x94, 0x03, 0x00, 0x00,
}
