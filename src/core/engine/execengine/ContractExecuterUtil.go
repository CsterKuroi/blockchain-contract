package execengine

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/contract"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/task"
)

func loadTask(p_contract *contract.CognitiveContract, p_component interface{}) error {
	var err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract LoadTask;")
	if p_component == nil {
		err = errors.New("Param[component] is null!")
		r_buf.WriteString("[Result]:　LoadTask fail;")
		r_buf.WriteString("[Error]: Param[component] is null;")
		logs.Warning(r_buf.String())
		return err
	}
	map_task, ok := p_component.(map[string]interface{})
	if !ok {
		logs.Error("assert error")
		return fmt.Errorf("assert error")
	}
	switch map_task["Ctype"] {
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Enquiry]:
		//1.反序列化
		p_task := task.NewEnquiry()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		//fmt.Println("333333333: ", p_task)
		//2 初始化
		p_task.InitEnquriy()
		//fmt.Println("444444444: ", p_task)
		//3 处理数组属性
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			r_buf.WriteString("[Result]:　loadTaskInnerComponent[Enquiry] fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//4 添加任务组件到component_table中
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Action]:
		p_task := task.NewAction()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitAction()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			r_buf.WriteString("[Result]:　loadTaskInnerComponent[Action] fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision]:
		p_task := task.NewDecision()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitDecision()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			r_buf.WriteString("[Result]:　loadTaskInnerComponent[Decision] fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Plan]:
		p_task := task.NewPlan()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitPlan()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			r_buf.WriteString("[Result]:　loadTaskInnerComponent[Decision] fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		p_contract.AddComponent(p_task)
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Unknown]:
		p_task := task.NewGeneralTask()
		byte_task, _ := json.Marshal(map_task)
		err = json.Unmarshal(byte_task, &p_task)
		p_task.InitGeneralTask()
		if err := loadTaskInnerComponent(p_contract, map_task, p_task); err != nil {
			r_buf.WriteString("[Result]:　loadTaskInnerComponent[Unknow] fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		p_contract.AddComponent(p_task)
	}
	r_buf.WriteString("[Cname]: " + map_task["Cname"].(string) + "[Ctype]: " + map_task["Ctype"].(string) + "[Result]: LoadTask success;")
	logs.Info(r_buf.String())
	return err
}

func loadTaskInnerComponent(p_contract *contract.CognitiveContract, m_task interface{}, p_task interface{}) error {
	var err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("loadTaskInnerComponent;")
	if m_task == nil || p_task == nil {
		r_buf.WriteString("[Result]: loadTaskInnerComponent fail;")
		r_buf.WriteString("[Error]: " + err.Error() + ";")
		logs.Warning(r_buf.String())
		return err
	}
	map_task, ok := m_task.(map[string]interface{})
	if !ok {
		logs.Error("assert error")
		return fmt.Errorf("assert error")
	}
	switch map_task["Ctype"] {
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Enquiry]:
		pre_conditions := map_task["PreCondition"]
		sl1, ok := pre_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl1 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Enquiry.PreCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		sl2, ok := comp_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl2 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Enquiry.CompleteCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		sl3, ok := digard_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl3 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Enquiry.DiscardCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		data_list := map_task["DataList"]
		sl4, ok := data_list.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl4 {
			if err := loadData(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadData[Enquiry.DataList] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		dataexpress_list := map_task["DataValueSetterExpressionList"]
		sl5, ok := dataexpress_list.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl5 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Enquiry.DataValueSetterExpressionList] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Action]:
		pre_conditions := map_task["PreCondition"]
		sl6, ok := pre_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl6 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Action.PreCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		sl7, ok := comp_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl7 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Action.CompleteCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		sl8, ok := digard_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl8 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Action.DiscardCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		data_list := map_task["DataList"]
		sl9, ok := data_list.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl9 {
			if err := loadData(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadData[Action.DataList] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		dataexpress_list := map_task["DataValueSetterExpressionList"]
		sl10, ok := dataexpress_list.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl10 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Action.DataValueSetterExpressionList] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Decision]:
		pre_conditions := map_task["PreCondition"]
		sl11, ok := pre_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl11 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Decision.PreCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		sl12, ok := comp_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl12 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Decision.CompleteCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		sl13, ok := digard_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl13 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Decision.DiscardCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		candidate_list := map_task["CandidateList"]
		sl14, ok := candidate_list.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl14 {
			if err := loadCandidate(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadCandidate[Decision.CandidateList] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.TaskType[constdef.Task_Plan]:
		pre_conditions := map_task["PreCondition"]
		sl16, ok := pre_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl16 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Plan.PreCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		sl100, ok := comp_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl100 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Plan.CompleteCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		sl17, ok := digard_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl17 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Plan.DiscardCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
	default:
		pre_conditions := map_task["PreCondition"]
		sl18, ok := pre_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl18 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Unknow.PreCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		comp_conditions := map_task["CompleteCondition"]
		sl19, ok := comp_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl19 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Unknow.CompleteCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
		digard_conditions := map_task["DiscardCondition"]
		sl20, ok := digard_conditions.([]interface{})
		if !ok {
			logs.Error("assert error")
			return fmt.Errorf("assert error")
		}
		for _, p_value := range sl20 {
			if err := loadExpression(p_contract, p_value, p_task); err != nil {
				r_buf.WriteString("[Result]: loadExpression[Unknow.DiscardCondition] fail;")
				r_buf.WriteString("[Error]: " + err.Error() + ";")
				logs.Warning(r_buf.String())
				return err
			}
		}
	}
	r_buf.WriteString("[Cname]: " + map_task["Cname"].(string) + "[Ctype]: " + map_task["Ctype"].(string) + "[Result]: loadTaskInnerComponent success;")
	logs.Info(r_buf.String())
	return err
}

func loadData(p_contract *contract.CognitiveContract, m_data interface{}, p_task interface{}) error {
	var err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("loadData;")
	if p_contract == nil || m_data == nil || p_task == nil {
		err = errors.New("Param[p_contract || m_data || p_task] is null!")
		r_buf.WriteString("[Result]: loadData fail;")
		r_buf.WriteString("[Error]: Param[p_contract || m_data || p_task] is null;")
		logs.Warning(r_buf.String())
		return err
	}
	map_data, ok := m_data.(map[string]interface{})
	if !ok {
		logs.Error("assert error")
		return fmt.Errorf("assert error")
	}
	switch map_data["Ctype"] {
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Int]:
		//1.反序列化
		p_data := data.NewIntData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Numeric_Int)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Numeric_Int )Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("44444444: ", p_data)
		//2 初始化
		err = p_data.InitIntData()
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Numeric_Int )InitIntData fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
		//4 添加data组件到task中
		//v_task := p_task.(inf.ITask)
		//v_task.AddData()
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Uint]:
		//1.反序列化
		p_data := data.NewUintData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Numeric_Uint)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Numeric_Uint)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("44444444: ", p_data)
		//2 初始化
		err = p_data.InitUintData()
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Numeric_Uint)InitUintData fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Numeric_Float]:
		//1.反序列化
		p_data := data.NewFloatData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Numeric_Float)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Numeric_Float)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("44444444: ", p_data)
		//2 初始化
		err = p_data.InitFloatData()
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Numeric_Uint)InitFloatData fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Text]:
		//1.反序列化
		p_data := data.NewTextData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Text)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Text)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("44444444: ", p_data)
		//2 初始化
		err = p_data.InitTextData()
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Text)InitTextData fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Date]:
		//1.反序列化
		p_data := data.NewDateData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Date)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Date)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("44444444: ", p_data)
		//2 初始化
		err = p_data.InitDateData()
		if err != nil {
			r_buf.WriteString("[Result]: loadData(Data_Date)InitDateData fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Array]:
		//1.反序列化
		p_data := data.NewArrayData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(Data_Array)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(Data_Array)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("44444444: ", p_data)
		//2 初始化
		err = p_data.InitArrayData()
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(Component_Data)InitArrayData fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Matrix]:
		//1.反序列化
		p_data := data.NewMatrixData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(Data_Matrix)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(Data_Matrix)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("44444444: ", p_data)
		//2 初始化
		err = p_data.InitMatrixData()
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(Data_Matrix)InitMatrixData fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	case constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Compound]:
		//1.反序列化
		p_data := data.NewCompoundData()
		byte_data, err := json.Marshal(map_data)
		r_buf.WriteString("[Result]: Component_Data(Data_Compound)Marshal fail;")
		r_buf.WriteString("[Error]: " + err.Error() + ";")
		logs.Warning(r_buf.String())
		return err
		err = json.Unmarshal(byte_data, &p_data)
		r_buf.WriteString("[Result]: Component_Data(Data_Compound)Unmarshal fail;")
		r_buf.WriteString("[Error]: " + err.Error() + ";")
		logs.Warning(r_buf.String())
		return err
		//fmt.Println("44444444: ", p_data)
		//2 初始化
		err = p_data.InitCompoundData()
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(Data_Compound)InitCompoundData fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	default:
		//1.反序列化
		p_data := data.NewGeneralData()
		byte_data, err := json.Marshal(map_data)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(GeneralData)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_data, &p_data)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(GeneralData)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("44444444: ", p_data)
		//2 初始化
		err = p_data.InitGeneralData()
		if err != nil {
			r_buf.WriteString("[Result]: Component_Data(GeneralData)InitGeneralData fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("55555555: ", p_data)
		//3 添加data组件到component_table中
		p_contract.AddComponent(p_data)
	}
	r_buf.WriteString("[Cname]: " + map_data["Cname"].(string) + "[Ctype]: " + map_data["Ctype"].(string) + "[Result]: loadData success;")
	logs.Info(r_buf.String())
	return err
}

