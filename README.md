
kratos new babytree -r https://gitee.com/go-kratos/kratos-layout.git

$ cd babytree
$ go generate ./...
$ go build -o ./bin/ ./...
$ ./bin/babytree -conf ./configs

添加pb文件
kratos proto add api/helloworld/v1/student.proto


生成 proto 对应代码#
通过 make 命令生成：

make api
或者通过 kratos cli 生成：

kratos proto client api/helloworld/v1/student.proto

生成 Service 代码#
通过 proto 文件，直接生成对应的 Service 代码。使用 -t 指定生成目录：

kratos proto server api/helloworld/v1/student.proto -t internal/service

参考地址:
https://www.cnblogs.com/jiujuan/p/16331967.html

https://xie.infoq.cn/article/64641cdf2e6dcd91c1c43971c



# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```


