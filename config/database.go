package config

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB is the MySQL database connection
	DB *gorm.DB
)