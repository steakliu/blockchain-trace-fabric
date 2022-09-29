# 基于Fabric的农产品溯源系统  

>让人类餐桌上的食物回归到农耕时代！

# 介绍  

>随着社会的不断发展，科技的不断进步，到了今天，你是否还记得故乡地里的那片只施农家肥的菜园，在小时候，物质虽然相对贫乏，但是吃到的
每一口在长大后再也没吃过那么香的，如今，为了生活，离开深爱的故乡，蜗居在城市的钢筋混泥土下，你是否会偶尔回味起哪些遗失的味道。

基于区块链的农产品溯源系统，分为pc端，h5端，一共分为五个组织节点，农户，原料厂商，生产厂商，物流厂商，零售商，在全流程走完后，消费者买到产品后，可对其
进行溯源，每个环节的数据都会进行上链，所以整个过程是透明的。  

本系统提供的是一种区块链应用的思想，因为在大部分人的认知里，区块链只是就是比特币，就是割韭菜，其实这是不对的，区块链使用到传统行业中，可能会有颠覆性的
作用，比如农产品，供应链溯源，电子存证，司法行业，金融行业，实物赋能，教育档案管理，医疗，公证等等场景和行业中，区块链都有其天然优势，其不可篡改，去中心化，
分布式，可追溯等特性，能在很多场景发挥作用。  

![输入图片说明](install-fabric-env/image.png)

# 系统模块和技术  

此系统有四个模块， blockchain-trace-bcnetwork ，blockchain-trace-applets ，blockchain-trace-pc ，blockchain-trace-basic-data，
本系统只做了一个数据上链的操作，比较简单，因为技术其实我们都是能想办法去实现的，但是产品思维是很珍贵的，能对在相应的场景中想到相应的解决方案也是比较难的，
有了想法后，再去深入一步看细节，比如农户，原料厂商这些节点应该都是可伸缩的，可以动态的加入节点，踢出节点，而且还应该对他们有一点的奖惩机制，比如信誉到达一个
值，该踢出就踢出，还有数据上链应该采取什么方法，每个操作都上链，那么就会导致链很庞大，不每个操作都上，那么可信度又怎么解决，环节又该怎么设计，是动态添加环节，
还是直接定制，这些都是需要我们去思考的。  

    blockchain-trace-bcnetwork：区块链网络，可直接将文件上传至服务器，然后启动里面的脚本

    blockchain-trace-applets：小程序端溯源，做得十分简单

    blockchain-trace-pc：PC端,使用的是RuoYi-Vue

    blockchain-trace-basic-data：系统基础数据后台，使用的是RuoYi

    前端：Vue.js , Element UI , mpvue

    后端：SpringBoot , Mybatis , FastDFS , Node.js , Redis , MySQL

    区块链：Fabric1.2

    智能合约：Golang语言

    环境：Ubuntu16.04 64位(建议2核 4G以上，2G也能运行)，Docker 18.09.7 , Docker-compose 1.29.1 , 

因为本系统的区块链部分是完全按照最简单的来，比如共识是自带的solo，orderer节点是单节点，这些肯定都是不合理的，所以这个版本就单纯传达一个上层应用思想吧，
后续可能会在此思想上重新搭建一套能够投入使用的。  

# 安装教程
## 一.fabric网络  
### 1.确保环境配置好  
> node.js 8* 

> docker

> docker-compose 

> Redis

> FastDFS

> Mysql8

> go语言环境  

### 2.拉取docker镜像  

#### pull 
``` 
docker pull hyperledger/fabric-peer:1.2.0 && 
docker pull hyperledger/fabric-orderer:1.2.0 && 
docker pull hyperledger/fabric-ca:1.2.0 && 
docker pull hyperledger/fabric-tools:1.2.0 && 
docker pull hyperledger/fabric-ccenv:1.2.0 && 
docker pull hyperledger/fabric-baseimage:0.4.10 && 
docker pull hyperledger/fabric-baseos:0.4.10 && 
docker pull hyperledger/fabric-couchdb:0.4.10
```

#### tag  
```
docker tag hyperledger/fabric-peer:1.2.0 hyperledger/fabric-peer && 
docker tag hyperledger/fabric-orderer:1.2.0 hyperledger/fabric-orderer && 
docker tag hyperledger/fabric-ca:1.2.0 hyperledger/fabric-ca && 
docker tag hyperledger/fabric-tools:1.2.0 hyperledger/fabric-tools && 
docker tag hyperledger/fabric-ccenv:1.2.0 hyperledger/fabric-ccenv && 
docker tag hyperledger/fabric-baseimage:0.4.10  hyperledger/fabric-baseimage && 
docker tag hyperledger/fabric-baseos:0.4.10 hyperledger/fabric-baseos && 
docker tag hyperledger/fabric-couchdb:0.4.10 hyperledger/fabric-couchdb 
```

