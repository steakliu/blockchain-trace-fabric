<template>
	<div class="app-container">
    <el-divider>待处理业务</el-divider>

		<el-table v-loading="loading" :data="cropsList" @selection-change="handleSelectionChange">
			<el-table-column type="selection" width="55" align="center" />
			<el-table-column label="id" prop="id" />
			<el-table-column label="作物编号" prop="cropsId" />
			<el-table-column label="原料厂商" prop="farmerDept" />
			<el-table-column label="作物名称" prop="cropsName" />
			<el-table-column label="原料厂商负责人" prop="farmerNickName" />
			<el-table-column label="加工状态" prop="maching">
				<template slot-scope="scope">
					<el-tag v-show="scope.row.machingStatus === null">未加工</el-tag>
					<el-tag type="danger" v-show="scope.row.machingStatus === 1">加工中(任务已经分发)</el-tag>
					<el-tag type="warning" v-show="scope.row.machingStatus === 2">加工完成</el-tag>
				</template>
			</el-table-column>
			<el-table-column label="库状态" prop="out_factory_status">
				<template slot-scope="scope">
					<el-tag v-show="scope.row.outFactoryStatus === null">未出库</el-tag>
					<el-tag v-show="scope.row.outFactoryStatus === 1">已出库</el-tag>
				</template>
			</el-table-column>
			<el-table-column label="操作" align="center" class-name="small-padding fixed-width">
				<template slot-scope="scope">
					<el-button v-show="scope.row.machingStatus !== 1 && scope.row.machingStatus !== 2" size="mini" type="text" @click="taskDistribution(scope.row)">任务分发</el-button>
					<el-button v-show="scope.row.machingStatus !== null" size="mini" type="text" @click="queryTaskByCropsId(scope.row)">查看任务分发信息</el-button>
					<el-button v-show="scope.row.productWriteStatus !== 1" size="mini" type="text" @click="addThing(scope.row)">填写产品基本信息</el-button>
					<el-button v-show="scope.row.productWriteStatus !== null" size="mini" type="text" @click="queryProductInfo(scope.row)">查看产品基本信息</el-button>
					<el-button size="mini" type="text" @click="staffWork(scope.row)">员工工作量</el-button>
					<el-button size="mini" type="text" @click="noticeTransport(scope.row)">通知运输</el-button>
					<el-button v-show="scope.row.outFactoryStatus === null" size="mini" type="text" @click="outFactory(scope.row)">产品出库</el-button>
				</template>
			</el-table-column>
		</el-table>

		<pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

		<el-dialog center title="任务分发" :visible.sync="open" width="700px" append-to-body>
			<el-form ref="form" :model="form" label-width="80px">
				<el-form-item label="任务标题"><el-input v-model="form.taskName" placeholder="请输入任务标题"></el-input></el-form-item>
				<el-form-item label="任务分工内容"><el-input v-model="form.taskRemark" type="textarea" placeholder="请输入内容"></el-input></el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button type="primary" @click="addTask">确 定</el-button>
				<el-button @click="cancel">取 消</el-button>
			</div>
		</el-dialog>

		<el-dialog center title="生产基本信息添加" :visible.sync="addThingDialog" width="700px" append-to-body>
			<el-form ref="form" label-width="80px" :model="recordForm">
			<el-row>
				<el-col :span="12">
					<el-form-item label="品名" prop="nickName"><el-input v-model="recordForm.ProductName" placeholder="请输入" /></el-form-item>
				</el-col>
				<el-col :span="12">
					<el-form-item label="配料" prop="nickName"><el-input v-model="recordForm.MixedIngredients" placeholder="请输入" /></el-form-item>
				</el-col>
			</el-row>
			<el-row>
				<el-col :span="12">
					<el-form-item label="负责人" prop="nickName"><el-input v-model="recordForm.Leader" placeholder="请输入" /></el-form-item>
				</el-col>
				<el-col :span="12">
					<el-form-item label="电话" prop="nickName"><el-input v-model="recordForm.LeaderTel" placeholder="请输入" /></el-form-item>
				</el-col>
			</el-row>
			<el-row>
				<el-col :span="12">
					<el-form-item label="生产车间" prop="nickName"><el-input v-model="recordForm.Workshop" placeholder="请输入" /></el-form-item>
				</el-col>
				<el-col :span="12">
					<el-form-item label="生产工时" prop="nickName"><el-input v-model="recordForm.WorkHours" placeholder="请输入" /></el-form-item>
				</el-col>
			</el-row>

			<el-row>
				<el-col :span="12">
					<el-form-item label="储藏方法" prop="nickName">
						 <el-input v-model="recordForm.KeepMethod" placeholder="请输入" />
					</el-form-item>
				</el-col>
				<el-col :span="12">
					<el-form-item label="净含量" prop="nickName">
						 <el-input v-model="recordForm.NetContent" placeholder="请输入" />
					</el-form-item>
				</el-col>
			</el-row>
			<el-form-item label="食用推荐"><el-input v-model="recordForm.CookingRecommend" type="textarea" placeholder="请输入内容"></el-input></el-form-item>
			<el-form-item label="备注"><el-input v-model="recordForm.Remarks" type="textarea" placeholder="请输入内容"></el-input></el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button type="primary" @click="submitProductInfo">确 定</el-button>
				<el-button @click="cancel">取 消</el-button>
			</div>
		</el-dialog>

		<el-dialog center title="任务下发详情" :visible.sync="taskInfoDialog" width="700px" append-to-body>
			<el-row>
				<el-col :span="12">任务ID：{{ taskInfo.taskId }}</el-col>
				<el-col :span="12">任务名称：{{ taskInfo.taskName }}</el-col>
			</el-row>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="12">作物ID：{{ taskInfo.cropsId }}</el-col>
				<el-col :span="12">发布时间：{{ taskInfo.time }}</el-col>
			</el-row>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="12">状态：{{ taskInfo.status }}</el-col>
			</el-row>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="24">详情内容：{{ taskInfo.taskRemark }}</el-col>
			</el-row>
		</el-dialog>

		<el-dialog center title="产品基本信息" :visible.sync="productInfoDialog" width="700px" append-to-body>
			<el-row>
				<el-col :span="24">作物ID：{{ productInfos.crops_id }}</el-col>
			</el-row>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="12">产品ID：{{ productInfos.product_id }}</el-col>
				<el-col :span="12">产品名称：{{ productInfos.product_name }}</el-col>
			</el-row>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="12">净含量：{{ productInfos.Net_Content }}</el-col>
				<el-col :span="12">保存方式：{{ productInfos.keep_mathod }}</el-col>
			</el-row>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="12">负责人：{{ productInfos.leader }}</el-col>
				<el-col :span="12">负责人电话：{{ productInfos.leader }}</el-col>
			</el-row>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="12">配料：{{ productInfos.mixed_ingredients }}</el-col>
				<el-col :span="12">工时：{{ productInfos.work_hours }}</el-col>
			</el-row>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="12">车间：{{ productInfos.workshop }}</el-col>
				<el-col :span="12">数据上链时间：{{ productInfos.time }}</el-col>
			</el-row>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="12">食用推荐：{{ productInfos.cooking_recommend }}</el-col>
				<el-col :span="12">备注：{{ productInfos.remarks }}</el-col>
			</el-row>
		</el-dialog>

		<el-dialog center title="员工工作量证明" :visible.sync="staffWorkDialog" width="1200px" append-to-body>
			<el-row :gutter="8" >
				<el-col :span="6" v-for="(info,index) in operationList" :key="index" style="padding-top: 10px;">
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
						<el-form-item label="零售商" prop="nickName">
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

	</div>
