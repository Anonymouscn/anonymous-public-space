#!/bin/bash

# 部署模式 [development / test / production]
mode=production
# 容器名称
container_name=public-system

yarn build --mode $mode
docker build -t $container_name .