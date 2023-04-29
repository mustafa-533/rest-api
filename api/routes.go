package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	articleApiV1Path  = "/api/v1/articles"
	categoryApiV1Path = "/api/v1/categories"
)

func (h *H) LoadRoutes() (http.Handler, error) {
	// Init Router
	g := gin.New()

	articleV1 := g.Group(articleApiV1Path)

	articleV1.GET("", h.getArticles)
	articleV1.GET("/:id", h.getArticle)
	articleV1.POST("", h.createArticle)
	articleV1.PUT("/:id", h.updateArticle)
	articleV1.DELETE("/:id", h.deleteArticle)

	categoryV1 := g.Group(categoryApiV1Path)

	categoryV1.GET("", listCategories)
	categoryV1.GET("/:id", getCategory)
	categoryV1.POST("", createCategory)
	categoryV1.PUT("/:id", updateCategory)
	categoryV1.DELETE("/:id", deleteCategory)

	return g, nil
}
