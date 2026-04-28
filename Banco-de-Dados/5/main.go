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
	ID           int `gorm:"primaryKey`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(10.10.10.193:3307)/curso_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	//create product
	product := []Product{
		{Name: "Mouse", Price: 99.99, CategoryID: 1},
		{Name: "Teclado", Price: 199.99, CategoryID: 1},
		{Name: "Monitor", Price: 999.99, CategoryID: 1},
		{Name: "Notebook", Price: 1999.99, CategoryID: 1},
		{Name: "Smartphone", Price: 2999.99, CategoryID: 1},
		{Name: "Tablet", Price: 1499.99, CategoryID: 1},
	}
	db.Create(&product)

	// create serial number
	db.Create(&[]SerialNumber{
		{Number: "123456789", ProductID: 1},
		{Number: "987654321", ProductID: 2},
		{Number: "456789123", ProductID: 3},
		{Number: "789123456", ProductID: 4},
		{Number: "321654987", ProductID: 5},
		{Number: "654987321", ProductID: 6},
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	//db.Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
