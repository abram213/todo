package cmd

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"todo/api"
	"todo/app"
)

func serveAPI(ctx context.Context, api *api.API) {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	router := chi.NewRouter()
	api.Init(router)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", api.Config.Port),
		Handler:     cors.Handler(router),
		ReadTimeout: 2 * time.Minute,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			fmt.Println(err)
		}
		close(done)
	}()

	fmt.Printf("serving at http://127.0.0.1:%d\n", api.Config.Port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println(err)
	}
	<-done
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the app",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := app.New("")
		if err != nil {
			return err
		}
		defer app.Database.CloseDB()

		api, err := api.New(app)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			<-ch
			fmt.Println("signal caught. shutting down...")
			cancel()
		}()

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer cancel()
			serveAPI(ctx, api)
		}()

		wg.Wait()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
