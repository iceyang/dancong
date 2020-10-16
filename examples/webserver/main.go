package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iceyang/dancong"
	"github.com/iceyang/dancong/pkg/runner"
	"google.golang.org/grpc"
)

func main() {
	opts := dancong.Options(
		dancong.WithBean(
			func() http.Handler {
				e := gin.New()
				e.Any("/", func(c *gin.Context) {
					c.JSON(200, &gin.H{
						"Hello": "Dancong App",
					})
				})
				return e
			},
			func() *grpc.Server {
				return grpc.NewServer()
			},
		),
		dancong.WithRunner(
			runner.HttpRunner,
			runner.GrpcRunner,
		),
		dancong.WithConfig("./config.yaml"),
	)

	dc := dancong.New(opts)

	dc.Run()
}
