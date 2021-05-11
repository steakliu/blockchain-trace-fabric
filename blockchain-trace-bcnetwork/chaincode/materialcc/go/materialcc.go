package main
import (
        "bytes"
        "encoding/json"
        "fmt"
       // "strconv"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Machining struct {
        //加工ID
        MachiningId  string `json:"machining_id"`
		//原料厂商负责人
		Leader string `json:"leader"`
		//货物ID
		CropsId string `json:"crops_id"`
		//原料厂商负责人tel
		LeaderTel string `json:"leader_tel"`
		//厂商名称
		FactoryName string `json:"factory_name"`
		//检测结果
		TestingResult string `json:"testing_result"`
		//入库时间
		InFactoryTime string `json:"in_factory_time"`
		//出库时间
		OutFactoryTime string `json:"out_factory_time"`
		//质检过程图片
		TestingPhotoUrl string `json:"testing_photo_url"`
		//备注
		Remarks string `json:"remarks"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
        return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
        function, args := APIstub.GetFunctionAndParameters()
        if function == "queryMachiningByCropsId" {
             return s.queryMachiningByCropsId(APIstub, args)
        }else if function == "initLedger" {
             return s.initLedger(APIstub)
        }else if function == "createMachining" {
             return s.createMachining(APIstub, args)
        }else if function == "queryMachiningById" { //根据id查询物流详情（司机端）
             return s.queryMachiningById(APIstub,args)
        }
        return shim.Error("Invalid Smart Contract function name.")
}

/**
 * 根据物流ID查询物流信息
 */
func (s *SmartContract) queryMachiningById(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        cropsAsBytes, _ := APIstub.GetState(args[0])
        return shim.Success(cropsAsBytes)
}

//根据crops_id查询原料详情
func (s *SmartContract) queryMachiningByCropsId(APIstub shim.ChaincodeStubInterface,args[]string) sc.Response{
        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        CropsId := args[0]
        queryString := fmt.Sprintf("{\"selector\":{\"crops_id\":{\"$eq\":\"%s\"}}}", CropsId)
        resultsIterator, err := APIstub.GetQueryResult(queryString)
        if err != nil {
                return shim.Error(err.Error())
        }
        defer resultsIterator.Close()

        // buffer is a JSON array containing QueryResults
        var buffer bytes.Buffer
        buffer.WriteString("[")

        bArrayMemberAlreadyWritten := false
        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()
                if err != nil {
                        return shim.Error(err.Error())
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

        fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

        return shim.Success(buffer.Bytes())
}


/**
 * 初始化原料厂商质检账本
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
        machining := []Machining{
                 Machining{MachiningId: "the first maching",
                 Leader: "刘牌",
                 LeaderTel: "13657746155",
                 FactoryName: "steakliu原料厂商",
                 CropsId: "123456",
                 TestingResult: "一切合格",
                 InFactoryTime: "2020.10.1",
				 OutFactoryTime: "2020.10.2",
				 TestingPhotoUrl: "质检过程图片",
                 Remarks: "原料厂商初始块"},
                        }
        i := 0
        for i < len(machining) {
                fmt.Println("i is ", i)
                machiningAsBytes, _ := json.Marshal(machining[i])
                APIstub.PutState(machining[i].MachiningId, machiningAsBytes)
                fmt.Println("Added", machining[i])
                i = i + 1
        }
        return shim.Success(nil)
}


/**
 *添加原料厂商质检信息
 */
func (s *SmartContract) createMachining(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
         if len(args) != 11 {
                return shim.Error("Incorrect number of arguments. Expecting 11")
         }
         var machining = Machining{MachiningId: args[1], Leader: args[2], LeaderTel: args[3],FactoryName: args[4],CropsId: args[5], TestingResult: args[6],InFactoryTime: args[7],OutFactoryTime: args[8],TestingPhotoUrl: args[9],Remarks: args[10]}
         machiningAsBytes, _ := json.Marshal(machining)
         APIstub.PutState(args[0], machiningAsBytes)
         return shim.Success(nil)
}

func main() {
        err := shim.Start(new(SmartContract))
        if err != nil {
                fmt.Printf("Error creating new Smart Contract: %s", err)
        }

}

