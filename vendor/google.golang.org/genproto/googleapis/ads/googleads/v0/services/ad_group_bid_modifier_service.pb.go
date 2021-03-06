// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/ad_group_bid_modifier_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

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

// Request message for [AdGroupBidModifierService.GetAdGroupBidModifier][google.ads.googleads.v0.services.AdGroupBidModifierService.GetAdGroupBidModifier].
type GetAdGroupBidModifierRequest struct {
	// The resource name of the ad group bid modifier to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupBidModifierRequest) Reset()         { *m = GetAdGroupBidModifierRequest{} }
func (m *GetAdGroupBidModifierRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupBidModifierRequest) ProtoMessage()    {}
func (*GetAdGroupBidModifierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_bid_modifier_service_fbd69257eb9b6c8d, []int{0}
}
func (m *GetAdGroupBidModifierRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupBidModifierRequest.Unmarshal(m, b)
}
func (m *GetAdGroupBidModifierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupBidModifierRequest.Marshal(b, m, deterministic)
}
func (dst *GetAdGroupBidModifierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupBidModifierRequest.Merge(dst, src)
}
func (m *GetAdGroupBidModifierRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupBidModifierRequest.Size(m)
}
func (m *GetAdGroupBidModifierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupBidModifierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupBidModifierRequest proto.InternalMessageInfo

func (m *GetAdGroupBidModifierRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupBidModifierService.MutateAdGroupBidModifiers][google.ads.googleads.v0.services.AdGroupBidModifierService.MutateAdGroupBidModifiers].
type MutateAdGroupBidModifiersRequest struct {
	// ID of the customer whose ad group bid modifiers are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual ad group bid modifiers.
	Operations           []*AdGroupBidModifierOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MutateAdGroupBidModifiersRequest) Reset()         { *m = MutateAdGroupBidModifiersRequest{} }
func (m *MutateAdGroupBidModifiersRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupBidModifiersRequest) ProtoMessage()    {}
func (*MutateAdGroupBidModifiersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_bid_modifier_service_fbd69257eb9b6c8d, []int{1}
}
func (m *MutateAdGroupBidModifiersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupBidModifiersRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupBidModifiersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupBidModifiersRequest.Marshal(b, m, deterministic)
}
func (dst *MutateAdGroupBidModifiersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupBidModifiersRequest.Merge(dst, src)
}
func (m *MutateAdGroupBidModifiersRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupBidModifiersRequest.Size(m)
}
func (m *MutateAdGroupBidModifiersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupBidModifiersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupBidModifiersRequest proto.InternalMessageInfo

func (m *MutateAdGroupBidModifiersRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupBidModifiersRequest) GetOperations() []*AdGroupBidModifierOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation (create, remove, update) on an ad group bid modifier.
type AdGroupBidModifierOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupBidModifierOperation_Create
	//	*AdGroupBidModifierOperation_Update
	//	*AdGroupBidModifierOperation_Remove
	Operation            isAdGroupBidModifierOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *AdGroupBidModifierOperation) Reset()         { *m = AdGroupBidModifierOperation{} }
func (m *AdGroupBidModifierOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupBidModifierOperation) ProtoMessage()    {}
func (*AdGroupBidModifierOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_bid_modifier_service_fbd69257eb9b6c8d, []int{2}
}
func (m *AdGroupBidModifierOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupBidModifierOperation.Unmarshal(m, b)
}
func (m *AdGroupBidModifierOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupBidModifierOperation.Marshal(b, m, deterministic)
}
func (dst *AdGroupBidModifierOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupBidModifierOperation.Merge(dst, src)
}
func (m *AdGroupBidModifierOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupBidModifierOperation.Size(m)
}
func (m *AdGroupBidModifierOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupBidModifierOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupBidModifierOperation proto.InternalMessageInfo

func (m *AdGroupBidModifierOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdGroupBidModifierOperation_Operation interface {
	isAdGroupBidModifierOperation_Operation()
}

type AdGroupBidModifierOperation_Create struct {
	Create *resources.AdGroupBidModifier `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupBidModifierOperation_Update struct {
	Update *resources.AdGroupBidModifier `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupBidModifierOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupBidModifierOperation_Create) isAdGroupBidModifierOperation_Operation() {}

func (*AdGroupBidModifierOperation_Update) isAdGroupBidModifierOperation_Operation() {}

func (*AdGroupBidModifierOperation_Remove) isAdGroupBidModifierOperation_Operation() {}

func (m *AdGroupBidModifierOperation) GetOperation() isAdGroupBidModifierOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupBidModifierOperation) GetCreate() *resources.AdGroupBidModifier {
	if x, ok := m.GetOperation().(*AdGroupBidModifierOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupBidModifierOperation) GetUpdate() *resources.AdGroupBidModifier {
	if x, ok := m.GetOperation().(*AdGroupBidModifierOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupBidModifierOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupBidModifierOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AdGroupBidModifierOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdGroupBidModifierOperation_OneofMarshaler, _AdGroupBidModifierOperation_OneofUnmarshaler, _AdGroupBidModifierOperation_OneofSizer, []interface{}{
		(*AdGroupBidModifierOperation_Create)(nil),
		(*AdGroupBidModifierOperation_Update)(nil),
		(*AdGroupBidModifierOperation_Remove)(nil),
	}
}

func _AdGroupBidModifierOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdGroupBidModifierOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AdGroupBidModifierOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *AdGroupBidModifierOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *AdGroupBidModifierOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("AdGroupBidModifierOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _AdGroupBidModifierOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdGroupBidModifierOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.AdGroupBidModifier)
		err := b.DecodeMessage(msg)
		m.Operation = &AdGroupBidModifierOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.AdGroupBidModifier)
		err := b.DecodeMessage(msg)
		m.Operation = &AdGroupBidModifierOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &AdGroupBidModifierOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _AdGroupBidModifierOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdGroupBidModifierOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AdGroupBidModifierOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupBidModifierOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupBidModifierOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for ad group bid modifiers mutate.
type MutateAdGroupBidModifiersResponse struct {
	// All results for the mutate.
	Results              []*MutateAdGroupBidModifierResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *MutateAdGroupBidModifiersResponse) Reset()         { *m = MutateAdGroupBidModifiersResponse{} }
func (m *MutateAdGroupBidModifiersResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupBidModifiersResponse) ProtoMessage()    {}
func (*MutateAdGroupBidModifiersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_bid_modifier_service_fbd69257eb9b6c8d, []int{3}
}
func (m *MutateAdGroupBidModifiersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupBidModifiersResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupBidModifiersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupBidModifiersResponse.Marshal(b, m, deterministic)
}
func (dst *MutateAdGroupBidModifiersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupBidModifiersResponse.Merge(dst, src)
}
func (m *MutateAdGroupBidModifiersResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupBidModifiersResponse.Size(m)
}
func (m *MutateAdGroupBidModifiersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupBidModifiersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupBidModifiersResponse proto.InternalMessageInfo

func (m *MutateAdGroupBidModifiersResponse) GetResults() []*MutateAdGroupBidModifierResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateAdGroupBidModifierResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupBidModifierResult) Reset()         { *m = MutateAdGroupBidModifierResult{} }
func (m *MutateAdGroupBidModifierResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupBidModifierResult) ProtoMessage()    {}
func (*MutateAdGroupBidModifierResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_bid_modifier_service_fbd69257eb9b6c8d, []int{4}
}
func (m *MutateAdGroupBidModifierResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupBidModifierResult.Unmarshal(m, b)
}
func (m *MutateAdGroupBidModifierResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupBidModifierResult.Marshal(b, m, deterministic)
}
func (dst *MutateAdGroupBidModifierResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupBidModifierResult.Merge(dst, src)
}
func (m *MutateAdGroupBidModifierResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupBidModifierResult.Size(m)
}
func (m *MutateAdGroupBidModifierResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupBidModifierResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupBidModifierResult proto.InternalMessageInfo

func (m *MutateAdGroupBidModifierResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupBidModifierRequest)(nil), "google.ads.googleads.v0.services.GetAdGroupBidModifierRequest")
	proto.RegisterType((*MutateAdGroupBidModifiersRequest)(nil), "google.ads.googleads.v0.services.MutateAdGroupBidModifiersRequest")
	proto.RegisterType((*AdGroupBidModifierOperation)(nil), "google.ads.googleads.v0.services.AdGroupBidModifierOperation")
	proto.RegisterType((*MutateAdGroupBidModifiersResponse)(nil), "google.ads.googleads.v0.services.MutateAdGroupBidModifiersResponse")
	proto.RegisterType((*MutateAdGroupBidModifierResult)(nil), "google.ads.googleads.v0.services.MutateAdGroupBidModifierResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupBidModifierServiceClient is the client API for AdGroupBidModifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupBidModifierServiceClient interface {
	// Returns the requested ad group bid modifier in full detail.
	GetAdGroupBidModifier(ctx context.Context, in *GetAdGroupBidModifierRequest, opts ...grpc.CallOption) (*resources.AdGroupBidModifier, error)
	// Creates, updates, or removes ad group bid modifiers.
	// Operation statuses are returned.
	MutateAdGroupBidModifiers(ctx context.Context, in *MutateAdGroupBidModifiersRequest, opts ...grpc.CallOption) (*MutateAdGroupBidModifiersResponse, error)
}

type adGroupBidModifierServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupBidModifierServiceClient(cc *grpc.ClientConn) AdGroupBidModifierServiceClient {
	return &adGroupBidModifierServiceClient{cc}
}

func (c *adGroupBidModifierServiceClient) GetAdGroupBidModifier(ctx context.Context, in *GetAdGroupBidModifierRequest, opts ...grpc.CallOption) (*resources.AdGroupBidModifier, error) {
	out := new(resources.AdGroupBidModifier)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdGroupBidModifierService/GetAdGroupBidModifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupBidModifierServiceClient) MutateAdGroupBidModifiers(ctx context.Context, in *MutateAdGroupBidModifiersRequest, opts ...grpc.CallOption) (*MutateAdGroupBidModifiersResponse, error) {
	out := new(MutateAdGroupBidModifiersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdGroupBidModifierService/MutateAdGroupBidModifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupBidModifierServiceServer is the server API for AdGroupBidModifierService service.
type AdGroupBidModifierServiceServer interface {
	// Returns the requested ad group bid modifier in full detail.
	GetAdGroupBidModifier(context.Context, *GetAdGroupBidModifierRequest) (*resources.AdGroupBidModifier, error)
	// Creates, updates, or removes ad group bid modifiers.
	// Operation statuses are returned.
	MutateAdGroupBidModifiers(context.Context, *MutateAdGroupBidModifiersRequest) (*MutateAdGroupBidModifiersResponse, error)
}

func RegisterAdGroupBidModifierServiceServer(s *grpc.Server, srv AdGroupBidModifierServiceServer) {
	s.RegisterService(&_AdGroupBidModifierService_serviceDesc, srv)
}

func _AdGroupBidModifierService_GetAdGroupBidModifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupBidModifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupBidModifierServiceServer).GetAdGroupBidModifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdGroupBidModifierService/GetAdGroupBidModifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupBidModifierServiceServer).GetAdGroupBidModifier(ctx, req.(*GetAdGroupBidModifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupBidModifierService_MutateAdGroupBidModifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupBidModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupBidModifierServiceServer).MutateAdGroupBidModifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdGroupBidModifierService/MutateAdGroupBidModifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupBidModifierServiceServer).MutateAdGroupBidModifiers(ctx, req.(*MutateAdGroupBidModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupBidModifierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.AdGroupBidModifierService",
	HandlerType: (*AdGroupBidModifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupBidModifier",
			Handler:    _AdGroupBidModifierService_GetAdGroupBidModifier_Handler,
		},
		{
			MethodName: "MutateAdGroupBidModifiers",
			Handler:    _AdGroupBidModifierService_MutateAdGroupBidModifiers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/ad_group_bid_modifier_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/ad_group_bid_modifier_service.proto", fileDescriptor_ad_group_bid_modifier_service_fbd69257eb9b6c8d)
}

var fileDescriptor_ad_group_bid_modifier_service_fbd69257eb9b6c8d = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x49, 0x8a, 0x86, 0xe6, 0xc0, 0xc5, 0x12, 0x52, 0x56, 0xa6, 0x11, 0x02, 0x87, 0xaa,
	0x87, 0xa4, 0x2a, 0x9a, 0x40, 0x9b, 0x82, 0x68, 0x0b, 0x74, 0x1c, 0xca, 0xa6, 0x20, 0xed, 0x30,
	0x15, 0x45, 0x6e, 0xed, 0x46, 0xd1, 0x9a, 0x38, 0xd8, 0x49, 0x2f, 0xd3, 0x84, 0xc4, 0x57, 0xe0,
	0x1b, 0x8c, 0x1b, 0x1f, 0x05, 0x89, 0x13, 0x07, 0x6e, 0x9c, 0xb8, 0xf0, 0x2d, 0x90, 0xe3, 0xb8,
	0x0c, 0xb5, 0x69, 0xd1, 0x76, 0x7b, 0x62, 0x3f, 0xfe, 0x3d, 0x2f, 0x7f, 0x3f, 0x0e, 0x78, 0x11,
	0x52, 0x1a, 0x4e, 0x89, 0x8b, 0x30, 0x77, 0xa5, 0x29, 0xac, 0x59, 0xcb, 0xe5, 0x84, 0xcd, 0xa2,
	0x31, 0xe1, 0x2e, 0xc2, 0x41, 0xc8, 0x68, 0x9e, 0x06, 0xa3, 0x08, 0x07, 0x31, 0xc5, 0xd1, 0x24,
	0x22, 0x2c, 0x28, 0xb7, 0x9d, 0x94, 0xd1, 0x8c, 0x42, 0x4b, 0x1e, 0x75, 0x10, 0xe6, 0xce, 0x9c,
	0xe2, 0xcc, 0x5a, 0x8e, 0xa2, 0xd4, 0xbd, 0xaa, 0x38, 0x8c, 0x70, 0x9a, 0xb3, 0xca, 0x40, 0x32,
	0x40, 0x7d, 0x5b, 0x1d, 0x4f, 0x23, 0x17, 0x25, 0x09, 0xcd, 0x50, 0x16, 0xd1, 0x84, 0x97, 0xbb,
	0x65, 0x78, 0xb7, 0xf8, 0x1a, 0xe5, 0x13, 0x77, 0x12, 0x91, 0x29, 0x0e, 0x62, 0xc4, 0x4f, 0xa5,
	0x87, 0xdd, 0x03, 0xdb, 0x7d, 0x92, 0x75, 0x70, 0x5f, 0x04, 0xe8, 0x46, 0x78, 0x50, 0xe2, 0x7d,
	0xf2, 0x3e, 0x27, 0x3c, 0x83, 0x0f, 0xc1, 0x1d, 0x95, 0x48, 0x90, 0xa0, 0x98, 0x98, 0x9a, 0xa5,
	0x35, 0x36, 0xfd, 0xdb, 0x6a, 0xf1, 0x0d, 0x8a, 0x89, 0x7d, 0xa1, 0x01, 0x6b, 0x90, 0x67, 0x28,
	0x23, 0x8b, 0x20, 0xae, 0x48, 0xf7, 0x81, 0x31, 0xce, 0x79, 0x46, 0x63, 0xc2, 0x82, 0x08, 0x97,
	0x1c, 0xa0, 0x96, 0x5e, 0x63, 0xf8, 0x0e, 0x00, 0x9a, 0x12, 0x26, 0x0b, 0x30, 0x75, 0xab, 0xd6,
	0x30, 0xda, 0x9e, 0xb3, 0xae, 0x81, 0xce, 0x62, 0xc8, 0x43, 0x45, 0xf1, 0x2f, 0x01, 0xed, 0xcf,
	0x3a, 0xb8, 0xb7, 0xc2, 0x17, 0xee, 0x03, 0x23, 0x4f, 0x31, 0xca, 0x48, 0xd1, 0x1e, 0xf3, 0xa6,
	0xa5, 0x35, 0x8c, 0x76, 0x5d, 0xc5, 0x57, 0x1d, 0x74, 0x5e, 0x89, 0x0e, 0x0e, 0x10, 0x3f, 0xf5,
	0x81, 0x74, 0x17, 0x36, 0x3c, 0x04, 0x1b, 0x63, 0x46, 0x50, 0x26, 0xfb, 0x63, 0xb4, 0x77, 0x2b,
	0xf3, 0x9e, 0xcb, 0xba, 0x24, 0xf1, 0x83, 0x1b, 0x7e, 0x89, 0x11, 0x40, 0x89, 0x37, 0xf5, 0x6b,
	0x02, 0x25, 0x06, 0x9a, 0x60, 0x83, 0x91, 0x98, 0xce, 0x88, 0x59, 0x13, 0x9d, 0x17, 0x3b, 0xf2,
	0xbb, 0x6b, 0x80, 0xcd, 0x79, 0x9b, 0xec, 0x0f, 0xe0, 0xc1, 0x0a, 0x25, 0x79, 0x4a, 0x13, 0x4e,
	0xe0, 0x09, 0xb8, 0xc5, 0x08, 0xcf, 0xa7, 0x99, 0x92, 0xe9, 0xf9, 0x7a, 0x99, 0xaa, 0xa8, 0x7e,
	0x01, 0xf2, 0x15, 0xd0, 0x7e, 0x09, 0x76, 0x56, 0xbb, 0xfe, 0xd7, 0x95, 0x6c, 0xff, 0xa8, 0x81,
	0xad, 0x45, 0xc2, 0x5b, 0x99, 0x0d, 0xfc, 0xa6, 0x81, 0xbb, 0x4b, 0xaf, 0x3d, 0x7c, 0xb6, 0xbe,
	0x92, 0x55, 0xf3, 0x52, 0xbf, 0x9a, 0x4e, 0xb6, 0xf7, 0xf1, 0xfb, 0xaf, 0x4f, 0xfa, 0x13, 0xb8,
	0x2b, 0x26, 0xff, 0xec, 0x9f, 0xf2, 0x3c, 0x35, 0x22, 0xdc, 0x6d, 0xba, 0x68, 0x51, 0x15, 0xb7,
	0x79, 0x0e, 0x7f, 0x6b, 0x60, 0xab, 0x52, 0x36, 0xd8, 0xbd, 0xba, 0x3a, 0x6a, 0x7a, 0xeb, 0xbd,
	0x6b, 0x31, 0xe4, 0xbd, 0xb1, 0x7b, 0x45, 0x95, 0x9e, 0xfd, 0x54, 0x54, 0xf9, 0xb7, 0xac, 0xb3,
	0x4b, 0xef, 0x82, 0xd7, 0x3c, 0x5f, 0x56, 0xe4, 0x5e, 0x5c, 0xc0, 0xf7, 0xb4, 0x66, 0xf7, 0xa7,
	0x06, 0x1e, 0x8d, 0x69, 0xbc, 0x36, 0x9f, 0xee, 0x4e, 0xa5, 0xfe, 0x47, 0x62, 0x98, 0x8f, 0xb4,
	0x93, 0x83, 0x92, 0x11, 0xd2, 0x29, 0x4a, 0x42, 0x87, 0xb2, 0xd0, 0x0d, 0x49, 0x52, 0x8c, 0xba,
	0x7a, 0x8b, 0xd3, 0x88, 0x57, 0xff, 0x02, 0xf6, 0x95, 0x71, 0xa1, 0xd7, 0xfa, 0x9d, 0xce, 0x17,
	0xdd, 0xea, 0x4b, 0x60, 0x07, 0x73, 0x47, 0x9a, 0xc2, 0x3a, 0x6e, 0x39, 0x65, 0x60, 0xfe, 0x55,
	0xb9, 0x0c, 0x3b, 0x98, 0x0f, 0xe7, 0x2e, 0xc3, 0xe3, 0xd6, 0x50, 0xb9, 0x8c, 0x36, 0x8a, 0x04,
	0x1e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x68, 0xc5, 0x0e, 0x82, 0x06, 0x00, 0x00,
}
