package apis

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	respRec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, respRec)

	rerror := NewHealthCheckAPI().HealthAPI(ctx)

	assert.Nil(t, rerror)
	assert.Equal(t, http.StatusOK, respRec.Code)
}
