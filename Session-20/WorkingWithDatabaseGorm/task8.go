package WorkingWithDatabaseGorm

import (
	"fmt"
	"gorm.io/gorm"
)

type user struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Email string `gorm:"size:255"`
}

func Tr(db *gorm.DB) {
	err := db.AutoMigrate(&user{})
	if err != nil {
		panic(err)
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user{
			ID:    1,
			Name:  "ali",
			Email: "Alihuseyn.999@gmail.com",
		}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Printf("Transaction succeeded")
	}
}
