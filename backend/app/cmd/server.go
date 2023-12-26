package main

import (
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/timeout"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/minomusigk/tech-blog-summary/backend/app/presentation/controller"
	"github.com/minomusigk/tech-blog-summary/backend/oapi"
	"go.uber.org/zap"
)

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://83c5b91265c7693ea10e3ad6609bd2af@o4506426584596480.ingest.sentry.io/4506426586497024",
		EnableTracing: true,
		// パフォーマンス計測に必要なサンプリング数を節約のため10％に指定
		TracesSampleRate: 0.1,
	}); err != nil {
		log.Printf("Sentry initialization failed: %v", err)
	}

	r := gin.New()

	logger, _ := zap.NewProduction()

	// TODO: RequestごとのIDを取得する
	// https://github.com/gin-contrib/zap?tab=readme-ov-file#custom-zap-fields
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	r.Use(timeoutMiddleware())
	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	))

	var pingController controller.PingController

	oapi.RegisterHandlers(r, &pingController)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func timeoutResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(500*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}
