
all: local

local:
	GOOS=linux GOARCH=amd64 go build  -o=scv .

build:
	docker build --no-cache . -t registry.cn-hangzhou.aliyuncs.com/geekcloud/scv

push:
	docker push registry.cn-hangzhou.aliyuncs.com/geekcloud/scv

clean:
	rm -f scv

