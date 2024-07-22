package router

import (
	"test-service/internal/service"

	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
)

type Router struct {
	router     *router.Router
	logger     *logrus.Logger
	msgService service.MessageService
}

func NewRouter(logger *logrus.Logger, msgService *service.MessageService) *Router {
	r := &Router{
		router:     router.New(),
		logger:     logger,
		msgService: *msgService,
	}
	RegisterMessage(r)
	return r
}
