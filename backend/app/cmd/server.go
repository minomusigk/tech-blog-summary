package main

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/minomusigk/tech-blog-summary/backend/app/presentation/controller"

	sentrygin "github.com/getsentry/sentry-go/gin"
)

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://83c5b91265c7693ea10e3ad6609bd2af@o4506426584596480.ingest.sentry.io/4506426586497024",
		EnableTracing: true,
		// パフォーマンス計測に必要なサンプリング数を節約のため10％に指定
		TracesSampleRate: 0.1,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}

	r := gin.Default()

	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	r.GET("/ping", controller.GetPing)
	r.Run() // listen and serve on 0.0.0.0:8080
}
