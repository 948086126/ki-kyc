

###  参数 ------------------------------------------------------

## 链码名称
GO_CC_NAME=("bookstorechain")
## 链码路径
GO_CC_SRC_PATH=("github.com/chaincode/sun")
## 链码语言
CC_GOLANG="golang"
## 链码版本
CC_VERSION="3.2"
## 通道名称
CHANNEL_NAME="bookchannel"

## 拼接参数
DOMAIN_NAME="bookstore.com"

## 排序节点地址： 默认请求的orderer节点地址
ORDERER1_ADDRESS="orderer1.bookstore.com:7050"

## org1 id 地址
ORG1_ADDRESS="peer0.org1.bookstore.com:7051"

## org2 id 地址
ORG2_ADDRESS="peer0.org2.bookstore.com:9051"

## 组织名称
ORG_NAME=("org1" "org2")

## crt
SERVER_CRT="server.crt"

## key
SERVER_KEY="server.key"

## ca
SERVER_CA="ca.crt"

## yaml
DOCKER_YAML="docker-compose-etcdraft2.yaml"

## cmd
CMD='{
"channelName":"bookchannel",
"chainCodeName":"bookstorechain",
"functionName":"invoke",
"data":["id23xiosjodfis11113","zhangsan3","393"]
}'

## tls +
TLS_PATH="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/"
## orderer + tls +
ORDERER_TLS_PATH="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/"

## orderer + ca
ORDERER_CAFILE="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookstore.com/orderers/orderer1.bookstore.com/msp/tlscacerts/tlsca.bookstore.com-cert.pem"


## cli
CLI_CLIENT="cli"

## tls
TLS="true"

## channel-artifacts
CHANNEL_ARTIFACTS="./channel-artifacts/"

### PEER
CORE_PEER_LOCALMSPID="Org1MSP"

##  CORE_PEER_MSPCONFIGPATH
CORE_PEER_MSPCONFIGPATH="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/users/Admin@org1.bookstore.com/msp"
##  CORE_PEER_ADDRESS
CORE_PEER_ADDRESS="peer0.org1.bookstore.com:7051"
##  CORE_PEER_TLS_CERT_FILE
CORE_PEER_TLS_CERT_FILE="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/server.crt"

##  CORE_PEER_TLS_KEY_FILE
CORE_PEER_TLS_KEY_FILE="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/server.key"

##  CORE_PEER_TLS_ROOTCERT_FILE
CORE_PEER_TLS_ROOTCERT_FILE="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/ca.crt"

##  org1 org2
ORG1_CAT="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/ca.crt"

## org2
ORG2_CAT="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.bookstore.com/peers/peer0.org2.bookstore.com/tls/ca.crt"

## 工具函数
get_mspid() {
    local org=$1
    case "$org" in
        org1)
            echo "Org1MSP"
            ;;
        org2)
            echo "Org2MSP"
            ;;
        *)
            echo "error org name $org"
            exit 1
            ;;
    esac
}

get_msp_channel-artifacts_path() {

    local org=$1
    local peer=$2

    if [[ "$org" = "Org1" ]] && [[ "$org" = "Org2" ]]; then
        echo "error org name $org"
        exit 1
    fi

    if [[ "$peer" = "peer0" ]] && [[ "$peer" = "peer1" ]]; then
        echo "error peer name $peer"
        exit 1
    fi

    echo "${TLS_PATH}peerOrganizations/$org.bookstore.com/users/Admin@$org.bookstore.com/msp"

}

get_peer_address() {
    local org=$1
    local peer=$2
    local port=$3
    if [[ "$org" != "org1" ]] && [[ "$org" != "org2" ]]; then
        echo "error org name $org"
        exit 1
    fi

    echo "${peer}.${org}.${DOMAIN_NAME}:$port"
}

get_peer_tls_cert(){
    local org=$1
    local peer=$2
    local type=$3
    if [[ "$org" != "org1" ]] && [[ "$org" != "org2" ]]; then
        echo "error org name $org"
        exit 1
    fi

    echo "${TLS_PATH}peerOrganizations/${org}.bookstore.com/peers/${peer}.${org}.bookstore.com/tls/$type"

}


