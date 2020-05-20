package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const (
	TxId string = "txid-"
)

type BookList struct {
	TxId  string `json:"txid"`  //txid
	Value []byte `json:"value"` //value
}
type Book struct {
}

type Book struct {
	LoginId       string `json:"loginid"`
	UserId        string `json:"userid"`
	UserName      string `json:"username"`
	FullName      string `json:"fullname"`
	UserPhone     string `json:"userphone"`
	CompanyId     string `json:"companyid"`
	CompanyName   string `json:"companyname"`
	LoginState    string `json:"loginstate"`
	LoginType     string `json:"logintype"`
	LoginTime     string `json:"logintime"`
	LoginIp       string `json:"loginip"`
	DeviceType    string `json:"devicetype"`
	LoginAddress  string `json:"loginaddress"`
	LoginLocation string `json:"loginlocation"`
}

type BookInfo struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

func (t *BookInfo) error(msg string) {
	t.Status = false
	t.Msg = msg
}
func (t *BookInfo) ok(msg string) {
	t.Status = true
	t.Msg = msg
}

func (t *BookInfo) response() pb.Response {
	resJson, err := json.Marshal(t)
	if err != nil {
		return shim.Error("Failed to generate json result " + err.Error())
	}
	return shim.Success(resJson)
}

type process func(shim.ChaincodeStubInterface, []string) *BookInfo

func (t *Book) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *Book) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	if function == "invoke" {
		return t.invoke(stub, args)
	} else if function == "delete" {
		return t.delete(stub, args)
	} else if function == "update" {
		return t.update(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	} else if function == "updateTxid" {
		return t.updateTxid(stub, args)
	} else if function == "history" {
		return t.history(stub, args)
	}else if function =="test"{
		return t.test(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"update\" \"query\"")
}

func (t *Book) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return t.handleProcess(stub, args, 14, func(shim.ChaincodeStubInterface, []string) *BookInfo {
		ri := &BookInfo{true, ""}
		_id := args[0]
		_userid := args[1]
		_username := args[2]
		_fullname := args[3]
		_userphone := args[4]
		_companyid := args[5]
		_companyname := args[6]
		_loginstate := args[7]
		_logintype := args[8]
		_logintime := args[9]
		_loginip := args[10]
		_devicetype := args[11]
		_loginaddress := args[12]
		_loginlocation := args[13]

		_Book := &Book{_id, _userid, _username,
			_fullname, _userphone, _companyid, _companyname,
			_loginstate, _logintype, _logintime, _loginip,
			_devicetype, _loginaddress, _loginlocation}

		_ejson, err := json.Marshal(_Book)

		if err != nil {
			ri.error(err.Error())
		} else {
			_old, err := stub.GetState(_id)
			if err != nil {
				ri.error(err.Error())
			} else if _old != nil {
				ri.error("the Book has exists")
			} else {
				err := stub.PutState(_id, _ejson)
				if err != nil {
					ri.error(err.Error())
				} else {
					// 追加 根据 key 查询
					_tjson, err := json.Marshal(stub.GetTxID())
					if err != nil {
						ri.error(err.Error())
					} else {
						err := stub.PutState(TxId+_id, _tjson)
						if err != nil {
							ri.error(err.Error())
						} else {
							ri.ok(stub.GetTxID())
						}
					}
				}
			}
		}
		return ri
	})
}

func (t *Book) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return t.handleProcess(stub, args, 1, func(shim.ChaincodeStubInterface, []string) *BookInfo {
		ri := &BookInfo{true, ""}
		_id := args[0]
		_Book, err := stub.GetState(_id)
		if err != nil {
			ri.error(err.Error())
		} else {
			if _Book == nil {
				ri.ok("Warnning Book does not exists")
			} else {
				err := stub.DelState(_id)
				if err != nil {
					ri.error(err.Error())
				} else {
					ri.ok(stub.GetTxID())
				}
			}
		}
		return ri
	})
}

