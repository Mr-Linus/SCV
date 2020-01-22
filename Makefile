
all: local

local:
	GOOS=linux GOARCH=amd64 go build  -o=scv .

build:
	sudo docker build --no-cache . -t registry.cn-hangzhou.aliyuncs.com/geekcloud/scv

push:
	sudo docker push registry.cn-hangzhou.aliyuncs.com/geekcloud/scv

format:
	sudo gofmt -l -w .
clean:
	sudo rm -f scv

