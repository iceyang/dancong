package runner

import (
	"context"
	"log"
	"net/http"

	"github.com/iceyang/dancong"
)

const HttpRunner = "httpRunner"

type httpRunner struct {
	BaseRunner

	handler http.Handler
	server  *http.Server
}

func init() {
	dancong.RegisterRunner(HttpRunner, &httpRunner{})
}

func (runner *httpRunner) Init(ctx *dancong.Context) interface{} {
	return func(handler http.Handler) {
		runner.handler = handler
	}
}

func (runner *httpRunner) Start(ctx *dancong.Context) error {
	v, _ := ctx.GetConfig("http.addr")
	addr := v.(string)
	log.Printf("[Dancong] Starting HTTP server. Listening at %s\n", addr)
	go func() {
		runner.server = &http.Server{
			Addr:    addr,
			Handler: runner.handler,
		}
		err := runner.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return nil
}

func (runner *httpRunner) Stop(ctx *dancong.Context) error {
	return runner.server.Shutdown(context.TODO())
}