func (t *Book) update(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return t.handleProcess(stub, args, 7, func(shim.ChaincodeStubInterface, []string) *BookInfo {
		ri := &BookInfo{true, ""}
		_id := args[0]
		_userid := args[1]
		_username := args[2]
		_fullname := args[3]
		_userphone := args[4]
		_companyid := args[5]
		_companyname := args[6]
		_loginstate := args[7]
		_logintype := args[8]
		_logintime := args[9]
		_loginip := args[10]
		_devicetype := args[11]
		_loginaddress := args[12]
		_loginlocation := args[13]

		newBook := &Book{_id, _userid, _username,
			_fullname, _userphone, _companyid, _companyname,
			_loginstate, _logintype, _logintime, _loginip,
			_devicetype, _loginaddress, _loginlocation}

		_Book, err := stub.GetState(_id)

		if err != nil {
			ri.error(err.Error())
		} else {
			if _Book == nil {
				ri.error("Error the Book does not exists")
			} else {
				_ejson, err := json.Marshal(newBook)
				if err != nil {
					ri.error(err.Error())
				} else {
					err := stub.PutState(_id, _ejson)
					if err != nil {
						ri.error(err.Error())
					} else {
						// 追加 根据 key 查询
						_tjson, err := json.Marshal(stub.GetTxID())
						if err != nil {
							ri.error(err.Error())
						} else {
							err := stub.PutState(TxId+_id, _tjson)
							if err != nil {
								ri.error(err.Error())
							} else {
								ri.ok(stub.GetTxID())
							}
						}
					}
				}
			}
		}
		return ri
	})
}

func (t *Book) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return t.handleProcess(stub, args, 1, func(shim.ChaincodeStubInterface, []string) *BookInfo {
		ri := &BookInfo{true, ""}
		queryString := args[0]

		value, err := stub.GetState(queryString)
		//rich查询，leveldb不支持
		//queryResults, err := getQueryResultForQueryString(stub, queryString)
		if err != nil {
			ri.error(err.Error())
		} else {
			ri.ok(string(value))
		}
		return ri
	})
}

func (t *Book) history(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return t.handleProcess(stub, args, 1, func(shim.ChaincodeStubInterface, []string) *BookInfo {
		ri := &BookInfo{true, ""}
		queryString := args[0]

		queryResults, err := getHistoryForKeyStrcts(stub, queryString)
		if err != nil {
			ri.error(err.Error())
		} else {
			ri.ok(string(queryResults))
		}
		return ri
	})
}
func getHistoryForKeyStrcts(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
	var list []BookList
	var row BookList
	resultsIterator, err := stub.GetHistoryForKey(queryString)
	if err != nil {
		return nil, err
	}

	defer resultsIterator.Close()
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		row.TxId = queryResponse.TxId
		row.Value = queryResponse.Value
		list = append(list, row)
	}
	BytesList, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return BytesList, nil
}

func getHistoryForKeyString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetHistoryForKey(queryString)
	if err != nil {
		return nil, err
	}

	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"txid\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"value\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *Book) handleProcess(stub shim.ChaincodeStubInterface, args []string, expectNum int, f process) pb.Response {
	res := &BookInfo{false, ""}
	err := t.checkArgs(args, expectNum)
	if err != nil {
		res.error(err.Error())
	} else {
		res = f(stub, args)
	}
	return res.response()
}

func (t *Book) checkArgs(args []string, expectNum int) error {
	if len(args) != expectNum {
		return fmt.Errorf("Incorrect number of arguments. Expecting  " + strconv.Itoa(expectNum))
	}
	for p := 0; p < len(args); p++ {
		if len(args[p]) <= 0 {
			return fmt.Errorf(strconv.Itoa(p+1) + "nd argument must be a non-empty string")
		}
	}
	return nil
}

func (t *Book) updateTxid(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return t.handleProcess(stub, args, 2, func(shim.ChaincodeStubInterface, []string) *BookInfo {
		ri := &BookInfo{true, ""}
		_id := args[0]
		_txid := args[1]
		err := stub.PutState(TxId+_id, []byte(_txid))
		if err != nil {
			ri.error(err.Error())
		} else {
			ri.ok(stub.GetTxID())
		}
		return ri
	})
}

func (t *Book) test(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return t.handleProcess(stub, args, 1, func(shim.ChaincodeStubInterface, []string) *BookInfo {
		ri := &BookInfo{true, ""}

		ri.ok("init server is ok ")

		return ri
	})
}
func main() {
	err := shim.Start(new(Book))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
