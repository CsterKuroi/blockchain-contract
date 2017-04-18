// Code generated by protoc-gen-go.
// source: contract_output.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	contract_output.proto
	contract.proto
	data.proto

It has these top-level messages:
	Asset
	ConditionDetails
	Condition
	ConditionsItem
	Fulfillment
	Metadata
	RelactionSignature
	Relaction
	Transaction
	ContractOutput
	ContractAttributes
	ContractSignature
	ContractAssert
	PlanTaskCondition
	Plan
	Task
	ContractComponents
	Contract
	ContractProto
	ContractData
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An Asset is a fungible unit to spend and lock with Transactions
type Asset struct {
	Data       *google_protobuf.Any `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Id         string               `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Divisible  bool                 `protobuf:"varint,3,opt,name=divisible" json:"divisible,omitempty"`
	Refillable bool                 `protobuf:"varint,4,opt,name=refillable" json:"refillable,omitempty"`
	Updatable  bool                 `protobuf:"varint,5,opt,name=updatable" json:"updatable,omitempty"`
}

func (m *Asset) Reset()                    { *m = Asset{} }
func (m *Asset) String() string            { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()               {}
func (*Asset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Asset) GetData() *google_protobuf.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Asset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Asset) GetDivisible() bool {
	if m != nil {
		return m.Divisible
	}
	return false
}

func (m *Asset) GetRefillable() bool {
	if m != nil {
		return m.Refillable
	}
	return false
}

func (m *Asset) GetUpdatable() bool {
	if m != nil {
		return m.Updatable
	}
	return false
}

type ConditionDetails struct {
	Bitmask   int32  `protobuf:"varint,1,opt,name=bitmask" json:"bitmask,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
	Type      string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	TypeId    int32  `protobuf:"varint,5,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
}

func (m *ConditionDetails) Reset()                    { *m = ConditionDetails{} }
func (m *ConditionDetails) String() string            { return proto.CompactTextString(m) }
func (*ConditionDetails) ProtoMessage()               {}
func (*ConditionDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConditionDetails) GetBitmask() int32 {
	if m != nil {
		return m.Bitmask
	}
	return 0
}

func (m *ConditionDetails) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *ConditionDetails) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *ConditionDetails) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ConditionDetails) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type Condition struct {
	Details *ConditionDetails `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	Uri     string            `protobuf:"bytes,2,opt,name=uri" json:"uri,omitempty"`
}

func (m *Condition) Reset()                    { *m = Condition{} }
func (m *Condition) String() string            { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()               {}
func (*Condition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Condition) GetDetails() *ConditionDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Condition) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type ConditionsItem struct {
	Amount      int64      `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	Cid         int32      `protobuf:"varint,2,opt,name=cid" json:"cid,omitempty"`
	Condition   *Condition `protobuf:"bytes,3,opt,name=condition" json:"condition,omitempty"`
	OwnersAfter []string   `protobuf:"bytes,4,rep,name=owners_after,json=ownersAfter" json:"owners_after,omitempty"`
	Isfreeze    bool       `protobuf:"varint,5,opt,name=isfreeze" json:"isfreeze,omitempty"`
}

func (m *ConditionsItem) Reset()                    { *m = ConditionsItem{} }
func (m *ConditionsItem) String() string            { return proto.CompactTextString(m) }
func (*ConditionsItem) ProtoMessage()               {}
func (*ConditionsItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ConditionsItem) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ConditionsItem) GetCid() int32 {
	if m != nil {
		return m.Cid
	}
	return 0
}

func (m *ConditionsItem) GetCondition() *Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *ConditionsItem) GetOwnersAfter() []string {
	if m != nil {
		return m.OwnersAfter
	}
	return nil
}

func (m *ConditionsItem) GetIsfreeze() bool {
	if m != nil {
		return m.Isfreeze
	}
	return false
}

