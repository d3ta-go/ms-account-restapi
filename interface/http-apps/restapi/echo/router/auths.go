package router

import (
	"github.com/d3ta-go/ms-account-restapi/interface/http-apps/restapi/echo/features/account"
	"github.com/labstack/echo/v4"
)

// SetAuths set Auths Router
func SetAuths(eg *echo.Group, f *account.FAuths) {

	gc := eg.Group("/auths")

	gc.POST("/register", f.RegisterUser)
	gc.GET("/registration/activate/:activationCode/:format", f.ActivateRegistration)
	gc.POST("/login", f.Login)
	gc.POST("/login-app", f.LoginApp)
}
