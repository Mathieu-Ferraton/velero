// Code generated by protoc-gen-go. DO NOT EDIT.
// source: BlockStore.proto

package generated

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateVolumeRequest struct {
	SnapshotID string `protobuf:"bytes,1,opt,name=snapshotID" json:"snapshotID,omitempty"`
	VolumeType string `protobuf:"bytes,2,opt,name=volumeType" json:"volumeType,omitempty"`
	VolumeAZ   string `protobuf:"bytes,3,opt,name=volumeAZ" json:"volumeAZ,omitempty"`
	Iops       int64  `protobuf:"varint,4,opt,name=iops" json:"iops,omitempty"`
}

func (m *CreateVolumeRequest) Reset()                    { *m = CreateVolumeRequest{} }
func (m *CreateVolumeRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateVolumeRequest) ProtoMessage()               {}
func (*CreateVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateVolumeRequest) GetSnapshotID() string {
	if m != nil {
		return m.SnapshotID
	}
	return ""
}

func (m *CreateVolumeRequest) GetVolumeType() string {
	if m != nil {
		return m.VolumeType
	}
	return ""
}

func (m *CreateVolumeRequest) GetVolumeAZ() string {
	if m != nil {
		return m.VolumeAZ
	}
	return ""
}

func (m *CreateVolumeRequest) GetIops() int64 {
	if m != nil {
		return m.Iops
	}
	return 0
}

type CreateVolumeResponse struct {
	VolumeID string `protobuf:"bytes,1,opt,name=volumeID" json:"volumeID,omitempty"`
}

func (m *CreateVolumeResponse) Reset()                    { *m = CreateVolumeResponse{} }
func (m *CreateVolumeResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateVolumeResponse) ProtoMessage()               {}
func (*CreateVolumeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CreateVolumeResponse) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

type GetVolumeInfoRequest struct {
	VolumeID string `protobuf:"bytes,1,opt,name=volumeID" json:"volumeID,omitempty"`
	VolumeAZ string `protobuf:"bytes,2,opt,name=volumeAZ" json:"volumeAZ,omitempty"`
}

func (m *GetVolumeInfoRequest) Reset()                    { *m = GetVolumeInfoRequest{} }
func (m *GetVolumeInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVolumeInfoRequest) ProtoMessage()               {}
func (*GetVolumeInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetVolumeInfoRequest) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

func (m *GetVolumeInfoRequest) GetVolumeAZ() string {
	if m != nil {
		return m.VolumeAZ
	}
	return ""
}

type GetVolumeInfoResponse struct {
	VolumeType string `protobuf:"bytes,1,opt,name=volumeType" json:"volumeType,omitempty"`
	Iops       int64  `protobuf:"varint,2,opt,name=iops" json:"iops,omitempty"`
}

func (m *GetVolumeInfoResponse) Reset()                    { *m = GetVolumeInfoResponse{} }
func (m *GetVolumeInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVolumeInfoResponse) ProtoMessage()               {}
func (*GetVolumeInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetVolumeInfoResponse) GetVolumeType() string {
	if m != nil {
		return m.VolumeType
	}
	return ""
}

func (m *GetVolumeInfoResponse) GetIops() int64 {
	if m != nil {
		return m.Iops
	}
	return 0
}

type IsVolumeReadyRequest struct {
	VolumeID string `protobuf:"bytes,1,opt,name=volumeID" json:"volumeID,omitempty"`
	VolumeAZ string `protobuf:"bytes,2,opt,name=volumeAZ" json:"volumeAZ,omitempty"`
}

func (m *IsVolumeReadyRequest) Reset()                    { *m = IsVolumeReadyRequest{} }
func (m *IsVolumeReadyRequest) String() string            { return proto.CompactTextString(m) }
func (*IsVolumeReadyRequest) ProtoMessage()               {}
func (*IsVolumeReadyRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *IsVolumeReadyRequest) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

func (m *IsVolumeReadyRequest) GetVolumeAZ() string {
	if m != nil {
		return m.VolumeAZ
	}
	return ""
}

type IsVolumeReadyResponse struct {
	Ready bool `protobuf:"varint,1,opt,name=ready" json:"ready,omitempty"`
}

