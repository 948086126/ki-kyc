package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// 资产更新
func updateAssetInternal(stub shim.ChaincodeStubInterface, args []string, asset interface{}) (string, error) {
	var key string
	var err error
	if len(args) != 3 {
		return key, errors.New("Incorrect number of arguments. Expecting 3")
	}
	assetJsonStr := args[1]
	//updateTime := args[2]
	logger.Debug("获取到asset:", asset)
	//１ 反序列化传递进来的的对象
	if err := json.Unmarshal([]byte(assetJsonStr), asset); err != nil {
		return key, errors.New("第一次反序列话的时候 出错," + err.Error())
	}
	//2 获取上链key
	key, err = getPkey(asset)
	if err != nil {
		return key, err
	}
	//3 检查是否上链
	assetAsBytes, err := stub.GetState(KYC + key)
	if err != nil {
		return key, err
	}
	if assetAsBytes == nil {
		return key, errors.New("The asset  is　not existed! the key :" + key)
	}
	//4 反序列链上对象
	if err := json.Unmarshal(assetAsBytes, asset); err != nil {
		return key, errors.New("4 反序列链上对象 出错," + err.Error())
	}
	// 9　上链
	assetAsBytes, err = json.Marshal(asset)
	if err != nil {
		return key, err
	}
	if err = stub.PutState(KYC+key, assetAsBytes); err != nil {
		return key, err
	}
	return key, nil
}

// 资产上链
func uploadAssetInternal(stub shim.ChaincodeStubInterface, args []string, assetStruct interface{}) (string, error) {
	var key string
	var err error
	if len(args) != 3 {
		return key, errors.New("Incorrect number of arguments. Expecting 3")
	}
	assetJsonStr := args[1]
	//updateTime := args[2]
	logger.Debug("开始步骤")
	//１ 反序列化
	if err := json.Unmarshal([]byte(assetJsonStr), assetStruct); err != nil {
		return key, errors.New("//１ 反序列化:" + err.Error())
	}
	//2 获取上链key
	key, err = getPkey(assetStruct)
	if err != nil {
		return key, errors.New("//2 获取上链:" + err.Error())
	}
	//3 检查是否上链
	existed, err := checkIsExisted(stub, KYC+key)
	if err != nil {
		return key, errors.New("//3 检查是否上链:" + err.Error())
	}
	if existed {
		return key, fmt.Errorf("The asset  is existed! the key : %s", key)
	}

	// 5　上链
	assetAsBytes, err := json.Marshal(assetStruct)
	if err != nil {
		return key, errors.New("// 5　上链:" + err.Error())
	}
	if err = stub.PutState(KYC+key, assetAsBytes); err != nil {
		return key, err
	}
	return key, nil
}

// 资产关联上链
func uploadAssetInternalRel(stub shim.ChaincodeStubInterface, args []string, assetStruct interface{}) (string, error) {
	var key string
	var err error
	if len(args) != 3 {
		return key, errors.New("Incorrect number of arguments. Expecting 3")
	}
	assetJsonStr := args[2]
	logger.Debug("开始步骤")
	//１ 反序列化
	if err := json.Unmarshal([]byte(assetJsonStr), assetStruct); err != nil {
		return key, errors.New("//3-1 反序列化:" + err.Error())
	}
	//2 获取上链key
	key, err = getPkey(assetStruct)
	if err != nil {
		return key, errors.New("//3-2 获取上链:" + err.Error())
	}
	//3 检查是否上链
	existed, err := checkIsExisted(stub, BYC+key)
	if err != nil {
		return key, errors.New("//3-3 检查是否上链:" + err.Error())
	}

	if existed {

	}

	// 标签上链
	assetAsBytes, err := json.Marshal(assetStruct)
	if err != nil {
		return key, errors.New("// 3-5　上链:" + err.Error())
	}
	if err = stub.PutState(BYC+key, assetAsBytes); err != nil {
		return key, err
	}

	// 关联关系处理 //6 添加数据关联关系 ||访问关系  ||对方  父子调换
	if err = addRoll(stub, assetJsonStr); err != nil {
		return key, errors.New("// 7 添加数据关联关系:" + err.Error())
	}
	return key, nil
}

