package handler

import (
	"webapp/pkg/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() chi.Router {
	s := chi.NewRouter()
	gr := s.Group(func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-up", h.signUp)
			r.Post("/sign-in", h.signIn)
		})

		r.Route("/api", func(r chi.Router) {
			r.Use(h.userIdentify)
			r.Route("/lists", func(r chi.Router) {
				r.Post("/", h.createList)
				r.Get("/", h.getAllLists)
				r.Get("/{id}", h.getListById)
				r.Delete("/{id}", h.deleteList)
				r.Route("/{id}/items", func(r chi.Router) {
					r.Post("/", h.createItem)
					r.Get("/", h.getAllItems)
				})
			})
			r.Route("/items", func(r chi.Router) {
				r.Get("/{id}", h.getItemById)
				r.Delete("/{id}", h.deleteItem)
			})
		})
	})
	return gr
}
