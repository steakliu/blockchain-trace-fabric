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

type Operation struct {
        //操作ID
        OperationId  string `json:"operation_id"`
		//作物ID（作为溯源的ID）
		CropsId string `json:"crops_id"`
		//操作人
		OperationPeopleName string `json:"operation_people_name"`
		//操作人电话
		OperationPeopleTel string `json:"operation_people_tel"`
		//时间
		Time string `json:"time"`
		//工作内容
		WorkContent string `json:"work_content"`
		//备注
		Remarks string `json:"remarks"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
        return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
        function, args := APIstub.GetFunctionAndParameters()
        if function == "queryOperationInfoById" {
             return s.queryOperationById(APIstub, args)
        }else if function == "initLedger" {
             return s.initLedger(APIstub)
        }else if function == "createOperation" {
             return s.createOperation(APIstub, args)
        }else if function == "queryOperationByCropsId" {
             return s.queryOperationByCropsId(APIstub,args)
        }
        return shim.Error("Invalid Smart Contract function name.")
}

/**
 * 
 */
func (s *SmartContract) queryOperationById(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        cropsAsBytes, _ := APIstub.GetState(args[0])
        return shim.Success(cropsAsBytes)
}

//根据crops_id查询商品信息
func (s *SmartContract) queryOperationByCropsId(APIstub shim.ChaincodeStubInterface,args[]string) sc.Response{
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
        operation := []Operation{
                 Operation{OperationId: "the first operation info",
                 CropsId: "crops_id",
                 OperationPeopleName: "steakliu",
                 OperationPeopleTel: "13657746155",
		 Time: "2020.10.1",
	         WorkContent: "工作内容",
		 Remarks: "原料厂商初始块"},
                        }
        i := 0
        for i < len(operation) {
                fmt.Println("i is ", i)
                operationAsBytes, _ := json.Marshal(operation[i])
                APIstub.PutState(operation[i].OperationId, operationAsBytes)
                fmt.Println("Added", operation[i])
                i = i + 1
        }
        return shim.Success(nil)
}


/**
 *添加操作
 */
func (s *SmartContract) createOperation(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
         if len(args) != 8 {
                return shim.Error("Incorrect number of arguments. Expecting 8")
         }
         var operation = Operation{OperationId: args[1], CropsId: args[2], OperationPeopleName: args[3],OperationPeopleTel: args[4],
		 Time: args[5], WorkContent: args[6],Remarks: args[7]}
         operationAsBytes, _ := json.Marshal(operation)
         APIstub.PutState(args[0], operationAsBytes)
         return shim.Success(nil)
}

func main() {
        err := shim.Start(new(SmartContract))
        if err != nil {
                fmt.Printf("Error creating new Smart Contract: %s", err)
        }

}

