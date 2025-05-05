package server

import (
	"net/http"

	"github.com/Cozzytree/comtroller/internal/server/ws"
	"github.com/Cozzytree/comtroller/internal/template"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (s *ServerStruct) RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		template.Home().Render(r.Context(), w)
	})

	r.Get("/app", func(w http.ResponseWriter, r *http.Request) {
		template.Socket().Render(r.Context(), w)
	})

	r.Get("/ws/{client}", func(w http.ResponseWriter, r *http.Request) {
		ws.InitNewClient(w, r, s.Hub)
	})

	return r
}
