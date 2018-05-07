// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/translator.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/translator.proto

It has these top-level messages:
	TranslateRequest
	TranslateResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Language int32

const (
	Language_en Language = 0
	Language_fr Language = 1
)

var Language_name = map[int32]string{
	0: "en",
	1: "fr",
}
var Language_value = map[string]int32{
	"en": 0,
	"fr": 1,
}

func (x Language) String() string {
	return proto1.EnumName(Language_name, int32(x))
}
func (Language) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TranslateRequest struct {
	Text     string   `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Language Language `protobuf:"varint,2,opt,name=language,enum=proto.Language" json:"language,omitempty"`
}

func (m *TranslateRequest) Reset()                    { *m = TranslateRequest{} }
func (m *TranslateRequest) String() string            { return proto1.CompactTextString(m) }
func (*TranslateRequest) ProtoMessage()               {}
func (*TranslateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TranslateRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *TranslateRequest) GetLanguage() Language {
	if m != nil {
		return m.Language
	}
	return Language_en
}

type TranslateResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *TranslateResponse) Reset()                    { *m = TranslateResponse{} }
func (m *TranslateResponse) String() string            { return proto1.CompactTextString(m) }
func (*TranslateResponse) ProtoMessage()               {}
func (*TranslateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TranslateResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto1.RegisterType((*TranslateRequest)(nil), "proto.TranslateRequest")
	proto1.RegisterType((*TranslateResponse)(nil), "proto.TranslateResponse")
	proto1.RegisterEnum("proto.Language", Language_name, Language_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Translator service

type TranslatorClient interface {
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error)
}

type translatorClient struct {
	cc *grpc.ClientConn
}

func NewTranslatorClient(cc *grpc.ClientConn) TranslatorClient {
	return &translatorClient{cc}
}

func (c *translatorClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error) {
	out := new(TranslateResponse)
	err := grpc.Invoke(ctx, "/proto.Translator/Translate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Translator service

type TranslatorServer interface {
	Translate(context.Context, *TranslateRequest) (*TranslateResponse, error)
}

func RegisterTranslatorServer(s *grpc.Server, srv TranslatorServer) {
	s.RegisterService(&_Translator_serviceDesc, srv)
}

func _Translator_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Translator/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Translator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Translator",
	HandlerType: (*TranslatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _Translator_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/translator.proto",
}

func init() { proto1.RegisterFile("proto/translator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0xce, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x03, 0x0b, 0x08,
	0xb1, 0x82, 0x29, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd,
	0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x22, 0xa5, 0x60, 0x2e,
	0x81, 0x10, 0xa8, 0xc6, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96,
	0x92, 0xd4, 0x8a, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x9b, 0x8b,
	0x23, 0x27, 0x31, 0x2f, 0xbd, 0x34, 0x31, 0x3d, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0xcf, 0x88,
	0x1f, 0x62, 0x82, 0x9e, 0x0f, 0x54, 0x38, 0x08, 0xae, 0x40, 0x49, 0x9d, 0x4b, 0x10, 0xc9, 0xd0,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x6c, 0xa6, 0x6a, 0x49, 0x71, 0x71, 0xc0, 0xb4, 0x0b, 0xb1,
	0x71, 0x31, 0xa5, 0xe6, 0x09, 0x30, 0x80, 0xe8, 0xb4, 0x22, 0x01, 0x46, 0xa3, 0x34, 0x2e, 0xae,
	0x10, 0xb8, 0x97, 0x84, 0x22, 0xb8, 0x38, 0xe1, 0x46, 0x0a, 0x89, 0x43, 0xad, 0x46, 0x77, 0xb9,
	0x94, 0x04, 0xa6, 0x04, 0xc4, 0x76, 0x25, 0x89, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x09, 0x29, 0xf1,
	0xea, 0x97, 0x19, 0xc2, 0x83, 0x2a, 0xd5, 0x8a, 0x51, 0x2b, 0x89, 0x0d, 0xac, 0xc5, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x66, 0x78, 0xa4, 0x86, 0x47, 0x01, 0x00, 0x00,
}