package openapis

import (
	"net/http"

	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features"
	"github.com/d3ta-go/system/system/handler"
	"github.com/labstack/echo/v4"
)

// NewOpenAPIs new FOpenAPIs
func NewOpenAPIs(h *handler.Handler) (*FOpenAPIs, error) {

	f := new(FOpenAPIs)
	f.SetHandler(h)

	return f, nil
}

// FOpenAPIs represent FOpenAPIs
type FOpenAPIs struct {
	features.BaseFeature
}

// SwaggerUI generate SwaggerUI html page
func (f *FOpenAPIs) SwaggerUI(c echo.Context) error {

	cfg, err := f.GetHandler().GetDefaultConfig()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"htmlTitle": cfg.Applications.Servers.RestAPI.Name,
	}

	return c.Render(http.StatusOK, "openapis/swagger-ui", data)
}

// GenOpenAPI generate openapi definition
func (f *FOpenAPIs) GenOpenAPI(c echo.Context) error {
	cfg, err := f.GetHandler().GetDefaultConfig()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"info.Title":              cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Title,
		"info.Description":        cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Description,
		"info.Contact.Name":       cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Contact.Name,
		"info.Contact.URL":        cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Contact.URL,
		"info.Contact.Email":      cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Contact.Email,
		"info.License.Name":       cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.License.Name,
		"info.License.URL":        cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.License.URL,
		"info.Version":            cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Version,
		"server.URL.Host.Default": c.Request().Host,
	}

	// return c.Blob(http.StatusOK, "text/plain; charset=utf-8", data)
	return c.Render(http.StatusOK, "openapi.yaml", data)
}
