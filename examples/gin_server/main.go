package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iceyang/dancong"
	"github.com/iceyang/dancong/pkg/runner"
)

func main() {
	opts := dancong.Options(
		// Add http Handler for http service
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
		),
		// Add http service
		dancong.WithRunner(
			runner.HttpRunner,
		),
		// Load config fle
		dancong.WithConfig("./dancong.yaml"),
	)

	dc := dancong.New(opts)

	dc.Run()
}
