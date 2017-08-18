// Code generated by protoc-gen-go.
// source: response_contracts.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResponseContracts struct {
	// response code, 9bit
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code"`
	// response message
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	// response data
	Result []*Contract `protobuf:"bytes,3,rep,name=result" json:"result"`
}

func (m *ResponseContracts) Reset()                    { *m = ResponseContracts{} }
func (m *ResponseContracts) String() string            { return proto.CompactTextString(m) }
func (*ResponseContracts) ProtoMessage()               {}
func (*ResponseContracts) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ResponseContracts) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseContracts) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ResponseContracts) GetResult() []*Contract {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseContracts)(nil), "protos.ResponseContracts")
}

func init() { proto.RegisterFile("response_contracts.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x8d, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x52, 0x7c, 0x30, 0x09, 0x88, 0xb8, 0x52,
	0x32, 0x97, 0x60, 0x10, 0x54, 0x8f, 0x33, 0x4c, 0x8b, 0x90, 0x10, 0x17, 0x4b, 0x72, 0x7e, 0x4a,
	0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x98, 0x2d, 0x24, 0xc0, 0xc5, 0x9c, 0x5b, 0x9c,
	0x2e, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x0a, 0x69, 0x70, 0xb1, 0x15, 0xa5, 0x16,
	0x97, 0xe6, 0x94, 0x48, 0x30, 0x2b, 0x30, 0x6b, 0x70, 0x1b, 0x09, 0x40, 0x8c, 0x2c, 0xd6, 0x83,
	0x19, 0x14, 0x04, 0x95, 0x77, 0xd2, 0xe3, 0x12, 0x49, 0xce, 0xcf, 0xd5, 0x2b, 0xcd, 0xcb, 0xcc,
	0x49, 0x4d, 0x49, 0x4f, 0x2d, 0x82, 0x2a, 0x74, 0x12, 0x0b, 0x00, 0xd1, 0x18, 0xf6, 0x27, 0x41,
	0x1c, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x1c, 0x6e, 0xf9, 0xcf, 0x00, 0x00, 0x00,
}