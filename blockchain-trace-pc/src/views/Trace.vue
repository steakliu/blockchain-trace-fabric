<template>
  <div style="height: 700px;overflow: auto;">
    <el-row :gutter="30">
      <el-col :span="8">
        <avue-data-tabs :option="option"></avue-data-tabs>
        <br />
        <div>
          <div>
            <el-row>
              <el-col :span="24">区块编号：{{ blockInfo.data.header.number }}</el-col>
            </el-row>
            <br />
            <el-row>
              <el-col :span="24">
                前置Hash：
                <el-link @click="queryInfo(blockInfo.data.header.previous_hash)" type="success">{{ blockInfo.data.header.previous_hash }}<span style="color: red">(点击查看区块信息)</span></el-link>
              </el-col>
            </el-row>
            <br />
            <el-row>
              <el-col :span="24">
                数据Hash：
                <el-link :underline="false" type="success">{{ blockInfo.data.header.data_hash }}</el-link>
              </el-col>
            </el-row>
          </div>
          <div align="center">
            <el-divider content-position="left"><i class="el-icon-top"></i></el-divider>
          </div>
          <div>
            <el-row>
              <el-col :span="24">区块编号：{{ previousBlockInfo.data.header.number }}</el-col>
            </el-row>
            <br />
            <el-row>
              <el-col :span="24">
                前置Hash：
                <el-link @click="queryInfo(previousBlockInfo.data.header.previous_hash)" type="success">{{ previousBlockInfo.data.header.previous_hash }}<span style="color: red">(点击查看区块信息)</span></el-link>
              </el-col>
            </el-row>
            <br />
            <el-row>
              <el-col :span="24">
                数据Hash：
                <el-link :underline="false" type="success">{{ previousBlockInfo.data.header.data_hash }}</el-link>
              </el-col>
            </el-row>
          </div>
          <div align="center">
            <el-divider content-position="left"><i class="el-icon-top"></i></el-divider>
          </div>

          <div>
            <el-row>
              <el-col :span="24">区块编号：{{ previousBlockInfo2.data.header.number }}</el-col>
            </el-row>
            <br />
            <el-row>
              <el-col :span="24">
                前置Hash：
                <el-link @click="queryInfo(previousBlockInfo2.data.header.previous_hash)" type="success">
                  {{ previousBlockInfo2.data.header.previous_hash }}<span style="color: red">(点击查看区块信息)</span>
                </el-link>
              </el-col>
            </el-row>
            <br />
            <el-row>
              <el-col :span="24">
                数据Hash：
                <el-link :underline="false" type="success">{{ previousBlockInfo2.data.header.data_hash }}</el-link>
              </el-col>
            </el-row>
          </div>
        </div>
      </el-col>
      <el-col :span="16">
        <div>
          <el-input style="width: 40%;" v-model="traceId" placeholder="溯源编号"></el-input>
          &nbsp;&nbsp;&nbsp;
          <el-button type="success" @click="query">查询</el-button>
        </div>
        <br />
        <div>
          <el-row>
            <el-col :span="6">
              <el-form>
                <el-form-item><span style="color: #13CE66;">产品基本信息</span></el-form-item>
                <el-form-item label="生产厂商负:">{{ productData.factory }}</el-form-item>
                <el-form-item label="生产厂商负责人:">{{ productData.leader }}</el-form-item>
                <el-form-item label="电话:">{{ productData.leader_tel }}</el-form-item>
                <el-form-item label="净含量:">{{ productData.Net_Content }}</el-form-item>
                <el-form-item label="生产车间:">{{ productData.workshop }}</el-form-item>
                <el-form-item label="生产工时">{{ productData.work_hours }}</el-form-item>
                <el-form-item label="保存方法">{{ productData.keep_mathod }}</el-form-item>
                <el-form-item label="食用推荐">{{ productData.cooking_recommend }}</el-form-item>
              </el-form>
            </el-col>
            <el-col :span="6">
              <el-form>
                <el-form-item><span style="color: #13CE66;">零售商信息</span></el-form-item>
                <el-form-item label="产品ID:">{{ retailerInfo.product_id }}</el-form-item>
                <el-form-item label="产品名称:">{{ productData.product_name }}</el-form-item>
                <el-form-item label="零售点:">{{ retailerInfo.retailer }}</el-form-item>
                <el-form-item label="负责人:">{{ retailerInfo.retailer_name }}</el-form-item>
                <el-form-item label="电话">{{ retailerInfo.retailer_tel }}</el-form-item>
              </el-form>
            </el-col>
            <el-col :span="12">
              <el-tabs v-model="activeName" @tab-click="handleClick">
                <el-tab-pane label="种植过程" name="1">
                  <el-row :gutter="8">
                    <el-col :span="24" v-for="(info, index) in cropsProcessDetailsArray" :key="index" style="padding-top: 10px;">
                      <el-card style="border-color: #42b983;">
                        <el-col :span="12">
                          <el-row>
                            <el-col :span="24">生长情况：{{ info.grow_status }}</el-col>
                          </el-row>
                          <br />
                          <el-row>
                            <el-col :span="24">光照情况：{{ info.illumination_status }}</el-col>
                          </el-row>
                          <br />
                          <el-row>
                            <el-col :span="24">水分：{{ info.water_content }}</el-col>
                          </el-row>
                          <br />
                          <el-row>
                            <el-col :span="24">温度：{{ info.temperature }}</el-col>
                          </el-row>
                          <br />
                          <el-row>
                            <el-col :span="24">记录时间：{{ info.record_time }}</el-col>
                          </el-row>
                          <br />
                          <el-row>
                            <el-col :span="24">备注：{{ info.remarks }}</el-col>
                          </el-row>
                        </el-col>
                        <el-col :span="12"><img style="width: 100%;height: 11.05rem;" :src="info.crops_grow_photo_url" /></el-col>
                      </el-card>
                    </el-col>
                    <br />
                  </el-row>
                </el-tab-pane>
                <el-tab-pane label="流通行为" name="2">
                  <div class="block">
                    <el-timeline v-for="(detail,index) in cropsDriverArray">
                      <el-timeline-item :timestamp="detail.transport_to_chain_time" placement="top" >
                        <el-card>
                          <h4 style="color: red;">经过地:{{detail.transport_to_address}}</h4>
                          <p>物流ID:{{detail.transport_id}}</p>
                          <p>司机:{{detail.driver_name}}</p>
                          <p>电话:{{detail.driver_tel}}</p>
                          <p>备注:{{detail.remarks}}</p>
                        </el-card>
                      </el-timeline-item>
                    </el-timeline>
                  </div>
                </el-tab-pane>
                <el-tab-pane label="厂商质检" name="3">
                  <el-card style="border-color: #42b983;">
                    <el-col :span="12">
                      <el-row>
                        <el-col :span="24">负责人：{{ machingInfo.leader }}</el-col>
                      </el-row>
                      <br /><br />
                      <el-row>
                        <el-col :span="24">电话：{{ machingInfo.leader_tel }}</el-col>
                      </el-row>
                      <br /><br />
                      <el-row>
                        <el-col :span="24">原料厂商：{{ machingInfo.factory_name }}</el-col>
                      </el-row>
                      <br /><br />
                      <el-row>
                        <el-col :span="24">入厂时间：{{ machingInfo.in_factory_time }}</el-col>
                      </el-row>
                      <br /><br />
                      <el-row>
                        <el-col :span="24">出厂时间：{{ machingInfo.out_factory_time }}</el-col>
                      </el-row>
                      <br /><br />
                      <el-row>
                        <el-col :span="24">检测结果：{{ machingInfo.testing_result }}</el-col>
                      </el-row>
                      <br /><br />
                      <el-row>
                        <el-col :span="24">备注：{{ machingInfo.remarks }}</el-col>
                      </el-row>
                      <br />
                    </el-col>
                    <el-col :span="12"><img style="width: 100%;height: 21.05rem;" :src="machingInfo.testing_photo_url" /></el-col>
                  </el-card>
                </el-tab-pane>
                <el-tab-pane label="生产行为" name="4">
                  <el-row :gutter="8" >
                    <el-col :span="24" v-for="(info,index) in operationArray" :key="index" style="padding-top: 10px;">
                      <el-card style="border-color: #42b983;">
                        <el-row>
                          <el-col :span="24">操作ID：{{ info.operation_id }}</el-col>
                        </el-row>
                        <br>
                        <el-row>
                          <el-col :span="24">操作人：{{ info.operation_people_name }}</el-col>
                        </el-row>
                        <br>
                        <el-row>
                          <el-col :span="24">操作人电话：{{ info.operation_people_tel }}</el-col>
                        </el-row>
                        <br>
                        <el-row>
                          <el-col :span="24">操作时间：{{ info.time }}</el-col>
                        </el-row>
                        <br>
                        <el-row>
                          <el-col :span="24">操作内容：{{ info.work_content }}</el-col>
                        </el-row>
                      </el-card>
                    </el-col>
                    <br>
                  </el-row>
                </el-tab-pane>
                <el-tab-pane label="农产信息" name="5">
                  <el-card style="border-color: #42b983;">
                    <el-row><el-col :span="24">农产名字：{{ cropsDetails.crops_name }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">种植户：{{ cropsDetails.farmer_name }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">电话：{{ cropsDetails.farmer_tel }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">地址：{{ cropsDetails.address }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">种植方式：{{ cropsDetails.plant_mode }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">年度：{{ cropsDetails.year }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">套袋方式：{{ cropsDetails.bagging_status }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">使用肥料：{{ cropsDetails.fertilizer_name }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">育苗周期：{{ cropsDetails.grow_seedlings_cycle }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">灌溉周期：{{ cropsDetails.irrigation_cycle }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">除草周期：{{ cropsDetails.weed_cycle }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">施肥周期：{{ cropsDetails.apply_fertilizer_cycle }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">注册时间：{{ cropsDetails.register_time }}</el-col></el-row><br><br>
                    <el-row><el-col :span="24">备注：{{ cropsDetails.remarks }}</el-col></el-row><br>
                  </el-card>
                </el-tab-pane>
              </el-tabs>
            </el-col>
          </el-row>
        </div>
      </el-col>
    </el-row>
    <BlockInfo ref="getBlockInfoRef"></BlockInfo>
  </div>
</template>

<script>
import axios from 'axios';
import BlockInfo from "./BlockInfo";

export default {
  components: {BlockInfo},
  data() {
    return {
      queryBlockByNumDrawer: false,
      channelBlockInfo: '',
      blockInfo: '',
      previousBlockInfo: '',
      previousBlockInfo2: '',
      open: false,
      blockInfoHash: '',
      option: {
        data: []
      },
      traceId: '',
      retailerInfo: '',
      productData: '',
      activeName: '1',
      cropsProcessDetailsArray: [],
      cropsDriverArray:[],
      machingInfo:'',
      operationArray:[],
      cropsDetails:''
    };
  },
  created() {
    this.$httpBlock
      .get(this.$httpUrl + '/blockexplorerapi/channelBlockInfo')
      .then(res => {
        this.channelBlockInfo = res;
        let data1 = {
          title: '区块高度',
          subtitle: '实时',
          count: res.data.height.low,
          text: '当前区块高度',
          color: 'rgb(27, 201, 142)',
          key: '高'
        };
        this.option.data.push(data1);
        //当前区块
        var number = Number(this.channelBlockInfo.data.height.low - 1);
        this.$httpBlock
          .get(this.$httpUrl + '/blockexplorerapi/queryBlockInfo?number=' + number)
          .then(res => {
            this.blockInfo = res;
          })
          .catch(err => {
            console.log('err ' + err);
          });

        //前一个区块
        var previous = Number(this.channelBlockInfo.data.height.low - 2);
        this.$httpBlock
          .get(this.$httpUrl + '/blockexplorerapi/queryBlockInfo?number=' + previous)
          .then(res => {
            this.previousBlockInfo = res;
          })
          .catch(err => {
            console.log('err ' + err);
          });

        //前2个区块
        var previous2 = Number(this.channelBlockInfo.data.height.low - 3);
        this.$httpBlock
          .get(this.$httpUrl + '/blockexplorerapi/queryBlockInfo?number=' + previous2)
          .then(res => {
            this.previousBlockInfo2 = res;
          })
          .catch(err => {
            console.log('err ' + err);
          });
      })
      .catch(err => {
        console.log('err ' + err);
      });

    //this.getHash();
  },
  methods: {
    handleClick() {
      if (this.activeName === '1') {
        this.$httpBlock
          .get(this.$httpUrl + '/farmerapi/queryCropsProcessByCropsId?id=' + this.productData.crops_id)
          .then(res => {
            const array = [];
            for (let i = 0; i < res.data.length; i++) {
              array.push(res.data[i].Record);
            }
            this.cropsProcessDetailsArray = array;
          })
          .catch(err => {});
      }
      if(this.activeName === '2'){
        this.$httpBlock
          .get(this.$httpUrl + '/driverapi/queryTransportByCropsId?id='+this.productData.crops_id)
          .then(res => {
            const array = [];
            this.len = res.data.length
            for (let i = 0; i < res.data.length; i++) {
              array.push(res.data[i].Record);
            }
            this.cropsDriverArray = array
          })
          .catch(err => {});
      }
      if(this.activeName === '3'){
        this.$httpBlock
          .get(this.$httpUrl + '/materialapi/queryMachiningByCropsId?id='+this.productData.crops_id)
          .then(res => {
            this.machingInfo = res.data[0].Record
          })
          .catch(err => {});
      }
      if(this.activeName === '4'){
        this.$httpBlock
          .get(this.$httpUrl + '/productapi/queryOperationByCropsId?id='+this.productData.crops_id)
          .then(res => {
            const array = [];
            for (let i = 0; i < res.data.length; i++) {
              array.push(res.data[i].Record);
            }
            this.operationArray = array;
          })
          .catch(err => {
            this.msgError('查询异常 ' + err);
          });
      }
      if(this.activeName === '5'){
        this.$httpBlock
          .get(this.$httpUrl + '/farmerapi/queryCropsById?id='+this.productData.crops_id)
          .then(res => {
            this.cropsDetails = res.data;
          })
          .catch(err => {});
      }
    },

    query() {
      this.$httpBlock
        .get(this.$httpUrl + '/retailerapi/queryRetailerById?id=' + this.traceId)
        .then(res => {
          this.retailerInfo = res.data; //零售
          this.$httpBlock
            .get(this.$httpUrl + '/productapi/queryProductInfoByCropsId?id=' + res.data.crops_id)
            .then(res => {
              this.productData = res.data[0].Record; //产品信息
              this.handleClick()
            })
            .catch(err => {
              this.msgError('查询异常 ' + err);
            });
        })
        .catch(err => {
          console.log('调用失败 ' + JSON.stringify(err));
        });

    },

    queryInfo(hash) {
      this.open = true;
      this.$httpBlock
        .get(this.$httpUrl + '/blockexplorerapi/queryBlockBuHash?hash=' + hash)
        .then(res => {
          this.blockInfoHash = res;
          this.$refs.getBlockInfoRef.init(res)
        })
        .catch(err => {
          console.log('err ' + err);
        });
    },

    getHash() {
      var number = Number(this.channelBlockInfo.data.height.low);
      this.$httpBlock
        .get(this.$httpUrl + '/blockexplorerapi/queryBlockInfo?number=' + number)
        .then(res => {
          this.blockInfo = res;
        })
        .catch(err => {
          console.log('err ' + err);
        });
    },

    queryBlockByNum() {
      this.queryBlockByNumDrawer = true;
    }
  }
};
</script>

<style lang="scss" scoped>
#app {
  width: 100%;
  padding: 20px 20px 10px;
  border-radius: 4px;
  margin: 0px auto;
  background: #ffffff;

  box-sizing: border-box;
}

.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

@media (max-width: 1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}

#input-box {
  display: block;
  width: 100%;
  height: 46px;
  margin-bottom: 15px;
  border: none;
  border-radius: 6px;
  background: #add8e6;
  padding: 0 5px;
  text-align: right;
  font-size: 20px;
  color: #ffffff;
  font-weight: bold;
  letter-spacing: 1px;
}
</style>
