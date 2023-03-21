#!/usr/bin/env bash

image="bbs-go"
imagePort="13180"

# 打包 go 的时候自带动态包
echo "-----开始打包-----"
echo "1、删除 main 文件"
rm -f main

echo "2、安装依赖"
go env -w GO111MODULE=on
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go get

echo "3、打包 main"
CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o main .

# 删除之前的镜像
echo "4、停止并删除之前的容器、镜像"
docker container stop "$image"
docker container rm "$image"
docker image rm "$image"

# 打包新镜像
echo "5、打包生成新的镜像"
docker build -t "$image" .
docker container run --name "$image" -dit -p "$imagePort":7001 -v /etc/localtime:/etc/localtime -v $(dirname "$PWD")/log-"$image":/log "$image"

echo "-----打包完成-----"
