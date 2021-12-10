<template>
	<div class="app-container">
		<el-divider>待处理工作业务</el-divider>
		<el-table :data="taskLists">
			<el-table-column type="selection" width="55" align="center" />
			<el-table-column label="任务编号" prop="taskId" />
			<el-table-column label="任务名称" prop="taskName" />
			<el-table-column label="发布人" prop="leaderName" />
			<el-table-column label="时间" prop="time" />
			<el-table-column label="具体内容" prop="taskRemark" />
			<el-table-column label="操作" align="center" class-name="small-padding fixed-width">
				<template slot-scope="scope">
					<el-button size="mini" type="text" @click="addWorkContent(scope.row)">上报工作内容</el-button>

				</template>
			</el-table-column>
		</el-table>

		<el-dialog center title="上报工作详情" :visible.sync="workContentDialog" width="700px" append-to-body>
			<el-form ref="form" label-width="80px" :model="recordForm">
			<el-row>
				<el-col :span="12">
					<el-form-item label="任务ID" prop="nickName"><el-input :disabled="true" v-model="taskInfo.taskId" placeholder="请输入" /></el-form-item>
				</el-col>
				<el-col :span="12">
					<el-form-item label="任务名称" prop="nickName"><el-input :disabled="true" v-model="taskInfo.taskName" placeholder="请输入" /></el-form-item>
				</el-col>
			</el-row>
			<el-row>
				<el-col :span="24">
					<el-form-item label="工作内容"><el-input type="textarea" v-model="recordForm.workContent" placeholder="请输入" /></el-form-item>
				</el-col>
			</el-row>
			<el-row>
				<el-col :span="24">
					<el-form-item label="备注"><el-input type="textarea" v-model="recordForm.remarks" placeholder="请输入" /></el-form-item>
				</el-col>
			</el-row>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button type="primary" @click="submitOperation">确 定</el-button>
				<el-button @click="cancel">取 消</el-button>
			</div>
		</el-dialog>
	</div>
</template>

<script>
import {queryTaskList} from '../../../api/trace/product';
import { formatDate } from '../../../utils/index.js';
export default {
	name: 'Record',
	data() {
		return {
			taskLists:[],
			workContentDialog:false,
			recordForm:{},
			taskInfo:'',
		};
	},
	created(){
		this.taskList();
	},
	methods:{
		submitOperation(){
			const operationArray = []
			const id = new this.$snowFlakeId().generate()
			operationArray.push(id)
			operationArray.push(id)
			operationArray.push(this.taskInfo.cropsId)
			operationArray.push(this.$store.getters.nickName)
			operationArray.push(this.$store.getters.phonenumber)
			operationArray.push(formatDate(new Date()))
			operationArray.push(this.recordForm.workContent)
			operationArray.push(this.recordForm.remarks)
			this.$httpBlock
				.post(this.$httpUrl + '/productapi/createOperation', { operationArray: operationArray })
				.then(res => {
					if (res.status === 200) {
						this.msgSuccess("数据上链成功");
						this.workContentDialog = false
					}
				})
				.catch(err => {
					this.msgError('存储异常 ' + err);
				});
		},

		addWorkContent(row){
			this.taskInfo = row
			this.workContentDialog = true
		},

		taskList(){
			queryTaskList().then(res => {
				this.taskLists = res.data
			}).catch(err => {
				this.msgError("获取任务失败")
			})
		}
	}

}
</script>

<style>
</style>
