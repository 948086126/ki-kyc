#!/usr/bin/env bash

###  参数 ------------------------------------------------------

## 链码名称
GO_CC_NAME=("bookstorechain")
## 链码路径
GO_CC_SRC_PATH=("github.com/chaincode/book")
## 链码版本
CC_VERSION="1.0"

### tls

TLS="true"

## 链码语言
CC_GOLANG="golang"
## 通道名称
CHANNEL_NAME="bookchannel"
DOMAIN_NAME="bookstore.com"
orderer1_ADDRESS="orderer1.bookstore.com:7050"
ORG1_ADDRESS="peer0.org1.bookstore.com:7051"
ORG2_ADDRESS="peer0.org2.bookstore.com:9051"

ORG_NAME=("org1" "org2")
TLS_PATH="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/"
ORDERER_TLS_PATH="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/"
ORDERER_CAFILE="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/bookstore.com/orderers/orderer1.bookstore.com/msp/tlscacerts/tlsca.bookstore.com-cert.pem"

##  org1 org2
ORG1_CAT="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/ca.crt"
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
    local channel=$1
    local org="org1"
    local peer="peer0"
    local port="7051"
    local cert="server.crt"
    local key="server.key"
    local rootcert="ca.crt"
    local orderer="orderer1"

    docker exec \
        -e "CORE_PEER_LOCALMSPID=Org1MSP" \
        -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/users/Admin@org1.bookstore.com/msp" \
        -e "CORE_PEER_ADDRESS=peer0.org1.bookstore.com:7051" \
        -e "CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/server.crt" \
        -e "CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/server.key" \
        -e "CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.bookstore.com/peers/peer0.org1.bookstore.com/tls/ca.crt" \
        cli \
        peer channel create \
                    -o $orderer1_ADDRESS \
                    -c $channel \
                    -f ./channel-artifacts/channelcopyright.tx \
                    --tls $TLS \
                    --cafile  $ORDERER_CAFILE
}

channel_join() {
    local channel=$1
    local org=$2
    local peer=$3
    local port=$4
    local cert=$5
    local key=$6
    local rootcert=$7
    ###
    docker exec \
        -e "CORE_PEER_LOCALMSPID=$(get_mspid $org)"\
        -e "CORE_PEER_MSPCONFIGPATH=$(get_msp_channel-artifacts_path $org $peer)"\
        -e "CORE_PEER_ADDRESS=$(get_peer_address $org $peer $port)"\
        -e "CORE_PEER_TLS_CERT_FILE=$(get_peer_tls_cert $org $peer $cert)"\
        -e "CORE_PEER_TLS_KEY_FILE=$(get_peer_tls_cert $org $peer $key)"\
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$(get_peer_tls_cert $org $peer $rootcert)"\
        cli \
          peer channel \
          join -b $channel.block

     echo "********************$org...$peer join channel successful***************"
}

install_and_instantiate() {

    local lang=$1
    local cc_name=($2)
    local cc_src_path=($3)

    chaincode_install $CHANNEL_NAME  "org1" "peer0" "7051" "server.crt" "server.key" "ca.crt" "orderer1" ${cc_name[0]} ${cc_src_path[0]} $lang "org1"

    chaincode_install $CHANNEL_NAME  "org1" "peer1" "8051" "server.crt" "server.key" "ca.crt" "orderer1" ${cc_name[0]} ${cc_src_path[0]} $lang "org1"

    chaincode_install $CHANNEL_NAME  "org2" "peer0" "9051" "server.crt" "server.key" "ca.crt" "orderer1" ${cc_name[0]} ${cc_src_path[0]} $lang "org1"

    chaincode_install $CHANNEL_NAME  "org2" "peer1" "10051" "server.crt" "server.key" "ca.crt" "orderer1" ${cc_name[0]} ${cc_src_path[0]} $lang "org1"

    ### 实例化
    chaincode_instantiate   $CHANNEL_NAME "org1" "peer0" "7051" "server.crt" "server.key" "ca.crt" "orderer1" ${cc_name[0]} ${cc_src_path[0]} $lang "org1"

    sleep 5
### 初始化
    chaincode_invoke $CHANNEL_NAME "org1" "org1" "peer0" "7051" "server.crt" "server.key" "ca.crt" "orderer1" ${cc_name[0]} "org2" "org2" "peer0" "9051" '{"function":"","Args":[""]}'
}

