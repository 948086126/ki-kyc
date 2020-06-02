package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const (
	Txid string = "txid-"
)

type GoodsChaincode struct {
}

type Goods struct {
	Uid          string `json:"uid"`
	Goods_Code   string `json:"goods_code"`
	Goods_Name   string `json:"goods_name"`
	Goods_type   string `json:"goods_type"`
	Goods_Price  string `json:"goods_price"`
	Goods_Status string `json:"goods_status"`
	Goods_Remark string `json:"goods_remark"`
	Create_Time  string `json:"create_time"`
}

//初始化
func (g *GoodsChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

//执行琏码
func (g *GoodsChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "upInfoToBlock" {
		return upInfoToBlock(stub, args)
	} else if function == "delInfoFromBlok" {

	}
	return shim.Error("please check request")
}

func upInfoToBlock(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	_Txid := args[0]
	_Goods_Code := args[1]
	_Goods_Name := args[2]
	_Goods_type := args[3]
	_Goods_Price := args[4]
	_Goods_Status := args[5]
	_Goods_Remark := args[6]
	_Create_Time := args[7]

	goods := &Goods{Uid: _Txid, Goods_Code: _Goods_Code, Goods_Name: _Goods_Name, Goods_type: _Goods_type, Goods_Price: _Goods_Price, Goods_Status: _Goods_Status, Goods_Remark: _Goods_Remark, Create_Time: _Create_Time}
	//接收参数,序列化json
	requestInfo, err2 := json.Marshal(goods)
	//判断是否存在该id数据
	state, err := stub.GetState(_Txid)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		if state != nil {
			return shim.Error("had same code ,please edit")
		} else {
			if err2 != nil {
				return shim.Error(err.Error())
			} else {
				stub.PutState(_Txid, requestInfo)
				return shim.Success(nil)
			}
		}
	}
	return shim.Error("add error")
}
