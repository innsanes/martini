package martini

import "github.com/gin-gonic/gin"

type Engine[Resp interface{}, E error] struct {
	engine     *gin.Engine
	logger     gin.HandlerFunc
	resHandler ResHandler[Resp, E]
	mode       string
}

func NewEngine[Resp interface{}, E error](resHandler ResHandler[Resp, E]) *Engine[Resp, E] {
	ret := &Engine[Resp, E]{
		engine:     gin.New(),
		logger:     gin.Logger(),
		resHandler: resHandler,
	}
	return ret
}

func (e *Engine[Resp, E]) BaseRouter(middlewares []gin.HandlerFunc, handles ...*RouterHandler[Resp, E]) *Router[Resp, E] {
	router := &Router[Resp, E]{
		relativePath: "",
		middlewares:  middlewares,
		handlers:     handles,
	}
	router.resHandler = e.resHandler
	router.routerGroup = e.engine.Group("")
	router.routerGroup.Use(router.middlewares...)
	router.handles()
	return router
}
