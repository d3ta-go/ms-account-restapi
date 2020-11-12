package account

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/d3ta-go/ddd-mod-account/modules/account/application"
	appDTOAuth "github.com/d3ta-go/ddd-mod-account/modules/account/application/dto/auth"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features"
	captcha "github.com/d3ta-go/system/interface/http-apps/restapi/echo/features/system/captcha"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/response"
	"github.com/d3ta-go/system/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFAuths new  FAuths
func NewFAuths(h *handler.Handler) (*FAuths, error) {
	var err error

	f := new(FAuths)
	f.SetHandler(h)

	if f.accountApp, err = application.NewAccountApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FAuths feature Auths
type FAuths struct {
	features.BaseFeature
	accountApp *application.AccountApp
}

// RegisterUser register user
func (f *FAuths) RegisterUser(c echo.Context) error {

	req := new(appDTOAuth.RegisterReqDTO)
	if err := c.Bind(req); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	if !f.InTestMode() {
		decodedCaptcha, err := captcha.DecodeCaptcha(req.CaptchaID, c.RealIP())
		if err != nil {
			return response.FailWithMessage(err.Error(), c)
		}

		if !captcha.VerifyString(decodedCaptcha, req.Captcha) {
			return response.FailWithMessage("Captcha verification code error", c)
		}
	}

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.accountApp.AuthenticationSvc.Register(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.CreatedWithData(resp, c)
}

// ActivateRegistration activater user registration
func (f *FAuths) ActivateRegistration(c echo.Context) error {

	//params
	format := strings.ToLower(c.Param("format"))

	req := new(appDTOAuth.ActivateRegistrationReqDTO)
	req.ActivationCode = c.Param("activationCode")

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.accountApp.AuthenticationSvc.ActivateRegistration(req, i)
	if err != nil {
		if format == "html" {
			data := map[string]interface{}{
				"message": err.Error(),
			}
			return c.Render(http.StatusBadRequest, "auths/activate.registration", data)
		}

		return f.TranslateErrorMessage(err, c)
	}

	if format == "html" {
		data := map[string]interface{}{
			"message": fmt.Sprintf("Your user [%s] are now active", resp.Email),
		}
		return c.Render(http.StatusOK, "auths/activate.registration", data)
	}

	return response.OKWithData(resp, c)
}

// Login user Login
func (f *FAuths) Login(c echo.Context) error {

	req := new(appDTOAuth.LoginReqDTO)
	if err := c.Bind(req); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	if !f.InTestMode() {
		decodedCaptcha, err := captcha.DecodeCaptcha(req.CaptchaID, c.RealIP())
		if err != nil {
			return response.FailWithMessage(err.Error(), c)
		}

		if !captcha.VerifyString(decodedCaptcha, req.Captcha) {
			return response.FailWithMessage("Captcha verification code error", c)
		}
	}

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.accountApp.AuthenticationSvc.Login(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	// set interface-session-jwt on cacher
	if err := f._setSession(resp.Token, resp.ExpiredAt); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// LoginApp login client app
func (f *FAuths) LoginApp(c echo.Context) error {

	req := new(appDTOAuth.LoginAppReqDTO)
	if err := c.Bind(req); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.accountApp.AuthenticationSvc.LoginApp(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	// set interface-session-jwt on cacher
	if err := f._setSession(resp.Token, resp.ExpiredAt); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

func (f *FAuths) _setSession(token string, expiredAt int64) error {
	sessionValue := token
	expiredAtDT := time.Unix(expiredAt/1000, 0) // ExpiredAt = UnixTimeStamp
	expiration := expiredAtDT.Sub(time.Now()).Seconds()
	if err := f.SetSession(sessionValue, int64(expiration)); err != nil {
		return err
	}
	return nil
}
