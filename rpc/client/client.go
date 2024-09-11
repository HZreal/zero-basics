package main

/**
 * @Author elastic·H
 * @Date 2024-09-05
 * @File: client.go
 * @Description:
 */

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"zero-basics/rpc/greet"
)

func main() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/greet-client.yaml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	client := greet.NewGreetClient(conn.Conn())
	resp, err := client.Ping(context.Background(), &greet.Request{})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resp)
}
