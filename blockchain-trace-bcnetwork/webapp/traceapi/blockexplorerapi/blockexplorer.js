const express = require("express");
const router = express.Router();

var Fabric_Client = require('fabric-client');
var path = require('path');
var util = require('util');
var os = require('os');

var fabric_client = new Fabric_Client();
var channel = fabric_client.newChannel('tracechannel');
var order = fabric_client.newOrderer('grpc://localhost:7050');
channel.addOrderer(order);
var peer = fabric_client.newPeer('grpc://localhost:7051');
channel.addPeer(peer);

var member_user = null;
var store_path = path.join(__dirname, '../../hfc-key-store');
var tx_id = null;


router.get("/channelBlockInfo",async function (req,res) {
try {
    var state_store = await Fabric_Client.newDefaultKeyValueStore({path: store_path});
    fabric_client.setStateStore(state_store);
    var crypto_suite = Fabric_Client.newCryptoSuite();
    // use the same location for the state store (where the users' certificate are kept)
    // and the crypto store (where the users' keys are kept)
    var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
    crypto_suite.setCryptoKeyStore(crypto_store);
    fabric_client.setCryptoSuite(crypto_suite);

    // get the enrolled user from persistence, this user will sign all requests
    var user_from_store = await fabric_client.getUserContext('user1', true);
    if (user_from_store && user_from_store.isEnrolled()) {
      console.log('Successfully loaded user1 from persistence');
      member_user = user_from_store;
    } else {
      throw new Error('Failed to get user1.... run registerUser.js');
    }
    var query_responses = await channel.queryInfo(peer);
    res.send(query_responses)
  }catch (e) {
    console.error('Failed to query successfully :: ' + e);
  }

});

//根据区块号获取区块的详细信息
router.get("/queryBlockInfo",async function(req,res){
console.log("number   "+req.query.number)
try {
    var state_store = await Fabric_Client.newDefaultKeyValueStore({path: store_path});
    fabric_client.setStateStore(state_store);
    var crypto_suite = Fabric_Client.newCryptoSuite();
    // use the same location for the state store (where the users' certificate are kept)
    // and the crypto store (where the users' keys are kept)
    var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
    crypto_suite.setCryptoKeyStore(crypto_store);
    fabric_client.setCryptoSuite(crypto_suite);
    
    // get the enrolled user from persistence, this user will sign all requests
    var user_from_store = await fabric_client.getUserContext('user1', true);
    if (user_from_store && user_from_store.isEnrolled()) {
      console.log('Successfully loaded user1 from persistence');
      member_user = user_from_store;
    } else {
      throw new Error('Failed to get user1.... run registerUser.js');
    }
    var query_responses = await channel.queryBlock(Number(req.query.number),peer,null);
    res.send(query_responses)
  }catch (e) {
    console.error('Failed to query successfully :: ' + e);
  }

});


//根据区块号获取区块的详细信息
router.get("/queryBlockBuHash",async function(req,res){
console.log("number   "+req.query.hash)
try {
    var state_store = await Fabric_Client.newDefaultKeyValueStore({path: store_path});
    fabric_client.setStateStore(state_store);
    var crypto_suite = Fabric_Client.newCryptoSuite();
    // use the same location for the state store (where the users' certificate are kept)
    // and the crypto store (where the users' keys are kept)
    var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
    crypto_suite.setCryptoKeyStore(crypto_store);
    fabric_client.setCryptoSuite(crypto_suite);

    // get the enrolled user from persistence, this user will sign all requests
    var user_from_store = await fabric_client.getUserContext('user1', true);
    if (user_from_store && user_from_store.isEnrolled()) {
      console.log('Successfully loaded user1 from persistence');
      member_user = user_from_store;
    } else {
      throw new Error('Failed to get user1.... run registerUser.js');
    }
    var query_responses = await channel.queryBlockByHash(new Buffer(req.query.hash,"hex"),peer);
    res.send(query_responses)
  }catch (e) {
    console.error('Failed to query successfully :: ' + e);
  }

});

module.exports = router;
