package WorkingWithDatabaseGorm

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint    `gorm:"primary_key"`
	FullName string  `gorm:"size:255"`
	Profile  Profile `gorm:"foreignkey:UserID"`
}
type Profile struct {
	ID     uint   `gorm:"primary_key"`
	UserID uint   `gorm:"unique"`
	Bio    string `gorm:"size:255"`
}

func AutoMigraterProAndUser(db *gorm.DB) {

	err := db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		fmt.Errorf("AutoMigrateProAndUser Error ", err)
	}
}
func CeateUser(db *gorm.DB) {
	db.Create(&User{
		FullName: "John Doe",
		Profile: Profile{
			Bio: "Software Developer",
		},
	})
}
func GetUserInfo(db *gorm.DB) {
	var users []User
	db.Preload("Profile").Find(&users)
	for _, user := range users {
		fmt.Printf("User:%s,Bio:%s", user.FullName, user.Profile.Bio)
	}

}
