package main

import (
	"bks/bffH5/handler/api"
	"bks/common/config"
	pbUsers "bks/proto/users"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bks/bffH5/basic/router"
	_ "bks/common/init"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接 gRPC 服务
	conn, err := grpc.Dial(config.Cfg.Services.UsersServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	api.Client = pbUsers.NewUsersClient(conn)

	r := gin.Default()
	router.Router(r)
	c := cron.New()
	c.AddFunc("30 * * * *", api.HandlerExpireOrder)
	c.Start()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
