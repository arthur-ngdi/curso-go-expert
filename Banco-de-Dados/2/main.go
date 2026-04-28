package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID    int `gorm:"primaryKey`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(10.10.10.193:3307)/curso_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//create
	/* product := Product{
		Name: "TV",
		Price: 4000.00,
	}
	db.Create(&product) */

	// create batch
	/* products :=  []Product{
		{Name: "TV", Price: 4000.00,},
		{Name:	"Notebook", Price: 7000.00},
		{Name:	"Liquidificador", Price: 500.00},
		{Name:	"Filtro de água", Price: 600.00},
	}
	db.Create(&products) */

	// Busca
	/* var product Product
	db.First(&product, 7)
	fmt.Println(product
	db.First(&product, "name = ?", "Notebook")
	fmt.Println(product) */

	// select All
	/* var product []Product
	db.Limit(2).Offset(2).Find(&products)
	db.Find(&product)
	for _, product := range products {
		fmt.Println(product)
	} */

	// where
	/* var products []Product
	db.Where("price > ?", 4000).Find(&products)
	fmt.Println(products)

	db.Where("name LIKE ?", "%c%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	} */

	// delete
	/* 	var product Product
	   	db.First(&product, 1)
	   	product.Name = "Television Wide Screen"
	   	db.Save(&product) */

	/* var product2 Product
	db.First(&product2, 2)
	fmt.Println(product2.Name)

	db.Delete(&product2)
	fmt.Printf("Produto %s foi removido\n", product2.Name) */

	// soft delete
}
