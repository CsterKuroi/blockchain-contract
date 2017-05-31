// Code generated by protoc-gen-go.
// source: contract.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContractSignature struct {
	OwnerPubkey   string `protobuf:"bytes,1,opt,name=OwnerPubkey" json:"OwnerPubkey"`
	Signature     string `protobuf:"bytes,2,opt,name=Signature" json:"Signature"`
	SignTimestamp string `protobuf:"bytes,3,opt,name=SignTimestamp" json:"SignTimestamp"`
}

func (m *ContractSignature) Reset()                    { *m = ContractSignature{} }
func (m *ContractSignature) String() string            { return proto.CompactTextString(m) }
func (*ContractSignature) ProtoMessage()               {}
func (*ContractSignature) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ContractSignature) GetOwnerPubkey() string {
	if m != nil {
		return m.OwnerPubkey
	}
	return ""
}

func (m *ContractSignature) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *ContractSignature) GetSignTimestamp() string {
	if m != nil {
		return m.SignTimestamp
	}
	return ""
}

type ContractAsset struct {
	AssetId     string  `protobuf:"bytes,1,opt,name=AssetId" json:"AssetId"`
	Name        string  `protobuf:"bytes,2,opt,name=Name" json:"Name"`
	Caption     string  `protobuf:"bytes,3,opt,name=Caption" json:"Caption"`
	Description string  `protobuf:"bytes,4,opt,name=Description" json:"Description"`
	Unit        string  `protobuf:"bytes,5,opt,name=Unit" json:"Unit"`
	Amount      float32 `protobuf:"fixed32,6,opt,name=Amount" json:"Amount"`
	MetaData    []byte  `protobuf:"bytes,7,opt,name=MetaData,proto3" json:"MetaData"`
}

func (m *ContractAsset) Reset()                    { *m = ContractAsset{} }
func (m *ContractAsset) String() string            { return proto.CompactTextString(m) }
func (*ContractAsset) ProtoMessage()               {}
func (*ContractAsset) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ContractAsset) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *ContractAsset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContractAsset) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ContractAsset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContractAsset) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *ContractAsset) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ContractAsset) GetMetaData() []byte {
	if m != nil {
		return m.MetaData
	}
	return nil
}

type ExpressionResult struct {
	Messsage string `protobuf:"bytes,1,opt,name=Messsage" json:"Messsage"`
	Code     string `protobuf:"bytes,2,opt,name=Code" json:"Code"`
}

func (m *ExpressionResult) Reset()                    { *m = ExpressionResult{} }
func (m *ExpressionResult) String() string            { return proto.CompactTextString(m) }
func (*ExpressionResult) ProtoMessage()               {}
func (*ExpressionResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ExpressionResult) GetMesssage() string {
	if m != nil {
		return m.Messsage
	}
	return ""
}

func (m *ExpressionResult) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ComponentsExpression struct {
	Cname            string            `protobuf:"bytes,1,opt,name=Cname" json:"Cname"`
	Ctype            string            `protobuf:"bytes,2,opt,name=Ctype" json:"Ctype"`
	Caption          string            `protobuf:"bytes,3,opt,name=Caption" json:"Caption"`
	Description      string            `protobuf:"bytes,4,opt,name=Description" json:"Description"`
	ExpressionStr    string            `protobuf:"bytes,5,opt,name=ExpressionStr" json:"ExpressionStr"`
	ExpressionResult *ExpressionResult `protobuf:"bytes,6,opt,name=ExpressionResult" json:"ExpressionResult"`
	LogicValue       string            `protobuf:"bytes,7,opt,name=LogicValue" json:"LogicValue"`
}

func (m *ComponentsExpression) Reset()                    { *m = ComponentsExpression{} }
func (m *ComponentsExpression) String() string            { return proto.CompactTextString(m) }
func (*ComponentsExpression) ProtoMessage()               {}
func (*ComponentsExpression) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ComponentsExpression) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ComponentsExpression) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *ComponentsExpression) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ComponentsExpression) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ComponentsExpression) GetExpressionStr() string {
	if m != nil {
		return m.ExpressionStr
	}
	return ""
}

