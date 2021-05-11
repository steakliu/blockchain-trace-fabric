const express = require("express");
const router = express.Router();

var co = require('co');
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
//channel.addPeer(peer2);

var member_user = null;
var store_path = path.join(__dirname, '../../hfc-key-store');


//溯源作物过程记录
router.get("/queryRetailerById", async function (req,res) {
console.log("arr  "+req.query.id)
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
    const request = {
      //targets : --- letting this default to the peers assigned to the channel
      chaincodeId: 'retailecc',
      fcn: 'queryRetailerById',
      args: [req.query.id]
    };

    var query_responses = await channel.queryByChaincode(request);
    console.log("Query has completed, checking results");
    // query_responses could have more than one  results if there multiple peers were used as targets
    if (query_responses && query_responses.length === 1) {
      if (query_responses[0] instanceof Error) {
        res.send(query_responses[0])
        console.error("error from query = ", query_responses[0]);
      } else {
        res.send(query_responses[0].toString())
        console.log("Response is ", query_responses[0].toString());
      }
    } else {
      console.log("No payloads were returned from query");
    }
  }catch (e) {
    console.error('Failed to query successfully :: ' + e);
  }
});


//溯源作物过程记录
router.get("/queryRetailerByCropsId", async function (req,res) {
console.log("arr  "+req.query.cropsId+" "+req.query.retailerId)
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
    const arry = []
    arry.push(req.query.cropsId)
    arry.push(req.query.retailerId)	
    const request = {
      //targets : --- letting this default to the peers assigned to the channel
      chaincodeId: 'retailecc',
      fcn: 'queryRetailerByCropsId',
      args: arry
    };

    var query_responses = await channel.queryByChaincode(request);
    console.log("Query has completed, checking results");
    // query_responses could have more than one  results if there multiple peers were used as targets
    if (query_responses && query_responses.length === 1) {
      if (query_responses[0] instanceof Error) {
        res.send(query_responses[0])
        console.error("error from query = ", query_responses[0]);
      } else {
        res.send(query_responses[0].toString())
        console.log("Response is ", query_responses[0].toString());
      }
    } else {
      console.log("No payloads were returned from query");
    }
  }catch (e) {
    console.error('Failed to query successfully :: ' + e);
  }
});


router.post("/createRetailer", async function(req,res){
           console.log("productInfo    "+JSON.stringify(req.body.retailerArray));
    try {
            var state_store = await Fabric_Client.newDefaultKeyValueStore({path: store_path});
            fabric_client.setStateStore(state_store);
            var crypto_suite = Fabric_Client.newCryptoSuite();
            var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
            crypto_suite.setCryptoKeyStore(crypto_store);
            fabric_client.setCryptoSuite(crypto_suite);
            var user_from_store = await fabric_client.getUserContext('user1', true);
            if (user_from_store && user_from_store.isEnrolled()) {
              console.log('Successfully loaded user1 from persistence');
              member_user = user_from_store;
            } else {
              throw new Error('Failed to get user1.... run registerUser.js');
            }
            let tx_id = fabric_client.newTransactionID();
         
            const request = {
              chaincodeId: 'retailecc',
              fcn: 'createRetailer',
              txId: tx_id,
              args: req.body.retailerArray,
              chainId: 'tracechannel'
            };
            let chaincodeinvokeresult  = await channel.sendTransactionProposal(request);
            //获取返回结果
            var proposalResponses = chaincodeinvokeresult[0];
            var proposal = chaincodeinvokeresult[1];
            var header = chaincodeinvokeresult[2];
            var all_good = true;
            for(var i in proposalResponses){
                let one_good = false;
        //检验提案响应信息是否正确
                if(proposalResponses && proposalResponses[0].response && proposalResponses[0].response.status === 200){
                        one_good = true;
                        console.info('transaction proposal was good');
                }else{
                        console.error('transaction proposal was bad');
                }
                all_good = all_good & one_good;
            }
            if(all_good){
                console.info(util.format(
                        'Successfully : Status - %s, message - %s, metadata - %s, endorsement signature - %s',
                        proposalResponses[0].response.status,proposalResponses[0].response.message,
                        proposalResponses[0].response.payload,proposalResponses[0].endorsement.signature
                ));
                var request2 = {
                        proposalResponses: proposalResponses,
                        proposal: proposal,
                        header: header
                };
                channel.sendTransaction(request2)
                const resBody = {
                        status:200,
                        msg:"数据上链成功"
                };
                res.send(resBody);
        }
          }catch (e) {
            console.error('Failed to query successfully :: ' + e);
          }
});


module.exports = router;

