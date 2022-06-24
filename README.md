# grpc-connect-go
grpc-connect-go

## サーバー起動
```shell
$  go run cmd/server/main.go 

2022/06/24 22:22:13 path /hello.v1.HelloService/

-- クライアント起動時に下記を出力するよ
2022/06/24 22:22:33 Request headers:  map[Accept-Encoding:[identity] Content-Type:[application/grpc-web+proto] Grpc-Accept-Encoding:[gzip] User-Agent:[grpc-go-connect/0.1.1 (go1.18)]]
```

## クライアント起動
```shell
$ go run cmd/client/main.go 

2022/06/24 22:22:33 Hello, Jane!
```
