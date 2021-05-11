echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "Build your first network (BYFN) end-to-end test"
echo
CHANNEL_NAME="$1"
DELAY="$2"
LANGUAGE="$3"
TIMEOUT="$4"
VERBOSE="$5"
: ${CHANNEL_NAME:="tracechannel"}
: ${DELAY:="3"}
: ${LANGUAGE:="node"}
: ${TIMEOUT:="10"}
: ${VERBOSE:="false"}
LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
COUNTER=1
MAX_RETRY=5

CC_SRC_PATH="github.com/chaincode/chaincode_example02/go/"
if [ "$LANGUAGE" = "node" ]; then
        CC_SRC_PATH="/opt/gopath/src/github.com/chaincode/chaincode_example02/node/"
fi

echo "Channel name : "$CHANNEL_NAME

# import utils
#. scripts/utils.sh

createChannel() {
        setGlobals 0 1

        if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
                set -x
                peer channel create -o orderer.trace.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx >&log.txt
                res=$?
                set +x
        else
                                set -x
                peer channel create -o orderer.trace.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
                res=$?
                                set +x
        fi
        cat log.txt
        verifyResult $res "Channel creation failed"
        echo "===================== Channel '$CHANNEL_NAME' created ===================== "
        echo
}

joinChannel () {
        for org in 1 2; do
            for peer in 0 1; do
                joinChannelWithRetry $peer $org
                echo "===================== peer${peer}.org${org} joined channel '$CHANNEL_NAME' ===================== "
                sleep $DELAY
                echo
            done
        done
}
