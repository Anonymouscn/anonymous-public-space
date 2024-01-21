#!/bin/bash
# 编译项目
GOOS=linux GOARCH=amd64 go build -o bin/public-system-app main.go
# docker 打包
docker build -t public-system-backend .