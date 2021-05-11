#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

docker-compose -f docker-compose.yml down

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/trace.com/orderers/orderer.trace.com/msp/tlscacerts/tlsca.trace.com-cert.pem
# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=10
#echo ${FABRIC_START_TIMEOUT}
sleep ${FABRIC_START_TIMEOUT}

docker-compose -f docker-compose.yml up -d

#
#
# 以下创建通道，加入通道，更新锚节点  分为使用TLS和不使用TLS ， 这里先不使用
#
#
# Create the channel

#Use TLS
#docker exec cli peer channel create -o orderer.trace.com:7050 -c tracechannel --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/trace.com/orderers/orderer.trace.com/msp/tlscacerts/tlsca.trace.com-cert.pem -f ./channel-artifacts/channel.tx

#UseLess TLS
docker exec cli peer channel create -o orderer.trace.com:7050 -c tracechannel -f ./channel-artifacts/channel.tx


# Join peer1 to the channel. #####################################################################################
#Use TLS
#docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/peers/peer0.org1.trace.com/tls/ca.crt" cli peer channel join --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/trace.com/orderers/orderer.trace.com/msp/tlscacerts/tlsca.trace.com-cert.pem -b tracechannel.block

#UseLess TLS
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051" cli peer channel join  -b tracechannel.block

#Update Anchors

#Use Tls
#docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/peers/peer0.org1.trace.com/tls/ca.crt" cli peer channel update -o orderer.trace.com:7050 -c tracechannel -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile $ORDERER_CA

#UseLess TLS
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051" cli peer channel update -o orderer.trace.com:7050 -c tracechannel -f ./channel-artifacts/Org1MSPanchors.tx


# Join peer2 to the channel.###########################################################################################################
#Use TLS
docker exec -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.trace.com/users/Admin@org2.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org2.trace.com:7051" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.trace.com/peers/peer0.org2.trace.com/tls/ca.crt" cli peer channel join -b tracechannel.block

#Useless TLS
docker exec -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.trace.com/users/Admin@org2.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org2.trace.com:7051" cli peer channel update -o orderer.trace.com:7050 -c tracechannel -f ./channel-artifacts/Org2MSPanchors.tx


# Join peer3 to the channel.#################################################################################################################
docker exec -e "CORE_PEER_LOCALMSPID=Org3MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.trace.com/users/Admin@org3.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org3.trace.com:7051" cli peer channel join -b tracechannel.block

docker exec -e "CORE_PEER_LOCALMSPID=Org3MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.trace.com/users/Admin@org3.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org3.trace.com:7051" cli peer channel update -o orderer.trace.com:7050 -c tracechannel -f ./channel-artifacts/Org3MSPanchors.tx

# Join peer4 to the channel.###############################################################################################################
docker exec -e "CORE_PEER_LOCALMSPID=Org4MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org4.trace.com/users/Admin@org4.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org4.trace.com:7051" cli peer channel join -b tracechannel.block

docker exec -e "CORE_PEER_LOCALMSPID=Org4MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org4.trace.com/users/Admin@org4.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org4.trace.com:7051" cli peer channel update -o orderer.trace.com:7050 -c tracechannel -f ./channel-artifacts/Org4MSPanchors.tx

# Join peer5 to the channel.#################################################################################################################
docker exec -e "CORE_PEER_LOCALMSPID=Org5MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org5.trace.com/users/Admin@org5.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org5.trace.com:7051" cli peer channel join -b tracechannel.block

docker exec -e "CORE_PEER_LOCALMSPID=Org5MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org5.trace.com/users/Admin@org5.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org5.trace.com:7051" cli peer channel update -o orderer.trace.com:7050 -c tracechannel -f ./channel-artifacts/Org5MSPanchors.tx


