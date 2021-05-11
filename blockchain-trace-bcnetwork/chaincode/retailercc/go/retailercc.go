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

type Retailer struct {
        //产品唯一ID（用于从零售商开始溯源）
        ProductId  string `json:"product_id"`
		//作物ID（作为溯源的ID）
		CropsId string `json:"crops_id"`
		//零售商负责人
		RetailerName string `json:"retailer_name"`
		//零售商电话
		RetailerTel string `json:"retailer_tel"`
                //零售商ID
		RetailerId string `json:"retailer_id"`
		//零售商店名
		Retailer string `json:"retailer"`
		//备注
		Remarks string `json:"remarks"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
        return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
        function, args := APIstub.GetFunctionAndParameters()
        if function == "queryRetailerById" {
             return s.queryRetailerById(APIstub, args)
        }else if function == "initLedger" {
             return s.initLedger(APIstub)
        }else if function == "createRetailer" {
             return s.createRetailer(APIstub, args)
        }else if function == "queryRetailerByCropsId" { //根据id查询物流详情（司机端）
             return s.queryRetailerByCropsId(APIstub,args)
        }
        return shim.Error("Invalid Smart Contract function name.")
}

/**
 * 根据商品ID查询商品信息
 */
func (s *SmartContract) queryRetailerById(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        cropsAsBytes, _ := APIstub.GetState(args[0])
        return shim.Success(cropsAsBytes)
}

//根据crops_id查询商品信息
func (s *SmartContract) queryRetailerByCropsId(APIstub shim.ChaincodeStubInterface,args[]string) sc.Response{
        if len(args) != 2 {
                return shim.Error("Incorrect number of arguments. Expecting 2")
        }
        CropsId := args[0]
	RetailerId := args[1]
        queryString := fmt.Sprintf("{\"selector\": {\"crops_id\": {\"$eq\": \"%s\"},\"retailer_id\": {\"$eq\":\"%s\"}}}", CropsId, RetailerId)
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
        retailer := []Retailer{
                 Retailer{ProductId: "the first product info",
                 CropsId: "crops_id",
                 RetailerName: "刘牌",
		 RetailerId: "steak",
                 RetailerTel: "13657746155",
                 Retailer: "小四零售部",
				 Remarks: "零售厂商初始块",
				 },
                        }
        i := 0
        for i < len(retailer) {
                fmt.Println("i is ", i)
                retailerAsBytes, _ := json.Marshal(retailer[i])
                APIstub.PutState(retailer[i].ProductId, retailerAsBytes)
                fmt.Println("Added", retailer[i])
                i = i + 1
        }
        return shim.Success(nil)
}


/**
 *添加商品信息
 */
func (s *SmartContract) createRetailer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
         if len(args) != 8 {
                return shim.Error("Incorrect number of arguments. Expecting 7")
         }
         var retailer = Retailer{ProductId: args[1], CropsId: args[2], RetailerName: args[3],RetailerId: args[4], RetailerTel: args[5],Retailer: args[6], Remarks: args[7]}
         retailerAsBytes, _ := json.Marshal(retailer)
         APIstub.PutState(args[0], retailerAsBytes)
         return shim.Success(nil)
}

func main() {
        err := shim.Start(new(SmartContract))
        if err != nil {
                fmt.Printf("Error creating new Smart Contract: %s", err)
        }

}

