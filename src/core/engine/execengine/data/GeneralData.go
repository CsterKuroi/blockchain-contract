package data

import (
	"strconv"
	"unicontract/src/core/engine/execengine/component"
	"unicontract/src/core/engine/execengine/inf"
	"fmt"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
)

type GeneralData struct{
	component.GeneralComponent
	Value interface{} `json:"Value"`
	DefaultValue interface{} `json:"DefaultValue"`
	Unit string `json:"Unit"`
	ModifyDate string `json:"ModifyDate"`
	HardConvType string `json:"HardConvType"`
	Mandatory bool `json:"Mandatory"`
	Category []string `json:"Category"`
	Options map[string]int `json:"Options"`
	Parent inf.IData `json:"Parent"`
}

const (
	_Value = "_Value"
	_DefaultValue = "_DefaultValue"
	_Unit = "_Unit"
	_ModifyDate = "_ModifyDate"
	_HardConvType = "_HardConvType"
	_Mandatory = "_Mandatory"
	_Category = "_Category"
	_Options = "_Options"
	_Parent = "_Parent"
	_DataRange = "_DataRange"
)

func NewGeneralData()*GeneralData{
	d := &GeneralData{}
	return d
}
//===============接口实现===================
func (gd GeneralData)GetName()string{
	if gd.PropertyTable[_Parent] != nil  {
		parent_property := gd.PropertyTable[_Parent].(property.PropertyT)
		if parent_property.GetValue() != nil {
			v_general_data := parent_property.GetValue().(inf.IData)
			if v_general_data.GetName() != "" {
				return v_general_data.GetName() + "." + gd.GetCname()
			} else {
				return gd.GetCname()
			}
		}
	}
	return gd.GetCname()
}

func (gd GeneralData) GetValue() interface{}{
	value_property := gd.PropertyTable[_Value].(property.PropertyT)
	if value_property.GetValue() != nil {
		return value_property.GetValue()
	} else {
		var v_default interface{}
		switch gd.GetDefaultValue().(type) {
		case string:
			v_default = gd.GetDefaultValue().(string)
			v_contract := gd.GeneralComponent.GetContract()
			v_default = v_contract.EvaluateExpression(gd.GetDefaultValue().(string))
		default:
			v_default = gd.GetDefaultValue()
		}
		return v_default
	}
}
func (gd GeneralData)SetContract(p_contract inf.ICognitiveContract) {
	gd.GeneralComponent.SetContract(p_contract)
}