// 资产关联上链
func uploadAssetInternalMaster(stub shim.ChaincodeStubInterface, args []string, assetStruct interface{}) (string, error) {
	var key string
	var err error
	if len(args) != 3 {
		return key, errors.New("Incorrect number of arguments. Expecting 3")
	}
	assetJsonStr := args[2]
	logger.Debug("开始步骤")
	//１ 反序列化
	if err := json.Unmarshal([]byte(assetJsonStr), assetStruct); err != nil {
		return key, errors.New("//3-1 反序列化:" + err.Error())
	}
	//2 获取上链key
	key, err = getPkey(assetStruct)
	if err != nil {
		return key, errors.New("//3-2 获取上链:" + err.Error())
	}
	//3 检查是否上链
	existed, err := checkIsExisted(stub, BYC+key)
	if err != nil {
		return key, errors.New("//3-3 检查是否上链:" + err.Error())
	}

	if existed {

	}

	// 标签上链
	assetAsBytes, err := json.Marshal(assetStruct)
	if err != nil {
		return key, errors.New("// 3-5　上链:" + err.Error())
	}
	if err = stub.PutState(BYC+key, assetAsBytes); err != nil {
		return key, err
	}
	return key, nil
}

//	更新关联资产
func updateAssociation(stub shim.ChaincodeStubInterface, args []string) (err error) {

	data := Byc{}
	// 数据序列化获取到
	if len(args) != 3 {
		return errors.New("Incorrect number of arguments. Expecting 3")
	}

	assetJsonStr := args[2]
	logger.Debug("开始步骤")
	//１ 反序列化
	err = json.Unmarshal([]byte(assetJsonStr), &data)

	if err != nil {
		return errors.New("//3-1 反序列化:" + err.Error())
	}
	// 校验参数
	if data.BycByKey == "" || data.BycByfMaster == "" {
		data.BycByKey = "default"
		data.BycByfMaster = "default"
	}

	// 数据本身查询关联
	Pg, err1 := json.Marshal(MasterInfo{Id: data.BycByKey, Master: data.BycByfMaster, TxId: stub.GetTxID(), Span: data.BycBySpan})

	if err1 != nil {
		return errors.New("//3-3 数据本身查询关联序列化")
	}
	//	外键查询关联
	Hg, err2 := json.Marshal(MasterInfo{Id: data.BycByfMaster, Master: data.BycByKey, TxId: stub.GetTxID(), Span: data.BycBySpan})
	if err2 != nil {
		return errors.New("//3-4 外键查询关联关联序列化")
	}
	//	存 k-v
	err = stub.PutState(Ain+data.BycByKey, Pg)
	if err != nil {
		return errors.New("//3-5 存 k-v")
	}

	err = stub.PutState(Mas+data.BycByfMaster, Hg)

	if err != nil {
		return errors.New("//3-6 存 v-k")
	}

	return nil

}

//
// 资产关联更新上链
func updateAssetInternalRel(stub shim.ChaincodeStubInterface, args []string, assetStruct interface{}) (string, error) {
	var key string
	var err error
	if len(args) != 3 {
		return key, errors.New("Incorrect number of arguments. Expecting 3")
	}
	assetJsonStr := args[2]
	logger.Debug("开始步骤")
	//１ 反序列化
	if err := json.Unmarshal([]byte(assetJsonStr), assetStruct); err != nil {
		return key, errors.New("//１ 反序列化:" + err.Error())
	}
	//2 获取上链key
	key, err = getPkey(assetStruct)
	if err != nil {
		return key, errors.New("//2 获取上链:" + err.Error())
	}
	//3 检查是否上链
	existed, err := checkIsExisted(stub, BYC+key)
	if err != nil {
		return key, errors.New("//3 检查是否上链:" + err.Error())
	}

	// 如果是第一次存,那就存进去标签，如果不是，那就更新 数据关系和 访问关系
	if !existed {
		// 5　上链
		assetAsBytes, err := json.Marshal(assetStruct)
		if err != nil {
			return key, errors.New("// 5　上链:" + err.Error())
		}
		if err = stub.PutState(BYC+key, assetAsBytes); err != nil {
			return key, err
		}
	}

	//6 添加数据关联关系 ||访问关系
	if err = addRelationShipForSubAssetData(stub, assetJsonStr, BYC+key); err != nil {
		return key, errors.New("// 6 添加数据关联关系:" + err.Error())
	}

	//6 添加数据关联关系 ||访问关系  ||对方  父子调换
	if err = addRelationShipForSubAssetDataRoll(stub, assetJsonStr); err != nil {
		return key, errors.New("// 6 添加数据关联关系:" + err.Error())
	}
	return key, nil
}

