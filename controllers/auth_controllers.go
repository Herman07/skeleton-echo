package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"skeleton-echo/request"
	"skeleton-echo/services"
	"skeleton-echo/utils/session"
)

type FrontAuthController struct {
	BaseFrontendController
	Controller
	service *services.AuthService
}

func NewAuthController(services *services.AuthService) FrontAuthController {
	return FrontAuthController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Login",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *FrontAuthController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Login",
		"link": "/login",
	}
	return Render(ctx, "Login", "auth/login.html", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}
func (c *FrontAuthController) Login(ctx echo.Context) error {
	var dataReq request.LoginRequest

	if err := ctx.Bind(&dataReq); err != nil {
		return ctx.Redirect(http.StatusFound, "/inventaris/v1/login")
	}
	data, err := c.service.Login(dataReq.Username)
	if err != nil {
		return ctx.JSON(400, echo.Map{"message": "Gagal Login"})
	}
	userInfo := session.UserInfo{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		TypeUser: data.TypeUser,
	}
	fmt.Println("masuk sini gak ", userInfo)
	if err := session.Manager.Set(ctx, session.SessionId, &userInfo)
	err != nil {
		return ctx.Redirect(http.StatusFound, "/inventaris/v1/login")
	}
	fmt.Println("masuk sini gak ", dataReq)
	return ctx.Redirect(http.StatusFound, "/inventaris/v1/admin")
}
func (c *FrontAuthController) Logout(ctx echo.Context) error {
	err := session.Manager.Delete(ctx, session.SessionId)
	if err != nil {
		return ctx.Redirect(302,  "/inventaris/v1/admin")
	}
	return ctx.Redirect(http.StatusFound,  "/login")
}