func (gd GeneralData)GetContract() inf.ICognitiveContract {
	return gd.GeneralComponent.GetContract()
}
func (gc GeneralData)GetCtype()string{
	if gc.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable["_Ctype"].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
//===============描述态=====================
func (gd *GeneralData)ToString() interface{}{
	value_property := gd.PropertyTable[_Value].(property.PropertyT)
	return value_property.GetValue()
}
//===============运行态=====================
func (gd *GeneralData) InitGeneralData()error {
	var err error = nil
	/*
	if gd.Value == nil {
		//TODO log
		err = errors.New("GeneralData need Value!")
		return err
	}
	if gd.DefaultValue == nil {
		//TODO log
		err = errors.New("GeneralData need DefaultValue!")
		return err
	}*/
	err = gd.InitGeneralComponent()
	if err != nil {
		//TODO log
		return err
	}
	gd.SetCtype(constdef.ComponentType[constdef.Component_Data])

	gd.AddProperty(gd, _Value, gd.Value)
	gd.AddProperty(gd, _DefaultValue, gd.DefaultValue)
	gd.AddProperty(gd, _HardConvType, gd.HardConvType)
	gd.AddProperty(gd, _Unit, gd.Unit)
	gd.AddProperty(gd, _ModifyDate, gd.ModifyDate)
	gd.AddProperty(gd, _Mandatory, gd.Mandatory)
	if gd.Category == nil {
		gd.AddProperty(gd, _Category, make([]string, 0))
	} else {
		gd.AddProperty(gd, _Category, gd.Category)
	}
	if gd.Options == nil {
		gd.AddProperty(gd, _Options, make(map[string]int, 0))
	} else {
		gd.AddProperty(gd, _Options, gd.Options)
	}
	gd.AddProperty(gd, _Parent, gd.Parent)

	return err
}

func (gd *GeneralData) Aquired()bool {
	value_property := gd.PropertyTable[_Value].(property.PropertyT)
	if value_property.GetValue() != nil {
		return true
	} else {
		return false
	}
}

func (gd *GeneralData) ResetOptions(p_listoption []string){
	option_property := gd.PropertyTable[_Options].(property.PropertyT)
	if option_property.GetValue() == nil {
		option_property.SetValue(make(map[string]int, 0))
	}
	option_map := option_property.GetValue().(map[string]int)
	if p_listoption != nil {
		for _,Value := range p_listoption {
			option_map[Value] = 0
		}
	}
	option_property.SetValue(option_map)
	gd.PropertyTable[_Options] = option_map
	gd.Options = option_map
}

func (gd *GeneralData) InputOptions(){
	option_property := gd.PropertyTable[_Options].(property.PropertyT)
	if option_property.GetValue() == nil {
		option_property.SetValue(make(map[string]int, 0))
	}
	option_map := option_property.GetValue().(map[string]int)
	fmt.Println("input 1->yes  0->no")
	var input string = ""
	var err error = nil
	for key,_ := range option_map {
		fmt.Println("do you have ", key, " ?")
		_,err_input := fmt.Scanln(&input)
		if err_input != nil {
			option_map[key],err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Get input[", key, "] Error!")
			}
		} else {
			fmt.Println("Get input[", key, "] Error!")
		}
	}
	option_property.SetValue(option_map)
	gd.PropertyTable[_Options] = option_map
	gd.Options = option_map
}

func (gd *GeneralData) Optionsum()int {
	option_property := gd.PropertyTable[_Options].(property.PropertyT)
	if option_property.GetValue() == nil {
		option_property.SetValue(make(map[string]int, 0))
	}
	option_map := option_property.GetValue().(map[string]int)
	var sum int = 0
	if gd.Options != nil {
		for key, _ := range option_map {
			sum = sum + option_map[key]
		}
	}
	return sum
}

func (gd *GeneralData) IsCategorical()bool{
	category_property := gd.PropertyTable[_Category].(property.PropertyT)
	var r_flag bool = false
	if category_property.GetValue() != nil && len(category_property.GetValue().([]string)) > 0 {
		r_flag = true
	} else {
		r_flag = false
	}
	return r_flag
}

func (gd *GeneralData) AddCategory(arr_Category []string){
	category_property := gd.PropertyTable[_Category].(property.PropertyT)
	if category_property.GetValue() == nil {
		category_property.SetValue(make([]string, 0))
	}
	category_arr := category_property.GetValue().([]string)
	if len(arr_Category) > 0 {
		for _,f_Value := range arr_Category {
			category_arr = append(category_arr, f_Value)
		}
	}
	category_property.SetValue(category_arr)
	gd.PropertyTable[_Category] = category_property
	gd.Category = category_property.GetValue().([]string)
}

func (gd *GeneralData) RemoveCategory(arr_Category []interface{}) {
	category_property := gd.PropertyTable[_Category].(property.PropertyT)
	if category_property.GetValue() == nil {
		category_property.SetValue(make([]interface{}, 0))
	}
	category_arr := category_property.GetValue().([]interface{})
	if category_arr != nil && len(category_arr) > 0 {
		if len(arr_Category) >0 {
			for _,f_Value := range arr_Category {
				for g_idx, g_value := range category_arr {
					if g_value == f_Value {
						category_arr = append(category_arr[:g_idx],category_arr[g_idx+1:]...)
					}
				}
			}
		}
	}
	category_property.SetValue(category_arr)
	gd.PropertyTable[_Category] = category_property
	gd.Category = category_property.GetValue().([]string)
}

