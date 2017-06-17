package task

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type DecisionCandidate struct {
	Enquiry
	SupportArguments []string `json:"SupportArguments"`
	AgainstArguments []string `json:"AgainstArguments"`
	Support          int      `json:"Support"`
	Text             []string `json:"Text"`
}

const (
	_SupportArguments = "_SupportArguments"
	_AgainstArguments = "_AgainstArguments"
	_Support          = "_Support"
	_Text             = "_Text"
)

func NewDecisionCandidate() *DecisionCandidate {
	d := &DecisionCandidate{}
	return d
}

//===============接口实现===================
func (dc DecisionCandidate) SetContract(p_contract inf.ICognitiveContract) {
	dc.Enquiry.SetContract(p_contract)
}

func (dc DecisionCandidate) GetContract() inf.ICognitiveContract {
	return dc.Enquiry.GetContract()
}

func (dc DecisionCandidate) CleanValueInProcess() {
	dc.Enquiry.CleanValueInProcess()
	dc.ResetSupport()
}

//===============描述态=====================

//===============运行态=====================
func (dc *DecisionCandidate) InitDecisionCandidate() error {
	var err error = nil
	err = dc.InitEnquriy()
	if err != nil {
		logs.Error("InitDecisionCandidate fail[" + err.Error() + "]")
		return err
	}
	dc.SetCtype(constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_DecisionCandidate])
	//supportArguments
	if dc.SupportArguments == nil {
		dc.SupportArguments = make([]string, 0)
	}
	map_supportArgument := make(map[string]string, 0)
	for _, p_support := range dc.SupportArguments {
		map_supportArgument[p_support] = p_support
	}
	common.AddProperty(dc, dc.PropertyTable, _SupportArguments, map_supportArgument)
	//againstArguments
	if dc.AgainstArguments == nil {
		dc.AgainstArguments = make([]string, 0)
	}
	map_againstArgument := make(map[string]string, 0)
	for _, p_against := range dc.AgainstArguments {
		map_againstArgument[p_against] = p_against
	}
	common.AddProperty(dc, dc.PropertyTable, _AgainstArguments, map_againstArgument)
	//support
	common.AddProperty(dc, dc.PropertyTable, _Support, 0)
	//text
	if dc.Text == nil {
		dc.Text = make([]string, 0)
	}
	map_Text := make(map[string]string, 0)
	for _, p_text := range dc.Text {
		map_Text[p_text] = p_text
	}
	common.AddProperty(dc, dc.PropertyTable, _Text, map_Text)
	return err
}

func (dc *DecisionCandidate) AddText(p_strarr []string) {
	if p_strarr != nil {
		text_property := dc.PropertyTable[_Text].(property.PropertyT)
		if text_property.GetValue() == nil {
			text_property.SetValue(make([]string, 0))
		}
		map_text := text_property.GetValue().(map[string]string)
		for _, v_Text := range p_strarr {
			map_text[v_Text] = v_Text
		}
		text_property.SetValue(map_text)
		dc.PropertyTable[_Text] = text_property
	}
}

func (dc *DecisionCandidate) ShowText() {
	text_property := dc.PropertyTable[_Text].(property.PropertyT)
	if text_property.GetValue() != nil {
		map_text := text_property.GetValue().(map[string]string)
		for _, v_Text := range map_text {
			fmt.Println(v_Text)
		}
	}
}

func (dc *DecisionCandidate) AddSupportArgument(p_Support string) {
	if p_Support != "" {
		supports_property := dc.PropertyTable[_SupportArguments].(property.PropertyT)
		if supports_property.GetValue() == nil {
			supports_property.SetValue(make(map[string]string, 0))
		}
		map_supports := supports_property.GetValue().(map[string]string)
		map_supports[p_Support] = p_Support
		supports_property.SetValue(map_supports)
		dc.PropertyTable[_SupportArguments] = supports_property
	}
}

func (dc *DecisionCandidate) AddAgainstArgument(p_against string) {
	if p_against != "" {
		against_property := dc.PropertyTable[_AgainstArguments].(property.PropertyT)
		if against_property.GetValue() == nil {
			against_property.SetValue(make(map[string]string, 0))
		}
		map_againsts := against_property.GetValue().(map[string]string)
		map_againsts[p_against] = p_against
		against_property.SetValue(map_againsts)
		dc.PropertyTable[_SupportArguments] = against_property
	}
}

func (dc *DecisionCandidate) ResetSupport() {
	dc.Support = 0
	support_property := dc.PropertyTable[_Support].(property.PropertyT)
	support_property.SetValue(0)
	dc.PropertyTable[_Support] = support_property
}

func (dc *DecisionCandidate) GetSupport() int {
	dc.Eval()
	return dc.Support
}

func (dc *DecisionCandidate) Eval() int {
	var Support_sum int = 0
	var against_sum int = 0
	var Support_bool interface{}
	var against_bool interface{}
	var err error = nil
	supports_property := dc.PropertyTable[_SupportArguments].(property.PropertyT)
	if supports_property.GetValue() != nil {
		for _, v_Support := range supports_property.GetValue().(map[string]string) {
			v_contract := dc.GetContract()
			Support_bool, err = v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], v_Support)
			if err != nil {
				logs.Warning("DecisionCandidate.Eval fail[" + err.Error() + "]")
			}
			if Support_bool.(bool) {
				Support_sum += 1
			}
		}
	}
	against_property := dc.PropertyTable[_AgainstArguments].(property.PropertyT)
	if against_property.GetValue() != nil {
		for _, v_against := range against_property.GetValue().(map[string]string) {
			v_contract := dc.GetContract()
			against_bool, err = v_contract.EvaluateExpression(constdef.ExpressionType[constdef.Expression_Condition], v_against)
			if err != nil {
				logs.Warning("DecisionCandidate.Eval fail[" + err.Error() + "]")
			}
			if against_bool.(bool) {
				against_sum += 1
			}
		}
	}
	dc.Support = Support_sum - against_sum
	return dc.Support
}
