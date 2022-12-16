package modules

import (
	"context"
	"gogenerate/bootstrap"
	"gogenerate/infrastructure"
	"time"
)

const idleTimeout = 5 * time.Second

func Webservice() {
	infrastructure.Load()
	ctx := context.Background()
	infrastructure.Open()

	bootstrap.Dispatch(ctx)
}
