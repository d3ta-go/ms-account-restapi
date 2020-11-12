package router

import (
	"github.com/d3ta-go/ms-account-restapi/interface/http-apps/restapi/echo/features/openapis"
	"github.com/labstack/echo/v4"
)

// SetOpenAPI set OpenAPI Router
func SetOpenAPI(eg *echo.Group, f *openapis.FOpenAPIs) {
	eg.GET("/docs/openapi.yaml", f.GenOpenAPI)
	eg.GET("/docs/index.html", f.SwaggerUI)
	eg.Static("/docs/swagger-ui/assets", "./www/public/swagger-ui/assets")
}
