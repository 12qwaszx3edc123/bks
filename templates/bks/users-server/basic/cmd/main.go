package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"bks/common/config"
	_ "bks/common/init"
	pbUsers "bks/proto/users"
	"bks/users-server/server"
)

func main() {
	addr := config.Cfg.Services.UsersServerAddr
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", addr)
	s := grpc.NewServer()
	pbUsers.RegisterUsersServer(s, &server.UsersService{})
	s.Serve(listener)
}
