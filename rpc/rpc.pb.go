// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	Engine
	Info
	Identity
	Frontend
	Empty
	EngineSpec
	EngineStatus
	StartEngineRequest
	StopEngineRequest
	GetEngineRequest
	EngineResponse
	ListEnginesResponse
*/
package rpc

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Engine struct {
	Binary   string   `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
	Replicas []string `protobuf:"bytes,2,rep,name=replicas" json:"replicas,omitempty"`
}

func (m *Engine) Reset()                    { *m = Engine{} }
func (m *Engine) String() string            { return proto.CompactTextString(m) }
func (*Engine) ProtoMessage()               {}
func (*Engine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Engine) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *Engine) GetReplicas() []string {
	if m != nil {
		return m.Replicas
	}
	return nil
}

type Info struct {
	Volume   string `protobuf:"bytes,1,opt,name=volume" json:"volume,omitempty"`
	Frontend string `protobuf:"bytes,2,opt,name=frontend" json:"frontend,omitempty"`
	Endpoint string `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Info) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *Info) GetFrontend() string {
	if m != nil {
		return m.Frontend
	}
	return ""
}

func (m *Info) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type Identity struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (m *Identity) String() string            { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Identity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Frontend struct {
	Frontend string `protobuf:"bytes,1,opt,name=Frontend" json:"Frontend,omitempty"`
}

func (m *Frontend) Reset()                    { *m = Frontend{} }
func (m *Frontend) String() string            { return proto.CompactTextString(m) }
func (*Frontend) ProtoMessage()               {}
func (*Frontend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Frontend) GetFrontend() string {
	if m != nil {
		return m.Frontend
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type EngineSpec struct {
	Name          string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Binary        string   `protobuf:"bytes,2,opt,name=binary" json:"binary,omitempty"`
	Args          []string `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	ReservedPorts []int32  `protobuf:"varint,4,rep,packed,name=reserved_ports,json=reservedPorts" json:"reserved_ports,omitempty"`
}

func (m *EngineSpec) Reset()                    { *m = EngineSpec{} }
func (m *EngineSpec) String() string            { return proto.CompactTextString(m) }
func (*EngineSpec) ProtoMessage()               {}
func (*EngineSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EngineSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EngineSpec) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *EngineSpec) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *EngineSpec) GetReservedPorts() []int32 {
	if m != nil {
		return m.ReservedPorts
	}
	return nil
}

type EngineStatus struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *EngineStatus) Reset()                    { *m = EngineStatus{} }
func (m *EngineStatus) String() string            { return proto.CompactTextString(m) }
func (*EngineStatus) ProtoMessage()               {}
func (*EngineStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EngineStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type StartEngineRequest struct {
	Spec *EngineSpec `protobuf:"bytes,1,opt,name=spec" json:"spec,omitempty"`
}

func (m *StartEngineRequest) Reset()                    { *m = StartEngineRequest{} }
func (m *StartEngineRequest) String() string            { return proto.CompactTextString(m) }
func (*StartEngineRequest) ProtoMessage()               {}
func (*StartEngineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *StartEngineRequest) GetSpec() *EngineSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type StopEngineRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StopEngineRequest) Reset()                    { *m = StopEngineRequest{} }
func (m *StopEngineRequest) String() string            { return proto.CompactTextString(m) }
func (*StopEngineRequest) ProtoMessage()               {}
func (*StopEngineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StopEngineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetEngineRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetEngineRequest) Reset()                    { *m = GetEngineRequest{} }
func (m *GetEngineRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEngineRequest) ProtoMessage()               {}
func (*GetEngineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetEngineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EngineResponse struct {
	Spec   *EngineSpec   `protobuf:"bytes,1,opt,name=spec" json:"spec,omitempty"`
	Status *EngineStatus `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *EngineResponse) Reset()                    { *m = EngineResponse{} }
func (m *EngineResponse) String() string            { return proto.CompactTextString(m) }
func (*EngineResponse) ProtoMessage()               {}
func (*EngineResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *EngineResponse) GetSpec() *EngineSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *EngineResponse) GetStatus() *EngineStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ListEnginesResponse struct {
	Engines []*EngineResponse `protobuf:"bytes,1,rep,name=engines" json:"engines,omitempty"`
}

func (m *ListEnginesResponse) Reset()                    { *m = ListEnginesResponse{} }
func (m *ListEnginesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEnginesResponse) ProtoMessage()               {}
func (*ListEnginesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListEnginesResponse) GetEngines() []*EngineResponse {
	if m != nil {
		return m.Engines
	}
	return nil
}

func init() {
	proto.RegisterType((*Engine)(nil), "Engine")
	proto.RegisterType((*Info)(nil), "Info")
	proto.RegisterType((*Identity)(nil), "Identity")
	proto.RegisterType((*Frontend)(nil), "Frontend")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*EngineSpec)(nil), "EngineSpec")
	proto.RegisterType((*EngineStatus)(nil), "EngineStatus")
	proto.RegisterType((*StartEngineRequest)(nil), "StartEngineRequest")
	proto.RegisterType((*StopEngineRequest)(nil), "StopEngineRequest")
	proto.RegisterType((*GetEngineRequest)(nil), "GetEngineRequest")
	proto.RegisterType((*EngineResponse)(nil), "EngineResponse")
	proto.RegisterType((*ListEnginesResponse)(nil), "ListEnginesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LonghornLauncherService service

type LonghornLauncherServiceClient interface {
	UpgradeEngine(ctx context.Context, in *Engine, opts ...grpc.CallOption) (*Empty, error)
	GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Info, error)
	StartFrontend(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Empty, error)
	ShutdownFrontend(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Empty, error)
	StartEngineFrontend(ctx context.Context, in *Frontend, opts ...grpc.CallOption) (*Empty, error)
	ShutdownEngineFrontend(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type longhornLauncherServiceClient struct {
	cc *grpc.ClientConn
}

func NewLonghornLauncherServiceClient(cc *grpc.ClientConn) LonghornLauncherServiceClient {
	return &longhornLauncherServiceClient{cc}
}

func (c *longhornLauncherServiceClient) UpgradeEngine(ctx context.Context, in *Engine, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/LonghornLauncherService/UpgradeEngine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longhornLauncherServiceClient) GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := grpc.Invoke(ctx, "/LonghornLauncherService/GetInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longhornLauncherServiceClient) StartFrontend(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/LonghornLauncherService/StartFrontend", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longhornLauncherServiceClient) ShutdownFrontend(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/LonghornLauncherService/ShutdownFrontend", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longhornLauncherServiceClient) StartEngineFrontend(ctx context.Context, in *Frontend, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/LonghornLauncherService/StartEngineFrontend", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longhornLauncherServiceClient) ShutdownEngineFrontend(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/LonghornLauncherService/ShutdownEngineFrontend", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LonghornLauncherService service

type LonghornLauncherServiceServer interface {
	UpgradeEngine(context.Context, *Engine) (*Empty, error)
	GetInfo(context.Context, *Empty) (*Info, error)
	StartFrontend(context.Context, *Identity) (*Empty, error)
	ShutdownFrontend(context.Context, *Identity) (*Empty, error)
	StartEngineFrontend(context.Context, *Frontend) (*Empty, error)
	ShutdownEngineFrontend(context.Context, *Empty) (*Empty, error)
}

func RegisterLonghornLauncherServiceServer(s *grpc.Server, srv LonghornLauncherServiceServer) {
	s.RegisterService(&_LonghornLauncherService_serviceDesc, srv)
}

func _LonghornLauncherService_UpgradeEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Engine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornLauncherServiceServer).UpgradeEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornLauncherService/UpgradeEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornLauncherServiceServer).UpgradeEngine(ctx, req.(*Engine))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonghornLauncherService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornLauncherServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornLauncherService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornLauncherServiceServer).GetInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonghornLauncherService_StartFrontend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornLauncherServiceServer).StartFrontend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornLauncherService/StartFrontend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornLauncherServiceServer).StartFrontend(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonghornLauncherService_ShutdownFrontend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornLauncherServiceServer).ShutdownFrontend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornLauncherService/ShutdownFrontend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornLauncherServiceServer).ShutdownFrontend(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonghornLauncherService_StartEngineFrontend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Frontend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornLauncherServiceServer).StartEngineFrontend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornLauncherService/StartEngineFrontend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornLauncherServiceServer).StartEngineFrontend(ctx, req.(*Frontend))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonghornLauncherService_ShutdownEngineFrontend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornLauncherServiceServer).ShutdownEngineFrontend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornLauncherService/ShutdownEngineFrontend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornLauncherServiceServer).ShutdownEngineFrontend(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LonghornLauncherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LonghornLauncherService",
	HandlerType: (*LonghornLauncherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpgradeEngine",
			Handler:    _LonghornLauncherService_UpgradeEngine_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _LonghornLauncherService_GetInfo_Handler,
		},
		{
			MethodName: "StartFrontend",
			Handler:    _LonghornLauncherService_StartFrontend_Handler,
		},
		{
			MethodName: "ShutdownFrontend",
			Handler:    _LonghornLauncherService_ShutdownFrontend_Handler,
		},
		{
			MethodName: "StartEngineFrontend",
			Handler:    _LonghornLauncherService_StartEngineFrontend_Handler,
		},
		{
			MethodName: "ShutdownEngineFrontend",
			Handler:    _LonghornLauncherService_ShutdownEngineFrontend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// Client API for LonghornEngineLauncherService service

type LonghornEngineLauncherServiceClient interface {
	StartEngine(ctx context.Context, in *StartEngineRequest, opts ...grpc.CallOption) (*EngineResponse, error)
	StopEngine(ctx context.Context, in *StopEngineRequest, opts ...grpc.CallOption) (*EngineResponse, error)
	GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*EngineResponse, error)
	ListEngines(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListEnginesResponse, error)
}

type longhornEngineLauncherServiceClient struct {
	cc *grpc.ClientConn
}

func NewLonghornEngineLauncherServiceClient(cc *grpc.ClientConn) LonghornEngineLauncherServiceClient {
	return &longhornEngineLauncherServiceClient{cc}
}

func (c *longhornEngineLauncherServiceClient) StartEngine(ctx context.Context, in *StartEngineRequest, opts ...grpc.CallOption) (*EngineResponse, error) {
	out := new(EngineResponse)
	err := grpc.Invoke(ctx, "/LonghornEngineLauncherService/StartEngine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longhornEngineLauncherServiceClient) StopEngine(ctx context.Context, in *StopEngineRequest, opts ...grpc.CallOption) (*EngineResponse, error) {
	out := new(EngineResponse)
	err := grpc.Invoke(ctx, "/LonghornEngineLauncherService/StopEngine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longhornEngineLauncherServiceClient) GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*EngineResponse, error) {
	out := new(EngineResponse)
	err := grpc.Invoke(ctx, "/LonghornEngineLauncherService/GetEngine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *longhornEngineLauncherServiceClient) ListEngines(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListEnginesResponse, error) {
	out := new(ListEnginesResponse)
	err := grpc.Invoke(ctx, "/LonghornEngineLauncherService/ListEngines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LonghornEngineLauncherService service

type LonghornEngineLauncherServiceServer interface {
	StartEngine(context.Context, *StartEngineRequest) (*EngineResponse, error)
	StopEngine(context.Context, *StopEngineRequest) (*EngineResponse, error)
	GetEngine(context.Context, *GetEngineRequest) (*EngineResponse, error)
	ListEngines(context.Context, *Empty) (*ListEnginesResponse, error)
}

func RegisterLonghornEngineLauncherServiceServer(s *grpc.Server, srv LonghornEngineLauncherServiceServer) {
	s.RegisterService(&_LonghornEngineLauncherService_serviceDesc, srv)
}

func _LonghornEngineLauncherService_StartEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornEngineLauncherServiceServer).StartEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornEngineLauncherService/StartEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornEngineLauncherServiceServer).StartEngine(ctx, req.(*StartEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonghornEngineLauncherService_StopEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornEngineLauncherServiceServer).StopEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornEngineLauncherService/StopEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornEngineLauncherServiceServer).StopEngine(ctx, req.(*StopEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonghornEngineLauncherService_GetEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornEngineLauncherServiceServer).GetEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornEngineLauncherService/GetEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornEngineLauncherServiceServer).GetEngine(ctx, req.(*GetEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LonghornEngineLauncherService_ListEngines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornEngineLauncherServiceServer).ListEngines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornEngineLauncherService/ListEngines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornEngineLauncherServiceServer).ListEngines(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LonghornEngineLauncherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LonghornEngineLauncherService",
	HandlerType: (*LonghornEngineLauncherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartEngine",
			Handler:    _LonghornEngineLauncherService_StartEngine_Handler,
		},
		{
			MethodName: "StopEngine",
			Handler:    _LonghornEngineLauncherService_StopEngine_Handler,
		},
		{
			MethodName: "GetEngine",
			Handler:    _LonghornEngineLauncherService_GetEngine_Handler,
		},
		{
			MethodName: "ListEngines",
			Handler:    _LonghornEngineLauncherService_ListEngines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdf, 0x6b, 0xdb, 0x30,
	0x10, 0x4e, 0x9c, 0xdf, 0x97, 0x25, 0x6b, 0x95, 0xd1, 0x19, 0xc3, 0x58, 0x10, 0x6b, 0x9a, 0x0d,
	0x26, 0x58, 0x4a, 0xdf, 0xf6, 0xb0, 0x87, 0x76, 0x25, 0x90, 0x87, 0xe1, 0xb0, 0xb1, 0xb7, 0xe1,
	0xda, 0x6a, 0x62, 0x68, 0x25, 0x4d, 0x92, 0x33, 0xf2, 0xb6, 0x7f, 0x78, 0xff, 0xc3, 0xb0, 0x6c,
	0xa9, 0x4e, 0x1a, 0xda, 0xb7, 0xfb, 0xee, 0xbe, 0x3b, 0xdd, 0x77, 0xfe, 0x30, 0xf4, 0xa4, 0x88,
	0x89, 0x90, 0x5c, 0x73, 0xfc, 0x19, 0xda, 0x57, 0x6c, 0x95, 0x32, 0x8a, 0x4e, 0xa0, 0x7d, 0x93,
	0xb2, 0x48, 0x6e, 0xfd, 0xfa, 0xb8, 0x3e, 0xed, 0x85, 0x25, 0x42, 0x01, 0x74, 0x25, 0x15, 0x77,
	0x69, 0x1c, 0x29, 0xdf, 0x1b, 0x37, 0xa6, 0xbd, 0xd0, 0x61, 0xfc, 0x03, 0x9a, 0x73, 0x76, 0xcb,
	0xf3, 0xde, 0x0d, 0xbf, 0xcb, 0xee, 0xa9, 0xed, 0x2d, 0x50, 0xde, 0x7b, 0x2b, 0x39, 0xd3, 0x94,
	0x25, 0xbe, 0x67, 0x2a, 0x0e, 0xe7, 0x35, 0xca, 0x12, 0xc1, 0x53, 0xa6, 0xfd, 0x46, 0x51, 0xb3,
	0x18, 0x07, 0xd0, 0x9d, 0x27, 0x94, 0xe9, 0x54, 0x6f, 0xd1, 0x10, 0xbc, 0xf9, 0x65, 0x39, 0xd7,
	0x9b, 0x5f, 0xe2, 0x09, 0x74, 0xbf, 0x56, 0x66, 0xd8, 0xb8, 0x64, 0x38, 0x8c, 0x3b, 0xd0, 0xba,
	0xba, 0x17, 0x7a, 0x8b, 0x15, 0x40, 0x21, 0x71, 0x29, 0x68, 0x8c, 0x10, 0x34, 0x59, 0xe4, 0x16,
	0x35, 0x71, 0x45, 0xba, 0xb7, 0x23, 0x1d, 0x41, 0x33, 0x92, 0x2b, 0xe5, 0x37, 0x8c, 0x6c, 0x13,
	0xa3, 0x53, 0x18, 0x4a, 0xaa, 0xa8, 0xdc, 0xd0, 0xe4, 0x97, 0xe0, 0x52, 0x2b, 0xbf, 0x39, 0x6e,
	0x4c, 0x5b, 0xe1, 0xc0, 0x66, 0xbf, 0xe5, 0x49, 0x3c, 0x81, 0x17, 0xe5, 0xa3, 0x3a, 0xd2, 0x99,
	0xca, 0x9f, 0x50, 0x26, 0xb2, 0x17, 0x2a, 0x10, 0xbe, 0x00, 0xb4, 0xd4, 0x91, 0xd4, 0x05, 0x39,
	0xa4, 0xbf, 0x33, 0xaa, 0x34, 0x7a, 0x0b, 0x4d, 0x25, 0x68, 0x6c, 0xb8, 0xfd, 0x59, 0x9f, 0x3c,
	0xec, 0x1f, 0x9a, 0x02, 0x3e, 0x83, 0xe3, 0xa5, 0xe6, 0x62, 0xb7, 0xeb, 0x80, 0x34, 0x3c, 0x81,
	0xa3, 0x6b, 0xaa, 0x9f, 0xe7, 0xfd, 0x84, 0xa1, 0x25, 0x29, 0xc1, 0x99, 0xa2, 0xcf, 0xee, 0x80,
	0x4e, 0x9d, 0x24, 0xcf, 0x50, 0x06, 0xa4, 0xaa, 0xd8, 0x29, 0xfc, 0x02, 0xa3, 0x45, 0xaa, 0xca,
	0x15, 0x94, 0x1b, 0xff, 0x1e, 0x3a, 0xb4, 0x48, 0xf9, 0xf5, 0x71, 0x63, 0xda, 0x9f, 0xbd, 0x24,
	0xbb, 0x0b, 0x84, 0xb6, 0x3e, 0xfb, 0xeb, 0xc1, 0xeb, 0x05, 0x67, 0xab, 0x35, 0x97, 0x6c, 0x11,
	0x65, 0x2c, 0x5e, 0x53, 0xb9, 0xa4, 0x72, 0x93, 0xc6, 0x14, 0x61, 0x18, 0x7c, 0x17, 0x2b, 0x19,
	0x25, 0xb4, 0xb4, 0x71, 0xa7, 0x1c, 0x13, 0xb4, 0x49, 0xf1, 0xf9, 0x6b, 0x28, 0x80, 0xce, 0x35,
	0xd5, 0xc6, 0xa8, 0x65, 0x32, 0x68, 0x91, 0x1c, 0xe2, 0x1a, 0x7a, 0x07, 0x03, 0x73, 0x7f, 0x67,
	0xa9, 0x1e, 0xb1, 0xce, 0xab, 0x4c, 0x38, 0x83, 0xa3, 0xe5, 0x3a, 0xd3, 0x09, 0xff, 0xc3, 0x9e,
	0x26, 0x7e, 0x80, 0x51, 0xe5, 0x73, 0x56, 0xb8, 0x36, 0xdc, 0xe1, 0x9e, 0xd8, 0xa1, 0x7b, 0x74,
	0xbb, 0xa5, 0xe3, 0xce, 0xfe, 0xd5, 0xe1, 0x8d, 0x3d, 0x41, 0x41, 0xde, 0x3f, 0xc4, 0x05, 0xf4,
	0x2b, 0x2f, 0xa3, 0x11, 0x79, 0x6c, 0xab, 0x60, 0xff, 0xc4, 0xb8, 0x86, 0xce, 0x01, 0x1e, 0x8c,
	0x84, 0x10, 0x79, 0xe4, 0xaa, 0x43, 0x4d, 0x9f, 0xa0, 0xe7, 0x4c, 0x85, 0x8e, 0xc9, 0xbe, 0xc1,
	0x0e, 0xb5, 0x7c, 0x84, 0x7e, 0xc5, 0x05, 0x4e, 0xe1, 0x2b, 0x72, 0xc0, 0x1b, 0xb8, 0x76, 0xd3,
	0x36, 0x7f, 0xa7, 0xf3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xf9, 0xf0, 0x9f, 0xaa, 0x04,
	0x00, 0x00,
}
