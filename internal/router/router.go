package router

import (
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
)

type Router struct {
	r *router.Router
	l *logrus.Logger
}

func NewRouter(logger *logrus.Logger) *Router {
	r := router.New()
	return &Router{
		r: r,
		l: logger,
	}
}
