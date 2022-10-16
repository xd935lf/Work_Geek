FROM golang:alpine AS builder

# 配置编译环境
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 拷贝源代码到镜像中

WORKDIR /build

COPY main.go .

RUN go build -o httpserver main.go

FROM alpine

WORKDIR /app

ENV VERSION=1.0.0

COPY --from=builder /build/httpserver /app/httpserver


# 声明暴露的端口
EXPOSE 80

# 设置服务入口
ENTRYPOINT [ "./httpserver" ]
