package router

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Database interface {
	SaveMessage()
	GetMessages()
}

func registerMessage(r *router.Router) {
	r.POST("/message", receiveMessage)
	r.GET("/messages/stats", getStatistics)
}

func receiveMessage(ctx *fasthttp.RequestCtx) {

}

func getStatistics(ctx *fasthttp.RequestCtx) {

}