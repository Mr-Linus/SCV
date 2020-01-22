
all: local

local:
	GOOS=linux GOARCH=amd64 go build  -o=scv .

build:
	sudo docker build --no-cache . -t registry.cn-hangzhou.aliyuncs.com/geekcloud/scv

push:
	sudo docker push registry.cn-hangzhou.aliyuncs.com/geekcloud/scv

clean:
	sudo rm -f scv

