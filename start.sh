#!/bin/bash
DIR="$(pwd)"
printf "\nCurrent directory: $DIR \n"

# --------------------- #
# ----- FUNCTIONS ----- #
# --------------------- #

clear_all()
{
    docker rm -f $(docker ps -a -q) 2>/dev/null
    printf "All docker containers removed\n"
    docker rmi $(docker images | grep "dev-" | awk '{print $1}') 2>/dev/null
    docker rmi $(docker images -qf "dangling=true") 2>/dev/null
    printf "All docker useless images removed\n"
}

# ----------------- #
# ----- START ----- #
# ----------------- #

clear_all

# run docker-compose
docker-compose up -d 2>/dev/null
#docker-compose -f docker-compose-4peer-pbft.yaml up -d 2>/dev/null
printf "Starting docker containers...\n"
sleep 10
printf "Docker containers up and running\n"

# Compile and build chaincode
cd $DIR/chaincode
go build

# Register Chaincode on peer.
printf "Launch chaincode\n"
export CORE_CHAINCODE_ID_NAME="mycc"
export CORE_PEER_ADDRESS="0.0.0.0:7051"
#/Users/marckx/go_workspace/src/github.com/chaincode_example02/chaincode_example02 &
$DIR/chaincode/chaincode &
sleep 1

# start server, catch ctrl+c to clean up
read -n1 -r -p "Press space to quit..." key
if [ "$key" = '' ]; then
    echo "Stop the process"
    clear_all
fi

exit 0
