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

type ProductInfo struct {
        //产品ID
        ProductId  string `json:"product_id"`
		//作物ID（作为溯源的ID）
		CropsId string `json:"crops_id"`
		//成品名称
		ProductName string `json:"product_name"`
		//配料
		MixedIngredients string `json:"mixed_ingredients"`
		//负责人
		Leader string `json:"leader"`
		//负责人电话
		LeaderTel string `json:"leader_tel"`
		//车间
		Workshop string `json:"workshop"`
		//工厂名称
		Factory string `json:"factory"`
		//工时
		WorkHours string `json:"work_hours"`
		//时间
		Time string `json:"time"`
		//保存方法
		KeepMethod string `json:"keep_mathod"`
		//净含量
		NetContent string `json:"Net_Content"`
		//食用推荐
		CookingRecommend string `json:"cooking_recommend"`
		//备注
		Remarks string `json:"remarks"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
        return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
        function, args := APIstub.GetFunctionAndParameters()
        if function == "queryProductInfoById" {
             return s.queryProductInfoById(APIstub, args)
        }else if function == "initLedger" {
             return s.initLedger(APIstub)
        }else if function == "createProductInfo" {
             return s.createProductInfo(APIstub, args)
        }else if function == "queryProductInfoByCropsId" { //根据id查询物流详情（司机端）
             return s.queryProductInfoByCropsId(APIstub,args)
        }
        return shim.Error("Invalid Smart Contract function name.")
}

/**
 * 根据商品ID查询商品信息
 */
func (s *SmartContract) queryProductInfoById(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        cropsAsBytes, _ := APIstub.GetState(args[0])
        return shim.Success(cropsAsBytes)
}

//根据crops_id查询商品信息
func (s *SmartContract) queryProductInfoByCropsId(APIstub shim.ChaincodeStubInterface,args[]string) sc.Response{
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
        productInfo := []ProductInfo{
                 ProductInfo{ProductId: "the first product info",
                 CropsId: "crops_id",
                 ProductName: "steak商品",
                 MixedIngredients: "配料",
                 Leader: "刘老板",
                 LeaderTel: "13657746155",
                 Workshop: "刘牌总公司第一车间",
		 Factory: "steakliu集团有限公司",
		 WorkHours: "36时",
		 Time: "2020.10.1",
		 KeepMethod: "放在通风处",
		 NetContent: "500克*10包",
		 CookingRecommend: "建议浸泡半个小时",
		 Remarks: "原料厂商初始块" },
                        }
        i := 0
        for i < len(productInfo) {
                fmt.Println("i is ", i)
                productInfoAsBytes, _ := json.Marshal(productInfo[i])
                APIstub.PutState(productInfo[i].ProductId, productInfoAsBytes)
                fmt.Println("Added", productInfo[i])
                i = i + 1
        }
        return shim.Success(nil)
}


/**
 *添加商品信息
 */
func (s *SmartContract) createProductInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
         if len(args) != 15 {
                return shim.Error("Incorrect number of arguments. Expecting 15")
         }
         var productInfo = ProductInfo{ProductId: args[1], CropsId: args[2], ProductName: args[3],MixedIngredients: args[4], Leader: args[5], LeaderTel: args[6],Workshop: args[7],Factory: args[8],WorkHours: args[9],Time: args[10],KeepMethod: args[11],NetContent: args[12],CookingRecommend: args[13],Remarks: args[14]}
         productInfoAsBytes, _ := json.Marshal(productInfo)
         APIstub.PutState(args[0], productInfoAsBytes)
         return shim.Success(nil)
}

func main() {
        err := shim.Start(new(SmartContract))
        if err != nil {
                fmt.Printf("Error creating new Smart Contract: %s", err)
        }

}

