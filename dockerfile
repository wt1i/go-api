# 使用官方 Golang 镜像作为基础镜像  
FROM golang:1.21-alpine3.19
  
# 设置工作目录  
WORKDIR /app  
  
# 将当前目录下的所有文件复制到容器中的 /app 目录下  
COPY . /app  
  
# 设置环境变量  
ENV CONFIG_FILE_PATH /app/config.yaml  
ENV GOPATH /go  
ENV PATH $PATH:/go/bin  
  
# 安装依赖   
RUN go mod tidy  
  
# 暴露容器端口（根据实际需求修改）  
EXPOSE 8080  
  
# 设置容器启动命令  
CMD ["go", "run", "main.go"]