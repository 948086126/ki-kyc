# ki-kyc



### 本案例是启动一个fabric服务  1.4.2 版本



### 机构名称： 北京图书馆

############### 服务器须知： 1. 阿里云服务器 
                            2. 按照时间付费，跑一遍不会花太多时间
                            3. 香港服务器，香港服务器，香港服务器，不要买国内的服务器，新加坡，日本，美国，新西兰随便买，就是不要买国内的。
                            4. 有任何问题直接 电联：18850272729 ,24小时解决问题。期望来电。



# 1.更新 yum 
yum update -y
# 2.安装 Git
yum install git
# 3.下载 Golang 源码包，版本为 1.14.3
wget https://studygolang.com/dl/golang/go1.14.3.linux-amd64.tar.gz
# 4.解压源码包至用户级的程序目录 /usr/local
sudo tar -C /usr/local -xzf go1.14.3.linux-amd64.tar.gz
# 5.打开配置文件
vi /etc/profile
# 6.写入
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$HOME/go/bin
# 7.刷新
source /etc/profile
# 8.验证 golang 是否安装成功
go version
# 9.安装 docker 
yum -y install docker
# 10.重启docker
systemctl start docker
# 11.验证 docker
docker version
# 12.添加用户组
# 如果还没有docker group就添加一个：
sudo groupadd docker
#安装好docker之后， 可能存在只有root用户才能调用的问题， 需要将当前用户加入docker组中：
sudo gpasswd -a ${USER} docker
sudo systemctl restart docker
# 13.重启docker
systemctl daemon-reload
systemctl restart docker.service
# 14.安装 docker-compose 版本为：1.8.0
curl -L https://github.com/docker/compose/releases/download/1.8.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
# 15.添加运行权限
chmod +x /usr/local/bin/docker-compose
# 16.验证 docker-compose
docker-compose version
# 17.安装 Hyperldger Fabric 
# 18.创建源代码目录
mkdir -p $GOPATH/src/github.com/hyperledger
# 19.切换到目录，下载 Hyperldger Fabric 源代码
cd $GOPATH/src/github.com/hyperledger && git clone https://github.com/hyperledger/fabric.git
# 切换版本
cd $GOPATH/src/github.com/hyperledger/fabric/
# 查看在哪个分支
git branch -a
# 查看分支
git tag
# 要用 1.4.4 版本，切换至1.4.4 
git checkout v1.4.4
# 20.验证是否切换成功
git branch
# 21.下载镜像 切换到目录
cd $GOPATH/src/github.com/hyperledger/fabric/scripts/
# 修改参数： bootrstrap.sh
DOCKER=true
SAMPLES=false
BINARIES=false
# 22.执行脚本下载镜像
./bootrstrap.sh
# 23.查看是否下载镜像
docker images
# 24.下载 bin 二进制工具.要什么版本记得修改参数。这里是1.4.4 ，如果是别的版本，请自己修改
wget https://github.com/hyperledger/fabric/releases/download/v1.4.4/hyperledger-fabric-linux-amd64-1.4.4.tar.gz
# 25.解压至  /usr/local/bin
tar -zxvf  hyperledger-fabric-linux-amd64-1.4.4.tar.gz 
cd  bin/
sudo cp *   /usr/local/bin
# 26.验证二进制工具版本
cd  /usr/local/bin
peer version
# 27.下载 fabric-simples 注意：cd 到根目录操作,不要在fabric目录下操作
git clone https://github.com/hyperledger/fabric-samples.git
# 28.一定要切换版本，和fabric 的版本一样,切换版本
git  branch -a
git tag 
git checkout xxxx
# 29.启动demo
cd  fabric-samples/first-network
# 30.生成证书
./byfn.sh generate
# 31.启动网络
./byfn.sh up
# 等待结果 出现 GOOD 表示网络没有问题
# 32.查看网络 容器状态是否都为 up
docker ps -a 
# 33.停止网络
./byfn.sh down
# 34.结束

