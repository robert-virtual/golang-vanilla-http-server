package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"vanila_http_server/server"
	// "vanila_http_server/server"
)

func main() {
	ctx := context.Background()
	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	server := server.New(":5000")
	// fmt.Println("Server running on port 5000...")
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	fmt.Println("Server running...")
	<-serverDoneChan
	server.Shutdown(ctx)
	fmt.Println("Server stopped")
}
