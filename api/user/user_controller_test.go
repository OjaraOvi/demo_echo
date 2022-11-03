package user

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	mockDB   = map[string]*User{}
	userJSON = `{"name":"Jon Snow","location":"chiguayork", "title": "GG", "password": "asddd"}`
)

func TestCreateUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "api/user/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{}

	// Assertions
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}

}
