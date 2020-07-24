package test

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/api"
	"todo/app"
	"todo/db"
)

func NewTestRouter(db *db.TestDatabase) (*chi.Mux, error) {
	tApi, err := api.New(&app.App{Database: db})
	if err != nil {
		return nil, err
	}
	router := chi.NewRouter()
	tApi.Init(router)
	return router, nil
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, []byte, error) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		return nil, []byte(""), errors.Wrap(err, "err in creating request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, []byte(""), errors.Wrap(err, "err in request")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, []byte(""), errors.Wrap(err, "err read from response")
	}
	defer resp.Body.Close()

	return resp, respBody, nil
}
