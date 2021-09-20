package infrastructure

import (
	controller "github.com/fizzfuzzHK/line_bot_fav/controller"
	echo "github.com/labstack/echo/v4"
)

type Router struct {
	e  *echo.Echo
	lc *controller.LineBotController
}

func NewRouter(e *echo.Echo, lc *controller.LineBotController) *Router {
	return &Router{e: e, lc: lc}
}

func (r *Router) Init() {
	r.e.POST("/callback", r.lc.HandleEvents())
}
