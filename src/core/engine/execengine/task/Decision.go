package task

import (
	"bytes"
	"fmt"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type Decision struct {
	Enquiry
	CandidateList []DecisionCandidate `json:"CandidateList"` //决策结果都在CandidateList中体现；每个决策候选集，一个决策结果
}

const (
	_CandidateList = "_CandidateList"
)

func NewDecision() *Decision {
	decision := &Decision{}
	return decision
}

//====================接口方法========================
func (d Decision) SetContract(p_contract inf.ICognitiveContract) {
	d.Enquiry.SetContract(p_contract)
}

func (d Decision) GetContract() inf.ICognitiveContract {
	return d.Enquiry.GetContract()
}

func (d Decision) GetName() string {
	return d.Enquiry.GetName()
}

func (d Decision) GetCtype() string {
	return d.Enquiry.GetCtype()
}

func (d Decision) GetDescription() string {
	return d.Enquiry.GetDescription()
}

func (d Decision) GetState() string {
	return d.Enquiry.GetState()
}

func (d Decision) SetState(p_state string) {
	d.Enquiry.SetState(p_state)
}

func (d Decision) GetNextTasks() []string {
	return d.Enquiry.GetNextTasks()
}

func (d Decision) UpdateState(nBrotherNum int) (int8, error) {
	return d.Enquiry.UpdateState(nBrotherNum)
}
func (d Decision) GetTaskId() string {
	return d.Enquiry.GetTaskId()
}

func (d Decision) GetTaskExecuteIdx() int {
	return d.Enquiry.GetTaskExecuteIdx()
}

func (d Decision) SetTaskId(str_taskId string) {
	d.Enquiry.SetTaskId(str_taskId)
}

func (d Decision) SetTaskExecuteIdx(int_idx int) {
	d.Enquiry.SetTaskExecuteIdx(int_idx)
}

func (d Decision) CleanValueInProcess() {
	d.Enquiry.CleanValueInProcess()
	d.ResetCandidate()
}

func (d Decision) UpdateStaticState() (interface{}, error) {
	return d.Enquiry.UpdateStaticState()
}

//====================描述态==========================

//====================运行态==========================
func (d *Decision) AddProperty(object interface{}, str_name string, value interface{}) property.PropertyT {
	var pro_object property.PropertyT
	if value == nil {
		pro_object = *property.NewPropertyT(str_name)
		d.PropertyTable[str_name] = pro_object
		return pro_object
	}
	switch value.(type) {
	case map[string]DecisionCandidate:
		pro_object = property.PropertyT{Name: str_name}
		pro_object.SetValue(value.(map[string]DecisionCandidate))
		d.PropertyTable[str_name] = pro_object
	}
	return pro_object
}

func (d *Decision) InitDecision() error {
	var err error = nil
	err = d.InitEnquriy()
	if err != nil {
		uniledgerlog.Error("InitDecision fail[" + err.Error() + "]")
		return err
	}
	d.SetCtype(constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision])
	//condidatelist arrar to map
	if d.CandidateList == nil {
		d.CandidateList = make([]DecisionCandidate, 0)
	}
	map_candidatelist := make(map[string]DecisionCandidate, 0)
	for _, p_cand := range d.CandidateList {
		map_candidatelist[p_cand.GetName()] = p_cand
	}
	common.AddProperty(d, d.PropertyTable, _CandidateList, map_candidatelist)
	return err
}

//====属性Get方法
//TODO： map本身是无序的，不需排序
func (d *Decision) GetCandidateList() map[string]DecisionCandidate {
	candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	candlist_value, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return nil
	}
	return candlist_value
}

func (d *Decision) GetCandidate(p_name string) DecisionCandidate {
	candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return DecisionCandidate{}
	}
	if candlist_property.GetValue() != nil {
		candlist_map, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return DecisionCandidate{}
		}
		candlist_value, ok := candlist_map[p_name]
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.NULL_ERROR, ""))
			return DecisionCandidate{}
		}
		return candlist_value
	}
	return DecisionCandidate{}
}