channel_create() {

    docker exec \
        -e "CORE_PEER_LOCALMSPID=$CORE_PEER_LOCALMSPID" \
        -e "CORE_PEER_MSPCONFIGPATH=$CORE_PEER_MSPCONFIGPATH" \
        -e "CORE_PEER_ADDRESS=$CORE_PEER_ADDRESS" \
        -e "CORE_PEER_TLS_CERT_FILE=$CORE_PEER_TLS_CERT_FILE" \
        -e "CORE_PEER_TLS_KEY_FILE=$CORE_PEER_TLS_KEY_FILE" \
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$CORE_PEER_TLS_ROOTCERT_FILE" \
        $CLI_CLIENT \
        peer channel create \
                    -o $ORDERER1_ADDRESS \
                    -c $CHANNEL_NAME \
                    -f $CHANNEL_ARTIFACTS$CHANNEL_NAME.tx \
                    --tls true \
                    --cafile  $ORDERER_CAFILE

     echo "***********************************************************************"
     echo "********************$CHANNEL_NAME create is ok !***********************"
     echo "***********************************************************************"
}

channel_join() {

    local channel=$1
    local org=$2
    local peer=$3
    local port=$4

    docker exec \
        -e "CORE_PEER_LOCALMSPID=$(get_mspid $org)"\
        -e "CORE_PEER_MSPCONFIGPATH=$(get_msp_channel-artifacts_path $org $peer)"\
        -e "CORE_PEER_ADDRESS=$(get_peer_address $org $peer $port)"\
        -e "CORE_PEER_TLS_CERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CRT)"\
        -e "CORE_PEER_TLS_KEY_FILE=$(get_peer_tls_cert $org $peer $SERVER_KEY)"\
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CA)"\
        $CLI_CLIENT \
                 peer channel \
                 join -b $channel.block

     echo "***********************************************************************"
     echo "********************$org...$peer join channel successful***************"
     echo "***********************************************************************"
}

install_and_instantiate() {

    local lang=$1
    local cc_name=($2)
    local cc_src_path=($3)

    chaincode_install $CHANNEL_NAME  "org1" "peer0" "7051"   ${cc_name[0]} ${cc_src_path[0]} $lang
    chaincode_install $CHANNEL_NAME  "org1" "peer1" "8051"   ${cc_name[0]} ${cc_src_path[0]} $lang
    chaincode_install $CHANNEL_NAME  "org2" "peer0" "9051"   ${cc_name[0]} ${cc_src_path[0]} $lang
    chaincode_install $CHANNEL_NAME  "org2" "peer1" "10051"  ${cc_name[0]} ${cc_src_path[0]} $lang

    chaincode_upgrade  $CHANNEL_NAME "org1" "peer0" "7051"  ${cc_name[0]} ${cc_src_path[0]} $lang

    sleep 10

    #chaincode_invoke $CHANNEL_NAME "org1"  "peer0" "7051"  ${cc_name[0]}  $CMD
}

chaincode_install() {

    local channel=$1
    local org=$2
    local peer=$3
    local port=$4
    local cc_name=$5
    local cc_src_path=$6
    local lang=$7

    docker exec \
        -e "CORE_PEER_LOCALMSPID=$(get_mspid $org)" \
        -e "CORE_PEER_MSPCONFIGPATH=$(get_msp_channel-artifacts_path $org $peer)" \
        -e "CORE_PEER_ADDRESS=$(get_peer_address $org $peer $port)" \
        -e "CORE_PEER_TLS_CERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CRT)"\
        -e "CORE_PEER_TLS_KEY_FILE=$(get_peer_tls_cert $org $peer $SERVER_KEY)"\
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CA)"\
        $CLI_CLIENT \
        peer chaincode install \
                -n $cc_name \
                -v $CC_VERSION \
                -l $lang \
                -p $cc_src_path

     echo "***********************************************************************"
     echo "***************$org...$peer chaincode installl successful**************"
     echo "***********************************************************************"
}

