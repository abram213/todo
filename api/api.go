package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"

	"todo/app"
)

type API struct {
	App    *app.App
	Config *Config
}

func New(a *app.App) (api *API, err error) {
	api = &API{App: a}
	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (a *API) Init(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Method("GET", "/", a.handler(func(w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(w, "Hello! This is TODO app!")
		return nil
	}))
}

func (a *API) handler(f func(http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 100*1024*1024)

		w.Header().Set("Content-Type", "application/json")

		if err := f(w, r); err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})
}