//====动态添加方法
func (d *Decision) AddCandidate(p_candidate interface{}) {
	if p_candidate != nil {
		candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
		if !ok {
			candlist_property = *property.NewPropertyT(_CandidateList)
		}
		if candlist_property.GetValue() == nil {
			candlist_property.SetValue(make(map[string]DecisionCandidate, 0))
		}
		v_candidate, ok := p_candidate.(DecisionCandidate)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		if d.GetContract() != nil {
			v_candidate.SetContract(d.GetContract())
		}
		map_candlist, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
		if !ok {
			map_candlist = make(map[string]DecisionCandidate, 0)
		}
		map_candlist[v_candidate.GetCname()] = v_candidate
		candlist_property.SetValue(map_candlist)
		d.PropertyTable[_CandidateList] = candlist_property
	}
}

func (d *Decision) RemoveCandidate(p_candidate interface{}) {
	if p_candidate != nil {
		candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return
		}
		if candlist_property.GetValue() != nil {
			v_candidate, ok := p_candidate.(DecisionCandidate)
			if !ok {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
				return
			}
			map_candlist, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
			if !ok {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
				return
			}
			delete(map_candlist, v_candidate.GetCname())
			candlist_property.SetValue(map_candlist)
			d.PropertyTable[_CandidateList] = candlist_property
		}
	}
}

func (d *Decision) evaluateCandidate() error {
	var err error = nil
	candlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return err
	}
	if candlist_property.GetValue() != nil {
		candlist_map, ok := candlist_property.GetValue().(map[string]DecisionCandidate)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return err
		}
		for _, v_value := range candlist_map {
			v_value.SetContract(d.GetContract())
			v_value.ResetDecisionCandidate()
			err := v_value.Eval()
			if err != nil {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.EXECUTE_ERROR, ""))
				return err
			}
		}
	}
	return err
}

func (d *Decision) ResetCandidate() {
	resultlist_property, ok := d.PropertyTable[_CandidateList].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return
	}
	if resultlist_property.GetValue() != nil {
		resultlist_property.SetValue(make(map[string]DecisionCandidate, 0))
	}
	d.PropertyTable[_CandidateList] = resultlist_property
}

//针对决策单独进行Start操作
func (gt *Decision) Start() (int8, error) {
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Task Process Runing:Dormant State.")
	r_buf.WriteString("[ContractID]: " + gt.GetContract().GetContractId() + ";")
	r_buf.WriteString("[ContractHashID]: " + gt.GetContract().GetId() + ";")
	r_buf.WriteString("[TaskName]: " + gt.GetName() + ";")
	uniledgerlog.Info(r_buf.String(), " begin....")
	var r_ret int8 = 0
	var r_err error = nil
	if gt.IsDormant() && gt.testPreCondition() {
		err := gt.evaluateCandidate()
		//执行失败，返回 -1
		if err != nil {
			r_ret = -1
			r_buf.WriteString("[Result]: Task execute fail;")
			r_buf.WriteString("[Error]: " + r_err.Error() + ";")
			r_buf.WriteString("fail....")
			uniledgerlog.Error(r_buf.String())
			return r_ret, r_err
		}
		r_buf.WriteString("[Result]: Task execute success;")
		uniledgerlog.Info(r_buf.String(), " Dormant to Inprocess....")
		gt.SetState(constdef.TaskState[constdef.TaskState_In_Progress])
	} else if gt.IsDormant() && !gt.testPreCondition() { //未达到执行条件，返回 0
		r_ret = 0
		r_buf.WriteString("[Result]: preCondition not true;")
		uniledgerlog.Warn(r_buf.String(), " exit....")
		return r_ret, r_err
	}
	//执行完动作后需要等待执行完成
	r_ret, r_err = gt.Complete()

	return r_ret, r_err
}
