package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//db.Create(&Product{Code: "D45", Price: 300})

	// Read
	var product Product
	var products []Product

	//db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "F19") // find product with code D42
	db.Find(&products, "code = ?", "D42") // find product with code D42

	for _, p := range products {
		fmt.Printf("ID: %d\n", p.ID)
		fmt.Printf("Price: %d\n", p.Price)
		fmt.Printf("Code: %s\n", p.Code)
		db.Model(&p).Update("Price", p.Price+1)
	}

	/*
		// Update - update product's price to 200
		db.Model(&product).Update("Price", 200)
		// Update - update multiple fields
		db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
		db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	*/
	// Delete - delete product
	db.Delete(&product, 1)
}
