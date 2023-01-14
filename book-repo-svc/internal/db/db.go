package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/heritechie/go-book-catalog/util"
)

func ConnectDB(config util.SharedConfig) *gorm.DB {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta", config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort, config.PGSSLMode)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{ // table name prefix, table for `User` would be `t_users`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	return conn
}
