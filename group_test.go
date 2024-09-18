package martini_test

import (
	"martini"
	"testing"
)

func TestRouter(t *testing.T) {
	//router := NewRouter("/test", nil, NewRouterHandler("GET", "/test", func(c *Context) (resp string, err error) {
	//	return "test", nil
	//}))
	//router.warp(router.handlers[0].handler)(&gin.Context{})
	engine := martini.New()
}
