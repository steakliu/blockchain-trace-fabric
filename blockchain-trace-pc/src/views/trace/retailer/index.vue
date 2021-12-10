<template>
	<div class="app-container">

		<el-divider>待处理业务</el-divider>

		<br />
		<el-table :data="transportList">
			<el-table-column type="selection" width="55" align="center" />
			<el-table-column label="产品编号" prop="cropsId" />
			<el-table-column label="产品名称" prop="cropsName" />
			<el-table-column label="产品来源" prop="farmerDept" />
			<el-table-column label="时间" prop="time" />
			<el-table-column label="签收状态" prop="retailerReceiveStatus">
				<template slot-scope="scope">
					<el-tag v-show="scope.row.retailerReceiveStatus === null">未签收</el-tag>
					<el-tag v-show="scope.row.retailerReceiveStatus === 1">已签收</el-tag>
				</template>
			</el-table-column>
			<el-table-column label="操作" align="center" class-name="small-padding fixed-width">
				<template slot-scope="scope">
					<el-button v-show="scope.row.retailerReceiveStatus === null" size="mini" type="text" @click="receive(scope.row)">签收</el-button>
<!--					<el-button size="mini" type="text" @click="noticeTransport(scope.row)">打印条形码</el-button>-->
					<el-button size="mini" type="text" @click="getTraceId(scope.row)">获取产品溯源ID</el-button>
				</template>
			</el-table-column>
		</el-table>


		<el-dialog :title="title" :visible.sync="open" width="700px" append-to-body>
			<el-form ref="form" :model="form" label-width="80px">
				<el-row>
					<el-col :span="12">
						<el-form-item label="检测结果" prop="nickName"><el-input v-model="form.testingResult" placeholder="请输入检测结果" /></el-form-item>
					</el-col>
				</el-row>
				<el-form-item label="备注"><el-input v-model="form.remarks" type="textarea" placeholder="请输入内容"></el-input></el-form-item>
				<el-row>
					<el-col :span="12">
						<el-form-item label="检测图片:">
							<el-upload class="avatar-uploader" :on-change="getFile" :show-file-list="false" :auto-upload="false">
								<img v-if="imageUrl" :src="imageUrl" class="avatar" />
								<i v-else class="el-icon-plus avatar-uploader-icon"></i>
							</el-upload>
						</el-form-item>
					</el-col>
				</el-row>

			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button type="primary" @click="createMaching">确 定</el-button>
				<el-button @click="cancel">取 消</el-button>
			</div>
		</el-dialog>

		<el-dialog center title="联系运输" :visible.sync="noticeDetaiDialog" width="500px" append-to-body>
			<el-form ref="form" :model="trasportForm" label-width="80px">
				<el-row>
					<el-col :span="24">
						<el-form-item label="选择司机" prop="nickName">
							<el-select v-model="trasportForm.userName" placeholder="请选择">
								<el-option v-for="user in driverList" :key="user.userName" :label="user.nickName" :value="user.userName"></el-option>
							</el-select>
						</el-form-item>
					</el-col>
				</el-row>
				<el-row>
					<el-col :span="24">
						<el-form-item label="生产厂商" prop="nickName">
							<el-select v-model="trasportForm.deptId" placeholder="请选择">
								<el-option v-for="dept in factoryList" :key="dept.deptId" :label="dept.deptName" :value="dept.deptId"></el-option>
							</el-select>
						</el-form-item>
					</el-col>
				</el-row>
				<el-form-item label="备注"><el-input v-model="trasportForm.remarks" type="textarea" placeholder="请输入内容"></el-input></el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button type="primary" @click="addNoticeTrasport">确 定</el-button>
				<el-button @click="cancel">取 消</el-button>
			</div>
		</el-dialog>

		<el-dialog center title="产品溯源ID" :visible.sync="traceIdDialog" width="500px" append-to-body>
			<json-viewer :value="retailerInfo.product_id" expand-depth=5 copyable boxed></json-viewer>
		</el-dialog>

	</div>
</template>

