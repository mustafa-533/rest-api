package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mustafa-533/rest-api/model"
	"go.uber.org/zap"
)

var articles = make(map[string]model.Article)

func (h *H) getArticle(c *gin.Context) {
	var (
		id     = c.Param("id")
		logger = h.logger.With(zap.String("api", "article"), zap.String("method", "get_article"), zap.String("article.id", id))
	)

	article, exists := articles[id]
	if !exists {
		logger.Info("Article not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *H) getArticles(c *gin.Context) {
	c.JSON(http.StatusOK, articles)
}

func (h *H) createArticle(c *gin.Context) {
	var (
		logger  = h.logger.With(zap.String("api", "article"), zap.String("method", "create_article"))
		article model.Article
	)

	if err := c.ShouldBindJSON(&article); err != nil {
		logger.Error("Error in binding json data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate ID
	article.ID = uuid.New().String()

	articles[article.ID] = article

	c.JSON(http.StatusCreated, article)
}

func (h *H) updateArticle(c *gin.Context) {
	var (
		id     = c.Param("id")
		logger = h.logger.With(zap.String("api", "article"), zap.String("method", "update_article"), zap.String("article.id", id))
	)

	article, exists := articles[id]
	if !exists {
		logger.Debug("Article not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	var updatedArticle model.Article
	if err := c.ShouldBindJSON(&updatedArticle); err != nil {
		logger.Debug("Bind requeest error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article.Title = updatedArticle.Title
	article.Content = updatedArticle.Content
	articles[id] = article

	c.JSON(http.StatusOK, article)
}

func (h *H) deleteArticle(c *gin.Context) {
	var (
		id     = c.Param("id")
		logger = h.logger.With(zap.String("api", "article"), zap.String("method", "delete_article"), zap.String("article.id", id))
	)

	_, exists := articles[id]
	if !exists {
		logger.Debug("Article not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	delete(articles, id)

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
