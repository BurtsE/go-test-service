package router

import (
	"test-service/internal/config"
	"test-service/internal/service"

	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Router struct {
	router     *router.Router
	srv        *fasthttp.Server
	logger     *logrus.Logger
	msgService service.MessageService
	port       string
}

func NewRouter(logger *logrus.Logger, msgService service.MessageService, cfg config.Config) *Router {
	router := router.New()
	r := &Router{
		router:     router,
		logger:     logger,
		msgService: msgService,
		port:       cfg.Service.Port,
		srv: &fasthttp.Server{
			Handler: router.Handler,
		},
	}
	RegisterMessage(r)
	r.router.GET("/status", statusHandler)
	return r
}

func (r *Router) Start() error {
	return r.srv.ListenAndServe(r.port)
}

func (r *Router) Shutdown() error {
	return r.srv.Shutdown()
}

func statusHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
}
