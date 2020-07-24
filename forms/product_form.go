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

// ProductAddForm 添加产品表单结构体
type ProductAddForm struct {
	Name   string   `form:"name" binding:"required,max=100"`
	Price  float32  `form:"price" binding:"required,numeric,gt=0,lt=100000000"`
	Pic    string   `form:"pic" binding:"required,url,max=255"`
	Images []string `form:"images" binding:"dive,url,max=255"`
}

// ValidateProductAddForm 产品列表表单验证
func ValidateProductAddForm(c *gin.Context) (form ProductAddForm, err error) {
	err = c.ShouldBindJSON(&form)
	return
}

// ProductUpdateForm 更新产品表单结构体
type ProductUpdateForm struct {
	ID int `form:"id" binding:"gt=0"`
	ProductAddForm
}

// ValidateProductAddForm 产品列表表单验证
func ValidateProductUpdateForm(c *gin.Context) (form ProductUpdateForm, err error) {
	err = c.ShouldBindJSON(&form)
	return
}
