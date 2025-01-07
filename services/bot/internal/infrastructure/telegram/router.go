package telegram

import "errors"

type HandlerType int

const (
	HandlerTypeCmd HandlerType = iota
	HandlerTypeText
)

type Option struct {
	Type  HandlerType
	Route string
}

type RouterHandlerFunc func(req *RequestCtx) error

type Router struct {
	handlers map[Option]RouterHandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[Option]RouterHandlerFunc),
	}
}

func (router *Router) AddHandler(handlerType HandlerType, route string, handler RouterHandlerFunc) {
	option := Option{
		Type:  handlerType,
		Route: route,
	}
	router.handlers[option] = handler
}

func (router *Router) GetHandler(cmd Option) (RouterHandlerFunc, error) {
	if handler, ok := router.handlers[cmd]; ok {
		return handler, nil
	}

	return nil, errors.New("command not found")
}