func (m *IsVolumeReadyResponse) Reset()                    { *m = IsVolumeReadyResponse{} }
func (m *IsVolumeReadyResponse) String() string            { return proto.CompactTextString(m) }
func (*IsVolumeReadyResponse) ProtoMessage()               {}
func (*IsVolumeReadyResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *IsVolumeReadyResponse) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

type ListSnapshotsRequest struct {
	TagFilters map[string]string `protobuf:"bytes,1,rep,name=tagFilters" json:"tagFilters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListSnapshotsRequest) Reset()                    { *m = ListSnapshotsRequest{} }
func (m *ListSnapshotsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSnapshotsRequest) ProtoMessage()               {}
func (*ListSnapshotsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ListSnapshotsRequest) GetTagFilters() map[string]string {
	if m != nil {
		return m.TagFilters
	}
	return nil
}

type ListSnapshotsResponse struct {
	SnapshotIDs []string `protobuf:"bytes,2,rep,name=snapshotIDs" json:"snapshotIDs,omitempty"`
}

func (m *ListSnapshotsResponse) Reset()                    { *m = ListSnapshotsResponse{} }
func (m *ListSnapshotsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSnapshotsResponse) ProtoMessage()               {}
func (*ListSnapshotsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ListSnapshotsResponse) GetSnapshotIDs() []string {
	if m != nil {
		return m.SnapshotIDs
	}
	return nil
}

type CreateSnapshotRequest struct {
	VolumeID string            `protobuf:"bytes,1,opt,name=volumeID" json:"volumeID,omitempty"`
	VolumeAZ string            `protobuf:"bytes,2,opt,name=volumeAZ" json:"volumeAZ,omitempty"`
	Tags     map[string]string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateSnapshotRequest) Reset()                    { *m = CreateSnapshotRequest{} }
func (m *CreateSnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSnapshotRequest) ProtoMessage()               {}
func (*CreateSnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *CreateSnapshotRequest) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

func (m *CreateSnapshotRequest) GetVolumeAZ() string {
	if m != nil {
		return m.VolumeAZ
	}
	return ""
}

func (m *CreateSnapshotRequest) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreateSnapshotResponse struct {
	SnapshotID string `protobuf:"bytes,1,opt,name=snapshotID" json:"snapshotID,omitempty"`
}

func (m *CreateSnapshotResponse) Reset()                    { *m = CreateSnapshotResponse{} }
func (m *CreateSnapshotResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateSnapshotResponse) ProtoMessage()               {}
func (*CreateSnapshotResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *CreateSnapshotResponse) GetSnapshotID() string {
	if m != nil {
		return m.SnapshotID
	}
	return ""
}

type DeleteSnapshotRequest struct {
	SnapshotID string `protobuf:"bytes,1,opt,name=snapshotID" json:"snapshotID,omitempty"`
}

func (m *DeleteSnapshotRequest) Reset()                    { *m = DeleteSnapshotRequest{} }
func (m *DeleteSnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSnapshotRequest) ProtoMessage()               {}
func (*DeleteSnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *DeleteSnapshotRequest) GetSnapshotID() string {
	if m != nil {
		return m.SnapshotID
	}
	return ""
}

type GetVolumeIDRequest struct {
	PersistentVolume []byte `protobuf:"bytes,1,opt,name=persistentVolume,proto3" json:"persistentVolume,omitempty"`
}

func (m *GetVolumeIDRequest) Reset()                    { *m = GetVolumeIDRequest{} }
func (m *GetVolumeIDRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVolumeIDRequest) ProtoMessage()               {}
func (*GetVolumeIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *GetVolumeIDRequest) GetPersistentVolume() []byte {
	if m != nil {
		return m.PersistentVolume
	}
	return nil
}

type GetVolumeIDResponse struct {
	VolumeID string `protobuf:"bytes,1,opt,name=volumeID" json:"volumeID,omitempty"`
}

func (m *GetVolumeIDResponse) Reset()                    { *m = GetVolumeIDResponse{} }
func (m *GetVolumeIDResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVolumeIDResponse) ProtoMessage()               {}
func (*GetVolumeIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *GetVolumeIDResponse) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

type SetVolumeIDRequest struct {
	PersistentVolume []byte `protobuf:"bytes,1,opt,name=persistentVolume,proto3" json:"persistentVolume,omitempty"`
	VolumeID         string `protobuf:"bytes,2,opt,name=volumeID" json:"volumeID,omitempty"`
}

func (m *SetVolumeIDRequest) Reset()                    { *m = SetVolumeIDRequest{} }
func (m *SetVolumeIDRequest) String() string            { return proto.CompactTextString(m) }
func (*SetVolumeIDRequest) ProtoMessage()               {}
func (*SetVolumeIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *SetVolumeIDRequest) GetPersistentVolume() []byte {
	if m != nil {
		return m.PersistentVolume
	}
	return nil
}

func (m *SetVolumeIDRequest) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

type SetVolumeIDResponse struct {
	PersistentVolume []byte `protobuf:"bytes,1,opt,name=persistentVolume,proto3" json:"persistentVolume,omitempty"`
}

func (m *SetVolumeIDResponse) Reset()                    { *m = SetVolumeIDResponse{} }
func (m *SetVolumeIDResponse) String() string            { return proto.CompactTextString(m) }
func (*SetVolumeIDResponse) ProtoMessage()               {}
func (*SetVolumeIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *SetVolumeIDResponse) GetPersistentVolume() []byte {
	if m != nil {
		return m.PersistentVolume
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateVolumeRequest)(nil), "generated.CreateVolumeRequest")
	proto.RegisterType((*CreateVolumeResponse)(nil), "generated.CreateVolumeResponse")
	proto.RegisterType((*GetVolumeInfoRequest)(nil), "generated.GetVolumeInfoRequest")
	proto.RegisterType((*GetVolumeInfoResponse)(nil), "generated.GetVolumeInfoResponse")
	proto.RegisterType((*IsVolumeReadyRequest)(nil), "generated.IsVolumeReadyRequest")
	proto.RegisterType((*IsVolumeReadyResponse)(nil), "generated.IsVolumeReadyResponse")
	proto.RegisterType((*ListSnapshotsRequest)(nil), "generated.ListSnapshotsRequest")
	proto.RegisterType((*ListSnapshotsResponse)(nil), "generated.ListSnapshotsResponse")
	proto.RegisterType((*CreateSnapshotRequest)(nil), "generated.CreateSnapshotRequest")
	proto.RegisterType((*CreateSnapshotResponse)(nil), "generated.CreateSnapshotResponse")
	proto.RegisterType((*DeleteSnapshotRequest)(nil), "generated.DeleteSnapshotRequest")
	proto.RegisterType((*GetVolumeIDRequest)(nil), "generated.GetVolumeIDRequest")
	proto.RegisterType((*GetVolumeIDResponse)(nil), "generated.GetVolumeIDResponse")
	proto.RegisterType((*SetVolumeIDRequest)(nil), "generated.SetVolumeIDRequest")
	proto.RegisterType((*SetVolumeIDResponse)(nil), "generated.SetVolumeIDResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlockStore service

type BlockStoreClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateVolumeFromSnapshot(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*CreateVolumeResponse, error)
	GetVolumeInfo(ctx context.Context, in *GetVolumeInfoRequest, opts ...grpc.CallOption) (*GetVolumeInfoResponse, error)
	IsVolumeReady(ctx context.Context, in *IsVolumeReadyRequest, opts ...grpc.CallOption) (*IsVolumeReadyResponse, error)
	ListSnapshots(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error)
	CreateSnapshot(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error)
	DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*Empty, error)
	GetVolumeID(ctx context.Context, in *GetVolumeIDRequest, opts ...grpc.CallOption) (*GetVolumeIDResponse, error)
	SetVolumeID(ctx context.Context, in *SetVolumeIDRequest, opts ...grpc.CallOption) (*SetVolumeIDResponse, error)
}

type blockStoreClient struct {
	cc *grpc.ClientConn
}

func NewBlockStoreClient(cc *grpc.ClientConn) BlockStoreClient {
	return &blockStoreClient{cc}
}

func (c *blockStoreClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/generated.BlockStore/Init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) CreateVolumeFromSnapshot(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*CreateVolumeResponse, error) {
	out := new(CreateVolumeResponse)
	err := grpc.Invoke(ctx, "/generated.BlockStore/CreateVolumeFromSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) GetVolumeInfo(ctx context.Context, in *GetVolumeInfoRequest, opts ...grpc.CallOption) (*GetVolumeInfoResponse, error) {
	out := new(GetVolumeInfoResponse)
	err := grpc.Invoke(ctx, "/generated.BlockStore/GetVolumeInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) IsVolumeReady(ctx context.Context, in *IsVolumeReadyRequest, opts ...grpc.CallOption) (*IsVolumeReadyResponse, error) {
	out := new(IsVolumeReadyResponse)
	err := grpc.Invoke(ctx, "/generated.BlockStore/IsVolumeReady", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) ListSnapshots(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error) {
	out := new(ListSnapshotsResponse)
	err := grpc.Invoke(ctx, "/generated.BlockStore/ListSnapshots", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) CreateSnapshot(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error) {
	out := new(CreateSnapshotResponse)
	err := grpc.Invoke(ctx, "/generated.BlockStore/CreateSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/generated.BlockStore/DeleteSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) GetVolumeID(ctx context.Context, in *GetVolumeIDRequest, opts ...grpc.CallOption) (*GetVolumeIDResponse, error) {
	out := new(GetVolumeIDResponse)
	err := grpc.Invoke(ctx, "/generated.BlockStore/GetVolumeID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) SetVolumeID(ctx context.Context, in *SetVolumeIDRequest, opts ...grpc.CallOption) (*SetVolumeIDResponse, error) {
	out := new(SetVolumeIDResponse)
	err := grpc.Invoke(ctx, "/generated.BlockStore/SetVolumeID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlockStore service

type BlockStoreServer interface {
	Init(context.Context, *InitRequest) (*Empty, error)
	CreateVolumeFromSnapshot(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error)
	GetVolumeInfo(context.Context, *GetVolumeInfoRequest) (*GetVolumeInfoResponse, error)
	IsVolumeReady(context.Context, *IsVolumeReadyRequest) (*IsVolumeReadyResponse, error)
	ListSnapshots(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error)
	CreateSnapshot(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error)
	DeleteSnapshot(context.Context, *DeleteSnapshotRequest) (*Empty, error)
	GetVolumeID(context.Context, *GetVolumeIDRequest) (*GetVolumeIDResponse, error)
	SetVolumeID(context.Context, *SetVolumeIDRequest) (*SetVolumeIDResponse, error)
}

func RegisterBlockStoreServer(s *grpc.Server, srv BlockStoreServer) {
	s.RegisterService(&_BlockStore_serviceDesc, srv)
}

func _BlockStore_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BlockStore/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_CreateVolumeFromSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).CreateVolumeFromSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BlockStore/CreateVolumeFromSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).CreateVolumeFromSnapshot(ctx, req.(*CreateVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_GetVolumeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolumeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).GetVolumeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BlockStore/GetVolumeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).GetVolumeInfo(ctx, req.(*GetVolumeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_IsVolumeReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsVolumeReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).IsVolumeReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BlockStore/IsVolumeReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).IsVolumeReady(ctx, req.(*IsVolumeReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_ListSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).ListSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BlockStore/ListSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).ListSnapshots(ctx, req.(*ListSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_CreateSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).CreateSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BlockStore/CreateSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).CreateSnapshot(ctx, req.(*CreateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_DeleteSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).DeleteSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BlockStore/DeleteSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).DeleteSnapshot(ctx, req.(*DeleteSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_GetVolumeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolumeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).GetVolumeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BlockStore/GetVolumeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).GetVolumeID(ctx, req.(*GetVolumeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_SetVolumeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVolumeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).SetVolumeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.BlockStore/SetVolumeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).SetVolumeID(ctx, req.(*SetVolumeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.BlockStore",
	HandlerType: (*BlockStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _BlockStore_Init_Handler,
		},
		{
			MethodName: "CreateVolumeFromSnapshot",
			Handler:    _BlockStore_CreateVolumeFromSnapshot_Handler,
		},
		{
			MethodName: "GetVolumeInfo",
			Handler:    _BlockStore_GetVolumeInfo_Handler,
		},
		{
			MethodName: "IsVolumeReady",
			Handler:    _BlockStore_IsVolumeReady_Handler,
		},
		{
			MethodName: "ListSnapshots",
			Handler:    _BlockStore_ListSnapshots_Handler,
		},
		{
			MethodName: "CreateSnapshot",
			Handler:    _BlockStore_CreateSnapshot_Handler,
		},
		{
			MethodName: "DeleteSnapshot",
			Handler:    _BlockStore_DeleteSnapshot_Handler,
		},
		{
			MethodName: "GetVolumeID",
			Handler:    _BlockStore_GetVolumeID_Handler,
		},
		{
			MethodName: "SetVolumeID",
			Handler:    _BlockStore_SetVolumeID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BlockStore.proto",
}

func init() { proto.RegisterFile("BlockStore.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0x56, 0x9a, 0xee, 0xd7, 0x7a, 0xba, 0xed, 0xaf, 0xdc, 0x76, 0x8a, 0x22, 0x51, 0x42, 0xae,
	0xaa, 0x49, 0x14, 0x28, 0x17, 0x1b, 0x48, 0x20, 0x06, 0xdd, 0x50, 0x45, 0x35, 0xa4, 0x64, 0x70,
	0x01, 0xdc, 0x04, 0x6a, 0xba, 0x6a, 0x6d, 0x1c, 0x6c, 0x77, 0x52, 0x1f, 0x80, 0x57, 0xe1, 0x59,
	0xb8, 0xe4, 0x91, 0x50, 0x12, 0x27, 0xb1, 0x53, 0xb7, 0x63, 0xea, 0x5d, 0x7d, 0x4e, 0xbe, 0xcf,
	0xdf, 0x39, 0x3e, 0xe7, 0x2b, 0x34, 0x5e, 0xcf, 0xc8, 0xb7, 0x6b, 0x9f, 0x13, 0x8a, 0x7b, 0x11,
	0x25, 0x9c, 0xa0, 0xda, 0x04, 0x87, 0x98, 0x06, 0x1c, 0x8f, 0xed, 0x3d, 0xff, 0x2a, 0xa0, 0x78,
	0x9c, 0x26, 0xdc, 0x9f, 0x06, 0x34, 0xdf, 0x50, 0x1c, 0x70, 0xfc, 0x91, 0xcc, 0x16, 0x73, 0xec,
	0xe1, 0x1f, 0x0b, 0xcc, 0x38, 0xea, 0x00, 0xb0, 0x30, 0x88, 0xd8, 0x15, 0xe1, 0xc3, 0x81, 0x65,
	0x38, 0x46, 0xb7, 0xe6, 0x49, 0x91, 0x38, 0x7f, 0x93, 0x00, 0x2e, 0x97, 0x11, 0xb6, 0x2a, 0x69,
	0xbe, 0x88, 0x20, 0x1b, 0x76, 0xd3, 0xd3, 0xe9, 0x27, 0xcb, 0x4c, 0xb2, 0xf9, 0x19, 0x21, 0xa8,
	0x4e, 0x49, 0xc4, 0xac, 0xaa, 0x63, 0x74, 0x4d, 0x2f, 0xf9, 0xed, 0xf6, 0xa1, 0xa5, 0xca, 0x60,
	0x11, 0x09, 0x99, 0xc4, 0x93, 0xab, 0xc8, 0xcf, 0xee, 0x05, 0xb4, 0xde, 0x62, 0x9e, 0x02, 0x86,
	0xe1, 0x77, 0x92, 0x69, 0xdf, 0x80, 0x51, 0x74, 0x55, 0x54, 0x5d, 0xee, 0x3b, 0x68, 0x97, 0xf8,
	0x84, 0x08, 0xb5, 0x58, 0x63, 0xa5, 0xd8, 0xac, 0xa0, 0x8a, 0x54, 0xd0, 0x05, 0xb4, 0x86, 0x2c,
	0x2b, 0x26, 0x18, 0x2f, 0xb7, 0x15, 0xf7, 0x10, 0xda, 0x25, 0x3e, 0x21, 0xae, 0x05, 0x3b, 0x34,
	0x0e, 0x24, 0x6c, 0xbb, 0x5e, 0x7a, 0x70, 0x7f, 0x19, 0xd0, 0x1a, 0x4d, 0x19, 0xf7, 0xc5, 0x93,
	0xb1, 0xec, 0xfe, 0xf7, 0x00, 0x3c, 0x98, 0x9c, 0x4f, 0x67, 0x1c, 0x53, 0x66, 0x19, 0x8e, 0xd9,
	0xad, 0xf7, 0x1f, 0xf5, 0xf2, 0xf1, 0xe8, 0xe9, 0x40, 0xbd, 0xcb, 0x1c, 0x71, 0x16, 0x72, 0xba,
	0xf4, 0x24, 0x0a, 0xfb, 0x05, 0xfc, 0x5f, 0x4a, 0xa3, 0x06, 0x98, 0xd7, 0x78, 0x29, 0xca, 0x8b,
	0x7f, 0xc6, 0x22, 0x6f, 0x82, 0xd9, 0x22, 0x9b, 0x94, 0xf4, 0xf0, 0xbc, 0x72, 0x62, 0xb8, 0xcf,
	0xa0, 0x5d, 0xba, 0x52, 0xd4, 0xe5, 0x40, 0xbd, 0x98, 0xb7, 0xb8, 0xb7, 0x66, 0xb7, 0xe6, 0xc9,
	0x21, 0xf7, 0xb7, 0x01, 0xed, 0x74, 0x68, 0x32, 0xf4, 0x96, 0x4d, 0x46, 0x2f, 0xa1, 0xca, 0x83,
	0x09, 0xb3, 0xcc, 0xa4, 0x2d, 0x47, 0x52, 0x5b, 0xb4, 0xf7, 0xc4, 0x7d, 0x11, 0x1d, 0x49, 0x70,
	0xf6, 0x31, 0xd4, 0xf2, 0xd0, 0x9d, 0xba, 0x70, 0x02, 0x87, 0xe5, 0x1b, 0x8a, 0xd9, 0xdb, 0xb4,
	0x88, 0xee, 0x31, 0xb4, 0x07, 0x78, 0x86, 0x57, 0x7b, 0x70, 0x1b, 0xf0, 0x15, 0xa0, 0x62, 0xda,
	0x07, 0x19, 0xea, 0x08, 0x1a, 0x11, 0xa6, 0x6c, 0xca, 0x38, 0x0e, 0x45, 0x32, 0xc1, 0xee, 0x79,
	0x2b, 0x71, 0xf7, 0x09, 0x34, 0x15, 0x86, 0x7f, 0x58, 0xd9, 0x2f, 0x80, 0xfc, 0xad, 0x2e, 0x55,
	0xd8, 0x2b, 0x25, 0xf6, 0x53, 0x68, 0xfa, 0x1a, 0x41, 0x77, 0xa0, 0xef, 0xff, 0xd9, 0x01, 0x28,
	0xdc, 0x13, 0x3d, 0x86, 0xea, 0x30, 0x9c, 0x72, 0x74, 0x28, 0x8d, 0x42, 0x1c, 0x10, 0xca, 0xed,
	0x86, 0x14, 0x3f, 0x9b, 0x47, 0x7c, 0x89, 0x3e, 0x83, 0x25, 0x1b, 0xd9, 0x39, 0x25, 0xf3, 0xec,
	0x65, 0x50, 0x67, 0x65, 0xa0, 0x14, 0xd3, 0xb5, 0xef, 0xaf, 0xcd, 0x8b, 0x4a, 0x3c, 0xd8, 0x57,
	0x1c, 0x0a, 0xc9, 0x08, 0x9d, 0x17, 0xda, 0xce, 0xfa, 0x0f, 0x0a, 0x4e, 0xc5, 0x58, 0x14, 0x4e,
	0x9d, 0x85, 0x29, 0x9c, 0x7a, 0x4f, 0xf2, 0x60, 0x5f, 0x59, 0x6a, 0x85, 0x53, 0xe7, 0x30, 0x0a,
	0xa7, 0xde, 0x0f, 0x3e, 0xc0, 0x81, 0xba, 0x22, 0xc8, 0xb9, 0x6d, 0x3f, 0xed, 0x07, 0x1b, 0xbe,
	0x10, 0xb4, 0x03, 0x38, 0x50, 0xf7, 0x47, 0xa1, 0xd5, 0xae, 0x96, 0xe6, 0xd5, 0x47, 0x50, 0x97,
	0x56, 0x01, 0xdd, 0xd3, 0x76, 0x3d, 0x9b, 0x77, 0xbb, 0xb3, 0x2e, 0x2d, 0x34, 0x8d, 0xa0, 0xee,
	0xaf, 0x61, 0xf3, 0x37, 0xb3, 0x69, 0xc6, 0xff, 0xeb, 0x7f, 0xc9, 0x3f, 0xfd, 0xd3, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xe7, 0x50, 0x0a, 0x1d, 0x16, 0x08, 0x00, 0x00,
}
