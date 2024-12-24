package WorkingWithDatabaseGorm

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

//many to many

type Userr struct {
	ID    uint     `gorm:"primary_key"`
	Name  string   `gorm:"size:255;not null"`
	Group []*Group `gorm:"many2many:userr_groups;"`
}
type Group struct {
	ID    uint     `gorm:"primary_key"`
	Name  string   `gorm:"size:255;not null"`
	Userr []*Userr `gorm:"many2many:userr_groups;"`
}

func AutoMigratorGrop(db *gorm.DB) {
	err := db.AutoMigrate(&Userr{}, &Group{})
	if err != nil {
		log.Fatal("AutoMigratorGrop err:", err)
	}

}
func CreateUserAndGroup(db *gorm.DB) {
	user := Userr{

		Name: "Ali",
		Group: []*Group{
			{

				Name: "Admin",
			},
			{
				Name: "Developer",
			},
		},
	}
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal("createUserAndGroup err:", err)
	}

}
func GetUserAndGrops(db *gorm.DB) {
	var users []Userr
	db.Preload("Group").Find(&users)
	for _, user := range users {
		//User: Ali Group: Admin Group: Developer
		fmt.Printf("user:%s,Group:%s,Group:%s", user.Name, user.Group[0].Name, user.Group[1].Name)
	}

}