</template>

<script>
import { listRole, getRole, delRole, addRole, updateRole, exportRole, dataScope, changeRoleStatus } from '@/api/system/role';
import { treeselect as menuTreeselect, roleMenuTreeselect } from '@/api/system/menu';
import { treeselect as deptTreeselect, roleDeptTreeselect } from '@/api/system/dept';
import {productOutFactory,updateProductWriteStatus,queryTaskByCropsId, queryCropsList,addTask,updateMachingStatus } from '../../../api/trace/product';
import { formatDate } from '../../../utils/index.js';
import {getFactoryByDeptId,getAllDriverByDeptId,addTransport} from '../../../api/trace/farmer';
export default {
	name: 'Role',

	data() {
		return {
			// 遮罩层
			loading: true,
			// 选中数组
			ids: [],
			// 非单个禁用
			single: true,
			// 非多个禁用
			multiple: true,
			// 显示搜索条件
			showSearch: true,
			// 总条数
			total: 0,
			// 角色表格数据
			cropsList: [],
			// 弹出层标题
			title: '',
			// 是否显示弹出层
			open: false,
			// 是否显示弹出层（数据权限）
			openDataScope: false,
			menuExpand: false,
			menuNodeAll: false,
			deptExpand: true,
			deptNodeAll: false,
			// 日期范围
			dateRange: [],
			// 状态数据字典
			statusOptions: ['在种', '停种'],

			beggingOptions: [],
			plantModeOptions: [], //种植方式
			cropsTypeOptions: [], //作物类型

			growDialog: false, //z作物生长记录弹出框
			cropsName: '',
			cropsDetaiDialog: false,

			cropsDetails: '',
			// 菜单列表
			menuOptions: [],
			// 部门列表
			deptOptions: [],
			// 查询参数
			queryParams: {
				pageNum: 1,
				pageSize: 10,
				roleName: undefined,
				roleKey: undefined,
				status: undefined
			},
			// 表单参数
			form: {},
			defaultProps: {
				children: 'children',
				label: 'label'
			},
			recordForm: {},
			imageUrl: '',
			imgBase64: [],
			cropsPhotoUrl: '',
			cropsInfo: '',
			growDialog: false,
			cropsProcessDetailsArray: [],
			cropsProcessDetaiDialog: false,
			noticeDetaiDialog: false,
			driverList: [],
			trasportForm: {
				cropsId: '',
				farmerUserName: '',
				farmerNickName: '',
				time: ''
			},
			trasportInfo: '',
			playTrackView: false,
			addThingDialog:false,
			taskInfo:'',
			taskInfoDialog:false,
			cropsInfos:'',
			productInfoDialog:false,
			productInfos:'',
			staffWorkDialog:false,
			operationList:[],
			driverList:[],
			factoryList:[],
		};
	},
	created() {
		this.getCropsList();
		this.getDicts('crops_bagging').then(response => {
			this.beggingOptions = response.data;
		});
		this.getDicts('crops_plant_type').then(res => {
			this.plantModeOptions = res.data;
		});
		this.getDicts('crops_type').then(res => {
			this.cropsTypeOptions = res.data;
		});
	},
	methods: {
		outFactory(row){
			this.$confirm('确定将ID为【'+row.cropsId+'】的产品出库吗', '提示', {
			          confirmButtonText: '确定',
			          cancelButtonText: '取消',
			          type: 'warning'
			        }).then(() => {
						productOutFactory(row.cropsId).then(() => {
							this.getCropsList();
							this.$message({
							  type: 'success',
							  message: '已出库!'
							});
						})

			        }).catch(() => {
			          this.$message({
			            type: 'info',
			            message: '已取消出库'
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
			getFactoryByDeptId(125)
				.then(res => {
					this.factoryList = res.data;
				})
				.catch(err => {});
		},

		staffWork(row){
			this.staffWorkDialog = true;
			this.$httpBlock
			    	.get(this.$httpUrl + '/productapi/queryOperationByCropsId?id='+row.cropsId)
			    	.then(res => {
						const array = [];
						for (let i = 0; i < res.data.length; i++) {
							array.push(res.data[i].Record);
						}
						this.operationList = array;
			    	})
			    	.catch(err => {
			    		this.msgError('查询异常 ' + err);
			    	});
		},

		queryProductInfo(row){
			this.productInfoDialog = true
			this.$httpBlock
			    	.get(this.$httpUrl + '/productapi/queryProductInfoByCropsId?id='+row.cropsId)
			    	.then(res => {
						this.productInfos = res.data[0].Record;
			    	})
			    	.catch(err => {
			    		this.msgError('查询异常 ' + err);
			    	});
		},

		submitProductInfo(){
			const productInfoArray = [];
			const id = new this.$snowFlakeId().generate();
			productInfoArray.push(id)
			productInfoArray.push(id)
			productInfoArray.push(this.cropsInfos.cropsId)
			productInfoArray.push(this.recordForm.ProductName)
			productInfoArray.push(this.recordForm.MixedIngredients)
			productInfoArray.push(this.recordForm.Leader)
			productInfoArray.push(this.recordForm.LeaderTel)
			productInfoArray.push(this.recordForm.Workshop)
			productInfoArray.push(this.$store.getters.deptName)
			productInfoArray.push(this.recordForm.WorkHours)
			productInfoArray.push(formatDate(new Date()))
			productInfoArray.push(this.recordForm.KeepMethod)
			productInfoArray.push(this.recordForm.NetContent)
			productInfoArray.push(this.recordForm.CookingRecommend)
			productInfoArray.push(this.recordForm.Remarks)
	        this.$httpBlock
	        	.post(this.$httpUrl + '/productapi/createProductInfo', { productInfoArray: productInfoArray })
	        	.then(res => {
	        		if (res.status === 200) {
						updateProductWriteStatus(this.cropsInfos.cropsId).then(res => {
							this.addThingDialog = false;
							this.msgSuccess("数据上链成功");
							this.getCropsList();
						}).catch(err => {
							this.msgError("出现异常");
						})
	        		}
	        	})
	        	.catch(err => {
	        		this.msgError('存储异常 ' + err);
	        	});
		},

		/**
		 * 根据cropsId查询任务详情
		 */
		queryTaskByCropsId(row){
			queryTaskByCropsId(row.cropsId).then(res => {
				this.taskInfo = res.data
				this.taskInfoDialog = true
			}).catch(err => {
				this.msgError("查询异常")
			})
		},

		addThing(row){
			this.cropsInfos = row
			this.addThingDialog = true
		},

		/**
		 * 确定发布
		 */
		addTask(){
			this.form.taskId = new this.$snowFlakeId().generate();
			this.form.cropsId = this.productInfo.cropsId;
			this.form.leaderId = this.$store.getters.name;
			this.form.time = formatDate(new Date());
			this.form.leaderName = this.$store.getters.nickName;
			this.form.status = 0;
			addTask(this.form).then(res => {
				const info = {
					cropsId:this.productInfo.cropsId,
					machingStatus:1
				}
				updateMachingStatus(info).then(res => {
					this.msgSuccess("产品进入加工状态")
					this.getCropsList();
				}).catch(err => {

				})
				this.msgSuccess("发布任务成功");
				this.open = false;
			}).catch(err => {
				this.msgError("添加任务出错")
			})
		},
		/**
		 * 任务分发
		 */
		taskDistribution(row){
			this.productInfo = row
			this.open = true
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
						this.cropsPhotoUrl = res.msg;
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


		handleRecord(row) {
			this.cropsName = row.cropsName;
			this.cropsInfo = row;
			this.growDialog = true;
		},

		/** 查询作物列表 */
		getCropsList() {
			this.loading = true;
			queryCropsList(this.$store.getters.deptId).then(response => {
				this.cropsList = response.data;
				this.total = response.total;
				this.loading = false;
			});
		},

		/**
		 * 作物详情
		 */
		cropsDetail(row) {
			this.$httpBlock
				.get(this.$httpUrl + '/farmerapi/queryCropsById?id=' + row.cropsId)
				.then(res => {
					this.cropsDetails = res.data;
					this.cropsDetaiDialog = true;
				})
				.catch(err => {});
		},

		/** 查询菜单树结构 */
		getMenuTreeselect() {
			menuTreeselect().then(response => {
				this.menuOptions = response.data;
			});
		},
		/** 查询部门树结构 */
		getDeptTreeselect() {
			deptTreeselect().then(response => {
				this.deptOptions = response.data;
			});
		},
		// 所有菜单节点数据
		getMenuAllCheckedKeys() {
			// 目前被选中的菜单节点
			let checkedKeys = this.$refs.menu.getHalfCheckedKeys();
			// 半选中的菜单节点
			let halfCheckedKeys = this.$refs.menu.getCheckedKeys();
			checkedKeys.unshift.apply(checkedKeys, halfCheckedKeys);
			return checkedKeys;
		},
		// 所有部门节点数据
		getDeptAllCheckedKeys() {
			// 目前被选中的部门节点
			let checkedKeys = this.$refs.dept.getHalfCheckedKeys();
			// 半选中的部门节点
			let halfCheckedKeys = this.$refs.dept.getCheckedKeys();
			checkedKeys.unshift.apply(checkedKeys, halfCheckedKeys);
			return checkedKeys;
		},
		/** 根据角色ID查询菜单树结构 */
		getRoleMenuTreeselect(roleId) {
			return roleMenuTreeselect(roleId).then(response => {
				this.menuOptions = response.menus;
				return response;
			});
		},
		/** 根据角色ID查询部门树结构 */
		getRoleDeptTreeselect(roleId) {
			return roleDeptTreeselect(roleId).then(response => {
				this.deptOptions = response.depts;
				return response;
			});
		},

		// 取消按钮
		cancel() {
			this.open = false;
			this.reset();
		},
		// 取消按钮（数据权限）
		cancelDataScope() {
			this.openDataScope = false;
			this.reset();
		},
		// 表单重置
		reset() {
			if (this.$refs.menu != undefined) {
				this.$refs.menu.setCheckedKeys([]);
			}
			(this.menuExpand = false),
				(this.menuNodeAll = false),
				(this.deptExpand = true),
				(this.deptNodeAll = false),
				(this.form = {
					roleId: undefined,
					roleName: undefined,
					roleKey: undefined,
					roleSort: 0,
					status: '0',
					menuIds: [],
					deptIds: [],
					menuCheckStrictly: true,
					deptCheckStrictly: true,
					remark: undefined
				});
			this.resetForm('form');
		},
		/** 搜索按钮操作 */
		handleQuery() {
			this.queryParams.pageNum = 1;
			this.getList();
		},
		/** 重置按钮操作 */
		resetQuery() {
			this.dateRange = [];
			this.resetForm('queryForm');
			this.handleQuery();
		},
		// 多选框选中数据
		handleSelectionChange(selection) {
			this.ids = selection.map(item => item.roleId);
			this.single = selection.length != 1;
			this.multiple = !selection.length;
		},
		// 树权限（展开/折叠）
		handleCheckedTreeExpand(value, type) {
			if (type == 'menu') {
				let treeList = this.menuOptions;
				for (let i = 0; i < treeList.length; i++) {
					this.$refs.menu.store.nodesMap[treeList[i].id].expanded = value;
				}
			} else if (type == 'dept') {
				let treeList = this.deptOptions;
				for (let i = 0; i < treeList.length; i++) {
					this.$refs.dept.store.nodesMap[treeList[i].id].expanded = value;
				}
			}
		},
		// 树权限（全选/全不选）
		handleCheckedTreeNodeAll(value, type) {
			if (type == 'menu') {
				this.$refs.menu.setCheckedNodes(value ? this.menuOptions : []);
			} else if (type == 'dept') {
				this.$refs.dept.setCheckedNodes(value ? this.deptOptions : []);
			}
		},
		// 树权限（父子联动）
		handleCheckedTreeConnect(value, type) {
			if (type == 'menu') {
				this.form.menuCheckStrictly = value ? true : false;
			} else if (type == 'dept') {
				this.form.deptCheckStrictly = value ? true : false;
			}
		},
		/** 新增按钮操作 */
		handleAdd() {
			this.reset();
			this.getMenuTreeselect();
			this.open = true;
			this.title = '记录作物';
		},
		/** 修改按钮操作 */
		handleUpdate(row) {
			this.reset();
			const roleId = row.roleId || this.ids;
			const roleMenu = this.getRoleMenuTreeselect(roleId);
			getRole(roleId).then(response => {
				this.form = response.data;
				this.open = true;
				this.$nextTick(() => {
					roleMenu.then(res => {
						this.$refs.menu.setCheckedKeys(res.checkedKeys);
					});
				});
				this.title = '修改作物';
			});
		},
		/** 分配数据权限操作 */
		handleDataScope(row) {
			this.reset();
			const roleDeptTreeselect = this.getRoleDeptTreeselect(row.roleId);
			getRole(row.roleId).then(response => {
				this.form = response.data;
				this.openDataScope = true;
				this.$nextTick(() => {
					roleDeptTreeselect.then(res => {
						this.$refs.dept.setCheckedKeys(res.checkedKeys);
					});
				});
				this.title = '分配数据权限';
			});
		},
		/** 提交按钮 */
		submitForm: function() {
			// this.$refs["form"].validate(valid => {
			//     if (valid) {
			saveCrops(this.form).then(res => {
				if (res.code === 200) {
					this.msgSuccess('修改成功');
					this.open = false;
				}
			});
			// if (this.form.roleId != undefined) {
			//     this.form.menuIds = this.getMenuAllCheckedKeys();
			//     updateRole(this.form).then(response => {
			//         if (response.code === 200) {
			//             this.msgSuccess("修改成功");
			//             this.open = false;
			//             this.getList();
			//         }
			//     });
			// } else {
			//     this.form.menuIds = this.getMenuAllCheckedKeys();
			//     addRole(this.form).then(response => {
			//         if (response.code === 200) {
			//             this.msgSuccess("新增成功");
			//             this.open = false;
			//             this.getList();
			//         }
			//     });
			// }
			//     }
			// });
		},
		/** 提交按钮（数据权限） */
		submitDataScope: function() {
			if (this.form.roleId != undefined) {
				this.form.deptIds = this.getDeptAllCheckedKeys();
				dataScope(this.form).then(response => {
					if (response.code === 200) {
						this.msgSuccess('修改成功');
						this.openDataScope = false;
						this.getList();
					}
				});
			}
		},
		/** 删除按钮操作 */
		handleDelete(row) {
			const roleIds = row.roleId || this.ids;
			this.$confirm('是否确认删除角色编号为"' + roleIds + '"的数据项?', '警告', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			})
				.then(function() {
					return delRole(roleIds);
				})
				.then(() => {
					this.getList();
					this.msgSuccess('删除成功');
				})
				.catch(function() {});
		},
		/** 导出按钮操作 */
		handleExport() {
			const queryParams = this.queryParams;
			this.$confirm('是否确认导出所有角色数据项?', '警告', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			})
				.then(function() {
					return exportRole(queryParams);
				})
				.then(response => {
					this.download(response.msg);
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
	width: 178px;
	height: 178px;
	display: block;
}

.image {
	width: 80%;
	display: block;
}
/**一定要给宽高，否则不显示*/
#position {
	height: 100%;
}
/**修改点标记图片的css, 图片给宽高，否则不显示 */
.map-page /deep/ {
	.amap-icon {
		width: 20px;
	}
	.amap-marker-label {
		border: 1px solid #ccc;
		font-size: 16px;
		display: inline-block;
		padding: 5px;
	}
}

#panel {
	position: fixed;
	background-color: white;
	max-height: 90%;
	overflow-y: auto;
	top: 10px;
	right: 10px;
	width: 280px;
}
#panel .amap-call {
	background-color: #009cf9;
	border-top-left-radius: 4px;
	border-top-right-radius: 4px;
}
#panel .amap-lib-driving {
	border-bottom-left-radius: 4px;
	border-bottom-right-radius: 4px;
	overflow: hidden;
}
html,
body,
#container {
	width: 100%;
	height: 100%;
}

.map {
	width: 100%;
	height: 500px;
}
</style>
