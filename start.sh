#!/usr/bin/env bash


### 1. 生成证书
cryptogen generate --config=./crypto-config.yaml

### 2. 生成配置文件
./generate.sh

### 3. 生成网络

# 初始化网络
./init.sh

# 实例化链码
./install.sh



