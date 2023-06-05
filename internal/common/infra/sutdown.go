package infra

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// WithGracefulShutdown runs the given function and waits for a SIGINT, SIGTERM or SIGHUP signal to gracefully shutdown
func WithGracefulShutdown(f func(ctx context.Context)) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		cancelFunc()
	}()

	f(ctx)
}
