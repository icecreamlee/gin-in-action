package forms

import "github.com/gin-gonic/gin"

// ProductForm 产品相关表单结构体
type ProductForm struct {
	Page     int    `json:"page" form:"page" binding:"gt=0"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"gt=0"`
	Name     string `json:"name" form:"name" binding:"max=100"`
}

// ValidateProductListForm 产品列表表单验证
func ValidateProductListForm(c *gin.Context) (ProductForm, error) {
	var productForm ProductForm
	err := c.ShouldBindQuery(&productForm)
	if err != nil {
		return productForm, err
	}

	if productForm.Page == 0 {
		productForm.Page = 1
	}
	if productForm.PageSize == 0 {
		productForm.PageSize = 10
	}
	return productForm, err
}