chaincode_install() {
    local channel=$1
    local org=$2
    local peer=$3
    local port=$4
    local cert=$5
    local key=$6
    local rootcert=$7
    local orderer=$8
    local cc_name=$9
    local cc_src_path=${10}
    local lang=${11}
    local Org=${12}


    docker exec \
        -e "CORE_PEER_LOCALMSPID=$(get_mspid $org)" \
        -e "CORE_PEER_MSPCONFIGPATH=$(get_msp_channel-artifacts_path $org $peer)" \
        -e "CORE_PEER_ADDRESS=$(get_peer_address $org $peer $port)" \
        -e "CORE_PEER_TLS_CERT_FILE=$(get_peer_tls_cert $org $peer $cert)" \
        -e "CORE_PEER_TLS_KEY_FILE=$(get_peer_tls_cert $org $peer $key)" \
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$(get_peer_tls_cert $org $peer $rootcert)" \
        cli \
        peer chaincode install \
        -n $cc_name \
        -v $CC_VERSION \
        -l $lang \
        -p $cc_src_path
}

chaincode_instantiate() {
    local channel=$1
    local org=$2
    local peer=$3
    local port=$4
    local cert=$5
    local key=$6
    local rootcert=$7
    local orderer=$8
    local cc_name=$9
    local cc_src_path=${10}
    local lang=${11}
    local Org=${12}

    docker exec \
        -e "CORE_PEER_LOCALMSPID=$(get_mspid $org)" \
        -e "CORE_PEER_MSPCONFIGPATH=$(get_msp_channel-artifacts_path $org $peer)" \
        -e "CORE_PEER_ADDRESS=$(get_peer_address $org $peer $port)" \
        -e "CORE_PEER_TLS_CERT_FILE=$(get_peer_tls_cert $org $peer $cert)" \
        -e "CORE_PEER_TLS_KEY_FILE=$(get_peer_tls_cert $org $peer $key)" \
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$(get_peer_tls_cert $org $peer $rootcert)" \
        cli \
        peer chaincode instantiate \
        -o $orderer1_ADDRESS \
        --tls $TLS \
        --cafile $ORDERER_CAFILE \
        -C $CHANNEL_NAME \
        -n $cc_name \
        -l $CC_GOLANG \
        -v $CC_VERSION \
        -c '{"Args":[""]}' \
        -P 'OR ('\''Org1MSP.member'\'','\''Org2MSP.member'\'')'

     echo "*******************************init chaincode is successful*********************************"
}

chaincode_invoke() {
    local channel=$1
    local org1=$2
    local Org1=$3
    local peer=$4
    local port=$5
    local cert=$6
    local key=$7
    local rootcert=$8
    local orderer=$9
    local cc_name=${10}
    local org2=${11}
    local Org2=${12}
    local Org2peer=${13}
    local Org2port=${14}
    local  cmd=${15}

    docker exec \
        -e "CORE_PEER_LOCALMSPID=$(get_mspid $org1)" \
        -e "CORE_PEER_MSPCONFIGPATH=$(get_msp_channel-artifacts_path $Org1 $peer)" \
        -e "CORE_PEER_ADDRESS=$(get_peer_address $org1 $peer $port)" \
        -e "CORE_PEER_TLS_CERT_FILE=$(get_peer_tls_cert $org1 $peer $cert)" \
        -e "CORE_PEER_TLS_KEY_FILE=$(get_peer_tls_cert $org1 $peer $key)" \
        -e "CORE_PEER_TLS_ROOTCERT_FILE=$(get_peer_tls_cert $org1 $peer $rootcert)" \
        cli \
        peer chaincode invoke \
        -o $orderer1_ADDRESS \
        --tls $TLS \
        --cafile  $ORDERER_CAFILE \
        -C $CHANNEL_NAME \
        -n $cc_name \
        --peerAddresses $ORG1_ADDRESS \
        --tlsRootCertFiles $ORG1_CAT \
        --peerAddresses $ORG2_ADDRESS \
        --tlsRootCertFiles $$ORG2_CAT \
        -c '{"Args":[""]}'

    echo "**********************************invoke chaincode*******$cc_name************************************************"
}


###  start --------------------------------------------------------------------------------------------
echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "....... Build your  Server ......."
echo
###


### 安装链码
install_and_instantiate $CC_GOLANG "${GO_CC_NAME[*]}" "${GO_CC_SRC_PATH[*]}"
echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo

