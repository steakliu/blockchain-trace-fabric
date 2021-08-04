基于Fabric的农产品溯源系统

#### 介绍
    
基于区块链的农产品溯源系统，分为pc端，h5端，后台采用SpringBoot，前端Vue，区块链网络采用Fabric，一共分为六个组织节点，农户，原料厂商，生产厂商，物流厂商，零售商，消费者，(若有什么问题，可提issue， **记得star一下哦**  :triumph:  :triumph:  :triumph:    体验地址：http://116.198.160.39/login

#### 软件架构

此系统有四个模块，已上传Gitee， blockchain-trace-bcnetwork ，blockchain-trace-applets ，blockchain-trace-pc ，blockchain-trace-basic-data。

    blockchain-trace-bcnetwork：区块链网络，可直接将文件上传至服务器，然后启动里面的脚本

    blockchain-trace-applets：小程序端溯源，做得十分简单

    blockchain-trace-pc：PC端,使用的是RuoYi-Vue

    blockchain-trace-basic-data：系统基础数据后台，使用的是RuoYi

    前端：Vue.js , Element UI , mpvue

    后端：SpringBoot , Mybatis , FastDFS , Node.js , Redis , MySQL

    区块链：Fabric1.2-1.4

    智能合约：Golang语言编写

    环境：Ubuntu16.04 64位(2核 4G 1Mbps，本人没钱，只能用乞丐服务器)，Docker 18.09.7 , Docker-compose 1.29.1 , 

#### 安装教程


1.先部署好区块链网络（blockchain-trace-bcnetwork）：
将traceNetwork上传至服务器（也可自己搭建），进入basic-network目录中，启动start.sh脚本（./start.sh）,启动成功后进入webapp目录，启动start.sh脚本（此脚本是安装只能合约，它里面包含了其他几个脚本，可以自己观看），启动成功后看一下docker容器，不出意外的话会安装了6个chaincode,安装成功后执行node enrollAdmin.js 和 node registerUser.js , 生成对应的密钥文件后，最后启动node服务，命令为 node app.js ， 如果需要让其常驻后台，需要安装pm2，然后执行启动 pm2 start app.js  , 停止  pm2 stop app.js ， （环境变量需要有node，npm , golang）
    
2.系统基础数据后端 （blockchain-trace-basic-data）
配置好Redis，MySQL , FastDFS , 端口为8088，（本系统全部所有服务都采用Docker部署）

3.PC端（blockchain-trace-pc）

    npm install --registry=https://registry.npm.taobao.org

    npm run dev

    npm run build:prod

4.小程序（blockchain-trace-applets）
自己使用开发者工具打开

    

#### **说明**

本系统过于简单，只适合做你的毕业设计，课程设计，学习使用，本系统只是提供一个大致的思想，简单的流程，粗糙的实现，
    
#### 界面
区块链浏览器
![区块链结构](https://images.gitee.com/uploads/images/2021/0510/100450_865a1f55_4775150.png "6.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/100501_6d258226_4775150.png "7.png")

基础数据
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/235544_3c60f0e6_4775150.png "a.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/235556_7b27b3c8_4775150.png "b.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/235607_f16f58ac_4775150.png "c.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/235621_c25cd622_4775150.png "d.png")


添加农作物
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/095409_1656ec9f_4775150.png "界面1.png")

农作物链上详情
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/095812_c090b8eb_4775150.png "2.png")

农作物种植过程记录
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/095938_b91c4a9e_4775150.png "3.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/100000_d3f09618_4775150.png "4.png")

物流追踪（因本系统，这里忘记去做了，只是模拟了一下）
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/100223_1c37229e_4775150.png "5.png")

司机运输定位（本系统采用PC端定位，实际上这是不行的，你想想，司机在开车过程中还要去打开浏览器进入系统去定位？这里只是模拟，让老师知道我要表达的意思，因为我也没钱去买设备来定位 :sweat:  :sweat: ）
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/101525_8ec61394_4775150.png "1.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/101537_479fa381_4775150.png "2.png")

原料厂商质检报告上链
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/101838_07d8b55c_4775150.png "3.png")

生产厂商部分截图
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/104012_a3d0a8f3_4775150.png "1.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/104029_3b281f09_4775150.png "2.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/104040_db0dd517_4775150.png "QQ图片20210510103842.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/104051_ded404ae_4775150.png "QQ图片20210510103846.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/104059_b6383ed4_4775150.png "5.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/104107_7bc0ed1e_4775150.png "6.png")


零售商生成产品唯一溯源ID
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/104115_bad29fb0_4775150.png "7.png")

pc溯源
![输入图片说明](https://images.gitee.com/uploads/images/2021/0707/152117_9beee42b_4775150.png "QQ图片20210707152056.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0707/152129_8e55e697_4775150.png "QQ图片20210707152102.png")

小程序端溯源

![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234302_29fe611b_4775150.png "7.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234335_43fbec55_4775150.png "1.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234347_e2572d95_4775150.png "3.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234359_b5bac058_4775150.png "4.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234413_33dd3e47_4775150.png "5.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234428_bc064965_4775150.png "6.png")

#### 参与贡献


#### 特技
