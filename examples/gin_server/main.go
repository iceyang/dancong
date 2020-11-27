package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iceyang/dancong"
	"github.com/iceyang/dancong/pkg/runner"
	"github.com/sirupsen/logrus"
)

func main() {
	// Custom Logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	opts := dancong.Options(
		// Add http Handler for http service
		dancong.WithBean(
			func() http.Handler {
				e := gin.Default()
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
		dancong.WithLogger(logger),
		// Load config fle
		dancong.WithConfig("./dancong.yaml"),
	)

	dc := dancong.New(opts)

	dc.Run()
}
