package account

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ht "github.com/d3ta-go/ms-account-restapi/interface/http-apps/restapi/echo/features/helper_test"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/d3ta-go/system/system/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuths_RegisterUser(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.interface-layer.features.register-user.request")

	unique := utils.GenerateUUID()
	// variables
	reqDTO := `{
		"username" : "` + fmt.Sprintf(testData["username"], unique) + `", 
		"password" : "` + testData["password"] + `",
		"email" : "` + fmt.Sprintf(testData["email"], unique) + `",
		"nickName" : "` + testData["nick-name"] + `",
		"captcha": "` + testData["captcha-value"] + `",
		"captchaID": "` + testData["captcha-id"] + `"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auths/register", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		panic(err)
	}

	auths, err := NewFAuths(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, auths.RegisterUser(c)) {
		// assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.ReadInConfig()
		viper.Set("test-data.account.auth.interface-layer.features.login.request.username", fmt.Sprintf(testData["username"], unique))
		viper.Set("test-data.account.auth.interface-layer.features.login.request.password", testData["password"])
		viper.Set("test-data.account.auth.interface-layer.features.login.request.captcha-value", testData["captcha-value"])
		viper.Set("test-data.account.auth.interface-layer.features.login.request.captcha-id", testData["captcha-id"])

		viper.Set("test-data.account.auth.interface-layer.features.register-user.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.auths.RegisterUser: %s", res.Body.String())
	}
}

func TestAuths_ActivateRegistration(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.interface-layer.features.activate-registration.request")

	// variables
	// via url path

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/auths/registration/activate/:activationCode", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("activationCode")
	c.SetParamValues(testData["activation-code"])

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		panic(err)
	}

	auths, err := NewFAuths(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, auths.ActivateRegistration(c)) {
		// assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.account.auth.interface-layer.features.activate-registration.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.auths.ActivateRegistration: %s", res.Body.String())
	}
}

func TestAuths_Login(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.interface-layer.features.login.request")

	// variables

	reqDTO := `{
		"username" : "` + testData["username"] + `", 
		"password" : "` + testData["password"] + `",
		"captcha": "` + testData["captcha-value"] + `",
		"captchaID": "` + testData["captcha-id"] + `"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auths/login", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		panic(err)
	}
	if err := initialize.OpenAllCacheConnection(h); err != nil {
		panic(err)
	}

	auths, err := NewFAuths(h)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, auths.Login(c)) {
		// assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.account.auth.interface-layer.features.login.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.auths.Login: %s", res.Body.String())
	}
}

func TestAuths_LoginApp(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.account.auth.interface-layer.features.login-app.request")

	// variables

	reqDTO := `{
		"clientKey" : "` + testData["client-key"] + `", 
		"secretKey" : "` + testData["secret-key"] + `"
	}`

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auths/login-app", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		panic(err)
	}
	if err := initialize.OpenAllCacheConnection(handler); err != nil {
		panic(err)
	}

	auths, err := NewFAuths(handler)
	if err != nil {
		panic(err)
	}

	// Assertions
	if assert.NoError(t, auths.LoginApp(c)) {
		// assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, resDTO, res.Body.String())
		// save to test-data
		// save result for next test
		viper.Set("test-data.account.auth.interface-layer.features.login-app.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.auths.LoginApp: %s", res.Body.String())
	}
}
