package controller

import (
	"product-services/models"

	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	products, err := models.GetProductList()
	if err != nil {
		c.JSON(500, map[string]interface{}{"msg": err.Error()})
		return
	}
	c.JSON(200, products)
}
