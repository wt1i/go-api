.PHONY: all build gotool clean help
BINARY=go-api
all: gotool build
build:
	go mod tidy
	go build -o ${BINARY}
install:
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init
	@go mod tidy
test:
	go test ./...
gotool:
	go fmt ./...
	golangci-lint run ./...
run:
	go run main.go
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
help:
	@echo "make build - 编译 Go 代码, 生成二进制文件（流水线使用）"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make install - 安装依赖包"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'golangci-lint'"