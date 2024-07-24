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
	r := &Router{
		router:     router.New(),
		logger:     logger,
		msgService: msgService,
		port:       cfg.Service.Port,
		srv:        &fasthttp.Server{},
	}
	RegisterMessage(r)
	return r
}

func (r *Router) Start() error {
	return r.srv.ListenAndServe(r.port)
}

func (r *Router) Shutdown() error {
	return r.srv.Shutdown()
}
