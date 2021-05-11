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

type Transport struct {
        //物流ID
        TransportId  string `json:"transport_id"`
		//司机ID
		DriverId string `json:"driver_id"`
		//司机名字
		DriverName string `json:"driver_name"`
		//司机电话
		DriverTel string `json:"driver_tel"`
		//所属部门
		DriverDept string `json:"driver_dept"`
		//货物ID
		CropsId string `json:"crops_id"`
		//物流信息上链时间
		TransportToChainTime string `json:"transport_to_chain_time"`
		//物流路过地址
		TransportToAddress string `json:"transport_to_address"`
		//备注（始发地，途中，目的地）
		Remarks string `json:"remarks"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
        return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
        function, args := APIstub.GetFunctionAndParameters()
        // Route to the appropriate handler function to interact with the ledger appropriately
                //根据作物id查询作物物流过程    
        if function == "queryTransportByCropsId" {
             return s.queryTransportByCropsId(APIstub, args)
        }else if function == "initLedger" {
             return s.initLedger(APIstub)
        }else if function == "createTransport" {
             return s.createTransport(APIstub, args)
        }else if function == "queryTransportById" { //根据id查询物流详情（司机端）
             return s.queryTransportById(APIstub,args)
        }
        return shim.Error("Invalid Smart Contract function name.")
}

/**
 * 根据物流ID查询物流信息
 */
func (s *SmartContract) queryTransportById(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        cropsAsBytes, _ := APIstub.GetState(args[0])
        return shim.Success(cropsAsBytes)
}

//根据crops_id溯源所有物流记录过程
func (s *SmartContract) queryTransportByCropsId(APIstub shim.ChaincodeStubInterface,args[]string) sc.Response{
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
 * 初始化物流账本
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
        transport := []Transport{
                 Transport{TransportId: "the first transportId",
                 DriverId: "2319492349",
                 DriverName: "steakliu",
                 DriverTel: "13657746155",
                 DriverDept: "steakliu集团有限公司",
                 CropsId: "123456",
                 TransportToChainTime: "2020.10.1",
                 TransportToAddress: "中华人民共和国",
                 Remarks: "物流初始块"},
                        }
        i := 0
        for i < len(transport) {
                fmt.Println("i is ", i)
                transportAsBytes, _ := json.Marshal(transport[i])
                APIstub.PutState(transport[i].TransportId, transportAsBytes)
                fmt.Println("Added", transport[i])
                i = i + 1
        }
        return shim.Success(nil)
}


/**
 *添加物流信息
 */
func (s *SmartContract) createTransport(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
         if len(args) != 10 {
                return shim.Error("Incorrect number of arguments. Expecting 10")
         }
         var transport = Transport{TransportId: args[1], DriverId: args[2], DriverName: args[3],DriverTel: args[4],DriverDept: args[5], TransportToChainTime: args[6],TransportToAddress: args[7],CropsId: args[8],Remarks: args[9]}
         transportAsBytes, _ := json.Marshal(transport)
         APIstub.PutState(args[0], transportAsBytes)
         return shim.Success(nil)
}

func main() {
        err := shim.Start(new(SmartContract))
        if err != nil {
                fmt.Printf("Error creating new Smart Contract: %s", err)
        }

}