//　检查是否上链了
func checkIsExisted(stub shim.ChaincodeStubInterface, key string) (bool, error) {
	// 校验key
	if key == "" {
		return false, errors.New("the pkey‘s value  is empty")
	}
	// 判断链上是否存在
	assetAsBytes, err := stub.GetState(key)
	if err != nil {
		return false, err
	}
	if assetAsBytes != nil {
		return true, nil
	}
	return false, nil
}

// 为结构体添加txID
func addUpdateTxID(updateTxID string, asset interface{}) error {
	v := reflect.ValueOf(asset).Elem()
	v.FieldByName("HisCurrentTx").SetString(updateTxID)
	return nil
}

// 为结构体添加更新时间
func addParentInfo(parentID, parentType string, asset interface{}) error {
	v := reflect.ValueOf(asset).Elem()
	v.FieldByName("ParentID").SetString(parentID)
	//v.FieldByName("ParentID")
	v.FieldByName("ParentType").SetString(parentType)
	return nil
}

// 获取上链对象的key
func getPkey(asset interface{}) (string, error) {
	t := reflect.TypeOf(asset).Elem()
	v := reflect.ValueOf(asset).Elem()
	fieldCount := t.NumField()
	for index := 0; index < fieldCount; index++ {
		// 有这个标识的　字段　对应的值
		_, hasKey := t.Field(index).Tag.Lookup("pkey")
		if hasKey {
			return v.FieldByName(t.Field(index).Name).String(), nil
		}
	}
	return "", errors.New("The asset has no field  about primary key!")
}

// 根据jsonObj 获取　更新时间
func getUpdateTime(str string) string {
	// logger.Debug("getUpdateTime 入参   ： ", str)
	resultStr := ""
	ss := strings.Split(str, ",")
	for _, value := range ss {
		if strings.Contains(value, "kycTime") {
			fieldValue := strings.Split(value, ":")[1]
			resultStr = strings.Split(fieldValue, "\"")[1]
			break
		}
	}
	// logger.Debug("getUpdateTime 结果   ： ", resultStr)
	return resultStr
}

// 更新数据关联关系
func addRelationShipForSubAssetData(stub shim.ChaincodeStubInterface, assetStruct string, key string) error {
	// 01	序列化
	currentData := Byc{}
	metadata := Byc{}

	//	02	获取数据
	// 判断链上是否存在
	assetAsBytes, err := stub.GetState(key)
	if err != nil {
		return err
	}
	if assetAsBytes != nil {
		return err
	}
	//	03	获取元数据
	json.Unmarshal(assetAsBytes, &metadata)
	//	04	获取本数据
	json.Unmarshal([]byte(assetStruct), &currentData)
	//	05	数据插入
	insertData(&metadata, &currentData)

	insertSign(&metadata, &currentData)
	//
	MarData, err := json.Marshal(&metadata)
	if err != nil {
		return err
	}
	err = stub.PutState(BYC+key, MarData)

	if err != nil {
		return err
	}
	//	05	返回

	return nil
}

// 更新数据关联关系
func addShip(stub shim.ChaincodeStubInterface, assetStruct string, key string) error {
	// 01	序列化
	currentData := Byc{} //  当前数据
	metadata := Byc{}    //	 原来标签

	//	02	获取数据
	// 判断链上是否存在
	assetAsBytes, err := stub.GetState(key)
	if err != nil {
		return err

	}
	if assetAsBytes == nil {
		return errors.New("关联数据|| asset is null")
	}

	//	03	获取元数据
	json.Unmarshal(assetAsBytes, &metadata)
	//	04	获取本数据
	json.Unmarshal([]byte(assetStruct), &currentData)
	//	05	数据插入

	//	06 追加关联ID
	for k, _ := range currentData.BycMes {
		if currentData.BycMes[k].BycID != "" {
			metadata.BycMes = append(metadata.BycMes, currentData.BycMes[k])
		}
	}
	fmt.Println("-34")
	//	07	 追加关联Key
	for k, _ := range currentData.BycSignList {
		if currentData.BycSignList[k].RelID != "" {
			metadata.BycSignList = append(metadata.BycSignList, currentData.BycSignList[k])
		}
	}
	//
	fmt.Println("-35")
	MarData, err := json.Marshal(&metadata)
	if err != nil {
		return err
	}
	err = stub.PutState(key, MarData)
	fmt.Println("-36")
	if err != nil {
		return err
	}
	//	05	返回

	return nil
}

