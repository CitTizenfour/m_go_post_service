#!/bin/bash
#CURRENT_DIR=$(pwd)
#echo $CURRENT_DIR
#for x in $(find ${CURRENT_DIR}/protos/* -type d); do
#  protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=plugins=grpc:${CURRENT_DIR} ${x}/*.proto
#done


CURRENT_DIR=$1
echo $CURRENT_DIR
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done