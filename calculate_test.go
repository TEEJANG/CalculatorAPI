package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusAB(t *testing.T) {
	t.Run("it should return httpCode 200 and correct plus result", func(t *testing.T) {
		expectedResult := fmt.Sprintf("%f", -1.0)
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/plus?A=-5.0&B=4.0", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotNil(t, w.Body)
		assert.Equal(t, expectedResult, w.Body.String())
	})

	t.Run("it shloud return httpCode 400 when input is not digit", func(t *testing.T) {
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/plus?A=+iueor90&B=kaaassdf", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestMinusAB(t *testing.T) {
	t.Run("it should return httpCode 200 and correct minus result", func(t *testing.T) {
		expectedResult := fmt.Sprintf("%f", -9.0)
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/minus?A=-5.0&B=4.0", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotNil(t, w.Body)
		assert.Equal(t, expectedResult, w.Body.String())
	})

	t.Run("it shloud return httpCode 400 when input is not digit", func(t *testing.T) {
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/minus?A=+iueor90&B=kaaassdf", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestMultiplyAB(t *testing.T) {
	t.Run("it should return httpCode 200 and correct multiply result", func(t *testing.T) {
		expectedResult := fmt.Sprintf("%f", -20.0)
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/multiply?A=-5.0&B=4.0", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotNil(t, w.Body)
		assert.Equal(t, expectedResult, w.Body.String())
	})

	t.Run("it shloud return httpCode 400 when input is not digit", func(t *testing.T) {
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/multiply?A=+iueor90&B=kaaassdf", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDivineAB(t *testing.T) {
	t.Run("it should return httpCode 200 and correct divine result", func(t *testing.T) {
		expectedResult := fmt.Sprintf("%f", -5.0)
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/divine?A=-20.0&B=4.0", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotNil(t, w.Body)
		assert.Equal(t, expectedResult, w.Body.String())
	})

	t.Run("it shloud return httpCode 400 when input's A is not digit", func(t *testing.T) {
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/divine?A=+iueor90&B=kaaassdf", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("it shloud return httpCode 400 when input's B is not digit", func(t *testing.T) {
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/divine?A=90&B=kaaassdf", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("it shloud return httpCode 400 when B is equal to 0", func(t *testing.T) {
		r := setupRouter()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/divine?A=-20.0&B=0.0", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}
