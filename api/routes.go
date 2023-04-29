package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	articleApiV1Path  = "/api/v1/articles"
	categoryApiV1Path = "/api/v1/categories"
	bookApiV1Path     = "/api/v1/books"
)

func (h *H) LoadRoutes() (http.Handler, error) {
	// Init Router
	g := gin.New()

	// Articles routes
	articleV1 := g.Group(articleApiV1Path)
	articleV1.GET("", h.getArticles)
	articleV1.GET("/:id", h.getArticle)
	articleV1.POST("", h.createArticle)
	articleV1.PUT("/:id", h.updateArticle)
	articleV1.DELETE("/:id", h.deleteArticle)

	// Categories routes
	categoryV1 := g.Group(categoryApiV1Path)
	categoryV1.GET("", h.listCategories)
	categoryV1.GET("/:id", h.getCategory)
	categoryV1.POST("", h.createCategory)
	categoryV1.PUT("/:id", h.updateCategory)
	categoryV1.DELETE("/:id", h.deleteCategory)

	// Books routes
	bookV1 := g.Group(bookApiV1Path)
	bookV1.GET("", h.getBooks)
	bookV1.GET("/:id", h.getBook)
	bookV1.POST("", h.createBook)
	bookV1.PUT("/:id", h.updateBook)
	bookV1.DELETE("/:id", h.deleteBook)

	return g, nil
}
