package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafa-533/rest-api/model"
)

const categoriesFile = "categories.json"

func listCategories(c *gin.Context) {
	categories, err := readCategoriesFromFile(categoriesFile)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

func getCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	categories, err := readCategoriesFromFile(categoriesFile)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, category := range categories {
		if category.ID == id {
			c.JSON(http.StatusOK, category)
			return
		}
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func createCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	categories, err := readCategoriesFromFile(categoriesFile)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Find the max ID in the existing categories
	maxID := 0
	for _, category := range categories {
		if category.ID > maxID {
			maxID = category.ID
		}
	}

	// Create a new category with the next ID
	category.ID = maxID + 1
	categories = append(categories, category)

	if err := writeCategoriesToFile(categoriesFile, categories); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, category)
}

func updateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	category.ID = id

	categories, err := readCategoriesFromFile(categoriesFile)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for i, existingCategory := range categories {
		if existingCategory.ID == id {
			categories[i] = category
			if err := writeCategoriesToFile(categoriesFile, categories); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusOK, category)
			return
		}
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func deleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	categories, err := readCategoriesFromFile(categoriesFile)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for i, category := range categories {
		if category.ID == id {
			// Remove the category from the slice
			categories = append(categories[:i], categories[i+1:]...)

			if err := writeCategoriesToFile(categoriesFile, categories); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			c.Status(http.StatusNoContent)
			return
		}
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func readCategoriesFromFile(filename string) ([]model.Category, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var categories []model.Category
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func writeCategoriesToFile(filename string, data []model.Category) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
