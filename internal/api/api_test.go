package api

import (
	"L0/internal/caching"
	"L0/internal/model"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var api Api

func TestMain(m *testing.M) {
	cache := caching.NewCache()
	api.InitRouter(cache)
	code := m.Run()
	os.Exit(code)
}

func TestApi_GetOrder(t *testing.T) {
	api.cache.Put(model.Order{OrderUID: "1"})

	url := fmt.Sprintf("/order/?id=1")
	req, _ := http.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()
	api.Router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	var order model.Order
	err := json.Unmarshal(response.Body.Bytes(), &order)
	assert.Equal(t, err, nil)
	assert.Equal(t, model.Order{OrderUID: "1"}, order)

	url = fmt.Sprintf("/order/?id=2")
	req, _ = http.NewRequest("GET", url, nil)
	badResponse := httptest.NewRecorder()
	api.Router.ServeHTTP(badResponse, req)
	assert.Equal(t, http.StatusNotFound, badResponse.Code)
	assert.Equal(t, "the order isn't in cache", badResponse.Body.String())

}
