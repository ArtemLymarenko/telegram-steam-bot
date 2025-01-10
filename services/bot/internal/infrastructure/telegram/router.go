package telegram

import (
	"errors"
	"log"
)

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
	handlers    map[Option]RouterHandlerFunc
	middlewares []RouterHandlerFunc
	inlineQuery RouterHandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers:    make(map[Option]RouterHandlerFunc),
		middlewares: make([]RouterHandlerFunc, 0),
	}
}

func (router *Router) HasHandler(option Option) bool {
	_, ok := router.handlers[option]
	return ok
}

func (router *Router) AddInlineQuery(handler RouterHandlerFunc) {
	if router.inlineQuery != nil {
		log.Fatal("AddInlineQuery already has a handler")
	}

	router.inlineQuery = handler
}

func (router *Router) GetInlineQuery() (RouterHandlerFunc, error) {
	if router.inlineQuery == nil {
		return nil, errors.New("GetInlineQuery not found in router")
	}

	return router.inlineQuery, nil
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

func (router *Router) UseMiddleware(middleware RouterHandlerFunc) {
	router.middlewares = append(router.middlewares, middleware)
}

func (router *Router) launchMiddlewares(ctx *RequestCtx) {
	ctx.currentIndex = 0
	ctx.abortProcessing = false

	for ctx.currentIndex < len(router.middlewares) {
		middleware := router.middlewares[ctx.currentIndex]
		err := middleware(ctx)
		if err != nil {
			log.Println(err)
		}

		if !ctx.abortProcessing {
			ctx.NextMiddleware()
		} else {
			break
		}
	}
}
