package main

import (
	"Session-20/WorkingWithDatabaseGorm"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB1 *gorm.DB
var DB2 *gorm.DB
var DB3 *gorm.DB
var DB4 *gorm.DB

func main() {
	db := WorkingWithDatabaseGorm.ConDb()
	DB = db
	DB1 = db
	DB2 = db
	DB3 = db
	DB4 = db
	DB1.Debug()
	DB2.Debug()
	DB3.Debug()
	//WorkingWithDatabaseGorm.AutoMigrater(DB)

	//WorkingWithDatabaseGorm.CrudNewBook(DB)
	//WorkingWithDatabaseGorm.CrudGetBook(DB)

	//task 4
	//WorkingWithDatabaseGorm.AutoMigraterProAndUser(DB1)
	//WorkingWithDatabaseGorm.CeateUser(DB1)
	//WorkingWithDatabaseGorm.GetUserInfo(DB1)

	//Task5
	//WorkingWithDatabaseGorm.AutoMigraterCustomer(DB2)
	//WorkingWithDatabaseGorm.CreateCustomer(DB2)
	//WorkingWithDatabaseGorm.GetAllCustomer(DB2)

	//Task7
	//WorkingWithDatabaseGorm.AutoMigratorGrop(DB3)
	//WorkingWithDatabaseGorm.CreateUserAndGroup(DB3)
	//WorkingWithDatabaseGorm.GetUserAndGrops(DB3)

	//Task8
	WorkingWithDatabaseGorm.Tr(DB4)

}

//Host: localhost
//Port: 5432
//User: postgres
//Password: postgres
//Database Name: testdb
