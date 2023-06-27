package categorycontroller

import (
	"net/http"

	"github.com/Timotius-Nugroho/learn-go-rest-api.git/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func Create(c *gin.Context) {
	var category models.Category
	var product models.Product
	validate = validator.New()
	tx := models.DB.Begin()

	var categoryBody struct {
		models.Category
		Products []models.Product `json:"products"`
	}

	err := c.ShouldBindJSON(&categoryBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// fmt.Printf("%+v\n", categoryBody)

	for _, item := range categoryBody.Products {
		product = item
		err = validate.Struct(product)
		if err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err = tx.Create(&product).Error
		if err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	category = categoryBody.Category
	err = validate.Struct(category)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = tx.Create(&category).Error
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"category": categoryBody})
}
