package models

import (
	"gin_in_action/core"
	"gin_in_action/forms"
	"github.com/gin-gonic/gin"
)

// Product 产品表结构体
type Product struct {
	ID        uint        `gorm:"primary_key" json:"id"`
	Name      string      `json:"name"`
	Price     float32     `json:"price"`
	Pic       string      `json:"pic"`
	Detail    string      `json:"detail"`
	IsDelete  int         `json:"is_delete"`
	CreatedAt core.MTime  `json:"created_at"`
	UpdatedAt core.MTime  `json:"updated_at"`
	DeletedAt *core.MTime `json:"-"`
}

func (p Product) TableName() string {
	return "products"
}

// ProductImage 产品图片表结构体
type ProductImage struct {
	ID        uint        `gorm:"primary_key" json:"id"`
	ProductId uint        `json:"product_id"`
	Image     string      `json:"image"`
	IsDelete  int         `json:"is_delete"`
	CreatedAt core.MTime  `json:"created_at"`
	UpdatedAt core.MTime  `json:"updated_at"`
	DeletedAt *core.MTime `json:"-"`
}

func (p ProductImage) TableName() string {
	return "product_images"
}

// GetRecentProducts 首页推荐产品列表
func GetRecentProducts() []Product {
	var products []Product
	DB.Limit(10).Find(&products)
	return products
}

// GetProduct 获取单个产品信息
func GetProduct(productId int) Product {
	var product Product
	DB.Where("id = ?", productId).First(&product)
	return product
}

// GetProductImages 获取指定产品的图片列表
func GetProductImages(productId int) []string {
	var productImages []ProductImage
	DB.Where("product_id = ?", productId).Find(&productImages)

	images := make([]string, 0)
	for _, image := range productImages {
		images = append(images, image.Image)
	}
	return images
}

// GetProductList 产品分页列表
func GetProductList(form forms.ProductForm) gin.H {
	var products []Product
	var count int
	var db = DB.Model(&Product{})
	if form.Name != "" {
		db = db.Where("name like ?", form.Name)
	}

	db.Offset((form.Page - 1) * form.PageSize).Limit(form.PageSize).Order("id").Find(&products)
	db.Count(&count)
	return gin.H{"count": count, "products": products}
}

// GetProductByName 根据名称获取产品信息
func GetProductByName(name string) (product Product) {
	DB.Where("name = ?", name).First(&product)
	return
}

// DeleteProductImages 删除指定产品的图片
func DeleteProductImages(productId uint) {
	DB.Where("id = ?", productId).Delete(&Product{})
}
