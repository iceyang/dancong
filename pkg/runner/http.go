package runner

import (
	"context"
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

func (runner *httpRunner) PreStart(dc *dancong.Dancong) interface{} {
	return func(handler http.Handler) {
		runner.handler = handler
	}
}

func (runner *httpRunner) Start(dc *dancong.Dancong) error {
	ctx := dc.GetContext()
	v, _ := ctx.GetConfig("http.addr")
	addr := v.(string)
	dc.GetLogger().Infof("[Dancong] Starting HTTP server. Listening at %s\n", addr)
	go func() {
		runner.server = &http.Server{
			Addr:    addr,
			Handler: runner.handler,
		}
		err := runner.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			dc.GetLogger().Fatalf("listen: %s\n", err)
		}
	}()
	return nil
}

func (runner *httpRunner) Stop(dc *dancong.Dancong) error {
	return runner.server.Shutdown(context.TODO())
}
