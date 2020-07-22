package controllers

import (
	"gin_in_action/core"
	"gin_in_action/forms"
	"gin_in_action/models"
	"github.com/IcecreamLee/goutils"
	"github.com/gin-gonic/gin"
)

// Product 产品详情接口
func Product(c *gin.Context) {
	id := goutils.ToInt(c.DefaultQuery("id", "0"))
	product := models.GetProduct(id)
	images := models.GetProductImages(id)

	c.JSON(200, core.ApiSuccess(0, "", gin.H{
		"id":     product.ID,
		"name":   product.Name,
		"price":  goutils.Float2money(product.Price),
		"pic":    product.Pic,
		"detail": product.Detail,
		"images": images,
	}))
}

// ProductList 产品列表接口
func ProductList(c *gin.Context) {
	productForm, err := forms.ValidateProductListForm(c)
	if err != nil {
		c.JSON(200, core.ApiFailure(err.Error()))
		return
	}
	products := models.GetProductList(productForm)
	c.JSON(200, core.ApiSuccess(0, "", products))
}
