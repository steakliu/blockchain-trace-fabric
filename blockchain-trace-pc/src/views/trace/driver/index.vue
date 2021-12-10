<template>
	<div class="app-container">
		<el-divider>待处理业务</el-divider>

		<br />
		<el-table :data="transportList">
			<el-table-column type="selection" width="55" align="center" />
			<el-table-column label="货物编号" prop="cropsId" />
			<el-table-column label="通知人" prop="farmerNickName" />
			<el-table-column label="部门" prop="farmerDept" />
			<el-table-column label="时间" prop="time" />
			<el-table-column label="备注" prop="remarks" />
			<el-table-column label="状态" prop="status">
				<template slot-scope="scope">
					<el-tag v-if="scope.row.status === 0">未运输</el-tag>
					<el-tag v-if="scope.row.status === 1">运输中</el-tag>
					<el-tag v-if="scope.row.status === 2">运输完成</el-tag>
				</template>
			</el-table-column>
			<el-table-column label="操作" align="center" class-name="small-padding fixed-width">
				<template slot-scope="scope">
					<el-button v-show="scope.row.status === 0" size="mini" type="text" @click="startTransport(scope.row)">开始配送</el-button>
					<el-button v-show="scope.row.status === 1" size="mini" type="text" @click="transportLocation(scope.row)">中途定位</el-button>
					<el-button v-show="scope.row.status === 1" size="mini" type="text" @click="isOk(scope.row)">配送完毕</el-button>
<!--					<el-button size="mini" type="text" @click="isOk(scope.row)">物流追踪</el-button>-->
				</template>
			</el-table-column>
		</el-table>

		<div v-show="false" id="tip"></div>
		<div v-show="false" id="lng"></div>
		<div v-show="false" id="lat"></div>
	</div>
</template>

