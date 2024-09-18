package martini

import "github.com/gin-gonic/gin"

type Router[Resp interface{}, E error] struct {
	routerGroup  *gin.RouterGroup
	middlewares  []gin.HandlerFunc
	resHandler   ResHandler[Resp, E]
	children     []*Router[Resp, E]
	relativePath string
	handlers     []*RouterHandler[Resp, E]
}

type ResHandler[Resp interface{}, E error] interface {
	ErrHandle(e E)
	RespHandle(resp Resp)
}

// NewRouter 默认是没有生成 router group 的 需要在上一级中传入
func NewRouter[Resp interface{}, E error](path string, middlewares []gin.HandlerFunc, handles ...*RouterHandler[Resp, E]) *Router[Resp, E] {
	router := &Router[Resp, E]{
		relativePath: path,
		middlewares:  middlewares,
		handlers:     handles,
	}
	return router
}

func Middlewares(middlewares ...gin.HandlerFunc) []gin.HandlerFunc {
	return middlewares
}

func (r *Router[Resp, E]) Children(children ...*Router[Resp, E]) {
	r.children = children
	for _, router := range r.children {
		router.routerGroup = r.routerGroup.Group(router.relativePath)
		// 向下传递
		router.resHandler = r.resHandler
	}
}

func (r *Router[Resp, E]) warp(h HandlerFunc[Resp, E]) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := h(&Context{
			Context: c,
		})
		if err != nil {
			r.resHandler.ErrHandle(err)
			return
		}
		r.resHandler.RespHandle(resp)
	}
}

func (r *Router[Resp, E]) handles() {
	for _, handle := range r.handlers {
		r.routerGroup.Handle(handle.method, handle.path, r.warp(handle.handler))
	}
}

type RouterHandler[Resp interface{}, E error] struct {
	method  string
	path    string
	handler HandlerFunc[Resp, E]
}

func NewRouterHandler[Resp interface{}, E error](method string, path string, handler HandlerFunc[Resp, E]) *RouterHandler[Resp, E] {
	return &RouterHandler[Resp, E]{
		method:  method,
		path:    path,
		handler: handler,
	}
}