chaincode_upgrade() {
    local channel=$1
    local org=$2
    local peer=$3
    local port=$4
    local cc_name=$5
    local cc_src_path=$6
    local lang=$7

    docker exec \
        -e "CORE_PEER_LOCALMSPID=$(get_mspid $org)" \
        -e "CORE_PEER_MSPCONFIGPATH=$(get_msp_channel-artifacts_path $org $peer)" \
        -e "CORE_PEER_ADDRESS=$(get_peer_address $org $peer $port)" \
        -e "CORE_PEER_TLS_CERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CRT)"\
        -e "CORE_PEER_TLS_KEY_FILE=$(get_peer_tls_cert $org $peer $SERVER_KEY)"\
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CA)"\
        $CLI_CLIENT \
        peer chaincode upgrade  \
        -o $ORDERER1_ADDRESS \
        --tls true \
        --cafile $ORDERER_CAFILE \
        -C $CHANNEL_NAME \
        -n $cc_name \
        -l golang \
        -v $CC_VERSION \
        -c '{"Args":[""]}' \
        -P 'OR ('\''Org1MSP.member'\'','\''Org2MSP.member'\'')'

     echo "*******************************update chaincode is successful*********************************"
}
chaincode_instantiate() {

    local channel=$1
    local org=$2
    local peer=$3
    local port=$4
    local cc_name=$5
    local cc_src_path=$6
    local lang=$7

    docker exec \
        -e "CORE_PEER_LOCALMSPID=$(get_mspid $org)" \
        -e "CORE_PEER_MSPCONFIGPATH=$(get_msp_channel-artifacts_path $org $peer)" \
        -e "CORE_PEER_ADDRESS=$(get_peer_address $org $peer $port)" \
        -e "CORE_PEER_TLS_CERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CRT)"\
        -e "CORE_PEER_TLS_KEY_FILE=$(get_peer_tls_cert $org $peer $SERVER_KEY)"\
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CA)"\
        $CLI_CLIENT \
        peer chaincode instantiate \
            -o $ORDERER1_ADDRESS \
            --tls $TLS \
            --cafile $ORDERER_CAFILE \
            -C $channel \
            -n $cc_name \
            -l $CC_GOLANG \
            -v $CC_VERSION \
            -c '{"Args":[""]}' \
            -P 'OR ('\''Org1MSP.member'\'','\''Org2MSP.member'\'')'
        ### 背书策略 ： OR
     echo "***********************************************************************"
     echo "************instantiate chaincode is successful************************"
     echo "***********************************************************************"
}

chaincode_invoke() {

    local channel=$1
    local org=$2
    local peer=$3
    local port=$4
    local cc_name=$5
    local cmd=$6

    docker exec \
        -e "CORE_PEER_LOCALMSPID=$(get_mspid $org)" \
        -e "CORE_PEER_MSPCONFIGPATH=$(get_msp_channel-artifacts_path $org $peer)" \
        -e "CORE_PEER_ADDRESS=$(get_peer_address $org $peer $port)" \
        -e "CORE_PEER_TLS_CERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CRT)"\
        -e "CORE_PEER_TLS_KEY_FILE=$(get_peer_tls_cert $org $peer $SERVER_KEY)"\
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$(get_peer_tls_cert $org $peer $SERVER_CA)"\
        $CLI_CLIENT \
        peer chaincode invoke \
                -o orderer1.bookstore.com:7050 \
                --tls true \
                --cafile  /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookstore.com/orderers/orderer1.bookstore.com/msp/tlscacerts/tlsca.bookstore.com-cert.pem \
                -C bookchannel \
                -n bookstorechain \
                --peerAddresses peer0.org1.bookstore.com:7051 \
                --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/ca.crt \
                --peerAddresses peer0.org2.bookstore.com:9051 \
                --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.bookstore.com/peers/peer0.org2.bookstore.com/tls/ca.crt \
                -c '{"Args":["test","abc"]}'

        # $CLI_CLIENT \
        # peer chaincode invoke \
        #         -o $ORDERER1_ADDRESS \
        #         --tls $TLS \
        #         --cafile  $ORDERER_CAFILE \
        #         -C $channel \
        #         -n $cc_name \
        #         --peerAddresses $ORG1_ADDRESS \
        #         --tlsRootCertFiles $ORG1_CAT \
        #         --peerAddresses $ORG2_ADDRESS \
        #         --tlsRootCertFiles $ORG2_CAT \
        #         -c $cmd

        # peer chaincode invoke -o orderer1.bookstore.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookstore.com/orderers/orderer1.bookstore.com/msp/tlscacerts/tlsca.bookstore.com-cert.pem -C bookchannel -n bookstorechain --peerAddresses peer0.org1.bookstore.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/ca.crt --peerAddresses peer0.org2.bookstore.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.bookstore.com/peers/peer0.org2.bookstore.com/tls/ca.crt -c '{"Args":["init"]}'
     echo "***********************************************************************"
     echo "*******************$cc_name invoke chaincode  successful***************"
     echo "***********************************************************************"
}


###  start --------------------------------------------------------------------------------------------
echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "Build your  Server......."
echo
###

## 启动容器
#docker-compose -f  \
#     $DOCKER_YAML  -p neet up -d

#sleep 5

### 创建通道
#channel_create

### 节点加入通道
#channel_join $CHANNEL_NAME "org1" "peer0" "7051"
#channel_join $CHANNEL_NAME "org1" "peer1" "8051"
#channel_join $CHANNEL_NAME "org2" "peer0" "9051"
#channel_join $CHANNEL_NAME "org2" "peer1" "10051"

install_and_instantiate $CC_GOLANG ${GO_CC_NAME[*]} ${GO_CC_SRC_PATH[*]}

echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo
