# Work_Geek
极客时间云原生训练营提交作业专用仓库

一、作业要求：
接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回 200
二、效果图

![效果](https://user-images.githubusercontent.com/110708332/194763632-5a0ad73f-293d-452f-b5ba-615e4ce56022.png)

三、存在问题
在testgetcode 中尝试获取http statuscode url为：http://www.baidu.com时可以获取 
url为：http://localhost/ 时报错
遂此功能尚未添加
