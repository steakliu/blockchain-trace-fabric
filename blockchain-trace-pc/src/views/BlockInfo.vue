<template>
<el-dialog center title="区块详情" :visible.sync="blockVisible" width="700px" append-to-body>
  <el-row :gutter="20">
    <el-col span="24">
      编号：<el-tag>{{blockInfo.data.header.number}}</el-tag>
    </el-col>
    <br>
    <br>
    <el-col span="24">
      前置hash：{{blockInfo.data.header.previous_hash}}
    </el-col>
    <br>
    <br>
    <el-col span="24">
      数据hash：{{blockInfo.data.header.data_hash}}
    </el-col>
    <br>
    <br>
    <el-col span="24">
      交易ID：{{blockInfo.data.data.data[0].payload.header.channel_header.tx_id}}
    </el-col>
    <br>
    <br>
    <el-col span="24">
      时间戳：{{blockInfo.data.data.data[0].payload.header.channel_header.timestamp}}
    </el-col>
    <br>
    <br>
    <el-col span="24">
      通道：<el-tag type="success">{{blockInfo.data.data.data[0].payload.header.channel_header.channel_id}}</el-tag>
    </el-col>
    <br>
    <br>
    <el-col span="24">
      智能合约：<el-tag type="info">{{blockInfo.data.data.data[0].payload.data.actions[0].payload.action.proposal_response_payload.extension.chaincode_id.name}}</el-tag>
    </el-col>
    <br>
    <br>
    <el-col span="24">
      智能合约版本：<el-tag type="warning">{{blockInfo.data.data.data[0].payload.data.actions[0].payload.action.proposal_response_payload.extension.chaincode_id.version}}</el-tag>
    </el-col>
<!--    <br>-->
<!--    <br>-->
<!--    <el-col span="24">-->
<!--      原始数据：{{dataDetail}}-->
<!--    </el-col>-->
  </el-row>
</el-dialog>
</template>

<script>
export default {
  name: "BlockInfo",
  data(){
    return{
      blockInfo:{},
      blockVisible:false,
      dataDetail:null
    }
  },

  methods:{
    init(blockInfo){
      this.blockVisible = true
      this.blockInfo = blockInfo
      this.dataDetail = JSON.parse(blockInfo.data.data.data[0].payload.data.actions[0].payload.action.proposal_response_payload.extension.results.ns_rwset[0].rwset.writes[0].value)
    }
  }
}
</script>

<style scoped>

</style>
