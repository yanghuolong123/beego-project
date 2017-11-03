#!/bin/bash

HOST="192.168.2.95"
USERNAME="yanghuolong"
PASSWD=""
PORT="22"
DIR="~"

echo $USERNAME@$HOST:$DIR

ssh -p $PORT $USERNAME@$HOST "killall go-web" 
echo "停止旧有项目!"


ssh -p $PORT $USERNAME@$HOST "rm -rf $DIR/*" 
echo "删除旧有目录完成!"

scp -P$PORT go-web $USERNAME@$HOST:$DIR
scp -P$PORT -r views $USERNAME@$HOST:$DIR
scp -P$PORT -r static $USERNAME@$HOST:$DIR
scp -P$PORT -r conf $USERNAME@$HOST:$DIR

echo "部署完成!"

ssh -p $PORT $USERNAME@$HOST "nohup $DIR/go-web > /dev/null &" 

echo "项目启动成功!"
