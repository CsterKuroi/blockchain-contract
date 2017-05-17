package task

import (
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/constdef"
	"github.com/astaxie/beego/logs"
)
//表示场景：查询一个数据，并且会根据查询结果进行分支处理
//    为保证多节点执行结果一致性，需要在结果赋值时，使用分支条件赋值
//如：查询账户余额mount, 1.mount>100分支  2.amount>50 分支  3 amount<50 分支
//    而多节点查询时机和查询结果可能不一致，无法共识，因此只要查询结果满足范围即可
type Enquiry struct {
	GeneralTask
}

func NewEnquiry() *Enquiry {
	e := &Enquiry{}
	return e
}

//===============接口实现===================
func (e Enquiry)SetContract(p_contract inf.ICognitiveContract){
	e.GeneralTask.SetContract(p_contract)
}

func (e Enquiry)GetContract() inf.ICognitiveContract{
	return e.GeneralTask.GetContract()
}

func (e Enquiry)GetName()string {
	return e.GeneralTask.GetCname()
}

func (e Enquiry)GetCtype()string{
	return e.GeneralTask.GetCtype()
}
func (e Enquiry)GetDescription()string {
	return e.GeneralTask.GetDescription()
}

func (e Enquiry)GetState()string {
	return e.GeneralTask.GetState()
}

func (e Enquiry)SetState(p_state string){
	e.GeneralTask.SetState(p_state)
}

func (e Enquiry)GetNextTasks() []string{
	return e.GeneralTask.GetNextTasks()
}

func (e Enquiry)UpdateState()(int8, error){
    return e.GeneralTask.UpdateState()
}
//===============描述态=====================


//===============运行态=====================
func (e *Enquiry) InitEnquriy() error{
	var err error = nil
	err = e.InitGeneralTask()
	if err != nil {
		logs.Error("InitEnquriy fail["+err.Error()+"]")
		return err
	}
	e.SetCtype(constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Enquiry])
	return err
}

//由于查询分支结果的不确定性，使用分支条件赋予预估值，使得多节点 不同时运行结果一致性
func (e *Enquiry)SetEstimatedDataValue(){
	//TODO
}