package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"reflect"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
	UUID  string `gorm:"unique;not null;type:varchar(100);default:null"`
	Count uint
}

func NewProduct() Product {
	product := Product{}
	product.Count = 0
	return product
}

func (p *Product) load(uuid string, db *gorm.DB) error {
	log.Printf("In load, trying to load product UUID %s", uuid)
	result := db.First(&p, "uuid = ?", uuid)
	log.Printf("Rows %d", result.RowsAffected)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("UUID Not Found")
		} else {
			fmt.Printf("Other error: %s\n", result.Error)
		}
		return result.Error
	} else {
		fmt.Println("Record Found")
	}
	log.Printf("In load, product UUID is %s", p.UUID)
	return nil
}

var db *gorm.DB

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()

	log.Println("The type of db is")
	log.Println(reflect.TypeOf(db))

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	uuid := uuid.NewString()
	prod := NewProduct()
	prod.UUID = uuid
	prod.Code = "a1"
	prod.Price = 122
	result := db.Create(&prod)
	if result.Error != nil {
		fmt.Println("error doing insert")
		return
	}
	fmt.Println("Product inserted")

	//db.Create(&Product{Code: "D45", Price: 300})

	// Read
	var product Product
	var products []Product

	err1 := product.load(uuid, db)
	if err1 != nil {
		fmt.Println("There was a problem loading the product")
		return
	}
	log.Printf("Product reloaded with UUID %s", product.UUID)

	for count := 0; count < 5; count++ {
		db.Model(&product).Update("Count", product.Count+1)
		log.Printf("The count for %s is %d", product.UUID, product.Count)
	}

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
