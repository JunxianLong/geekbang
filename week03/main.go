package main

import (
	"context"
	"fmt"
	pkgErrors "github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	quit := make(chan struct{})

	server := http.Server{Addr: "127.0.0.1:8080"}

	// 启动Web服务
	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("quit")
		case <-quit:
			log.Println("server will out...")
		}
		return server.Shutdown(ctx)
	})

	g.Go(func() error {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			fmt.Println("quit")
		case sig := <-signals:
			fmt.Printf("receive signal:%v\n", sig)
		}
		return pkgErrors.New("QUIT")
	})

	err := g.Wait()
	fmt.Println(err)
	time.Sleep(10 * time.Second)

}