type Fulfillment struct {
	Fid          int32                `protobuf:"varint,1,opt,name=fid" json:"fid,omitempty"`
	OwnersBefore []string             `protobuf:"bytes,2,rep,name=owners_before,json=ownersBefore" json:"owners_before,omitempty"`
	Fulfillment  string               `protobuf:"bytes,3,opt,name=fulfillment" json:"fulfillment,omitempty"`
	Input        *google_protobuf.Any `protobuf:"bytes,4,opt,name=input" json:"input,omitempty"`
}

func (m *Fulfillment) Reset()                    { *m = Fulfillment{} }
func (m *Fulfillment) String() string            { return proto.CompactTextString(m) }
func (*Fulfillment) ProtoMessage()               {}
func (*Fulfillment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Fulfillment) GetFid() int32 {
	if m != nil {
		return m.Fid
	}
	return 0
}

func (m *Fulfillment) GetOwnersBefore() []string {
	if m != nil {
		return m.OwnersBefore
	}
	return nil
}

func (m *Fulfillment) GetFulfillment() string {
	if m != nil {
		return m.Fulfillment
	}
	return ""
}

func (m *Fulfillment) GetInput() *google_protobuf.Any {
	if m != nil {
		return m.Input
	}
	return nil
}

type Metadata struct {
	Id   string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Data *google_protobuf.Any `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Metadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Metadata) GetData() *google_protobuf.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type RelactionSignature struct {
	ContractNodePubkey string `protobuf:"bytes,1,opt,name=contract_node_pubkey,json=contractNodePubkey" json:"contract_node_pubkey,omitempty"`
	Signature          string `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
}

func (m *RelactionSignature) Reset()                    { *m = RelactionSignature{} }
func (m *RelactionSignature) String() string            { return proto.CompactTextString(m) }
func (*RelactionSignature) ProtoMessage()               {}
func (*RelactionSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RelactionSignature) GetContractNodePubkey() string {
	if m != nil {
		return m.ContractNodePubkey
	}
	return ""
}

func (m *RelactionSignature) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

// 合约&交易关系信息
type Relaction struct {
	ContractId string                `protobuf:"bytes,1,opt,name=contract_id,json=contractId" json:"contract_id,omitempty"`
	TaskId     string                `protobuf:"bytes,2,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	Voters     []string              `protobuf:"bytes,3,rep,name=voters" json:"voters,omitempty"`
	Signatures []*RelactionSignature `protobuf:"bytes,4,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *Relaction) Reset()                    { *m = Relaction{} }
func (m *Relaction) String() string            { return proto.CompactTextString(m) }
func (*Relaction) ProtoMessage()               {}
func (*Relaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Relaction) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *Relaction) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *Relaction) GetVoters() []string {
	if m != nil {
		return m.Voters
	}
	return nil
}

func (m *Relaction) GetSignatures() []*RelactionSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type Transaction struct {
	Asset        *Asset            `protobuf:"bytes,1,opt,name=asset" json:"asset,omitempty"`
	Conditions   []*ConditionsItem `protobuf:"bytes,2,rep,name=conditions" json:"conditions,omitempty"`
	Fulfillments []*Fulfillment    `protobuf:"bytes,3,rep,name=fulfillments" json:"fulfillments,omitempty"`
	Metadata     *Metadata         `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
	Operation    string            `protobuf:"bytes,5,opt,name=operation" json:"operation,omitempty"`
	Timestamp    string            `protobuf:"bytes,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Relaction    *Relaction        `protobuf:"bytes,7,opt,name=relaction" json:"relaction,omitempty"`
	Contracts    *ContractProto    `protobuf:"bytes,8,opt,name=contracts" json:"contracts,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Transaction) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *Transaction) GetConditions() []*ConditionsItem {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *Transaction) GetFulfillments() []*Fulfillment {
	if m != nil {
		return m.Fulfillments
	}
	return nil
}

func (m *Transaction) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Transaction) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *Transaction) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Transaction) GetRelaction() *Relaction {
	if m != nil {
		return m.Relaction
	}
	return nil
}

func (m *Transaction) GetContracts() *ContractProto {
	if m != nil {
		return m.Contracts
	}
	return nil
}

type ContractOutput struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Transaction *Transaction `protobuf:"bytes,2,opt,name=transaction" json:"transaction,omitempty"`
	Version     int32        `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
}

