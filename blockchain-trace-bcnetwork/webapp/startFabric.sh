#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -e

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
LANGUAGE=golang
CC_SRC_PATH=github.com/chaincode/farmercc/go/
DRIVERCC_PATH=
# clean the keystore
rm -rf ./hfc-key-store

# launch network; create channel and join peer to channel
cd ../basic-network
./start.sh

docker-compose -f ./docker-compose.yml up -d cli

                #######################################      #install chaincode      ####################################################
#使用TLS
#docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/peers/peer0.org1.trace.com/tls/ca.crt" cli peer chaincode install -n farmercc -v 1.0 -l $LANGUAGE -p $CC_SRC_PATH
#不使用TLS
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051" cli peer chaincode install -n farmercc -v 1.0 -l $LANGUAGE -p $CC_SRC_PATH

#docker exec -e "CORE_PEER_LOCALMSPID=Org5MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org5.trace.com/users/Admin@org5.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org5.trace.com:7051" cli peer chaincode install -n drivercc -v 1.0 -l $LANGUAGE -p $CC_SRC_PATH


               ##########################################      #instantiate chaincode    ########################################
#使用TLS
#docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/peers/peer0.org1.trace.com/tls/ca.crt" cli peer chaincode instantiate -o orderer.trace.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/trace.com/orderers/orderer.trace.com/msp/tlscacerts/tlsca.trace.com-cert.pem -C tracechannel -n farmercc -l $LANGUAGE -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
#不使用TLS
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051" cli peer chaincode instantiate -o orderer.trace.com:7050 -C tracechannel -n farmercc -l $LANGUAGE -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')"

#docker exec -e "CORE_PEER_LOCALMSPID=Org5MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org5.trace.com/users/Admin@org5.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org5.trace.com:7051" cli peer chaincode instantiate -o orderer.trace.com:7050 -C tracechannel -n drivercc -l $LANGUAGE -v 1.0 -c '{"Args":[""]}' -P "OR ('Org5MSP.member','Org2MSP.member')"
#实例化智能合约后睡眠5秒钟
sleep 5

               ############################################        #invoke chaincode       #######################################   
#使用TLS
#docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051" -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/peers/peer0.org1.trace.com/tls/ca.crt"  cli peer chaincode invoke -o orderer.trace.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/trace.com/orderers/orderer.trace.com/msp/tlscacerts/tlsca.trace.com-cert.pem -C tracechannel -n farmercc -c '{"function":"initLedger","Args":[""]}'
#不使用TLS
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.trace.com/users/Admin@org1.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.trace.com:7051"  cli peer chaincode invoke -o orderer.trace.com:7050 -C tracechannel -n farmercc -c '{"function":"initLedger","Args":[""]}'

#docker exec -e "CORE_PEER_LOCALMSPID=Org5MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org5.trace.com/users/Admin@org5.trace.com/msp" -e "CORE_PEER_ADDRESS=peer0.org5.trace.com:7051"  cli peer chaincode invoke -o orderer.trace.com:7050 -C tracechannel -n drivercc -c '{"function":"initLedger","Args":[""]}'

printf "\nTotal setup execution time : $(($(date +%s) - starttime)) secs ...\n\n\n"
printf "Start by installing required packages run 'npm install'\n"
printf "Then run 'node enrollAdmin.js', then 'node registerUser'\n\n"
printf "The 'node invoke.js' will fail until it has been updated with valid arguments\n"
printf "The 'node query.js' may be run at anytime once the user has been registered\n\n"

