package util

import (
	"io"
	"net/http"
	"net/http/httptest"

	config "expense-logger/configs"
	db "expense-logger/web/app/models"
)

func PerformRequest(r http.Handler, method string, path string, payload io.Reader) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, payload)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w, err
}

func NewDBConnection() {
	config.Init()
	db.Init()
}

func EndDBConnection() {
	defer db.Close()
}
