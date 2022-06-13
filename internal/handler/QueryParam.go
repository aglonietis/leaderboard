package handler

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func QueryParamInt(ctx echo.Context, name string, defaultValue int) int {
	page, err := strconv.Atoi(ctx.QueryParam(name))

	if err != nil {
		page = defaultValue
	}

	return page
}
