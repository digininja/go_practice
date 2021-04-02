package main

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database")
		return nil, err
	}
	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})
	return db, nil
}

type Product struct {
	gorm.Model
	Code   string
	Price  uint
	UUID   string `gorm:"unique;not null;type:varchar(100);default:null"`
	Count  uint
	Status string
	// sqlite3 doesn't have enums, but for when I need them, this is how to do it
	// Status string `gorm:"type:enum('unknown', 'processing', 'complete', 'error')"`

	// Database connection string, maybe shouldn't be in here, not sure
	db *gorm.DB `gorm:"-"`
}

func NewProduct(db *gorm.DB) Product {
	product := Product{}
	product.db = db
	product.Count = 0
	product.Status = "processing"
	return product
}

func (p *Product) SetDatabase(db *gorm.DB) {
	p.db = db
}

func (p *Product) load(uuid string) error {
	log.Printf("In load, trying to load product UUID %s", uuid)
	result := p.db.First(&p, "uuid = ?", uuid)
	log.Printf("Rows %d", result.RowsAffected)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("UUID Not Found")
		} else {
			log.Printf("Other error: %s\n", result.Error)
		}
		return result.Error
	} else {
		log.Println("Record Found")
	}
	log.Printf("In load, product UUID is %s", p.UUID)
	return nil
}

func main() {
	var db *gorm.DB
	var err error

	/*
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		// defer db.Close()

		// Migrate the schema
		db.AutoMigrate(&Product{})
	*/

	db, err = ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Create
	uuid := uuid.NewString()
	prod := NewProduct(db)
	prod.UUID = uuid
	prod.Code = "a1"
	prod.Price = 122
	prod.Status = "processing"
	result := db.Create(&prod)
	if result.Error != nil {
		log.Println("error doing insert")
		return
	}
	log.Println("Product inserted")

	//db.Create(&Product{Code: "D45", Price: 300})

	// Read
	var product Product
	product.SetDatabase(db)
	var products []Product

	err = product.load(uuid)
	if err != nil {
		log.Println("There was a problem loading the product")
	}
	log.Printf("Product reloaded with UUID %s", product.UUID)

	err = product.load(uuid + "x")
	if err != nil {
		log.Println("There was a problem loading the product with invalid UUID")
	}
	log.Printf("Product reloaded with UUID %s", product.UUID)

	for count := 0; count < 5; count++ {
		db.Model(&product).Update("Count", product.Count+1)
		if count == 4 {
			db.Model(&product).Update("Status", "complete")
		}
		log.Printf("The count for %s is %d", product.UUID, product.Count)
	}

	//db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "F19") // find product with code D42
	db.Find(&products, "code = ?", "D42") // find product with code D42

	for _, p := range products {
		log.Printf("ID: %d\n", p.ID)
		log.Printf("Price: %d\n", p.Price)
		log.Printf("Code: %s\n", p.Code)
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
