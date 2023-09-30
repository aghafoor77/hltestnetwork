#!/bin/bash
sudo apt-get update
sudo apt-get install git curl docker-compose -y
sudo usermod -a -G docker $USER
sudo systemctl start docker
sudo systemctl enable docker
MYHOME=${PWD}
docker --version
docker-compose --version

HLHOME=${PWD}/hldeploy
mkdir -p $HLHOME/go/src/github.com/abdulghafoor77
cp go1.21.1.linux-amd64.tar.gz $HLHOME/go/src/github.com/abdulghafoor77/
cp -r ./mvcchaincodes $HLHOME/go/src/github.com/abdulghafoor77/
cd $HLHOME/go/src/github.com/abdulghafoor77
tar -xzf go1.21.1.linux-amd64.tar.gz
export GOPATH=$HLHOME/go/src/github.com/abdulghafoor77/go
export PATH=$PATH:$GOPATH/bin
sudo apt-get install jq -y
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
./install-fabric.sh d s b
cp -r ./mvcchaincodes/mvcdidemail-go $HLHOME/go/src/github.com/abdulghafoor77/fabric-samples/asset-transfer-basic/
cp -r ./mvcchaincodes/mvcdidkeyids-go $HLHOME/go/src/github.com/abdulghafoor77/fabric-samples/asset-transfer-basic/
cp -r ./mvcchaincodes/mvckeyidpubkey-go $HLHOME/go/src/github.com/abdulghafoor77/fabric-samples/asset-transfer-basic/
cp -r ./mvcchaincodes/mvcmaildid-go $HLHOME/go/src/github.com/abdulghafoor77/fabric-samples/asset-transfer-basic/

cd $HLHOME/go/src/github.com/abdulghafoor77/fabric-samples/test-network
./network.sh down


# Deployment Starts
EMAIL_DID="mvcemaildid"
DID_EMAIL="mvcdidemail"
DID_KEYID="mvcdidkeyid"
KEYID_PUB="mvckeyidpub"

./network.sh up -ca
./network.sh createChannel
cp $HLHOME/go/src/github.com/abdulghafoor77/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/connection-org1.yaml $MYHOME/connection-org1.yaml

export FABRIC_CFG_PATH=../config/

export PATH=$HLHOME/go/src/github.com/abdulghafoor77/fabric-samples/bin:$PATH

export CORE_PEER_TLS_ENABLED=true

export CORE_PEER_LOCALMSPID="Org1MSP"

export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

export CORE_PEER_ADDRESS=localhost:7051

echo "Deploying mvcmaildid-go [$EMAIL_DID] . . ."
./network.sh deployCC -ccn $EMAIL_DID -ccp ../asset-transfer-basic/mvcmaildid-go/ -ccl go
echo "Completed : Deployed mvcmaildid-go [$EMAIL_DID]! "

echo "Deploying mvcmaildid-go [$DID_EMAIL] . . ."
./network.sh deployCC -ccn $DID_EMAIL -ccp ../asset-transfer-basic/mvcdidemail-go/ -ccl go
echo "Completed : Deployed mvcdidemail-go [$DID_EMAIL]! "

echo "Deploying mvcdidkeyids-go [$DID_KEYID] . . ."
./network.sh deployCC -ccn $DID_KEYID -ccp ../asset-transfer-basic/mvcdidkeyids-go/ -ccl go
echo "Completed : Deployed mvcdidkeyids-go [$DID_KEYID]! "

echo "Deploying mvckeyidpubkey-go [$KEYID_PUB] . . ."
./network.sh deployCC -ccn $KEYID_PUB -ccp ../asset-transfer-basic/mvckeyidpubkey-go/ -ccl go
echo "Completed : Deployed mvckeyidpubkey-go $KEYID_PUB! "

echo "Inocation calls . . ."

echo "Invoking mvcmaildid-go [$EMAIL_DID] . . ."
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n $EMAIL_DID --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'
echo "Completed : Invoked mvcmaildid-go [$EMAIL_DID] ! "

echo "Invoking mvcmaildid-go [$DID_EMAIL] . . ."
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n $DID_EMAIL --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'
echo "Completed : Invoked mvcdidemail-go $DID_EMAIL ! "

echo "Invoking mvcdidkeyids-go [$DID_KEYID] . . ."
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n $DID_KEYID --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'
echo "Completed : Invoked mvcdidkeyids-go $DID_KEYID ! "

echo "Invoking mvckeyidpubkey-go $KEYID_PUB . . ."
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n $KEYID_PUB --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'
echo "Completed : Invoked mvckeyidpubkey-go $KEYID_PUB ! "


echo "Fetching all assets from mvcmaildid-go [$EMAIL_DID] . . ."
peer chaincode query -C mychannel -n $EMAIL_DID -c '{"Args":["GetAll"]}'
echo "Completed : Fetched all assets from mvcmaildid-go [$EMAIL_DID]! "
echo "Fetching all assets from  mvcmaildid-go [$DID_EMAIL] . . ."
peer chaincode query -C mychannel -n $DID_EMAIL -c '{"Args":["GetAll"]}'
echo "Completed : Fetched all assets from  mvcdidemail-go [$DID_EMAIL]! "
echo "Fetching all assets from  mvcdidkeyids-go [$DID_KEYID] . . ."
peer chaincode query -C mychannel -n $DID_KEYID -c '{"Args":["GetAll"]}'
echo "Completed : Fetched mvcdidkeyids-go [$DID_KEYID]! "
echo "Fetching all assets from  mvckeyidpubkey-go [$KEYID_PUB] . . ."
peer chaincode query -C mychannel -n $KEYID_PUB -c '{"Args":["GetAll"]}'
echo "Completed : Fetched mvckeyidpubkey-go $KEYID_PUB! "
echo "Completed ! "
