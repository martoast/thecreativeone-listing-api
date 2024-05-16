package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dbusername, dbusername_ok := os.LookupEnv("MYSQL_USERNAME")
	dbname, dbname_ok := os.LookupEnv("MYSQL_DATABASE")
	rootpass, rootpass_ok := os.LookupEnv("MYSQL_ROOT_PASSWORD")
	dbhost, dbhost_ok := os.LookupEnv("MYSQL_HOST")
	dbport, dbport_ok := os.LookupEnv("MYSQL_PORT")
	if dbusername_ok && dbname_ok && rootpass_ok && dbhost_ok && dbport_ok {
		fmt.Println(rootpass, dbname, dbhost)
	} else {
		fmt.Println("Environment variables are not set. Please set environment variables.")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbusername, rootpass, dbhost, dbport, dbname)

	fmt.Println(dsn)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("successfully connected to the database!!")

	db = d
}

func GetDB() *gorm.DB {
	return db
}
