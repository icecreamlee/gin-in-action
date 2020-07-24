package logics

import (
	"gin_in_action/core"
	"gin_in_action/forms"
	"gin_in_action/models"
	"github.com/gin-gonic/gin"
)

// AddProduct 添加产品
func AddProduct(form forms.ProductAddForm) core.ApiError {
	product := models.GetProductByName(form.Name)
	if product.ID > 0 {
		return core.ApiFailure("product name exists")
	}

	tx := models.DB.Begin()
	product = models.Product{
		Name:  form.Name,
		Price: form.Price,
		Pic:   form.Pic,
	}
	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return core.ApiFailure(core.CodeFailure, "failed to create product", err.Error())
	}

	for _, image := range form.Images {
		productImage := models.ProductImage{
			ProductId: product.ID,
			Image:     image,
		}
		if err := tx.Create(&productImage).Error; err != nil {
			tx.Rollback()
			return core.ApiFailure(core.CodeFailure, "failed to create product image", err.Error())
		}
	}
	tx.Commit()
	return core.ApiSuccess(core.CodeSuccess, "", gin.H{"id": product.ID})
}

// 更新产品
func UpdateProduct(form forms.ProductUpdateForm) core.ApiError {
	product := models.GetProduct(form.ID)
	if product.ID <= 0 {
		return core.ApiFailure("product not exists")
	}

	productTmp := models.GetProductByName(form.Name)
	if productTmp.ID != 0 && product.ID != productTmp.ID {
		return core.ApiFailure("product name exists")
	}

	tx := models.DB.Begin()
	product.Name = form.Name
	product.Price = form.Price
	product.Pic = form.Pic
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return core.ApiFailure(core.CodeFailure, "failed to update product", err.Error())
	}

	tx.Where("product_id = ?", product.ID).Delete(&models.ProductImage{})
	for _, image := range form.Images {
		productImage := models.ProductImage{
			ProductId: product.ID,
			Image:     image,
		}
		if err := tx.Create(&productImage).Error; err != nil {
			tx.Rollback()
			return core.ApiFailure(core.CodeFailure, "failed to create product image", err.Error())
		}
	}
	tx.Commit()
	return core.ApiSuccess(core.CodeSuccess, "", gin.H{"id": product.ID})
}
