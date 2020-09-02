# 以golang官方镜像为基础镜像
FROM golang
# 设置工作目录，命令都在该目录下执行
WORKDIR /resource-pool
# 把项目拷贝到镜像中
COPY . /resource-pool
# 初始化当前目录module
RUN go mod init resource-pool
# 启用go module支持
RUN go env -w GO111MODULE=on
# 设置代理
RUN go env -w GOPROXY=https://goproxy.io,direct
# 暴露端口，声明容器内服务使用端口为80
EXPOSE 80
# 启动服务
ENTRYPOINT ["go","run", "webserver.go"]