func (m *ContractOutput) Reset()                    { *m = ContractOutput{} }
func (m *ContractOutput) String() string            { return proto.CompactTextString(m) }
func (*ContractOutput) ProtoMessage()               {}
func (*ContractOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ContractOutput) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContractOutput) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *ContractOutput) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*Asset)(nil), "protos.Asset")
	proto.RegisterType((*ConditionDetails)(nil), "protos.ConditionDetails")
	proto.RegisterType((*Condition)(nil), "protos.Condition")
	proto.RegisterType((*ConditionsItem)(nil), "protos.ConditionsItem")
	proto.RegisterType((*Fulfillment)(nil), "protos.Fulfillment")
	proto.RegisterType((*Metadata)(nil), "protos.Metadata")
	proto.RegisterType((*RelactionSignature)(nil), "protos.RelactionSignature")
	proto.RegisterType((*Relaction)(nil), "protos.Relaction")
	proto.RegisterType((*Transaction)(nil), "protos.Transaction")
	proto.RegisterType((*ContractOutput)(nil), "protos.ContractOutput")
}

func init() { proto.RegisterFile("contract_output.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0x93, 0x38, 0x89, 0x8f, 0xdb, 0xaa, 0x77, 0x6e, 0xdb, 0xeb, 0x1b, 0xf1, 0x13, 0xdc,
	0x4d, 0x84, 0x50, 0x8a, 0x52, 0x01, 0x12, 0xbb, 0x42, 0x85, 0x14, 0x21, 0xa0, 0x0c, 0xec, 0xa3,
	0x49, 0x3c, 0xa9, 0x46, 0x75, 0x3c, 0xc6, 0x33, 0x2e, 0x0a, 0x8f, 0xc1, 0x02, 0x89, 0x05, 0x2f,
	0xc0, 0xeb, 0xf1, 0x02, 0x68, 0x7e, 0xe3, 0x36, 0x42, 0xac, 0x32, 0xe7, 0x6f, 0xce, 0xe7, 0xf3,
	0x7d, 0x67, 0x02, 0x87, 0x0b, 0x5e, 0xc8, 0x8a, 0x2c, 0xe4, 0x8c, 0xd7, 0xb2, 0xac, 0xe5, 0xb8,
	0xac, 0xb8, 0xe4, 0xa8, 0xab, 0x7f, 0xc4, 0xe0, 0xff, 0x4b, 0xce, 0x2f, 0x73, 0x7a, 0xa2, 0xcd,
	0x79, 0xbd, 0x3c, 0x21, 0xc5, 0xda, 0xa4, 0x0c, 0xf6, 0x5c, 0xa5, 0xb1, 0xd3, 0x1f, 0x01, 0x84,
	0x67, 0x42, 0x50, 0x89, 0x46, 0xd0, 0xc9, 0x88, 0x24, 0x49, 0x30, 0x0c, 0x46, 0xf1, 0xe4, 0x60,
	0x6c, 0xee, 0x18, 0xbb, 0x3b, 0xc6, 0x67, 0xc5, 0x1a, 0xeb, 0x0c, 0xb4, 0x07, 0x2d, 0x96, 0x25,
	0xad, 0x61, 0x30, 0x8a, 0x70, 0x8b, 0x65, 0xe8, 0x0e, 0x44, 0x19, 0xbb, 0x66, 0x82, 0xcd, 0x73,
	0x9a, 0xb4, 0x87, 0xc1, 0xa8, 0x8f, 0x37, 0x0e, 0x74, 0x0f, 0xa0, 0xa2, 0x4b, 0x96, 0xe7, 0x44,
	0x85, 0x3b, 0x3a, 0xdc, 0xf0, 0xa8, 0xea, 0xba, 0x54, 0xf7, 0xaa, 0x70, 0x68, 0xaa, 0xbd, 0x23,
	0xfd, 0x16, 0xc0, 0xfe, 0x4b, 0x5e, 0x64, 0x4c, 0x32, 0x5e, 0x9c, 0x53, 0x49, 0x58, 0x2e, 0x50,
	0x02, 0xbd, 0x39, 0x93, 0x2b, 0x22, 0xae, 0x34, 0xda, 0x10, 0x3b, 0x13, 0xdd, 0x05, 0x28, 0xeb,
	0x79, 0xce, 0x16, 0xb3, 0x2b, 0xba, 0xb6, 0x10, 0x23, 0xe3, 0x79, 0x4d, 0xd7, 0xaa, 0x97, 0x60,
	0x97, 0x05, 0x91, 0x75, 0x65, 0x90, 0x46, 0x78, 0xe3, 0x40, 0x08, 0x3a, 0x72, 0x5d, 0x1a, 0x8c,
	0x11, 0xd6, 0x67, 0xf4, 0x1f, 0xf4, 0xd4, 0xef, 0x8c, 0x65, 0x1a, 0x5b, 0x88, 0xbb, 0xca, 0x9c,
	0x66, 0xe9, 0x7b, 0x88, 0x3c, 0x2e, 0x34, 0x81, 0x5e, 0x66, 0xb0, 0xd9, 0xf1, 0x25, 0x66, 0x6e,
	0x62, 0x7c, 0x1b, 0x3b, 0x76, 0x89, 0x68, 0x1f, 0xda, 0x75, 0xc5, 0x2c, 0x46, 0x75, 0x4c, 0x7f,
	0x06, 0xb0, 0xe7, 0xf3, 0xc5, 0x54, 0xd2, 0x15, 0x3a, 0x82, 0x2e, 0x59, 0xf1, 0xba, 0x90, 0xfa,
	0xde, 0x36, 0xb6, 0x96, 0x2a, 0x5e, 0x58, 0x0e, 0x42, 0xac, 0x8e, 0xe8, 0x04, 0xa2, 0x85, 0xab,
	0xd5, 0x9f, 0x16, 0x4f, 0xfe, 0xd9, 0x02, 0x81, 0x37, 0x39, 0xe8, 0x01, 0xec, 0xf0, 0xcf, 0x05,
	0xad, 0xc4, 0x8c, 0x2c, 0x25, 0xad, 0x92, 0xce, 0xb0, 0x3d, 0x8a, 0x70, 0x6c, 0x7c, 0x67, 0xca,
	0x85, 0x06, 0xd0, 0x67, 0x62, 0x59, 0x51, 0xfa, 0xc5, 0x31, 0xe3, 0xed, 0xf4, 0x6b, 0x00, 0xf1,
	0xab, 0x3a, 0x57, 0x34, 0xae, 0xa8, 0x41, 0xb4, 0x64, 0x99, 0xe5, 0x43, 0x1d, 0xd1, 0x31, 0xec,
	0xda, 0x06, 0x73, 0xba, 0xe4, 0x15, 0x4d, 0x5a, 0xba, 0x83, 0xed, 0xfa, 0x42, 0xfb, 0xd0, 0x10,
	0xe2, 0xe5, 0xe6, 0x16, 0xcb, 0x49, 0xd3, 0x85, 0x1e, 0x42, 0xc8, 0x8a, 0xb2, 0x96, 0x9a, 0x96,
	0x3f, 0x09, 0xd3, 0xa4, 0xa4, 0xe7, 0xd0, 0x7f, 0x43, 0x25, 0x69, 0xa8, 0x34, 0xf0, 0x2a, 0x75,
	0xfa, 0x6e, 0xfd, 0x4d, 0xdf, 0x69, 0x06, 0x08, 0xd3, 0x9c, 0x2c, 0xd4, 0x98, 0x3e, 0x78, 0x75,
	0x3c, 0x86, 0x03, 0xbf, 0x75, 0x05, 0xcf, 0xe8, 0xac, 0xac, 0xe7, 0x4a, 0x64, 0xa6, 0x03, 0x72,
	0xb1, 0xb7, 0x3c, 0xa3, 0x17, 0x3a, 0x72, 0x53, 0x6d, 0xad, 0x5b, 0x6a, 0x4b, 0xbf, 0x07, 0x10,
	0xf9, 0x36, 0xe8, 0x3e, 0xc4, 0xfe, 0x76, 0x0f, 0x1b, 0x9c, 0x6b, 0x9a, 0x69, 0x21, 0x12, 0x71,
	0x35, 0xf3, 0x9b, 0xd7, 0x55, 0xe6, 0x34, 0x53, 0x12, 0xb9, 0xe6, 0x92, 0x56, 0x22, 0x69, 0xeb,
	0xf9, 0x5a, 0x0b, 0x3d, 0x07, 0xf0, 0xcd, 0x84, 0x66, 0x37, 0x9e, 0x0c, 0x9c, 0x22, 0xb6, 0xbf,
	0x0f, 0x37, 0xb2, 0xd3, 0x5f, 0x2d, 0x88, 0x3f, 0x56, 0xa4, 0x10, 0x16, 0xdd, 0x31, 0x84, 0x44,
	0x3d, 0x12, 0x56, 0xdd, 0xbb, 0xee, 0x1a, 0xfd, 0x72, 0x60, 0x13, 0x43, 0x4f, 0x01, 0xbc, 0xba,
	0x84, 0x26, 0x3b, 0x9e, 0x1c, 0x6d, 0x49, 0x50, 0xeb, 0x1a, 0x37, 0x32, 0xd1, 0x33, 0xd8, 0x69,
	0xf0, 0x6d, 0x3e, 0x23, 0x9e, 0xfc, 0xeb, 0x2a, 0x1b, 0x22, 0xc3, 0x37, 0x12, 0xd1, 0x23, 0xe8,
	0xaf, 0x2c, 0xdb, 0x56, 0x1c, 0xfb, 0xae, 0xc8, 0xa9, 0x00, 0xfb, 0x0c, 0xc5, 0x06, 0x2f, 0x69,
	0x45, 0xf4, 0x82, 0x84, 0x86, 0x0d, 0xef, 0x50, 0x51, 0xc9, 0x56, 0x54, 0x48, 0xb2, 0x2a, 0x93,
	0xae, 0x89, 0x7a, 0x87, 0x5a, 0xae, 0xca, 0x4d, 0x2c, 0xe9, 0xdd, 0x5c, 0x2e, 0x3f, 0x4a, 0xbc,
	0xc9, 0x41, 0xa7, 0x7a, 0x1b, 0x35, 0x77, 0x22, 0xe9, 0xeb, 0x82, 0xc3, 0xc6, 0x28, 0x74, 0xe0,
	0x42, 0x99, 0x78, 0x93, 0x97, 0x7e, 0xd2, 0xeb, 0xaf, 0x8d, 0x77, 0xfa, 0x59, 0xdf, 0xd2, 0xf0,
	0x13, 0x88, 0xe5, 0x86, 0x16, 0x2b, 0x65, 0x3f, 0xa9, 0x06, 0x63, 0xb8, 0x99, 0xa7, 0xde, 0xcb,
	0x6b, 0x5a, 0x09, 0xf7, 0x32, 0x84, 0xd8, 0x99, 0x73, 0xf3, 0x8f, 0x71, 0xfa, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x9c, 0x27, 0x2a, 0xa8, 0x51, 0x06, 0x00, 0x00,
}
