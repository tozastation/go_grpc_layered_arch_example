# gRPC-Trainning
## 準備(Prepare)
### Protocol Compilerのインストール
- GIT_TAG="v1.2.0" # change as needed
- `go get -d -u github.com/golang/protobuf/protoc-gen-go`
- `git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG`
- `go install github.com/golang/protobuf/protoc-gen-go`
### コードの自動生成
- `sh protoc.sh`
## 実行
### Docker
- `docker-compose build; docker-compose up db`
> DBが立ち上がったら
- `docker-compose up api`
### Kubernetes
- `kubectl create -f k8s`