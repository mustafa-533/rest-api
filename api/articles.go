package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mustafa-533/rest-api/model"
)

var articles = make(map[string]model.Article)

func (h *H) getArticle(c *gin.Context) {
	id := c.Param("id")

	article, exists := articles[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *H) getArticles(c *gin.Context) {
	c.JSON(http.StatusOK, articles)
}

func (h *H) createArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate ID
	article.ID = uuid.New().String()

	articles[article.ID] = article

	c.JSON(http.StatusCreated, article)
}

func (h *H) updateArticle(c *gin.Context) {
	id := c.Param("id")

	article, exists := articles[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	var updatedArticle model.Article
	if err := c.ShouldBindJSON(&updatedArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article.Title = updatedArticle.Title
	article.Content = updatedArticle.Content
	articles[id] = article

	c.JSON(http.StatusOK, article)
}

func (h *H) deleteArticle(c *gin.Context) {
	id := c.Param("id")

	_, exists := articles[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	delete(articles, id)

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
