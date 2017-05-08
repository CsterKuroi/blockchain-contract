package rethinkdb

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/astaxie/beego/logs"
	r "gopkg.in/gorethink/gorethink.v3"

	"unicontract/src/common"
	"unicontract/src/core/model"
)

func Get(db string, name string, id string) *r.Cursor {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Run(session)
	if err != nil {
		logs.Error(err.Error())
	}
	return res
}

func Insert(db string, name string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Insert(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		logs.Error(err.Error())
	}
	return res
}

func Update(db string, name string, id string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Update(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		logs.Error(err.Error())
	}
	return res
}

func Delete(db string, name string, id string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(name).Get(id).Delete().RunWrite(session)
	if err != nil {
		logs.Error(err.Error())
	}
	return res
}

/*----------------------------unicontract ops-------------------------------------*/

/*----------------------------- contracts start---------------------------------------*/
// contract serialize contract string
func InsertContract(contract string) bool {
	if contract == "" {
		return false
	}
	res := Insert(DBNAME, TABLE_CONTRACTS, contract)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetContractById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONTRACTS, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 根据合约[id]获取具有相同contractId的合约
func GetContractsByContractId(contractId string) (string, error) {
	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Filter(r.Row.Field("ContractBody").Field("ContractId").Eq(contractId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
	}

	return common.Serialize(blo), nil
}

//根据 contract.id 获取合约处理主节点
func GetContractMainPubkeyByContract(id string) (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACTS).Get(id).Count().Default(0).Run(session)
	if err != nil {
		return "", err
	}

	var blo int
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	if blo == 0 {
		return "", nil
	}

	// continue ...
	res, err = r.Table(TABLE_CONTRACTS).Get(id).Field("ContractHead").Field("MainPubkey").Run(session)
	if err != nil {
		return "", err
	}

	var blo2 string
	err = res.One(&blo2)
	if err != nil {
		return "", err
	}
	return blo2, nil
}

/*----------------------------- contracts end---------------------------------------*/

/*----------------------------- votes start---------------------------------------*/
// vote serialize vote string
func InsertVote(vote string) bool {
	if vote == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_VOTES, vote)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetVoteById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}
	res := Get(DBNAME, TABLE_VOTES, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 根据合约[id]获取所有 vote
func GetVotesByContractId(contractId string) (string, error) {

	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_VOTES).Filter(r.Row.Field("Vote").Field("VoteFor").Eq(contractId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

/*----------------------------- votes end---------------------------------------*/

/*----------------------------- contractOutputs start---------------------------------------*/
// contractOutput serialize contractOutput string
func InsertContractOutput(contractOutput string) bool {
	if contractOutput == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONTRACT_OUTPUTS, contractOutput)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetContractOutputById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONTRACT_OUTPUTS, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

// 根据合约 整体的[id] transaction.contract.id 或者 relation.contractId 获取所有 contractOutput
func GetContractOutputByContractPrimaryId(id string) (string, error) {

	if id == "" {
		return "", errors.New("id blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_OUTPUTS).Filter(r.Row.Field("transaction").Field("Contract").Field("id").Eq(id)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err = res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

/*----------------------------- contractOutputs end---------------------------------------*/

/*----------------------------- contractTask start---------------------------------------*/
// contractTask serialize contractTask string
func InsertContractTask(contractTask string) bool {
	if contractTask == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONTRACT_TASKS, contractTask)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetContractTaskById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONTRACT_TASKS, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

func GetContractTasksByContractId(contractId string) (string, error) {
	if contractId == "" {
		return "", errors.New("contractId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONTRACT_TASKS).Filter(r.Row.Field("ContractId").Eq(contractId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	return common.Serialize(blo), nil
}

/*----------------------------- contractTask end---------------------------------------*/

/*----------------------------- consensusFailures start---------------------------------------*/
// consensusFailures serialize consensusFailures string
func InsertConsensusFailure(consensusFailures string) bool {
	if consensusFailures == "" {
		return false
	}

	res := Insert(DBNAME, TABLE_CONSENSUS_FAILURES, consensusFailures)
	if res.Inserted >= 1 {
		return true
	}
	return false
}

func GetConsensusFailureById(id string) (string, error) {
	if id == "" {
		return "", errors.New("id blank")
	}

	res := Get(DBNAME, TABLE_CONSENSUS_FAILURES, id)
	if res.IsNil() {
		return "", nil
	}

	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		return "", err
	}
	return common.Serialize(blo), nil
}

func GetConsensusFailuresByConsensusId(consensusId string) (string, error) {
	if consensusId == "" {
		return "", errors.New("consensusId blank")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_CONSENSUS_FAILURES).Filter(r.Row.Field("ConsensusId").Eq(consensusId)).Run(session)
	if err != nil {
		return "", err
	}
	if res.IsNil() {
		return "", nil
	}

	var blo []map[string]interface{}
	err = res.All(&blo)
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	return common.Serialize(blo), nil
}

/*----------------------------- consensusFailures end---------------------------------------*/

/*----------------------------- SendFailingRecords start---------------------------------------*/
func GetAllRecords(db string, name string) ([]string, error) {
	session := ConnectDB(db)
	ids, err := r.Table(name).Field("id").Run(session)
	if err != nil {
		logs.Error(err.Error())
	}
	var idlist []string
	err = ids.All(&idlist)
	if err != nil {
		logs.Error(err.Error())
		return nil, errors.New(err.Error())
	}
	return idlist, nil
}

/*----------------------------- SendFailingRecords end---------------------------------------*/

/*TaskSchedule start-------------------------------------------------------*/
// 插入一个nodepublickey的task方法
func InsertTaskSchedule(strTaskSchedule string) error {
	res := Insert(DBNAME, TABLE_TASK_SCHEDULE, strTaskSchedule)
	if res.Inserted >= 1 {
		return nil
	} else {
		return fmt.Errorf("insert failed")
	}
}

//---------------------------------------------------------------------------
// 根据nodePubkey和contractID获得表内ID
func _GetID(strNodePubkey, strContractID string) (string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("ContractId").Eq(strContractID)).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Run(session)
	if err != nil {
		return "", err
	}

	if res.IsNil() {
		return "", nil
	}

	var tasks map[string]interface{}
	err = res.One(&tasks)
	if err != nil {
		return "", err
	}

	return tasks["id"].(string), nil
}

//---------------------------------------------------------------------------
// 根据ID获取starttime和endtime
func _GetValidTime(strID string) (string, string, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("id").Eq(strID)).
		Run(session)
	if err != nil {
		return "", "", err
	}

	if res.IsNil() {
		return "", "", nil
	}

	var tasks map[string]interface{}
	err = res.One(&tasks)
	if err != nil {
		return "", "", err
	}

	return tasks["StartTime"].(string), tasks["EndTime"].(string), nil
}

//---------------------------------------------------------------------------
// 设置SendFlag字段，发送为1,未发送为0
func _SetTaskScheduleFlag(strID string, alreadySend bool) error {
	var send int
	if alreadySend {
		send = 1
	} else {
		send = 0
	}
	strJSON := fmt.Sprintf("{\"SendFlag\":%d,\"LastExecuteTime\":\"%s\"}",
		send, common.GenTimestamp())

	res := Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res.Replaced|res.Unchanged >= 1 {
		return nil
	} else {
		return fmt.Errorf("update failed")
	}
}

//---------------------------------------------------------------------------
// 设置FailedCount(或者SuccessCount)字段加一
func _SetTaskScheduleCount(strID string, success bool) (int, error) {
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("id").Eq(strID)).Run(session)
	if err != nil {
		return -1, err
	}

	if res.IsNil() {
		return -1, fmt.Errorf("set is null")
	}

	var tasks map[string]interface{}
	err = res.One(&tasks)
	if err != nil {
		return -1, err
	}

	failedCount := tasks["FailedCount"].(float64)
	failedCount += 1

	successCount := tasks["SuccessCount"].(float64)
	successCount += 1

	var strJSON string
	if success {
		strJSON = fmt.Sprintf("{\"SuccessCount\":%f,\"LastExecuteTime\":\"%s\"}",
			successCount, common.GenTimestamp())
	} else {
		strJSON = fmt.Sprintf("{\"FailedCount\":%f,\"LastExecuteTime\":\"%s\"}",
			failedCount, common.GenTimestamp())
	}

	res1 := Update(DBNAME, TABLE_TASK_SCHEDULE, strID, strJSON)
	if res1.Replaced|res1.Unchanged >= 1 {
		if success {
			return int(successCount), nil
		} else {
			return int(failedCount), nil
		}

	} else {
		return -1, fmt.Errorf("update failed")
	}
}

//---------------------------------------------------------------------------
// 设置发送标志为“已发送”，在将任务插入待执行队列后调用
func UpdateMonitorSend(strID string) error {
	if len(strID) == 0 {
		return fmt.Errorf("id is null")
	}
	return _SetTaskScheduleFlag(strID, true)
}

//---------------------------------------------------------------------------
// 执行失败：1.更新strContractID 的SendFlag = 0, FailedCount + 1, LastExecuteTime
// 返回FailedCount(或者SuccessCount)和error
func UpdateMonitorFail(strNodePubkey, strContractID string) (int, error) {
	if len(strNodePubkey) == 0 || len(strContractID) == 0 {
		return -1, fmt.Errorf("pubkey or contractid is null")
	}

	strID, err := _GetID(strNodePubkey, strContractID)
	if err != nil {
		return -1, err
	}

	if len(strID) == 0 {
		return -1, fmt.Errorf("not find")
	}

	err = _SetTaskScheduleFlag(strID, false)
	if err != nil {
		return -1, err
	}

	return _SetTaskScheduleCount(strID, false)
}

//---------------------------------------------------------------------------
// 执行条件不满足：1.更新strNodePubkey  的SendFlag = 0, LastExecuteTime
func UpdateMonitorWait(strNodePubkey, strContractID string) error {
	if len(strNodePubkey) == 0 || len(strContractID) == 0 {
		return fmt.Errorf("pubkey or contractid is null")
	}

	strID, err := _GetID(strNodePubkey, strContractID)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("not find")
	}

	return _SetTaskScheduleFlag(strID, false)
}

//---------------------------------------------------------------------------
// 执行成功：1.更新strContractIDold 的SendFlag=1, SuccessCount + 1, LastExecuteTime
//         2.将strContractIDnew 插入到扫描监控表中
func UpdateMonitorSucc(strNodePubkey, strContractIDold, strContractIDnew string) error {
	if len(strNodePubkey) == 0 ||
		len(strContractIDold) == 0 ||
		len(strContractIDnew) == 0 {
		return fmt.Errorf("pubkey or contractid is null")
	}

	strID, err := _GetID(strNodePubkey, strContractIDold)
	if err != nil {
		return err
	}

	if len(strID) == 0 {
		return fmt.Errorf("old contract id not find")
	}

	err = _SetTaskScheduleFlag(strID, true)
	if err != nil {
		return err
	}

	_, err = _SetTaskScheduleCount(strID, true)
	if err != nil {
		return err
	}

	startTime, endTime, err := _GetValidTime(strID)
	if err != nil {
		return err
	}

	if len(startTime) == 0 || len(endTime) == 0 {
		return fmt.Errorf("old contract valid time not find")
	}

	var taskSchedule model.TaskSchedule
	taskSchedule.Id = common.GenerateUUID()
	taskSchedule.ContractId = strContractIDnew
	taskSchedule.NodePubkey = strNodePubkey
	taskSchedule.StartTime = startTime
	taskSchedule.EndTime = endTime
	slJson, _ := json.Marshal(taskSchedule)
	return InsertTaskSchedule(string(slJson))
}

//---------------------------------------------------------------------------
// 获取所有未发送的任务，用于放在待执行队列中
func GetTaskSchedulesNoSend(strNodePubkey string) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("pubkey is null")
	}

	now := common.GenTimestamp()
	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Filter(r.Row.Field("StartTime").Le(now)).
		Filter(r.Row.Field("EndTime").Ge(now)).
		Filter(r.Row.Field("SendFlag").Eq(0)).
		Run(session)
	if err != nil {
		return "", err
	}

	if res.IsNil() {
		return "", nil
	}

	var tasks []map[string]interface{}
	err = res.All(&tasks)
	if err != nil {
		return "", err
	}
	return common.Serialize(tasks), nil
}

//---------------------------------------------------------------------------
// 获取已经执行成功后的任务，用于清理数据
func GetTaskSchedulesSuccess(strNodePubkey string) (string, error) {
	if len(strNodePubkey) == 0 {
		return "", fmt.Errorf("pubkey is null")
	}

	session := ConnectDB(DBNAME)
	res, err := r.Table(TABLE_TASK_SCHEDULE).
		Filter(r.Row.Field("SuccessCount").Ge(1)).
		Filter(r.Row.Field("NodePubkey").Eq(strNodePubkey)).
		Run(session)
	if err != nil {
		return "", err
	}

	if res.IsNil() {
		return "", nil
	}

	var tasks []map[string]interface{}
	err = res.All(&tasks)
	if err != nil {
		return "", err
	}
	return common.Serialize(tasks), nil
}

//---------------------------------------------------------------------------
// 删除一系列id的任务
func DeleteTaskSchedules(slID []string) (int, []error) {
	var nDeleteNum int
	var slerr []error
	for _, value := range slID {
		res := Delete(DBNAME, TABLE_TASK_SCHEDULE, value)
		if res.Deleted >= 1 {
			nDeleteNum += res.Deleted
		} else {
			slerr = append(slerr, fmt.Errorf(value))
		}
	}
	return nDeleteNum, slerr
}

/*TaskSchedule end---------------------------------------------------------*/
