# PATH 설정
```
export PATH="$PATH:$(go env GOPATH)/bin"
```
# Proto compile
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```