<script>
import { listDept, getDept, delDept, addDept, updateDept, listDeptExcludeChild } from '@/api/system/dept';
import { listTransport, saveAndUpdateTransportInfoToDb ,updateTransportStatus} from '../../../api/trace/driver';
import Treeselect from '@riophae/vue-treeselect';
import '@riophae/vue-treeselect/dist/vue-treeselect.css';
import { formatDate } from '../../../utils/index.js';
export default {
	name: 'Dept',
	components: { Treeselect },
	data() {
		return {
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
			lat: Number
		};
	},
	created() {
		this.getList();
	},

	mounted() {
		AMap.plugin('AMap.Geolocation', function() {
			var geolocation = new AMap.Geolocation({
				// 是否使用高精度定位，默认：true
				enableHighAccuracy: true,
				// 设置定位超时时间，默认：无穷大
				timeout: 10000,
				// 定位按钮的停靠位置的偏移量，默认：Pixel(10, 20)
				buttonOffset: new AMap.Pixel(10, 20),
				//  定位成功后调整地图视野范围使定位位置及精度范围视野内可见，默认：false
				zoomToAccuracy: true,
				//  定位按钮的排放位置,  RB表示右下
				buttonPosition: 'RB'
			});
			geolocation.getCurrentPosition();
			AMap.event.addListener(geolocation, 'complete', onComplete);
			AMap.event.addListener(geolocation, 'error', onError);
			function onComplete(data) {
				this.cityDetail = data.addressComponent.province + data.addressComponent.city + data.addressComponent.district + data.addressComponent.township;
				this.lng = data.position.lng;
				this.lat = data.position.lat;
				const str = [];
				str.push(data.position.lng);
				str.push(data.position.lat);
				document.getElementById('lng').innerHTML = data.position.lng;
				document.getElementById('lat').innerHTML = data.position.lat;
				document.getElementById('tip').innerHTML = this.cityDetail;
			}
			function onError(data) {
				console.log('定位错误   ' + JSON.stringify(data));
			}
		});
	},

	methods: {
		isOk(row){
			const transportArray = [];
			const id = new this.$snowFlakeId().generate();
			transportArray.push(id);
			transportArray.push(id);
			transportArray.push(this.$store.getters.name);
			transportArray.push(this.$store.getters.nickName);
			transportArray.push(this.$store.getters.phonenumber);
			transportArray.push(this.$store.getters.deptName);
			transportArray.push(formatDate(new Date()));
			transportArray.push(document.getElementById('tip').innerHTML);
			transportArray.push(row.cropsId);
			transportArray.push('终点站');
			this.$confirm('是否确认配送完毕', '配送', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			})
				.then(() => {
					this.$httpBlock
						.post(this.$httpUrl + '/driverapi/createTransport', { transportArray: transportArray })
						.then(res => {
							if (res.status === 200) {
								this.msgSuccess('数据上链成功');
								this.growDialog = false;
								const transportInfo = {
									lng: Number(document.getElementById('lng').innerHTML),
									lat: Number(document.getElementById('lat').innerHTML),
									cropsId: row.cropsId,
									status: 2,
									id: row.id,
									outFactoryStatus:0
								};
								saveAndUpdateTransportInfoToDb(transportInfo)
									.then(res => {
										this.getList();
									})
									.catch(err => {

									});
							}
						})
						.catch(err => {
							this.msgError('存储异常 ' + err);
						});
					this.$message({
						type: 'success',
						message: '配送成功!'
					});
				})
				.catch(() => {
					this.$message({
						type: 'info',
						message: '取消'
					});
				});
		},

		transportLocation(row) {
			const transportArray = [];
			const id = new this.$snowFlakeId().generate();
			transportArray.push(id);
			transportArray.push(id);
			transportArray.push(this.$store.getters.name);
			transportArray.push(this.$store.getters.nickName);
			transportArray.push(this.$store.getters.phonenumber);
			transportArray.push(this.$store.getters.deptName);
			transportArray.push(formatDate(new Date()));
			transportArray.push(document.getElementById('tip').innerHTML);
			transportArray.push(row.cropsId);
			transportArray.push('中途定位');
			this.$confirm(document.getElementById('tip').innerHTML, '定位', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			})
				.then(() => {
					this.$httpBlock
						.post(this.$httpUrl + '/driverapi/createTransport', { transportArray: transportArray })
						.then(res => {
							if (res.status === 200) {
								this.msgSuccess('数据上链成功');
								this.growDialog = false;
								const transportInfo = {
									lng: Number(document.getElementById('lng').innerHTML),
									lat: Number(document.getElementById('lat').innerHTML),
									cropsId: row.cropsId,
									status: 1,
								};
								saveAndUpdateTransportInfoToDb(transportInfo)
									.then(res => {
										this.getList();
									})
									.catch(err => {

									});
							}
						})
						.catch(err => {
							this.msgError('存储异常 ' + err);
						});

				})
				.catch(() => {
					this.$message({
						type: 'info',
						message: '取消定位'
					});
				});
		},
		/**
		 * @param {Object} row开始配送
		 */
		startTransport(row) {
			const transportArray = [];
			const id = new this.$snowFlakeId().generate();
			transportArray.push(id);
			transportArray.push(id);
			transportArray.push(this.$store.getters.name);
			transportArray.push(this.$store.getters.nickName);
			transportArray.push(this.$store.getters.phonenumber);
			transportArray.push(this.$store.getters.deptName);
			transportArray.push(formatDate(new Date()));
			transportArray.push(document.getElementById('tip').innerHTML);
			transportArray.push(row.cropsId);
			transportArray.push('始发地');
			this.$confirm('是否开始运送货物', '提示', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			})
				.then(() => {
					this.$httpBlock
						.post(this.$httpUrl + '/driverapi/createTransport', { transportArray: transportArray })
						.then(res => {
							if (res.status === 200) {
								this.msgSuccess('数据上链成功');
								this.growDialog = false;
								const transportInfo = {
									lng: Number(document.getElementById('lng').innerHTML),
									lat: Number(document.getElementById('lat').innerHTML),
									cropsId: row.cropsId,
									id: row.id,
									status: 1,
								};
								saveAndUpdateTransportInfoToDb(transportInfo)
									.then(res => {

										this.getList();
									})
									.catch(err => {

									});
							}
						})
						.catch(err => {
							this.msgError('存储异常 ' + err);
						});
					this.$message({
						type: 'success',
						message: '开始运送!'
					});
				})
				.catch(() => {
					this.$message({
						type: 'info',
						message: '已取消运送'
					});
				});
		},

		getList() {
			listTransport(this.$store.getters.name).then(res => {
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
