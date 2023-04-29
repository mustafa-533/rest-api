package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafa-533/rest-api/db"
	"github.com/mustafa-533/rest-api/model"
	"go.uber.org/zap"
)

func (h *H) getBook(c *gin.Context) {
	var (
		idParam = c.Param("id")
		logger  = h.logger.With(zap.String("api", "book"), zap.String("method", "get_book"), zap.String("book.id", idParam))
	)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error("Error in converting string id to int", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	book, err := h.db.GetByID(id)
	if err != nil {
		if err == db.ErrNotFound {
			logger.Error("Error in get book by id", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *H) getBooks(c *gin.Context) {
	logger := h.logger.With(zap.String("api", "book"), zap.String("method", "get_books"))

	books, err := h.db.GetAll()
	if err != nil {
		logger.Error("Get all books from db error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *H) createBook(c *gin.Context) {
	var (
		book   model.Book
		logger = h.logger.With(zap.String("api", "book"), zap.String("method", "create_book"))
	)

	if err := c.BindJSON(&book); err != nil {
		logger.Error("Bind json data error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook, err := h.db.Create(&book)
	if err != nil {
		logger.Error("Create book error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBook)

}

func (h *H) updateBook(c *gin.Context) {

	var (
		idParam = c.Param("id")
		logger  = h.logger.With(zap.String("api", "book"), zap.String("method", "update_book"), zap.String("book.id", idParam))
	)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error("Error in converting string id to int", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	var book model.Book
	if err := c.BindJSON(&book); err != nil {
		logger.Error("Bind json error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Update(&book, id); err != nil {
		logger.Error("Update book error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

func (h *H) deleteBook(c *gin.Context) {
	var (
		idParam = c.Param("id")
		logger  = h.logger.With(zap.String("api", "book"), zap.String("method", "delete_book"), zap.String("book.id", idParam))
	)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error("Error in converting string id to int", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	if err := h.db.Delete(id); err != nil {
		logger.Error("Delete book error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