<script>
import { listDept, getDept, delDept, addDept, updateDept, listDeptExcludeChild } from '@/api/system/dept';
import {addTransport, listCrops,changeInToOut} from '../../../api/trace/material';
import Treeselect from '@riophae/vue-treeselect';
import '@riophae/vue-treeselect/dist/vue-treeselect.css';
import { formatDate } from '../../../utils/index.js';
import {uplodImagesBase64,getAllDriverByDeptId,getFactoryByDeptId} from '../../../api/trace/farmer';
import {queryCropsList } from '../../../api/trace/product';
import {updateReceiveStatus } from '../../../api/trace/retailer';
export default {
	name: 'Dept',
	components: { Treeselect },
	data() {
		return {
			imageUrl:'',
			noticeDetaiDialog: false,
			// 遮罩层
			loading: true,
			// 显示搜索条件
			showSearch: true,
			// 表格树数据
			deptList: [],
			// 部门树选项
			deptOptions: [],
			// 弹出层标题
			title: '',
			// 是否显示弹出层
			open: false,
			// 状态数据字典
			statusOptions: [],
			// 查询参数
			queryParams: {
				deptName: undefined,
				status: undefined
			},
			// 表单参数
			form: {},

			transportList: [],
			cityDetail: '',
			lng: Number,
			lat: Number,
			imgBase64:'',
			TestingPhotoUrl:'',
			checkInfo:'',
			driverList:[],
			noticeTransportDialog:false,
			trasportForm:{},
			factoryList:[],
			traceIdDialog:false,
			retailerInfo:''
		};
	},
	created() {
		this.getList();
	},

	mounted() {
		//this.qrcode();
	},

	methods: {
		//获取产品溯源ID
		getTraceId(row){
			const retailerArray = []
			retailerArray.push(row.cropsId)
			retailerArray.push(this.$store.getters.name)
			this.$httpBlock
				.get(this.$httpUrl + '/retailerapi/queryRetailerByCropsId?cropsId='+row.cropsId+'&retailerId='+this.$store.getters.name)
				.then(res => {
          if (res.data.length < 1){
            this.msgError("请先签收后再获取");
          }else {
            this.traceIdDialog = true
            this.retailerInfo = res.data[0].Record;
          }

				})
				.catch(err => {
					this.msgError('存储异常 ' + err);
				});
		},

		//签收
		receive(row){
			const retailerArray = []
			const id = new this.$snowFlakeId().generate();
			retailerArray.push(id)
			retailerArray.push(id)
			retailerArray.push(row.cropsId)
			retailerArray.push(this.$store.getters.nickName)
			retailerArray.push(this.$store.getters.name)
			retailerArray.push(this.$store.getters.phonenumber)
			retailerArray.push(this.$store.getters.deptName)
			retailerArray.push('零售商售卖')
			this.$confirm('是否确认签收'+row.cropsName, '签收', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			})
				.then(() => {
					this.$httpBlock
						.post(this.$httpUrl + '/retailerapi/createRetailer', { retailerArray: retailerArray })
						.then(res => {
							if (res.status === 200) {
								const info = {
									cropsId:row.cropsId,
									factoryId:this.$store.getters.deptId,
								}
								updateReceiveStatus(info).then(() => {
									this.msgSuccess("签收成功");
									this.getList();
								}).catch(err => {
									this.msgError("签收失败")
								})
							}
						})
						.catch(err => {
							this.msgError('存储异常 ' + err);
						});
			})
				.catch(() => {
					this.$message({
						type: 'info',
						message: '取消'
					});
				});
		},

		addNoticeTrasport(){
			this.trasportForm.cropsId = this.checkInfo.cropsId;
			this.trasportForm.farmerUserName = this.$store.getters.name;
			this.trasportForm.farmerNickName = this.$store.getters.nickName;
			this.trasportForm.time = new Date();
			this.trasportForm.driverId = this.trasportForm.userName;
			this.trasportForm.status = 0;
			this.trasportForm.farmerTel = this.$store.getters.phonenumber;
			this.trasportForm.farmerDept = this.$store.getters.deptName;
			this.trasportForm.factoryId = this.trasportForm.deptId
			addTransport(this.trasportForm)
				.then(res => {
					this.msgSuccess('联系运输成功');
					this.noticeDetaiDialog = false;
				})
				.catch(err => {
					this.msgError('联系运输失败，发生异常');
				});
		},

		noticeTransport(row){
			this.checkInfo = row;
			this.noticeDetaiDialog = true
			getAllDriverByDeptId(this.$store.getters.deptId)
				.then(res => {
					this.driverList = res.data;
				})
				.catch(err => {});
			getFactoryByDeptId(123)
				.then(res => {
					this.factoryList = res.data;
				})
		},

		/**
		 * 出库
		 */
		outFactory(row){
			this.checkInfo = row;
			this.open = true
		},

		createMaching(){
		    const checkInfoArray = []
			const id = new this.$snowFlakeId().generate();
			checkInfoArray.push(id)
			checkInfoArray.push(id)
			checkInfoArray.push(this.$store.getters.nickName)
			checkInfoArray.push(this.$store.getters.phonenumber);
			checkInfoArray.push(this.$store.getters.deptName);
			checkInfoArray.push(this.checkInfo.cropsId);
			checkInfoArray.push(this.form.testingResult);
			checkInfoArray.push(this.checkInfo.time);
			checkInfoArray.push(formatDate(new Date()));
			checkInfoArray.push(this.TestingPhotoUrl);
			checkInfoArray.push(this.form.remarks);
			this.$httpBlock
				.post(this.$httpUrl + '/materialapi/createMaching', { checkInfoArray: checkInfoArray })
				.then(res => {
					if (res.status === 200) {
						changeInToOut(this.checkInfo.cropsId).then(res => {
							this.msgSuccess("出库成功")
							this.getList();
						}).catch(err => {
							this.msgError("出库失败")
						})
						this.open = false;
					}
				})
				.catch(err => {
					this.msgError('存储异常 ' + err);
				});
		},

		getFile(file) {
			this.imageUrl = URL.createObjectURL(file.raw);
			this.getBase64(file.raw).then(res => {
				this.imgBase64 = res;
				const datas = {
					imageBase64: this.imgBase64
				};
				uplodImagesBase64(datas)
					.then(res => {
						this.TestingPhotoUrl = res.msg;
						this.msgSuccess('图片上传成功');
					})
					.catch(err => {
						this.msgError('图片上传失败');
					});
			});
		},

		getBase64(file) {
			return new Promise(function(resolve, reject) {
				const reader = new FileReader();
				let imgResult = '';
				reader.readAsDataURL(file);
				reader.onload = function() {
					imgResult = reader.result;
				};
				reader.onerror = function(error) {
					reject(error);
				};
				reader.onloadend = function() {
					resolve(imgResult);
				};
			});
		},


		getList() {
			queryCropsList(this.$store.getters.deptId).then(res => {
				this.transportList = res.data;
			});
		},
		/** 转换部门数据结构 */
		normalizer(node) {
			if (node.children && !node.children.length) {
				delete node.children;
			}
			return {
				id: node.deptId,
				label: node.deptName,
				children: node.children
			};
		},
		// 字典状态字典翻译
		statusFormat(row, column) {
			return this.selectDictLabel(this.statusOptions, row.status);
		},
		// 取消按钮
		cancel() {
			this.open = false;
			this.reset();
		},
		// 表单重置
		reset() {
			this.form = {
				deptId: undefined,
				parentId: undefined,
				deptName: undefined,
				orderNum: undefined,
				leader: undefined,
				phone: undefined,
				email: undefined,
				status: '0'
			};
			this.resetForm('form');
		},
		/** 搜索按钮操作 */
		handleQuery() {
			this.getList();
		},
		/** 重置按钮操作 */
		resetQuery() {
			this.resetForm('queryForm');
			this.handleQuery();
		},
		/** 新增按钮操作 */
		handleAdd(row) {
			this.reset();
			if (row != undefined) {
				this.form.parentId = row.deptId;
			}
			this.open = true;
			this.title = '添加部门';
			listDept().then(response => {
				this.deptOptions = this.handleTree(response.data, 'deptId');
			});
		},
		/** 修改按钮操作 */
		handleUpdate(row) {
			this.reset();
			getDept(row.deptId).then(response => {
				this.form = response.data;
				this.open = true;
				this.title = '修改部门';
			});
			listDeptExcludeChild(row.deptId).then(response => {
				this.deptOptions = this.handleTree(response.data, 'deptId');
			});
		},
		/** 提交按钮 */
		submitForm: function() {
			this.$refs['form'].validate(valid => {
				if (valid) {
					if (this.form.deptId != undefined) {
						updateDept(this.form).then(response => {
							if (response.code === 200) {
								this.msgSuccess('修改成功');
								this.open = false;
								this.getList();
							}
						});
					} else {
						addDept(this.form).then(response => {
							if (response.code === 200) {
								this.msgSuccess('新增成功');
								this.open = false;
								this.getList();
							}
						});
					}
				}
			});
		},
		/** 删除按钮操作 */
		handleDelete(row) {
			this.$confirm('是否确认删除名称为"' + row.deptName + '"的数据项?', '警告', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			})
				.then(function() {
					return delDept(row.deptId);
				})
				.then(() => {
					this.getList();
					this.msgSuccess('删除成功');
				})
				.catch(function() {});
		}
	}
};
</script>
<style>
.avatar-uploader .el-upload {
	border: 1px dashed #d9d9d9;
	border-radius: 6px;
	cursor: pointer;
	position: relative;
	overflow: hidden;
}
.avatar-uploader .el-upload:hover {
	border-color: #409eff;
}
.avatar-uploader-icon {
	font-size: 28px;
	color: #8c939d;
	width: 178px;
	height: 178px;
	line-height: 178px;
	text-align: center;
}
.avatar {
	width: 300px;
	height: 270px;
	display: block;
}

.image {
	width: 80%;
	display: block;
}
</style>
