FROM golang:1.21

LABEL maintainer="Ving <ving@nivin.cn>"

# 设置 Go Proxy，避免墙的问题
ENV GOPROXY=https://goproxy.cn,direct

# 切换到 root 用户以便执行安装命令
USER root

# 设置国内的 apt 镜像源，安装必要的软件包
RUN echo "deb http://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main" > /etc/apt/sources.list && \
    echo "deb http://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main" >> /etc/apt/sources.list && \
    echo "deb http://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main" >> /etc/apt/sources.list && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    git \
    vim && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

# 安装 goctl
RUN GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest

# 安装 protoc（Protocol Buffers）
RUN curl -L -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip && \
    unzip -d /tmp/protoc /tmp/protoc.zip && \
    mv /tmp/protoc/bin/protoc $GOPATH/bin

# 安装 protoc-gen-go 插件
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 安装 protoc-gen-go-grpc 插件
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 将 $GOPATH/bin 添加到环境变量 PATH 中
ENV PATH=$GOPATH/bin:$PATH

# 设置工作目录
WORKDIR /usr/src/code

# 暴露多个端口
EXPOSE 8000
EXPOSE 8001
EXPOSE 8002
EXPOSE 8003
EXPOSE 9000
EXPOSE 9001
EXPOSE 9002
EXPOSE 9003

