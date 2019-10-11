// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_hook.proto

package gorr

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GrpcHookRequest struct {
	ReqId                int32    `protobuf:"varint,1,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	ReqName              string   `protobuf:"bytes,2,opt,name=req_name,json=reqName,proto3" json:"req_name,omitempty"`
	ReqData              string   `protobuf:"bytes,3,opt,name=req_data,json=reqData,proto3" json:"req_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcHookRequest) Reset()         { *m = GrpcHookRequest{} }
func (m *GrpcHookRequest) String() string { return proto.CompactTextString(m) }
func (*GrpcHookRequest) ProtoMessage()    {}
func (*GrpcHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a09e6bcbc888de9, []int{0}
}

func (m *GrpcHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcHookRequest.Unmarshal(m, b)
}
func (m *GrpcHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcHookRequest.Marshal(b, m, deterministic)
}
func (m *GrpcHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcHookRequest.Merge(m, src)
}
func (m *GrpcHookRequest) XXX_Size() int {
	return xxx_messageInfo_GrpcHookRequest.Size(m)
}
func (m *GrpcHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcHookRequest proto.InternalMessageInfo

func (m *GrpcHookRequest) GetReqId() int32 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *GrpcHookRequest) GetReqName() string {
	if m != nil {
		return m.ReqName
	}
	return ""
}

func (m *GrpcHookRequest) GetReqData() string {
	if m != nil {
		return m.ReqData
	}
	return ""
}

type GrpcHookResponse struct {
	RspId                int32    `protobuf:"varint,1,opt,name=rsp_id,json=rspId,proto3" json:"rsp_id,omitempty"`
	ReqId                int32    `protobuf:"varint,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	RspName              string   `protobuf:"bytes,3,opt,name=rsp_name,json=rspName,proto3" json:"rsp_name,omitempty"`
	RspData              string   `protobuf:"bytes,4,opt,name=rsp_data,json=rspData,proto3" json:"rsp_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcHookResponse) Reset()         { *m = GrpcHookResponse{} }
func (m *GrpcHookResponse) String() string { return proto.CompactTextString(m) }
func (*GrpcHookResponse) ProtoMessage()    {}
func (*GrpcHookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a09e6bcbc888de9, []int{1}
}

func (m *GrpcHookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcHookResponse.Unmarshal(m, b)
}
func (m *GrpcHookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcHookResponse.Marshal(b, m, deterministic)
}
func (m *GrpcHookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcHookResponse.Merge(m, src)
}
func (m *GrpcHookResponse) XXX_Size() int {
	return xxx_messageInfo_GrpcHookResponse.Size(m)
}
func (m *GrpcHookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcHookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcHookResponse proto.InternalMessageInfo

func (m *GrpcHookResponse) GetRspId() int32 {
	if m != nil {
		return m.RspId
	}
	return 0
}

func (m *GrpcHookResponse) GetReqId() int32 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *GrpcHookResponse) GetRspName() string {
	if m != nil {
		return m.RspName
	}
	return ""
}

func (m *GrpcHookResponse) GetRspData() string {
	if m != nil {
		return m.RspData
	}
	return ""
}

func init() {
	proto.RegisterType((*GrpcHookRequest)(nil), "grpc_hook.GrpcHookRequest")
	proto.RegisterType((*GrpcHookResponse)(nil), "grpc_hook.GrpcHookResponse")
}

func init() { proto.RegisterFile("grpc_hook.proto", fileDescriptor_9a09e6bcbc888de9) }

var fileDescriptor_9a09e6bcbc888de9 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2f, 0x2a, 0x48,
	0x8e, 0xcf, 0xc8, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xc5, 0x71, 0xf1, 0xbb, 0x17, 0x15, 0x24, 0x7b, 0xe4, 0xe7, 0x67, 0x07, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0x89, 0x72, 0xb1, 0x15, 0xa5, 0x16, 0xc6, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0xb0, 0x06, 0xb1, 0x16, 0xa5, 0x16, 0x7a, 0xa6, 0x08, 0x49, 0x72, 0x71, 0x80, 0x84, 0xf3,
	0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xd8, 0x8b, 0x52, 0x0b, 0xfd, 0x12,
	0x73, 0x53, 0x61, 0x52, 0x29, 0x89, 0x25, 0x89, 0x12, 0xcc, 0x70, 0x29, 0x97, 0xc4, 0x92, 0x44,
	0xa5, 0x32, 0x2e, 0x01, 0x84, 0xf9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x60, 0x0b, 0x8a, 0x0b,
	0x90, 0x2d, 0x28, 0x2e, 0xf0, 0x4c, 0x41, 0xb2, 0x97, 0x09, 0xdd, 0xde, 0xe2, 0x02, 0x88, 0xbd,
	0x30, 0xc3, 0x8b, 0x0b, 0xe0, 0xf6, 0x16, 0x17, 0x40, 0xec, 0x65, 0x81, 0x4b, 0x81, 0xec, 0x35,
	0x8a, 0x40, 0xf8, 0x2b, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x95, 0x8b, 0x23, 0x38,
	0x3f, 0x37, 0xd5, 0x39, 0x31, 0x27, 0x47, 0x48, 0x4a, 0x0f, 0x11, 0x26, 0x68, 0xfe, 0x97, 0x92,
	0xc6, 0x2a, 0x07, 0x71, 0xbb, 0x12, 0x43, 0x12, 0x1b, 0x38, 0x0c, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0x1e, 0xb4, 0x44, 0x56, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GrpcHookServiceClient is the client API for GrpcHookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcHookServiceClient interface {
	SomeCall(ctx context.Context, in *GrpcHookRequest, opts ...grpc.CallOption) (*GrpcHookResponse, error)
}

type grpcHookServiceClient struct {
	cc *grpc.ClientConn
}

func NewGrpcHookServiceClient(cc *grpc.ClientConn) GrpcHookServiceClient {
	return &grpcHookServiceClient{cc}
}

func (c *grpcHookServiceClient) SomeCall(ctx context.Context, in *GrpcHookRequest, opts ...grpc.CallOption) (*GrpcHookResponse, error) {
	out := new(GrpcHookResponse)
	err := c.cc.Invoke(ctx, "/grpc_hook.GrpcHookService/SomeCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcHookServiceServer is the server API for GrpcHookService service.
type GrpcHookServiceServer interface {
	SomeCall(context.Context, *GrpcHookRequest) (*GrpcHookResponse, error)
}

// UnimplementedGrpcHookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcHookServiceServer struct {
}

var _GrpcHookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_hook.GrpcHookService",
	HandlerType: (*GrpcHookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SomeCall",
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_hook.proto",
}