func (gd *GeneralData) CheckRange(p_Value interface{}) bool {
	category_property := gd.PropertyTable[_Category].(property.PropertyT)
	if category_property.GetValue() == nil {
		category_property.SetValue(make([]interface{}, 0))
	}
	category_arr := category_property.GetValue().([]interface{})
	var r_flag bool = false
	if category_arr == nil || len(category_arr) == 0 {
		r_flag = false
	} else {
		for _,f_Value := range category_arr {
			if f_Value == p_Value {
				r_flag = true
				break
			}
		}
	}
	return r_flag
}

//====属性Get方法
func (gd *GeneralData) GetDefaultValue()interface{}{
	defaultvalue_property := gd.PropertyTable[_DefaultValue].(property.PropertyT)
	return defaultvalue_property.GetValue()
}

func (gd *GeneralData) GetUnit() string{
	unit_property := gd.PropertyTable[_Unit].(property.PropertyT)
	return unit_property.GetValue().(string)
}
func (gd *GeneralData) GetModifyDate()string{
	modifydate_property := gd.PropertyTable[_ModifyDate].(property.PropertyT)
	return modifydate_property.GetValue().(string)
}

func (gd *GeneralData) GetHardConvType()string{
	hardconvtype_property := gd.PropertyTable[_HardConvType].(property.PropertyT)
	return hardconvtype_property.GetValue().(string)
}

func (gd *GeneralData) GetCategory()[]string{
	category_property := gd.PropertyTable[_Category].(property.PropertyT)
	return category_property.GetValue().([]string)
}

func (gd *GeneralData) GetParent()inf.IData{
	parent_property := gd.PropertyTable[_Parent].(property.PropertyT)
	return parent_property.GetValue().(inf.IData)
}

func (gd *GeneralData) GetMandatory()bool{
	mandatory_property := gd.PropertyTable[_Mandatory].(property.PropertyT)
	return mandatory_property.GetValue().(bool)
}
//====属性Set方法
func (gd *GeneralData) SetDefaultValue(p_DefaultValue interface{}){
	gd.DefaultValue = p_DefaultValue
	defaultvalue_property := gd.PropertyTable[_DefaultValue].(property.PropertyT)
	defaultvalue_property.SetValue(p_DefaultValue)
	gd.PropertyTable[_DefaultValue] = defaultvalue_property
}

func (gd *GeneralData) SetValue(p_Value interface{}){
	gd.Value = p_Value
	value_property := gd.PropertyTable[_Value].(property.PropertyT)
	value_property.SetValue(p_Value)
	gd.PropertyTable[_Value] = value_property
}

func (gd *GeneralData) SetUnit(p_Unit string){
	gd.Unit = p_Unit
	unit_property := gd.PropertyTable[_Unit].(property.PropertyT)
	unit_property.SetValue(p_Unit)
	gd.PropertyTable[_Unit] = unit_property
}

func (gd *GeneralData) SetHardConvType(p_convtype string){
	gd.HardConvType = p_convtype
	hardconvtype_property := gd.PropertyTable[_HardConvType].(property.PropertyT)
	hardconvtype_property.SetValue(p_convtype)
	gd.PropertyTable[_HardConvType] = hardconvtype_property
}

func (gd *GeneralData) SetModifyDate(p_date string){
	if p_date == "" {
		p_date = common.GenDate()
	}
	gd.ModifyDate = p_date
	modifydate_property := gd.PropertyTable[_ModifyDate].(property.PropertyT)
	modifydate_property.SetValue(p_date)
	gd.PropertyTable[_ModifyDate] = modifydate_property
}

func (gd *GeneralData) SetParent(p_Parent inf.IData){
	gd.Parent = p_Parent
	parent_property := gd.PropertyTable[_Parent].(property.PropertyT)
	parent_property.SetValue(p_Parent)
	gd.PropertyTable[_Parent] = parent_property
}

func (gd *GeneralData) SetMandatory(p_Mandatory bool){
	gd.Mandatory = p_Mandatory
	mandatory_property := gd.PropertyTable[_Mandatory].(property.PropertyT)
	mandatory_property.SetValue(p_Mandatory)
	gd.PropertyTable[_Mandatory] = p_Mandatory
}