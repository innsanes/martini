package martini

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
}

type HandlerFunc[Resp any, E error] func(c *Context) (resp Resp, err E)
