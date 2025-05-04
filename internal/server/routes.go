package server

import (
	"net/http"

	"github.com/Cozzytree/comtroller/internal/templete"
	"github.com/go-chi/chi/v5"
)

func (s *ServerStruct) RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		layout := templete.Layout("Hello world")
		layout.Render(r.Context(), w)
	})

	return r
}
