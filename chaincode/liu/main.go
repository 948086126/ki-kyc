package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//定义常量
const (
	Txid string = "txid-"
)

//定义资产链码结构体
type AssetChaincode struct {
}
type ConstrcChaincode struct {
}

//定义资产结构体
type Asset struct {
	AssetId      string `json:"Asset_id"`
	CargoName    string `json:"cargo_name"`
	CargoPrice   string `json:"cargo_price"`
	CargoAmount  string `json:"cargo_amount"`
	ContractId   string `json:"contract_id"`
	ProviderId   string `json:"provider_id"`
	ProviderName string `json:"provider_name"`
	CreateTime   string `json:"create_time"`
	CreateUser   string `json:"create_user"`
	UserDetail   string `json:"user_detail"`
	CargoAddress string `json:"cargo_address"`
}
type Contract struct {
	ContractId    string `json:"contract_id"`
	PartId        string `json:"part_id"`
	ContractName  string `json:"contract_name"`
	contractType  string `json:"contract_type"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	contractPrice string `json:"contract_price"`
	AssetId       string `json:"asset_id"`
}

func (a *AssetChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (a *AssetChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	if function == "add" {
		AssetId := args[0]
		CargoName := args[1]
		CargoPrice := args[2]
		CargoAmount := args[3]
		ContractId := args[4]
		ProviderId := args[5]
		ProviderName := args[6]
		CreateTime := args[7]
		CreateUser := args[8]
		UserDetail := args[9]
		CargoAddress := args[10]

		_a := &Asset{AssetId,
			CargoName,
			CargoPrice,
			CargoAmount,
			ContractId,
			ProviderId,
			ProviderName,
			CreateTime,
			CreateUser,
			UserDetail,
			CargoAddress,
		}
		marshal, err := json.Marshal(_a)
		if err == nil {
			err := stub.PutState(AssetId, marshal)
			if err != nil {
				return shim.Error(err.Error())
			}
		} else {
			return shim.Error("xushiliehuash")
		}
	} else if function == "del" {
		_assetId := args[0]
		if _assetId != "" {
			err := stub.DelState(_assetId)
			if err != nil {
				return shim.Error(err.Error())
			} else {
				return shim.Success(nil)
			}
		} else {
			return shim.Error("主键为nil")
		}
	} else if function == "update" {
		AssetId := args[0]
		CargoName := args[1]
		CargoPrice := args[2]
		CargoAmount := args[3]
		ContractId := args[4]
		ProviderId := args[5]
		ProviderName := args[6]
		CreateTime := args[7]
		CreateUser := args[8]
		UserDetail := args[9]
		CargoAddress := args[10]

		_a := &Asset{AssetId,
			CargoName,
			CargoPrice,
			CargoAmount,
			ContractId,
			ProviderId,
			ProviderName,
			CreateTime,
			CreateUser,
			UserDetail,
			CargoAddress,
		}
		marshal, err := json.Marshal(_a)
		if err == nil {
			err := stub.PutState(AssetId, marshal)
			if err != nil {
				return shim.Error(err.Error())
			}
		} else {
			return shim.Error("xushiliehuash")
		}
	} else {
		if function == "query" {
			_assetId := args[0]
			if _assetId != "" {
				state, err := stub.GetState(_assetId)
				if err != nil {
					return shim.Error(err.Error())
				} else {
					if err != nil {
						return shim.Error(err.Error())
					} else {
						return shim.Success(state)
					}
				}
			}
		} else {
		}
	}
	return shim.Error("方法异常")
}

func main() {

	err := shim.Start(new(AssetChaincode))
	if err == nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}

}