// 更新对方 关联关系
func addRelationShipForSubAssetDataRoll(stub shim.ChaincodeStubInterface, assetStruct string) error {
	// 01	序列化
	currentData := Byc{}
	metadata := Byc{}

	//	04	获取本数据
	json.Unmarshal([]byte(assetStruct), &currentData)

	for k, _ := range currentData.BycMes {
		//获取关联数据标签表
		resultdata, err := stub.GetState(BYC + currentData.BycMes[k].ParBycID)
		if err != nil {
			// 拼接 关系数据  父子对调
			cdata := BycMes{
				BycID:       currentData.BycMes[k].ParBycID,
				BycKey:      currentData.BycMes[k].ParBycKey,
				ParBycID:    currentData.BycMes[k].BycID,
				ParBycKey:   currentData.BycMes[k].BycKey,
				BycRelation: exchange(currentData.BycMes[k].BycRelation),
				BySol:       currentData.BycMes[k].BySol,
			}
			// 反序列化	关联数据标签表
			json.Unmarshal(resultdata, &metadata)
			// 数据追加 数据关联关系
			metadata.BycMes = append(metadata.BycMes, cdata)
			//保存
			MarData, err := json.Marshal(&metadata)
			if err != nil {
				return err
			}
			err = stub.PutState(BYC+currentData.BycMes[k].ParBycID, MarData)
			//
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 更新对方 关联关系 test
func addRoll(stub shim.ChaincodeStubInterface, assetStruct string) error {
	// 01	序列化
	currentData := Byc{} // 当前数据
	metadata := Byc{}    //  元数据
	//	04	获取本数据
	json.Unmarshal([]byte(assetStruct), &currentData)

	for k, _ := range currentData.BycMes {
		//获取关联数据标签表
		resultdata, err := stub.GetState(BYC + currentData.BycMes[k].ParBycID)
		if err == nil {
			// 拼接 关系数据  父子对调
			cdata := BycMes{
				BycID:       currentData.BycMes[k].ParBycID,
				BycKey:      currentData.BycMes[k].ParBycKey,
				ParBycID:    currentData.BycMes[k].BycID,
				ParBycKey:   currentData.BycMes[k].BycKey,
				BycRelation: exchange(currentData.BycMes[k].BycRelation),
				BySol:       currentData.BycMes[k].BySol,
			}
			//
			// 反序列化	关联数据标签表
			json.Unmarshal(resultdata, &metadata)
			// 数据追加 数据关联关系
			metadata.BycMes = append(metadata.BycMes, cdata)
			//保存
			MarData, err := json.Marshal(metadata)
			if err != nil {
				return err
			}
			err = stub.PutState(BYC+currentData.BycMes[k].ParBycID, MarData)
			//
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

// 循环插入
func insertData(metadata *Byc, currentData *Byc) {
	//
	for k, _ := range currentData.BycMes {
		if currentData.BycMes[k].BycID != "" {
			metadata.BycMes = append(metadata.BycMes, currentData.BycMes[k])
		}
	}
}

// 循环插入
func insertSign(metadata *Byc, currentData *Byc) {
	//
	for k, _ := range currentData.BycSignList {
		if currentData.BycSignList[k].RelID != "" {
			metadata.BycSignList = append(metadata.BycSignList, currentData.BycSignList[k])
		}
	}
}

// 父子关系 对调
func exchange(typeName string) (rollName string) {

	switch typeName {
	case Title_Parent:
		return Title_Sun
	case Title_Sun:
		return Title_Parent
	default:
		return Title_Tong
	}

}

// 操作资产前的　参数检查 3
func checkArgs(args []string) error {
	return checkArgsForCount(args, 3)
}

// 操作资产前的　参数检查 2
func checkArgsByOne(args []string) error {
	return checkArgsForCount(args, 1)
}

//	校验||参数检查 4
func checkArgsRec(args []string) error {
	return checkArgsForCount(args, 4)
}

//	校验||参数检查 4
func checkArgsForCount(args []string, count int) error {
	if len(args) != count {
		return fmt.Errorf("Incorrect number of arguments. Expecting :  %v", count)
	}
	// 验空
	for index := 0; index < count; index++ {
		if len(args[index]) <= 0 {
			return fmt.Errorf("index :%v  argument must be a non-empty string", index)
		}
	}
	return nil
}

//	结构体||根据类型返回对应结构体
func getAssetStructByType(assetType string) (interface{}, error) {
	var assetStruct interface{}
	switch assetType {
	case KYC:
		assetStruct = &Kyc{}
	case BYC:
		assetStruct = &Byc{}
	default:
		return nil, errors.New("The assetType is not supported:" + assetType)
	}

	return assetStruct, nil
}

// 获取 Kyc
func getKyc() *Kyc {
	data := Kyc{}
	return &data
}

// 获取 []Kyc
func getKycList() *[]Kyc {
	data := []Kyc{}
	return &data
}

// 获取 Byc
func getByc() *Byc {
	data := Byc{}
	return &data
}

// 获取 []Byc
func getBycList() *[]Byc {
	data := []Byc{}
	return &data
}

// 获取相关数据
func getAssetSunList(stub shim.ChaincodeStubInterface, List []BycMes) (Kyc *[]Kyc, state bool) {
	//
	KycList := getKycList()

	if len(List) > 0 {
		for k, _ := range List {
			if List[k].ParBycID != "" {
				KycData := getKyc()
				//查询关联 元数据   kyc
				assetAsBytes, err := stub.GetState(KYC + List[k].ParBycID)
				if err != nil {
					return nil, false
				}
				if assetAsBytes == nil {
					return nil, false
				}
				// 序列化
				json.Unmarshal(assetAsBytes, &KycData)
				*KycList = append(*KycList, *KycData)
			}
		}
	}
	return KycList, true
}

// -- 授权
func SignByData(stub shim.ChaincodeStubInterface, args []string) (err error) {

	// 参数检查
	if err := checkArgsByOne(args); err != nil {
		return errors.New("参数只有一个")
	}
	assetStruct, err := getAssetStructByType(CYC)
	if err != nil {
		return err

	}
	// -- 校验调用者是否有写的权限
	return updateThePermissions(stub, args, assetStruct)
}

//  更新权限
func updateThePermissions(stub shim.ChaincodeStubInterface, args []string, assetStruct interface{}) error {
	// -- 校验数据
	// -- 获取数据
	// -- 追加数据
	// -- 完成状态
	var key string
	var bycData Byc
	var sycData Syc
	var err error
	if len(args) != 1 {
		return errors.New("Incorrect number of arguments. Expecting 3")
	}
	assetJsonStr := args[0]
	logger.Debug("开始步骤")
	//１ 反序列化
	if err := json.Unmarshal([]byte(assetJsonStr), &sycData); err != nil {
		return errors.New("//3-1 反序列化:" + err.Error())
	}
	//2 获取上链key
	key, err = getPkey(assetStruct)
	if err != nil {
		return errors.New("//3-2 获取上链:" + err.Error())
	}

	// 	// -- 校验数据
	result, err := stub.GetState(BYC + key)
	if err != nil {
		return err

	}
	//
	err = json.Unmarshal(result, &bycData)

	if err != nil {
		return err
	}

	//校验权限
	if sycData.SignID != bycData.BycSign.Self {
		//
		return err
	}
	// nil
	if sycData.SignID == "" || sycData.KycID == "" || sycData.ImpUserID == "" || sycData.ImpUserType == "" {
		//
		return errors.New(" label is nil")
	}

	bycData.BycSignList = append(bycData.BycSignList, SignList{
		RelID:   sycData.ImpUserID,
		RelType: sycData.ImpUserType,
		RelSol:  sycData.ImpUserRol,
		RelNum:  sycData.ImpUserNum,
		RelTime: sycData.ImpUserTime,
	})
	Pg, err := json.Marshal(bycData)
	if err != nil {
		return err
	}
	err = stub.PutState(BYC+key, Pg)
	if err != nil {
		return err
	}
	return nil
}
