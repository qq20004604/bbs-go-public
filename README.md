# bbs-go

用go写的极简论坛（内部版）

## 1、运行前配置

### 1.1 数据库准备

1. 需要准备两个数据库，MySQL和Redis。这里默认你会玩这2个，不会的话不管埋；
2. 具体配置请更改 ``config/config.yml`` 里对应的内容，测试环境和生产模式视为两个不同的环境，支持根据配置加载不同链接；
3. 然后根据密码，修改并执行 ``db/数据库初始化.md`` 这个文件里的 mysql 脚本，创建 database；
4. 从而 MySQL 数据库创建完毕。
5. Redis 数据库无需特殊配置；

> 其他

1. 数据库表会自动创建；
2. 默认情况下，dev会自动创建超级管理员用户，

### 1.2 安装项目依赖

1. down下来所有项目文件
2. 设置阿里源：``go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/``
3. 安装依赖：``go get``
4. 运行项目
5. 其他：如果需要打包，执行 ``go build``，打包后如果需要在生产环境运行，输入 ``ENVIRONMENT=PROD ./main``

### 1.3 普通环境运行

1. 使用 GoLand ，直接运行 ``./main.go`` 即可
2. 手动命令行运行，略

### 1.4 打包后运行（生产环境）

1. 默认通过 Docker 来运行，clone 项目后，执行 ``./build.sh`` 即可自动打包；
2. 剩下的通过 Docker 来启动容器即可；

### 2、交流群

> QQ群：387017550

进群后联系群主即可

### 3、代码

> 前端：

https://github.com/qq20004604/web-for-bbs-go

> 后端

https://github.com/qq20004604/bbs-go-public