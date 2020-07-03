package main

import (
	"github.com/go-chi/chi"
)

//function with the routers, returning the router with new available functionalities
func routers() *chi.Mux {
    router.Get("/posts", AllPosts)
    router.Get("/posts/{id}", DetailPost)
    router.Post("/posts", CreatePost)
    router.Put("/posts/{id}", UpdatePost)
    router.Delete("/posts/{id}", DeletePost)
    
    return router
}