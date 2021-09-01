package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	_ "golang.org/x/net/context"
	"net/http"
)
import dapr "github.com/dapr/go-sdk/client"

var Client dapr.Client

func InitDapr() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	Client = client
}

func main() {
	InitDapr()

	r := gin.Default()
	v1 := r.Group("/api")
	Request(v1.Group("/"))
	Health(v1.Group("/"))

	_ = r.Run(":40002")
}

func Request(router *gin.RouterGroup) {
	ctx:=context.Background()
	router.POST("/request", func(context *gin.Context) {
		content := &dapr.DataContent{
			ContentType: "text/plain",
			Data:        []byte("hello"),
		}
		res, err := Client.InvokeMethodWithContent(ctx, "dapr-sample-java", "/sampleJavaApp/name", "post", content)
		if err != nil {
			log.Error().Err(err).Msg("")
			return
		}
		context.JSON(http.StatusOK, gin.H{"response": string(res)})
	})
}

func Health(router *gin.RouterGroup) {
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"response": "healthy"})
	})
}
