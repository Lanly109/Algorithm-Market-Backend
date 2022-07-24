# Algorithm-Market-Backend

base on [Singo](https://github.com/Gourouting/singo)

## Dependence

后端评测依赖[Judger](https://github.com/QingdaoU/Judger)，须事先安装
```bash
sudo apt-get install gcc g++ make cmake libseccomp-dev

git clone https://github.com/QingdaoU/Judger
cd Judger
mkdir build && cd build && cmake .. && make && sudo make install
``` 

`svg`转`png`依赖`python3`的`cairosvg`库，须事先为`root`用户安装，且必须在`root`的环境变量里。
```bash
sudo pip3 install cairosvg
``` 

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行

```shell
sudo go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)

## 编译运行

因评测机需要更改程序运行用户组，需要`root`权限

```bash
go build -o main .
sudo ./main
``` 
