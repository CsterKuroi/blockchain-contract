// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Response struct {
	// response code, 9bit
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code"`
	// response message
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	// response data
	Result string `protobuf:"bytes,3,opt,name=result" json:"result"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "protos.Response")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a,
	0x1e, 0x5c, 0x1c, 0x41, 0x50, 0x19, 0x21, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38, 0x5d, 0x82, 0x49,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x12, 0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29,
	0x91, 0x60, 0x06, 0x0b, 0x42, 0x79, 0x4e, 0xaa, 0x5c, 0x22, 0xc9, 0xf9, 0xb9, 0x7a, 0xa5, 0x79,
	0x99, 0x39, 0xa9, 0x29, 0xe9, 0xa9, 0x45, 0x10, 0x8b, 0x8a, 0x9d, 0x78, 0x03, 0x40, 0x34, 0xcc,
	0x92, 0x24, 0x88, 0xc5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x02, 0xf4, 0x31, 0x91,
	0x00, 0x00, 0x00,
}
