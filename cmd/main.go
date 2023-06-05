package main

import (
	"context"
	"navy/internal/common/infra"
	"navy/internal/common/log"
	"navy/internal/port/api"
)

func main() {
	log.Init()

	infra.WithGracefulShutdown(func(ctx context.Context) {
		if err := api.Run(ctx); err != nil {
			panic(err)
		}
	})
}
