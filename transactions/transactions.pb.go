// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transactions.proto

package transactions

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Transactions struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount               string   `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transactions) Reset()         { *m = Transactions{} }
func (m *Transactions) String() string { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()    {}
func (*Transactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b72849cf10e9c77, []int{0}
}

func (m *Transactions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transactions.Unmarshal(m, b)
}
func (m *Transactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transactions.Marshal(b, m, deterministic)
}
func (m *Transactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transactions.Merge(m, src)
}
func (m *Transactions) XXX_Size() int {
	return xxx_messageInfo_Transactions.Size(m)
}
func (m *Transactions) XXX_DiscardUnknown() {
	xxx_messageInfo_Transactions.DiscardUnknown(m)
}

var xxx_messageInfo_Transactions proto.InternalMessageInfo

func (m *Transactions) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Transactions) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Transactions) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type TransactionsRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Branch               string   `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionsRequest) Reset()         { *m = TransactionsRequest{} }
func (m *TransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionsRequest) ProtoMessage()    {}
func (*TransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b72849cf10e9c77, []int{1}
}

func (m *TransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionsRequest.Unmarshal(m, b)
}
func (m *TransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionsRequest.Marshal(b, m, deterministic)
}
func (m *TransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionsRequest.Merge(m, src)
}
func (m *TransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionsRequest.Size(m)
}
func (m *TransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionsRequest proto.InternalMessageInfo

func (m *TransactionsRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *TransactionsRequest) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

type TransactionsReply struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Transactions         *Transactions `protobuf:"bytes,2,opt,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TransactionsReply) Reset()         { *m = TransactionsReply{} }
func (m *TransactionsReply) String() string { return proto.CompactTextString(m) }
func (*TransactionsReply) ProtoMessage()    {}
func (*TransactionsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b72849cf10e9c77, []int{2}
}

func (m *TransactionsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionsReply.Unmarshal(m, b)
}
func (m *TransactionsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionsReply.Marshal(b, m, deterministic)
}
func (m *TransactionsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionsReply.Merge(m, src)
}
func (m *TransactionsReply) XXX_Size() int {
	return xxx_messageInfo_TransactionsReply.Size(m)
}
func (m *TransactionsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionsReply.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionsReply proto.InternalMessageInfo

func (m *TransactionsReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TransactionsReply) GetTransactions() *Transactions {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*Transactions)(nil), "transactions.Transactions")
	proto.RegisterType((*TransactionsRequest)(nil), "transactions.TransactionsRequest")
	proto.RegisterType((*TransactionsReply)(nil), "transactions.TransactionsReply")
}

func init() { proto.RegisterFile("transactions.proto", fileDescriptor_0b72849cf10e9c77) }

var fileDescriptor_0b72849cf10e9c77 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0x6d, 0x95, 0x13, 0xdf, 0x1d, 0x8a, 0x4f, 0x90, 0x70, 0x8b, 0x67, 0x26, 0xa7, 0x43,
	0xea, 0xee, 0xda, 0xbd, 0xb8, 0x08, 0x2e, 0x69, 0x12, 0x34, 0x50, 0x93, 0x98, 0xa4, 0x43, 0xff,
	0x7b, 0x69, 0x9a, 0x42, 0xb2, 0x74, 0xcb, 0xf7, 0xbe, 0x8f, 0xdf, 0xe3, 0x7d, 0x01, 0x0c, 0x8e,
	0x69, 0xcf, 0x78, 0x50, 0x46, 0xfb, 0xb3, 0x75, 0x26, 0x18, 0x3c, 0xe4, 0x33, 0xfa, 0x05, 0x87,
	0x8f, 0x4c, 0x23, 0xc2, 0x95, 0x60, 0x41, 0x92, 0xea, 0x54, 0xbd, 0xdc, 0x74, 0xf1, 0x8d, 0x27,
	0xd8, 0x0b, 0xe9, 0xb9, 0x53, 0x76, 0xce, 0x90, 0x3a, 0x5a, 0xf9, 0x08, 0x1f, 0x61, 0xc7, 0x7e,
	0xcd, 0xa8, 0x03, 0xb9, 0x8c, 0x66, 0x52, 0xb4, 0x85, 0x87, 0x9c, 0xde, 0xc9, 0xbf, 0x51, 0xfa,
	0x80, 0x04, 0xae, 0x19, 0xe7, 0x31, 0xbf, 0xec, 0x59, 0xe5, 0x0c, 0xea, 0x1d, 0xd3, 0xfc, 0x27,
	0x6d, 0x49, 0x8a, 0x72, 0xb8, 0x2f, 0x41, 0x76, 0x98, 0xf0, 0x16, 0x6a, 0x25, 0x12, 0xa1, 0x56,
	0x02, 0xdf, 0xa1, 0xb8, 0x2d, 0x22, 0xf6, 0xcd, 0xf1, 0x5c, 0x94, 0x50, 0x60, 0x8a, 0x7c, 0xf3,
	0x0d, 0xb0, 0xba, 0xc6, 0xe1, 0x27, 0xdc, 0xb5, 0x32, 0x14, 0xe5, 0x3c, 0x6f, 0xa0, 0x96, 0xd3,
	0x8e, 0x4f, 0x5b, 0x11, 0x3b, 0x4c, 0xf4, 0xe2, 0xb5, 0xea, 0x77, 0xf1, 0x27, 0xde, 0xfe, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x97, 0xb4, 0x1a, 0x9f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransactorClient is the client API for Transactor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactorClient interface {
	// Get transactions
	GetTransactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (Transactor_GetTransactionsClient, error)
}

type transactorClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactorClient(cc grpc.ClientConnInterface) TransactorClient {
	return &transactorClient{cc}
}

func (c *transactorClient) GetTransactions(ctx context.Context, in *TransactionsRequest, opts ...grpc.CallOption) (Transactor_GetTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Transactor_serviceDesc.Streams[0], "/transactions.Transactor/GetTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &transactorGetTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Transactor_GetTransactionsClient interface {
	Recv() (*TransactionsReply, error)
	grpc.ClientStream
}

type transactorGetTransactionsClient struct {
	grpc.ClientStream
}

func (x *transactorGetTransactionsClient) Recv() (*TransactionsReply, error) {
	m := new(TransactionsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransactorServer is the server API for Transactor service.
type TransactorServer interface {
	// Get transactions
	GetTransactions(*TransactionsRequest, Transactor_GetTransactionsServer) error
}

// UnimplementedTransactorServer can be embedded to have forward compatible implementations.
type UnimplementedTransactorServer struct {
}

func (*UnimplementedTransactorServer) GetTransactions(req *TransactionsRequest, srv Transactor_GetTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}

func RegisterTransactorServer(s *grpc.Server, srv TransactorServer) {
	s.RegisterService(&_Transactor_serviceDesc, srv)
}

func _Transactor_GetTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransactorServer).GetTransactions(m, &transactorGetTransactionsServer{stream})
}

type Transactor_GetTransactionsServer interface {
	Send(*TransactionsReply) error
	grpc.ServerStream
}

type transactorGetTransactionsServer struct {
	grpc.ServerStream
}

func (x *transactorGetTransactionsServer) Send(m *TransactionsReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Transactor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transactions.Transactor",
	HandlerType: (*TransactorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTransactions",
			Handler:       _Transactor_GetTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transactions.proto",
}
