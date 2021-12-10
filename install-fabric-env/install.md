

1.下载cryptogen工具和configtxgen工具

在fabric源码下面执行make cryptogen , make configtxgen , 会在./build/bin下面生成工具，我在centos下执行出错了，在ubuntu下可以，应该是golang路径引起的(暂时没有考证)

2.生成对应的证书
编写cryto-config.yml，然后执行 cryptogen generate --config=./crypto-config.yaml 生成个证书
注:如果在crytogen下面执行的话识别不出cryptogen，应该./cryptogen , 如：./cryptogen generate --config=./crypto-config.yaml
将会生成文件夹crypto-config，里面就是对应的证书

3.生成创世区块
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block

注：TwoOrgsOrdererGenesis是configtx.yaml里面的那个名字

将会在channel-artifacts目录下生成创世区块(genesis.block)

4.创建通道channel
channel名字
export CHANNEL_NAME=certchannel

OneOrgsChannel是configtx.yaml里面的那个名字
./configtxgen -profile OneOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME


将会在channel-artifacts目录下生成通道文件(channel.tx)




