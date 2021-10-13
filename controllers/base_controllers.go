package controllers

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseFrontendController struct {
	Title       string
	Menu        string
	BreadCrumbs []map[string]interface{}
}

func Render(ctx echo.Context, title string, view string, activeMenu string,
	breadcrumbs []map[string]interface{}, data interface{}) error {
	return echoview.Render(ctx, http.StatusOK, view, echo.Map{
		"title":        title,
		"activeMenu":   activeMenu,
		"breadCrumbs":  breadcrumbs,
		"ctx":          ctx,
		"data":         data,
	})
}