### 3.上传代码到linux服务器/或者直接git拉取  
`blockchain-trace-bcnetwork`

### 4.运行basic_network目录下的start.sh文件  
> chmod -R 777 start.sh

>./start.sh 

此处会安装可能会出现问题，具体问题可以百度。

### 5.运行webapp目录下的./start.sh  
先给webapp目录下的所有sh文件授权，如下
```
chmod -R 777 startFarmerCC.sh  
```

### 6.执行npm install安装依赖  
> npm install  

可能node版本问题安装会出一些问题，node8应该是没问题的。   

### 7.安装用户密钥  
> node enrollAdmin.js  

> node registerUser.js  

执行node registerUser.js可能会安装失败，请删除一下hfc-key-store后重新执行，如果还是失败，可能就是npm install出问题。  

### 8.启动node服务(node服务就是一个中间件，连接前端和区块链网络)  
> node app.js  

如果需要让其常驻后台，需要安装pm2。
> 启动：pm2 start app.js

> 停止：pm2 stop app.js  

到这里，区块链网络就部署完成  

## 二.系统基础数据后台blockchain-trace-basic-data，是一个SpringBoot项目  

### 1.修改application.yml文件中的Redis地址和fastdfs地址  
```yaml
  # redis 配置
  redis:
    # 地址
    host: 127.0.0.1
    # 端口，默认为6379
    port: 6379
    # 密码
    password: 

  fdfs:
      so-timeout: 1501
      connect-timeout: 601
      thumb-image:
        width: 60
        height: 60
      tracker-list: 127.0.0.1:22122
      address: http://127.0.0.1:8888/
```
### 2.修改application-druid.yml文件中mysql地址  
```yaml
# 数据源配置
spring:
    datasource:
        type: com.alibaba.druid.pool.DruidDataSource
        driverClassName: com.mysql.cj.jdbc.Driver
        druid:
            # 主库数据源
            master:
                url: jdbc:mysql://127.0.0.1:3306/blockchain?useUnicode=true&characterEncoding=utf8&zeroDateTimeBehavior=convertToNull&useSSL=true&serverTimezone=GMT%2B8
                username: root
                password: root
```

## 三.PC端（blockchain-trace-pc） 

### 1.安装依赖  
> npm install --registry=https://registry.npm.taobao.org

### 2.修改连接区块链网络地址
main.js，修改为区块链网络所在服务器地址  
```yaml
Vue.prototype.$httpUrl = "http://localhost:8080/route";
```
### 3.启动项目  
> npm run dev  


## 四.小程序（blockchain-trace-applets）
自己使用开发者工具打开
    

## **扩展**  
农产品溯源，其目的是给实物赋能，不单单是农产品，艺术品，工艺品等这些如果能给其赋能，那么其价值会得到一定程度的提升。  

NFT这个概念想必了解区块链的同学都会知道它，我们称之为数字藏品，它其实是有很高的价值的，当然，如果你为了炒它，那么将毫无价值，如果我们将实物和NFT，供应链
联合起来，就会碰撞出闪亮的火花，让NFT不单单只是一张图片，而是一种生活中的消费卷，是一种情感的见证卡片，是一种契约，更是一种连接人与人的桥梁。  

我们也推出了NFT技术实现，在了解NFT的价值后，我们可以将其与实物结合，实现实物的赋能和价值提升，有兴趣可以去看看，学习一下NFT平台。  

浏览地址：http://43.143.132.167/h5/#/

>如果本项目帮助到你，请给个免费的star  
    
## 界面
区块链浏览器
![区块链结构](https://images.gitee.com/uploads/images/2021/0510/100450_865a1f55_4775150.png "6.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/100501_6d258226_4775150.png "7.png")


添加农作物
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/095409_1656ec9f_4775150.png "界面1.png")

农作物链上详情
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/095812_c090b8eb_4775150.png "2.png")

农作物种植过程记录
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/095938_b91c4a9e_4775150.png "3.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/100000_d3f09618_4775150.png "4.png")

物流追踪（因本系统，这里忘记去做了，只是模拟了一下）
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/100223_1c37229e_4775150.png "5.png")

司机运输定位（本系统采用PC端定位，实际上这是不行的，你想想，司机在开车过程中还要去打开浏览器进入系统去定位？这里只是模拟 ）
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
![img.png](install-fabric-env/static/img.png)
小程序端溯源

![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234302_29fe611b_4775150.png "7.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234335_43fbec55_4775150.png "1.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234347_e2572d95_4775150.png "3.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234359_b5bac058_4775150.png "4.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234413_33dd3e47_4775150.png "5.png")
![输入图片说明](https://images.gitee.com/uploads/images/2021/0510/234428_bc064965_4775150.png "6.png")

