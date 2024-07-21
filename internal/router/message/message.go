package router

import (
	"test-service/internal/models"
	"test-service/internal/router"

	"github.com/valyala/fasthttp"
)

type Database interface {
	SaveMessage() error
	GetMessages() ([]models.Message, error)
}

func RegisterMessage(r *router.Router) {
	r.r.POST("message", receiveMessage)
	r.r.GET("/messages/stats", getStatistics)
}

func receiveMessage(ctx *fasthttp.RequestCtx) {

}

func getStatistics(ctx *fasthttp.RequestCtx) {

}
