package router

import (
	"github.com/go-chi/chi"
	"go_examples/api_chi_sql/handlers"
)

//function with the routers, returning the router with new available functionalities
func SetEndpoints(router *chi.Mux) *chi.Mux {
	router.Get("/posts", handlers.AllPosts)
	router.Get("/posts/{id}", handlers.DetailPost)
	router.Post("/posts", handlers.CreatePost)
	router.Put("/posts/{id}", handlers.UpdatePost)
	router.Delete("/posts/{id}", handlers.DeletePost)

	return router
}
