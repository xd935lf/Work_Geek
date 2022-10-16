使用说明

本程序提供一个httpserver的docker镜像，运行起来达到以下效果：

接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回 200
作业步骤




# lvfeng0935/httpsever为自己在docker镜像仓库中创建的仓库
docker build  -t lvfeng0935/httpsever .


执行下面命令推送到镜像仓库

# 需要先通过docker login登录自己的账户
docker push lvfeng0935/httpsever

![docker Hub](https://user-images.githubusercontent.com/110708332/196046168-74b1b4f5-81c6-48d3-a4ed-c519e39e9f3c.png)

my image in docker hub：https://hub.docker.com/repository/docker/lvfeng0935/httpsever

本地运行镜像

docker run -d -p 8080:80 lvfeng0935/httpsever

![status code 200](https://user-images.githubusercontent.com/110708332/196046198-b88da257-2528-4eef-99ae-d48adf0b23de.png)

使用curl验证httpserver应用已经生效

curl -v 127.0.0.1:8080/test
curl -v 127.0.0.1:8080/healthz
![healthz](https://user-images.githubusercontent.com/110708332/196046211-e26bb8bd-f054-496f-9e5a-41c6909e906d.png)


获取容器pid并且通过nsenter查看容器网络配置

# 获取容器Pid
docker inspect $(docker ps | grep http-server | awk '{print $1}') | grep Pid
# 根据获取到的pid查看容器的网络配置,其中17478为前面查询到的Pid
nsenter -t 17478 -n ip addr