func loadExpression(p_contract *contract.CognitiveContract, p_expression interface{}, p_task interface{}) error {
	var err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("loadExpression...;")
	if p_contract == nil || p_task == nil || p_expression == nil {
		err = errors.New("Param[p_contract or p_task or expression] is null!")
		r_buf.WriteString("[Result]: loadExpression fail;")
		r_buf.WriteString("[Error]:　" + err.Error() + ";")
		logs.Warning(r_buf.String())
		return err
	}
	map_expression, ok := p_expression.(map[string]interface{})
	if !ok {
		logs.Error("assert error")
		return fmt.Errorf("assert error")
	}

	switch map_expression["Ctype"] {
	case constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Condition]:
		object_expression := expression.NewLogicArgument()
		byte_task, err := json.Marshal(map_expression)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Expression(Expression_Condition)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_task, &object_expression)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Expression(Expression_Condition)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("777777777: ", object_expression)
		err = object_expression.InitLogicArgument()
		if err != nil {
			r_buf.WriteString("[Result]: Component_Expression(Expression_Condition)InitLogicArgument fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		p_contract.AddComponent(object_expression)
	case constdef.ComponentType[constdef.Component_Expression] + "." + constdef.ExpressionType[constdef.Expression_Function]:
		object_expression := expression.NewFunction()
		byte_task, err := json.Marshal(map_expression)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Expression(Expression_Function)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_task, &object_expression)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Expression(Expression_Function)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("777777777: ", object_expression)
		err = object_expression.InitFunction()
		if err != nil {
			r_buf.WriteString("[Result]: Component_Expression(Expression_Function)InitFunction fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		p_contract.AddComponent(object_expression)
	default:
		object_expression := &expression.GeneralExpression{}
		byte_task, err := json.Marshal(map_expression)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Expression(Expression_Unknow)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_task, &object_expression)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Expression(Expression_Unknow)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		//fmt.Println("777777777: ", object_expression)
		err = object_expression.InitExpression()
		if err != nil {
			r_buf.WriteString("[Result]: Component_Expression(Expression_Unknow)InitExpression fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		p_contract.AddComponent(object_expression)
	}
	r_buf.WriteString("[Cname]: " + map_expression["Cname"].(string) + "[Ctype]: " + map_expression["Ctype"].(string) + "[Result]: loadExpression success;")
	logs.Info(r_buf.String())
	return err
}

func loadCandidate(p_contract *contract.CognitiveContract, p_candidate interface{}, p_task interface{}) error {
	var err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("loadCandidate...;")
	if p_contract == nil || p_task == nil || p_candidate == nil {
		err = errors.New("Param[p_contract or p_task or p_candidate] is null!")
		r_buf.WriteString("[Result]: loadCandidate fail;")
		r_buf.WriteString("[Error]:　" + err.Error() + ";")
		logs.Warning(r_buf.String())
		return err
	}
	map_candidate := p_candidate.(map[string]interface{})
	switch map_candidate["Ctype"] {
	case constdef.ComponentType[constdef.Component_Task] + "." + constdef.ExpressionType[constdef.Task_DecisionCandidate]:
		object_candidate := task.NewDecisionCandidate()
		byte_task, err := json.Marshal(map_candidate)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Task(Task_DecisionCandidate)Marshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = json.Unmarshal(byte_task, &object_candidate)
		if err != nil {
			r_buf.WriteString("[Result]: Component_Task(Task_DecisionCandidate)Unmarshal fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		err = object_candidate.InitDecisionCandidate()
		if err != nil {
			r_buf.WriteString("[Result]: Component_Task(Task_DecisionCandidate)InitDecisionCandidate fail;")
			r_buf.WriteString("[Error]: " + err.Error() + ";")
			logs.Warning(r_buf.String())
			return err
		}
		p_contract.AddComponent(object_candidate)
	}
	r_buf.WriteString("[Cname]: " + map_candidate["Cname"].(string) + "[Ctype]: " + map_candidate["Ctype"].(string) + "[Result]: loadCandidate success;")
	logs.Info(r_buf.String())
	return err
}
