package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iceyang/dancong"
	"github.com/iceyang/dancong/pkg/runner"
	"google.golang.org/grpc"
)

func main() {
	e := gin.New()
	e.Any("/", func(c *gin.Context) {
		c.JSON(200, &gin.H{
			"Hello": "Dancong App",
		})
	})

	opts := dancong.Options(
		dancong.WithRunner(&runner.HttpRunner{
			Handler: e,
		}),
		dancong.WithRunner(&runner.GrpcRunner{
			Server: grpc.NewServer(),
			Addr:   ":10000",
		}),
		dancong.WithConfig("./config.yaml"),
	)

	dc := dancong.New(opts)

	dc.Run()
}
