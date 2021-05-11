<template>
	<div style="padding-top: 50px;">
		<l-card type="primary" l-img-class="right-card" position="left" image="https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=154562898,2539807908&fm=15&gp=0.jpg">
			<view class="content">基于区块链的农产品溯源平台，凭借着区块链的不可篡改，去中心化，可追溯特性，让每件农产品从农田到餐桌整个过程透明化</view>
		</l-card>

		<div style="padding-top: 10px;">
			<l-card
				type="primary"
				l-img-class="right-card"
				position="right"
				image="https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=2720209114,2971830483&fm=26&gp=0.jpg"
			>
				<view class="content">
					区块高度:
					<i-tag class="i-tags" name="区块高度" color="blue">{{ blockNumber }}</i-tag>
				</view>
				<br>
				<view class="content">
					参与方:
					<i-tag class="i-tags" name="区块高度" color="green">种植户</i-tag>
					&nbsp;&nbsp;
					<i-tag class="i-tags" name="区块高度" color="blue">原料厂商</i-tag>
					&nbsp;&nbsp;
					<i-tag class="i-tags" name="区块高度" color="red">生产厂商</i-tag>
					&nbsp;&nbsp;
					<i-tag class="i-tags" name="区块高度" color="yellow">零售商</i-tag>
					&nbsp;&nbsp;
					<i-tag class="i-tags" name="区块高度" color="default">物流公司</i-tag>
				</view>
				<br>
				<view class="content">
					技术：
					<i-tag class="i-tags" color="default">Node.js</i-tag>&nbsp;&nbsp;
					<i-tag class="i-tags" color="blue">Vue</i-tag>&nbsp;&nbsp;
					<i-tag class="i-tags" color="red">Hyperledger Fabric</i-tag>&nbsp;&nbsp;
					<i-tag class="i-tags" color="green">Express</i-tag>
				</view>
			</l-card>
		</div>
		
		<div style="padding-top:15px;padding-bottom: 15px;">
			<l-notice-bar :show="true" 
			              type="roll"
						  speed="10000">
			              最新区块:{{blockInfo.data.header.data_hash }}
			</l-notice-bar>
		</div>

		<div style="padding-bottom: 15px;">
			<div><l-search-bar shape="circle" @linconfirm="queryInfo" @linchange="getTraceId" placeholder="溯源ID" :show-cancel="false" /></div>
		</div>

		<!-- <div align="center">
			<l-loading :show="true" type="flip"></l-loading>
		</div> -->
	</div>
</template>

<script>
export default {
	data() {
		return {
			value6: 123,
			value7: 3454,
			cropsInfo: '',
			visible1: false,
			traceId: '',
			retailerInfo: '', //零售点信息
			productInfo: '', //产品信息
			blockNumber: '',
			blockInfo: ''
		};
	},

	methods: {
		getTraceId(trace) {
			console.log('res   ' + JSON.stringify(trace));
			this.traceId = trace.mp.detail.value;
		},

		queryInfo() {
			this.$httpBlock
				.get(this.$httpUrl + '/retailerapi/queryRetailerById?id=' + this.traceId)
				.then(res => {
					this.retailerInfo = res.data; //零售
					this.$httpBlock
						.get(this.$httpUrl + '/productapi/queryProductInfoByCropsId?id=' + res.data.crops_id)
						.then(res => {
							this.productInfo = res.data[0].Record; //产品信息
							const totalData = {
								product_id: this.retailerInfo.product_id,
								crops_id: this.retailerInfo.crops_id,
								product_name: this.productInfo.product_name,
								retailer: this.retailerInfo.retailer,
								retailer_name: this.retailerInfo.retailer_name,
								retailer_tel: this.retailerInfo.retailer_tel,
								factory: this.productInfo.factory,
								Net_Content: this.productInfo.Net_Content,
								workshop: this.productInfo.workshop,
								mixed_ingredients: this.productInfo.mixed_ingredients,
								keep_mathod: this.productInfo.keep_mathod,
								cooking_recommend: this.productInfo.cooking_recommend,
								work_hours: this.productInfo.work_hours,
								leader: this.productInfo.leader,
								leader_tel: this.productInfo.leader_tel,
								remarks: this.productInfo.remarks
							};
							const data = JSON.stringify(totalData);
							wx.navigateTo({
								url: '/pages/product/main?data=' + encodeURIComponent(data)
							});
						})
						.catch(err => {
							this.msgError('查询异常 ' + err);
						});

					wx.navigateTo({
						url: '/pages/product/main?data=' + encodeURIComponent(data)
					});
				})
				.catch(err => {
					console.log('调用失败 ' + JSON.stringify(err));
				});
		}
	},

	created() {
		this.$httpBlock.get(this.$httpUrl + '/blockexplorerapi/channelBlockInfo').then(res => {
			this.blockNumber = res.data.height.low + 10000;
			var number = Number(res.data.height.low - 1);
			this.$httpBlock
				.get(this.$httpUrl + '/blockexplorerapi/queryBlockInfo?number=' + number)
				.then(res => {
					this.blockInfo = res;
				})
				.catch(err => {
					console.log('err ' + err);
				});
		});
	}
};
</script>

<style scoped>
.userinfo {
	display: flex;
	flex-direction: column;
	align-items: center;
}

.userinfo-avatar {
	width: 128rpx;
	height: 128rpx;
	margin: 20rpx;
	border-radius: 50%;
}

.userinfo-nickname {
	color: #aaa;
}

.usermotto {
	margin-top: 150px;
}

.form-control {
	display: block;
	padding: 0 12px;
	margin-bottom: 5px;
	border: 1px solid #ccc;
}
.all {
	width: 7.5rem;
	height: 1rem;
	background-color: blue;
}
.all:after {
	display: block;
	content: '';
	clear: both;
}
.left {
	float: left;
	width: 3rem;
	height: 1rem;
	background-color: red;
}

.right {
	float: left;
	width: 4.5rem;
	height: 1rem;
	background-color: green;
}
</style>
