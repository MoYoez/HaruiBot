// Package database, store user handler information here, also user can register for storing data here.
package database

import (
	"github.com/MoYoez/HaruiBot"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB gorm.DB

func init() {
	// open database here.
	DB, err := gorm.Open(sqlite.Open(HaruiBot.ApplicationSets.DataBaseLocation), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(UserModel{})
	if err != nil {
		panic(err)
	}

}
