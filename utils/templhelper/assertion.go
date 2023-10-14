package templhelper

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderAssert(c echo.Context, statusCode int, name string, component templ.Component) error {
	return c.Render(statusCode, name, component)
}
