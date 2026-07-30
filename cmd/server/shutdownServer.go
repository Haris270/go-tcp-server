package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
)

func shutdownServer(listener net.Listener) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()
	err := listener.Close()
	fmt.Printf("listener.Close returned : %v", err)

}
