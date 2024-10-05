```
    goctl model mysql ddl -src ./model/pay.sql -dir ./model -c
    goctl api go -api ./api/pay.api -dir ./api
    goctl rpc protoc ./rpc/pay.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc
```
