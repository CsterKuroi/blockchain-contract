package expression

import (
	"github.com/astaxie/beego/logs"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
)

type Function struct {
	GeneralExpression
}

func NewFunction() *Function {
	nf := &Function{}
	return nf
}

//===============接口实现===================
func (f Function) SetContract(p_contract inf.ICognitiveContract) {
	f.GeneralExpression.SetContract(p_contract)
}
func (f Function) GetContract() inf.ICognitiveContract {
	return f.GeneralExpression.GetContract()
}
func (f Function) GetName() string {
	return f.GeneralExpression.GetCname()
}
func (f Function) GetCtype() string {
	return f.GeneralExpression.GetCtype()
}
func (f Function) GetExpressionStr() string {
	return f.GeneralExpression.GetExpressionStr()
}
func (f Function) SetExpressionResult(p_expresult interface{}) {
	f.GeneralExpression.SetExpressionResult(p_expresult)
}
func (f Function) CleanValueInProcess() {
	f.GeneralExpression.CleanValueInProcess()
}

//===============描述态=====================
func (f *Function) ToString() string {
	return f.GetCname()
}

//===============运行态=====================
func (f *Function) InitFunction() error {
	var err error = nil
	err = f.InitExpression()
	if err != nil {
		logs.Warning("InitExpression fail[" + err.Error() + "]")
		return err
	}
	f.SetCtype(constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Function])

	return err
}