func (m *ComponentsExpression) GetExpressionResult() *ExpressionResult {
	if m != nil {
		return m.ExpressionResult
	}
	return nil
}

func (m *ComponentsExpression) GetLogicValue() string {
	if m != nil {
		return m.LogicValue
	}
	return ""
}

type ComponentData struct {
	Cname        string `protobuf:"bytes,1,opt,name=Cname" json:"Cname"`
	Ctype        string `protobuf:"bytes,2,opt,name=Ctype" json:"Ctype"`
	Caption      string `protobuf:"bytes,3,opt,name=Caption" json:"Caption"`
	Description  string `protobuf:"bytes,4,opt,name=Description" json:"Description"`
	ModifyDate   string `protobuf:"bytes,5,opt,name=ModifyDate" json:"ModifyDate"`
	HardConvType string `protobuf:"bytes,6,opt,name=HardConvType" json:"HardConvType"`
	//    map<string,?> Category = 7; // map[string]interface{} Category
	Parent    *ComponentData `protobuf:"bytes,8,opt,name=Parent" json:"Parent"`
	Mandatory bool           `protobuf:"varint,9,opt,name=Mandatory" json:"Mandatory"`
	//    google.protobuf.Any DefaultValue = 10; // interface{} DefaultValue
	Unit string `protobuf:"bytes,11,opt,name=Unit" json:"Unit"`
	//    google.protobuf.Any Value = 12; // interface{} Value
	Options map[string]int32 `protobuf:"bytes,13,rep,name=Options" json:"Options" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	//    repeated google.protobuf.Any DataRange = 14; // []interface{} DataRange
	Format string `protobuf:"bytes,15,opt,name=Format" json:"Format"`
}

func (m *ComponentData) Reset()                    { *m = ComponentData{} }
func (m *ComponentData) String() string            { return proto.CompactTextString(m) }
func (*ComponentData) ProtoMessage()               {}
func (*ComponentData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ComponentData) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ComponentData) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *ComponentData) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ComponentData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ComponentData) GetModifyDate() string {
	if m != nil {
		return m.ModifyDate
	}
	return ""
}

func (m *ComponentData) GetHardConvType() string {
	if m != nil {
		return m.HardConvType
	}
	return ""
}

func (m *ComponentData) GetParent() *ComponentData {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *ComponentData) GetMandatory() bool {
	if m != nil {
		return m.Mandatory
	}
	return false
}

func (m *ComponentData) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *ComponentData) GetOptions() map[string]int32 {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *ComponentData) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

type SelectBranchExpression struct {
	BranchExpressionStr   string `protobuf:"bytes,1,opt,name=BranchExpressionStr" json:"BranchExpressionStr"`
	BranchExpressionValue string `protobuf:"bytes,2,opt,name=BranchExpressionValue" json:"BranchExpressionValue"`
}

func (m *SelectBranchExpression) Reset()                    { *m = SelectBranchExpression{} }
func (m *SelectBranchExpression) String() string            { return proto.CompactTextString(m) }
func (*SelectBranchExpression) ProtoMessage()               {}
func (*SelectBranchExpression) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SelectBranchExpression) GetBranchExpressionStr() string {
	if m != nil {
		return m.BranchExpressionStr
	}
	return ""
}

func (m *SelectBranchExpression) GetBranchExpressionValue() string {
	if m != nil {
		return m.BranchExpressionValue
	}
	return ""
}

type ContractComponent struct {
	Cname                         string                  `protobuf:"bytes,1,opt,name=Cname" json:"Cname"`
	Ctype                         string                  `protobuf:"bytes,2,opt,name=Ctype" json:"Ctype"`
	Caption                       string                  `protobuf:"bytes,3,opt,name=Caption" json:"Caption"`
	Description                   string                  `protobuf:"bytes,4,opt,name=Description" json:"Description"`
	State                         string                  `protobuf:"bytes,5,opt,name=State" json:"State"`
	PreCondition                  []*ComponentsExpression `protobuf:"bytes,6,rep,name=PreCondition" json:"PreCondition"`
	CompleteCondition             []*ComponentsExpression `protobuf:"bytes,7,rep,name=CompleteCondition" json:"CompleteCondition"`
	DiscardCondition              []*ComponentsExpression `protobuf:"bytes,8,rep,name=DiscardCondition" json:"DiscardCondition"`
	NextTasks                     []string                `protobuf:"bytes,9,rep,name=NextTasks" json:"NextTasks"`
	DataList                      []*ComponentData        `protobuf:"bytes,10,rep,name=DataList" json:"DataList"`
	DataValueSetterExpressionList []*ComponentsExpression `protobuf:"bytes,11,rep,name=DataValueSetterExpressionList" json:"DataValueSetterExpressionList"`
	CandidateList                 *ContractComponent      `protobuf:"bytes,12,opt,name=CandidateList" json:"CandidateList"`
	DecisionResult                *ContractComponent      `protobuf:"bytes,13,opt,name=DecisionResult" json:"DecisionResult"`
	TaskList                      []string                `protobuf:"bytes,14,rep,name=TaskList" json:"TaskList"`
	SupportArguments              []string                `protobuf:"bytes,15,rep,name=SupportArguments" json:"SupportArguments"`
	AgainstArguments              []string                `protobuf:"bytes,16,rep,name=AgainstArguments" json:"AgainstArguments"`
	Support                       int32                   `protobuf:"varint,17,opt,name=Support" json:"Support"`
	Text                          []string                `protobuf:"bytes,18,rep,name=Text" json:"Text"`
	// add date: 2017-05-11 任务执行索引次数 int
	TaskExecuteIdx int32  `protobuf:"varint,19,opt,name=TaskExecuteIdx" json:"TaskExecuteIdx"`
	TaskId         string `protobuf:"bytes,20,opt,name=TaskId" json:"TaskId"`
	// 2017-05-27 17:10:00 add
	SelectBranches []*SelectBranchExpression `protobuf:"bytes,21,rep,name=SelectBranches" json:"SelectBranches"`
}

func (m *ContractComponent) Reset()                    { *m = ContractComponent{} }
func (m *ContractComponent) String() string            { return proto.CompactTextString(m) }
func (*ContractComponent) ProtoMessage()               {}
func (*ContractComponent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ContractComponent) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ContractComponent) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *ContractComponent) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ContractComponent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContractComponent) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ContractComponent) GetPreCondition() []*ComponentsExpression {
	if m != nil {
		return m.PreCondition
	}
	return nil
}

func (m *ContractComponent) GetCompleteCondition() []*ComponentsExpression {
	if m != nil {
		return m.CompleteCondition
	}
	return nil
}

func (m *ContractComponent) GetDiscardCondition() []*ComponentsExpression {
	if m != nil {
		return m.DiscardCondition
	}
	return nil
}

func (m *ContractComponent) GetNextTasks() []string {
	if m != nil {
		return m.NextTasks
	}
	return nil
}

func (m *ContractComponent) GetDataList() []*ComponentData {
	if m != nil {
		return m.DataList
	}
	return nil
}

func (m *ContractComponent) GetDataValueSetterExpressionList() []*ComponentsExpression {
	if m != nil {
		return m.DataValueSetterExpressionList
	}
	return nil
}

func (m *ContractComponent) GetCandidateList() *ContractComponent {
	if m != nil {
		return m.CandidateList
	}
	return nil
}

func (m *ContractComponent) GetDecisionResult() *ContractComponent {
	if m != nil {
		return m.DecisionResult
	}
	return nil
}

func (m *ContractComponent) GetTaskList() []string {
	if m != nil {
		return m.TaskList
	}
	return nil
}

func (m *ContractComponent) GetSupportArguments() []string {
	if m != nil {
		return m.SupportArguments
	}
	return nil
}

func (m *ContractComponent) GetAgainstArguments() []string {
	if m != nil {
		return m.AgainstArguments
	}
	return nil
}

func (m *ContractComponent) GetSupport() int32 {
	if m != nil {
		return m.Support
	}
	return 0
}

func (m *ContractComponent) GetText() []string {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *ContractComponent) GetTaskExecuteIdx() int32 {
	if m != nil {
		return m.TaskExecuteIdx
	}
	return 0
}

func (m *ContractComponent) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ContractComponent) GetSelectBranches() []*SelectBranchExpression {
	if m != nil {
		return m.SelectBranches
	}
	return nil
}

type ContractBody struct {
	ContractId         string               `protobuf:"bytes,1,opt,name=ContractId" json:"ContractId"`
	Cname              string               `protobuf:"bytes,2,opt,name=Cname" json:"Cname"`
	Ctype              string               `protobuf:"bytes,3,opt,name=Ctype" json:"Ctype"`
	Caption            string               `protobuf:"bytes,4,opt,name=Caption" json:"Caption"`
	Description        string               `protobuf:"bytes,5,opt,name=Description" json:"Description"`
	ContractState      string               `protobuf:"bytes,6,opt,name=ContractState" json:"ContractState"`
	Creator            string               `protobuf:"bytes,7,opt,name=Creator" json:"Creator"`
	CreateTime         string               `protobuf:"bytes,8,opt,name=CreateTime" json:"CreateTime"`
	StartTime          string               `protobuf:"bytes,9,opt,name=StartTime" json:"StartTime"`
	EndTime            string               `protobuf:"bytes,10,opt,name=EndTime" json:"EndTime"`
	ContractOwners     []string             `protobuf:"bytes,11,rep,name=ContractOwners" json:"ContractOwners"`
	ContractAssets     []*ContractAsset     `protobuf:"bytes,12,rep,name=ContractAssets" json:"ContractAssets"`
	ContractSignatures []*ContractSignature `protobuf:"bytes,13,rep,name=ContractSignatures" json:"ContractSignatures"`
	ContractComponents []*ContractComponent `protobuf:"bytes,14,rep,name=ContractComponents" json:"ContractComponents"`
	// add date: 2017-05-11 map[string]interface{} 合约属性MetaData
	MetaAttribute []byte `protobuf:"bytes,15,opt,name=MetaAttribute,proto3" json:"MetaAttribute"`
}

func (m *ContractBody) Reset()                    { *m = ContractBody{} }
func (m *ContractBody) String() string            { return proto.CompactTextString(m) }
func (*ContractBody) ProtoMessage()               {}
func (*ContractBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ContractBody) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *ContractBody) GetCname() string {
	if m != nil {
		return m.Cname
	}
	return ""
}

func (m *ContractBody) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *ContractBody) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *ContractBody) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContractBody) GetContractState() string {
	if m != nil {
		return m.ContractState
	}
	return ""
}

func (m *ContractBody) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ContractBody) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *ContractBody) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *ContractBody) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *ContractBody) GetContractOwners() []string {
	if m != nil {
		return m.ContractOwners
	}
	return nil
}

func (m *ContractBody) GetContractAssets() []*ContractAsset {
	if m != nil {
		return m.ContractAssets
	}
	return nil
}

func (m *ContractBody) GetContractSignatures() []*ContractSignature {
	if m != nil {
		return m.ContractSignatures
	}
	return nil
}

func (m *ContractBody) GetContractComponents() []*ContractComponent {
	if m != nil {
		return m.ContractComponents
	}
	return nil
}

func (m *ContractBody) GetMetaAttribute() []byte {
	if m != nil {
		return m.MetaAttribute
	}
	return nil
}

type ContractHead struct {
	MainPubkey string `protobuf:"bytes,1,opt,name=MainPubkey" json:"MainPubkey"`
	Version    int32  `protobuf:"varint,2,opt,name=Version" json:"Version"`
	// 指派处理时间 add 2017-05-27 17:10:0
	AssignTime string `protobuf:"bytes,3,opt,name=AssignTime" json:"AssignTime"`
	// add date: 2017-05-11 合约执行时间戳
	// 操作时间,记录状态改变时间, Timestamp修改而来 2017-05-27 17:10:0
	OperateTime string `protobuf:"bytes,4,opt,name=OperateTime" json:"OperateTime"`
}

func (m *ContractHead) Reset()                    { *m = ContractHead{} }
func (m *ContractHead) String() string            { return proto.CompactTextString(m) }
func (*ContractHead) ProtoMessage()               {}
func (*ContractHead) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ContractHead) GetMainPubkey() string {
	if m != nil {
		return m.MainPubkey
	}
	return ""
}

func (m *ContractHead) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ContractHead) GetAssignTime() string {
	if m != nil {
		return m.AssignTime
	}
	return ""
}

func (m *ContractHead) GetOperateTime() string {
	if m != nil {
		return m.OperateTime
	}
	return ""
}

type Contract struct {
	Id           string        `protobuf:"bytes,1,opt,name=id" json:"id"`
	ContractHead *ContractHead `protobuf:"bytes,2,opt,name=ContractHead" json:"ContractHead"`
	ContractBody *ContractBody `protobuf:"bytes,3,opt,name=ContractBody" json:"ContractBody"`
}

func (m *Contract) Reset()                    { *m = Contract{} }
func (m *Contract) String() string            { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()               {}
func (*Contract) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *Contract) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Contract) GetContractHead() *ContractHead {
	if m != nil {
		return m.ContractHead
	}
	return nil
}

func (m *Contract) GetContractBody() *ContractBody {
	if m != nil {
		return m.ContractBody
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractSignature)(nil), "protos.ContractSignature")
	proto.RegisterType((*ContractAsset)(nil), "protos.ContractAsset")
	proto.RegisterType((*ExpressionResult)(nil), "protos.ExpressionResult")
	proto.RegisterType((*ComponentsExpression)(nil), "protos.ComponentsExpression")
	proto.RegisterType((*ComponentData)(nil), "protos.ComponentData")
	proto.RegisterType((*SelectBranchExpression)(nil), "protos.SelectBranchExpression")
	proto.RegisterType((*ContractComponent)(nil), "protos.ContractComponent")
	proto.RegisterType((*ContractBody)(nil), "protos.ContractBody")
	proto.RegisterType((*ContractHead)(nil), "protos.ContractHead")
	proto.RegisterType((*Contract)(nil), "protos.Contract")
}

func init() { proto.RegisterFile("contract.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1091 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0x86, 0x24, 0x5b, 0x96, 0x46, 0x3f, 0x71, 0x18, 0x25, 0x60, 0x83, 0xd4, 0x10, 0x16, 0x69,
	0x61, 0x14, 0xa8, 0xd1, 0xba, 0x3d, 0x18, 0x41, 0x8b, 0x56, 0x96, 0x1c, 0xd8, 0x45, 0x9c, 0x18,
	0x2b, 0xd7, 0x77, 0x5a, 0xcb, 0xaa, 0x8b, 0x58, 0x5c, 0x81, 0xe4, 0xa6, 0xd6, 0xad, 0xe7, 0x1c,
	0xfa, 0x28, 0x7d, 0x87, 0xde, 0x7b, 0xe9, 0x1b, 0x15, 0x33, 0xcb, 0xfd, 0x95, 0xe2, 0x1a, 0x28,
	0x90, 0x93, 0x76, 0xbe, 0x99, 0xf9, 0x86, 0x1c, 0xce, 0x90, 0x23, 0xe8, 0xcf, 0x22, 0x65, 0xb5,
	0x98, 0xd9, 0x83, 0xa5, 0x8e, 0x6c, 0xc4, 0x9a, 0xf4, 0x63, 0xbc, 0x15, 0x3c, 0x1c, 0x3b, 0xcd,
	0x34, 0x9c, 0x2b, 0x61, 0x63, 0x2d, 0xd9, 0x10, 0x3a, 0x6f, 0x7e, 0x53, 0x52, 0x5f, 0xc4, 0xd7,
	0x6f, 0xe5, 0x8a, 0xd7, 0x86, 0xb5, 0xfd, 0xb6, 0x5f, 0x84, 0xd8, 0x33, 0x68, 0x67, 0xe6, 0xbc,
	0x4e, 0xfa, 0x1c, 0x60, 0xcf, 0xa1, 0x87, 0xc2, 0x65, 0xb8, 0x90, 0xc6, 0x8a, 0xc5, 0x92, 0x37,
	0xc8, 0xa2, 0x0c, 0x7a, 0x7f, 0xd5, 0xa0, 0x97, 0xc6, 0x1e, 0x19, 0x23, 0x2d, 0xe3, 0xb0, 0x43,
	0x1f, 0x67, 0x81, 0x8b, 0x99, 0x8a, 0x8c, 0xc1, 0xd6, 0x6b, 0xb1, 0x48, 0x43, 0xd1, 0x37, 0x5a,
	0x8f, 0xc5, 0xd2, 0x86, 0x91, 0x72, 0xfc, 0xa9, 0x88, 0xeb, 0x9f, 0x48, 0x33, 0xd3, 0x61, 0xa2,
	0xdd, 0x4a, 0xd6, 0x5f, 0x80, 0x90, 0xef, 0x67, 0x15, 0x5a, 0xbe, 0x9d, 0xf0, 0xe1, 0x37, 0x7b,
	0x02, 0xcd, 0xd1, 0x22, 0x8a, 0x95, 0xe5, 0xcd, 0x61, 0x6d, 0xbf, 0xee, 0x3b, 0x89, 0x3d, 0x85,
	0xd6, 0xb9, 0xb4, 0x62, 0x22, 0xac, 0xe0, 0x3b, 0xc3, 0xda, 0x7e, 0xd7, 0xcf, 0x64, 0xef, 0x18,
	0x76, 0x4f, 0x6e, 0x97, 0x5a, 0x1a, 0x13, 0x46, 0xca, 0x97, 0x26, 0xbe, 0x71, 0xf6, 0xc6, 0x18,
	0x31, 0x97, 0x6e, 0x1b, 0x99, 0x8c, 0x71, 0xc7, 0x51, 0x90, 0xed, 0x03, 0xbf, 0xbd, 0xf7, 0x75,
	0x18, 0x8c, 0xa3, 0xc5, 0x32, 0x52, 0x52, 0x59, 0x93, 0xd3, 0xb1, 0x01, 0x6c, 0x8f, 0x15, 0xee,
	0x3a, 0x61, 0x49, 0x04, 0x42, 0xed, 0x6a, 0x99, 0x72, 0x24, 0xc2, 0xff, 0x4a, 0xc6, 0x73, 0xe8,
	0xe5, 0x51, 0xa7, 0x56, 0xbb, 0xac, 0x94, 0x41, 0x36, 0x59, 0xdf, 0x2a, 0x25, 0xaa, 0x73, 0xc8,
	0x93, 0x9a, 0x32, 0x07, 0x55, 0xbd, 0xbf, 0x9e, 0x9c, 0x3d, 0x80, 0x57, 0xd1, 0x3c, 0x9c, 0x5d,
	0x89, 0x9b, 0x58, 0x52, 0x3a, 0xdb, 0x7e, 0x01, 0xf1, 0xfe, 0x6c, 0x60, 0x51, 0xb8, 0x64, 0x60,
	0x8a, 0x3f, 0x5a, 0x16, 0xf6, 0x00, 0xce, 0xa3, 0x20, 0xfc, 0x65, 0x35, 0x11, 0x56, 0xba, 0x14,
	0x14, 0x10, 0xe6, 0x41, 0xf7, 0x54, 0xe8, 0x60, 0x1c, 0xa9, 0x77, 0x97, 0x18, 0xb8, 0x49, 0x16,
	0x25, 0x8c, 0x7d, 0x09, 0xcd, 0x0b, 0xa1, 0xa5, 0xb2, 0xbc, 0x45, 0x99, 0x79, 0x9c, 0x66, 0xa6,
	0xb4, 0x25, 0xdf, 0x19, 0x61, 0x17, 0x9d, 0x0b, 0x15, 0x08, 0x1b, 0xe9, 0x15, 0x6f, 0x0f, 0x6b,
	0xfb, 0x2d, 0x3f, 0x07, 0xb2, 0x1a, 0xed, 0x14, 0x6a, 0xf4, 0x3b, 0xd8, 0x79, 0x43, 0xcb, 0x35,
	0xbc, 0x37, 0x6c, 0xec, 0x77, 0x0e, 0xbd, 0x8d, 0x11, 0x0e, 0x9c, 0xd1, 0x89, 0xb2, 0x7a, 0xe5,
	0xa7, 0x2e, 0x58, 0xe1, 0x2f, 0x23, 0xbd, 0x10, 0x96, 0x3f, 0x20, 0x4e, 0x27, 0x3d, 0x7d, 0x01,
	0xdd, 0xa2, 0x03, 0xdb, 0x85, 0x46, 0xde, 0xf7, 0xf8, 0x89, 0xe9, 0x7e, 0x47, 0x27, 0x86, 0xe9,
	0xde, 0xf6, 0x13, 0xe1, 0x45, 0xfd, 0xa8, 0xe6, 0xfd, 0x5e, 0x83, 0x27, 0x53, 0x79, 0x23, 0x67,
	0xf6, 0x58, 0x0b, 0x35, 0xfb, 0xb5, 0x50, 0xbf, 0x5f, 0xc1, 0xa3, 0x2a, 0x86, 0xd5, 0x95, 0xd0,
	0x6e, 0x52, 0xb1, 0x6f, 0xe1, 0x71, 0x15, 0xbe, 0xca, 0xc2, 0xb6, 0xfd, 0xcd, 0x4a, 0xef, 0xef,
	0x9d, 0xfc, 0x12, 0xcb, 0xd2, 0xf0, 0xd1, 0xea, 0x66, 0x00, 0xdb, 0x53, 0x9b, 0x97, 0x4c, 0x22,
	0xb0, 0x1f, 0xa1, 0x7b, 0xa1, 0xe5, 0x38, 0x52, 0x41, 0x48, 0x8e, 0x4d, 0x3a, 0xad, 0x67, 0x6b,
	0xa7, 0x55, 0xe8, 0x77, 0xbf, 0xe4, 0xc1, 0x7e, 0xc2, 0x4d, 0x2d, 0x96, 0x37, 0xd2, 0x16, 0x68,
	0x76, 0xee, 0x41, 0xb3, 0xee, 0xc6, 0x4e, 0x61, 0x77, 0x12, 0x9a, 0x59, 0x52, 0xaa, 0x8e, 0xaa,
	0x75, 0x0f, 0xaa, 0x35, 0x2f, 0x2c, 0xd9, 0xd7, 0xf2, 0xd6, 0x5e, 0x0a, 0xf3, 0xd6, 0xf0, 0xf6,
	0xb0, 0x81, 0x17, 0x7f, 0x06, 0xb0, 0xaf, 0xa1, 0x85, 0xe5, 0xf7, 0x2a, 0x34, 0x96, 0x03, 0xf1,
	0x7f, 0xa0, 0x03, 0x32, 0x33, 0x76, 0x0d, 0x9f, 0xe2, 0x37, 0x9d, 0xe4, 0x54, 0x5a, 0x2b, 0x75,
	0x1e, 0x9f, 0x78, 0x3a, 0xf7, 0x58, 0xe7, 0xdd, 0x14, 0xec, 0x07, 0xe8, 0x8d, 0x85, 0x0a, 0xc2,
	0x40, 0x58, 0x49, 0x9c, 0x5d, 0xea, 0xce, 0x4f, 0x72, 0xce, 0x4a, 0xf1, 0xf8, 0x65, 0x7b, 0x36,
	0x82, 0xfe, 0x44, 0xce, 0xc2, 0xc2, 0xcd, 0xd7, 0xfb, 0x2f, 0x86, 0x8a, 0x03, 0xbe, 0x0a, 0x98,
	0x23, 0x0a, 0xdf, 0xa7, 0xbc, 0x65, 0x32, 0xfb, 0x02, 0x76, 0xa7, 0xf1, 0x72, 0x19, 0x69, 0x3b,
	0xd2, 0xf3, 0x78, 0x81, 0x9b, 0xe3, 0x0f, 0xc8, 0x66, 0x0d, 0x47, 0xdb, 0xd1, 0x5c, 0x84, 0xca,
	0x14, 0x6c, 0x77, 0x13, 0xdb, 0x2a, 0x8e, 0x65, 0xed, 0xfc, 0xf9, 0x43, 0xea, 0xdb, 0x54, 0xc4,
	0xbb, 0xe5, 0x52, 0xde, 0x5a, 0xce, 0xc8, 0x93, 0xbe, 0xd9, 0xe7, 0xd0, 0xc7, 0x15, 0x9d, 0xdc,
	0xca, 0x59, 0x6c, 0xe5, 0x59, 0x70, 0xcb, 0x1f, 0x91, 0x53, 0x05, 0xc5, 0x5b, 0x04, 0x91, 0xb3,
	0x80, 0x0f, 0x92, 0x5b, 0x24, 0x91, 0xd8, 0x4b, 0xe8, 0x17, 0x2f, 0x02, 0x69, 0xf8, 0x63, 0x3a,
	0xba, 0xbd, 0x34, 0x49, 0x9b, 0xaf, 0x09, 0xbf, 0xe2, 0xe5, 0xfd, 0xb3, 0x05, 0xdd, 0x34, 0x9f,
	0xc7, 0x51, 0xb0, 0xc2, 0x9b, 0x39, 0x95, 0xb3, 0xc9, 0xa0, 0x80, 0xe4, 0x9d, 0x5e, 0xdf, 0xd8,
	0xe9, 0x8d, 0x0f, 0x74, 0xfa, 0xd6, 0x9d, 0x9d, 0xbe, 0xbd, 0xf1, 0x9d, 0xcc, 0x66, 0x25, 0xea,
	0xf8, 0xe4, 0x09, 0x28, 0x83, 0x14, 0x41, 0x4b, 0xbc, 0xc2, 0xdd, 0xf3, 0x96, 0x8a, 0xb4, 0x0f,
	0xfc, 0x94, 0x38, 0x03, 0xd1, 0x0b, 0x81, 0xfb, 0xc8, 0x10, 0x1a, 0xaa, 0xac, 0xd0, 0x96, 0xd4,
	0x6d, 0x37, 0x54, 0xa5, 0x00, 0xf2, 0x9e, 0xa8, 0x80, 0x74, 0x90, 0xf0, 0x3a, 0x11, 0x0f, 0x2e,
	0x5d, 0x02, 0xcd, 0x68, 0x86, 0x7a, 0xa6, 0xed, 0x57, 0x50, 0xf6, 0x7d, 0x6e, 0x47, 0x73, 0x95,
	0xe1, 0xdd, 0x6a, 0x8f, 0x16, 0xb4, 0x7e, 0xc5, 0x98, 0x9d, 0x01, 0x5b, 0x1b, 0x15, 0xd3, 0x67,
	0x68, 0xad, 0x11, 0x32, 0x0b, 0x7f, 0x83, 0x53, 0x91, 0x2a, 0xef, 0x67, 0x6a, 0x8b, 0x3b, 0x7b,
	0x6a, 0x83, 0x13, 0x1e, 0x0a, 0x4e, 0x63, 0x23, 0x6b, 0x75, 0x78, 0x1d, 0x5b, 0x49, 0x4f, 0x5b,
	0xd7, 0x2f, 0x83, 0xde, 0xfb, 0x5a, 0x5e, 0x53, 0xa7, 0x52, 0x04, 0xf4, 0xda, 0x8b, 0x50, 0x95,
	0x26, 0xdc, 0x02, 0x82, 0xd9, 0xbe, 0x92, 0x1a, 0xeb, 0xd3, 0x3d, 0x79, 0xa9, 0x88, 0x9e, 0x23,
	0x63, 0xdc, 0x24, 0xeb, 0x8a, 0xab, 0x80, 0xd0, 0xf0, 0xbc, 0x94, 0x3a, 0x3d, 0x66, 0xf7, 0x62,
	0x14, 0x20, 0xef, 0x8f, 0x1a, 0xb4, 0xd2, 0xc5, 0xb0, 0x3e, 0xd4, 0xc3, 0xb4, 0xa8, 0xeb, 0x61,
	0xc0, 0x8e, 0xca, 0x0b, 0xa5, 0xe8, 0x9d, 0xc3, 0x41, 0x35, 0x29, 0xa8, 0xf3, 0xcb, 0x5b, 0x3a,
	0x2a, 0xb7, 0x0d, 0x2d, 0x6d, 0x83, 0x27, 0xea, 0xfc, 0x92, 0xe5, 0xf1, 0x67, 0x30, 0x98, 0x45,
	0x8b, 0x83, 0x58, 0x85, 0x37, 0x32, 0x98, 0x4b, 0xed, 0x5c, 0x8e, 0x7b, 0x17, 0xf8, 0x9b, 0x9a,
	0x5e, 0x27, 0xff, 0x19, 0xbe, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x12, 0xab, 0xa5, 0x4c,
	0x0c, 0x00, 0x00,
}
