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

type Crops struct {
	//作物ID
        CropsId   string `json:"crops_id"`
        //作物名称
        CropsName  string `json:"crops_name"`
	//所在地
	Address string `json:"address"`
	//生长开始日期
	RegisterTime string `json:"register_time"`
        //年度
	Year string `json:"year"`
	//农户名字
        FarmerName string `json:"farmer_name"`
	//农户ID
	FarmerID string `json:"farmer_id"`
        //联系电话
	FarmerTel string `json:"farmer_tel"`
	//肥料名称
	FertilizerName string `json:"fertilizer_name"`
	//种植方式
	PlatMode string `json:"plant_mode"`
	//是否套袋种植
	BaggingStatus string `json:"bagging_status"`
	//育苗周期
	GrowSeedlingsCycle string `json:"grow_seedlings_cycle"`
	//灌溉周期
	IrrigationCycle string `json:"irrigation_cycle"`
	//施肥周期
	ApplyFertilizerCycle string `json:"apply_fertilizer_cycle"`
	//除草周期
	WeedCycle string `json:"weed_cycle"`
	//备注
	Remarks string `json:"remarks"`

}

type CropsGrowInfo struct {
        //生长情况唯一ID
	CropsGrowId string `json:"crops_grow_id"`
        //作物ID
        CropsBakId string `json:"crops_bak_id"`
        //记录时间
        RecordTime string `json:"record_time"`
        //作物生长图片URL
        CropsGrowPhotoUrl string `json:"crops_grow_photo_url"`
        //温度
        Temperature string `json:"temperature"`
	//生长情况
	GrowStatus string `json:"grow_status"`
	//水分
	WaterContent string `json:"water_content"`
	//光照情况
	IlluminationStatus string `json:"illumination_status"`
        //备注
        Remarks string `json:"remarks"`

}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
        return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
        function, args := APIstub.GetFunctionAndParameters()
        // Route to the appropriate handler function to interact with the ledger appropriately
		//根据作物id查询作物	
        if function == "queryCropsById" {
             return s.queryCropsById(APIstub, args)
        }else if function == "initLedger" {
             return s.initLedger(APIstub)
        }else if function == "createCrops" {
             return s.createCrops(APIstub, args)
        }else if function == "queryCropsProcessByCropsId" {
             return s.queryCropsProcessByCropsId(APIstub,args)
        }else if function == "recordCropsGrow" {
             return s.recordCropsGrow(APIstub, args)
        }
        return shim.Error("Invalid Smart Contract function name.")
}

/**
 * 根据作物ID查询作物信息
 */
func (s *SmartContract) queryCropsById(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        cropsAsBytes, _ := APIstub.GetState(args[0])
        return shim.Success(cropsAsBytes)
}

//根据crops_id溯源所有记录过程
func (s *SmartContract) queryCropsProcessByCropsId(APIstub shim.ChaincodeStubInterface,args[]string) sc.Response{
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	CropsBakId := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"crops_bak_id\":{\"$eq\":\"%s\"}}}", CropsBakId)
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
 * 初始化账本
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
        crops := []Crops{
		 Crops{CropsId: "steak_liu_first_record",
                 CropsName: "Steakliu",
                 Address: "GuiZhou",
                 RegisterTime: "2020.10.1",
		 Year: "2020",
                 FarmerName: "LiUPAI",
                 FarmerID: "2319492349",
                 FertilizerName: "化肥",
                 PlatMode: "l露天",
                 BaggingStatus: "套袋",
                 GrowSeedlingsCycle: "30天",
                 IrrigationCycle: "4天",
                 ApplyFertilizerCycle: "无施肥",
                 WeedCycle: "5天",
                 Remarks: "天然有机农产品"},
                        }
        i := 0
        for i < len(crops) {
                fmt.Println("i is ", i)
                cropsAsBytes, _ := json.Marshal(crops[i])
                APIstub.PutState(crops[i].CropsId, cropsAsBytes)
                fmt.Println("Added", crops[i])
                i = i + 1
        }
        return shim.Success(nil)
}

/**
 *记录作物生长
 */
func (s *SmartContract) recordCropsGrow(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) !=  10{
                return shim.Error("Incorrect number of arguments. Expecting 10")
        }
        var cropsGrowInfo = CropsGrowInfo{CropsGrowId: args[1], CropsBakId: args[2], RecordTime: args[3], CropsGrowPhotoUrl: args[4], Temperature: args[5], GrowStatus: args[6], WaterContent: args[7], IlluminationStatus: args[8], Remarks: args[9]}
        cropsGrowInfoAsBytes, _ := json.Marshal(cropsGrowInfo)
        APIstub.PutState(args[0],cropsGrowInfoAsBytes)
        return shim.Success(nil)
}

/**
 *添加作物
 */
func (s *SmartContract) createCrops(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 if len(args) != 16 {
	        return shim.Error("Incorrect number of arguments. Expecting 15")
	 }
         var crops = Crops{CropsId: args[1], CropsName: args[2], Address: args[3],FarmerName: args[4],FertilizerName: args[5], PlatMode: args[6],BaggingStatus: args[7],GrowSeedlingsCycle: args[8],IrrigationCycle: args[9],ApplyFertilizerCycle: args[10],WeedCycle: args[11],Remarks: args[12],RegisterTime: args[13],Year: args[14],FarmerTel: args[15]}
	 cropsAsBytes, _ := json.Marshal(crops)
         APIstub.PutState(args[0], cropsAsBytes)
         return shim.Success(nil)
}

func (s *SmartContract) queryAllCropsByFarmerID(APIstub shim.ChaincodeStubInterface,args string) sc.Response {


        startKey := "CAR0"
        endKey := "CAR999"

        resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
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


func main() {
        err := shim.Start(new(SmartContract))
        if err != nil {
                fmt.Printf("Error creating new Smart Contract: %s", err)
        }

}

