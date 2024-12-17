package WorkingWithDatabaseGorm

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Customer struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"unique"`
	Order []Order
}
type Order struct {
	ID         uint `gorm:"primary_key"`
	CustomerID uint
	ItemID     string //Item olmalidi bosh
}

func AutoMigraterCustomer(db *gorm.DB) {
	err := db.AutoMigrate(&Customer{}, &Order{})
	if err != nil {
		fmt.Errorf("Automigrate Error ", err)
	}

}

func CreateCustomer(db *gorm.DB) {
	customer := Customer{
		Name: "John Doe",
		Order: []Order{
			{

				ItemID: "Laptop",
			},
			{
				ItemID: "Phone",
			},
		},
	}
	err := db.Create(&customer).Error
	if err != nil {
		log.Fatal("customer create err:", err)
	}

}
func GetAllCustomer(db *gorm.DB) {
	var customers []Customer

	err := db.Preload("Order").Find(&customers).Error
	if err != nil {
		log.Fatal("customer get err:", err)
	}
	for _, k := range customers {
		//Customer: Ali Order: Laptop Order: Phone
		fmt.Printf("Customer: %s,Order:%s Oreder:%s \n", k.Name, k.Order[0].ItemID, k.Order[1].ItemID)
	}

}
