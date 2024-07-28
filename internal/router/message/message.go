package router

import (
	"test-service/internal/converter"

	"github.com/valyala/fasthttp"
)

type messageImpl struct {
	r *Router
}

func RegisterMessage(r *Router) {
	msgImpl := messageImpl{r}
	r.router.POST("/message", msgImpl.receiveMessage)
	r.router.GET("/messages/stats", msgImpl.getStatistics)
}

func (m *messageImpl) receiveMessage(ctx *fasthttp.RequestCtx) {
	data := ctx.Request.Body()
	msg, err := converter.JsonToMessage(data)
	if err != nil {
		m.r.logger.Errorf("could not read input message")
		ctx.SetStatusCode(500)
		return
	}

	err = m.r.msgService.SaveMessage(msg)
	if err != nil {
		m.r.logger.Errorf("could not save message: %s", err.Error())
		ctx.SetStatusCode(500)
		return
	}
}

func (m *messageImpl) getStatistics(ctx *fasthttp.RequestCtx) {
	stats, err := m.r.msgService.GetStatistics()
	if err != nil {
		m.r.logger.Errorf("could not get statistics: %s", err.Error())
		ctx.SetStatusCode(500)
		return
	}

	data, err := converter.StatisticsToJson(stats)
	if err != nil {
		m.r.logger.Errorf("could not create response body: %s", err.Error())
		ctx.SetStatusCode(500)
		return
	}
	ctx.Response.AppendBody(data)
}
