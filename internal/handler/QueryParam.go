package handler

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func QueryParamInt(ctx echo.Context, name string, defaultValue int) (int, error) {
	var page int = defaultValue

	pageValue := ctx.QueryParam(name)

	if pageValue == "" {
		return page, nil
	}

	page, err := strconv.Atoi(pageValue)

	if err != nil {
		page = defaultValue
	}

	return page, err
}
