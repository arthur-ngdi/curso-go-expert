package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(10.10.10.193:3307)/curso_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create category
	/* category := Category{Name: "Eletronicos"}
	db.Create(&category) */

	//create product
	product := Product{
		Name: "Mouse",
		Price: 99.99,
		CategoryID: 1,
	}
	db.Create(&product)

	var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
