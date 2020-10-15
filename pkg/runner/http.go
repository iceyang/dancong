package runner

import (
	"context"
	"log"
	"net/http"

	"github.com/iceyang/dancong"
)

type HttpRunner struct {
	Handler http.Handler

	server *http.Server

	// listening address
	addr string
}

func (runner *HttpRunner) Start(ctx *dancong.Context) error {
	v, _ := ctx.GetConfig("http.addr")
	addr := v.(string)
	log.Printf("Starting HTTP server. Listening at %s\n", addr)
	go func() {
		runner.server = &http.Server{
			Addr:    addr,
			Handler: runner.Handler,
		}

		err := runner.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return nil
}

func (runner *HttpRunner) Stop(ctx *dancong.Context) error {
	return runner.server.Shutdown(context.TODO())
}
