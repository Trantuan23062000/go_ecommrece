package main

import (
	"database/sql"
	"log"

	"github.com/Trantuan23062000/go_ecommrece/cmd/api"
	"github.com/Trantuan23062000/go_ecommrece/config"
	"github.com/Trantuan23062000/go_ecommrece/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	// Khởi tạo cấu hình từ package config
	cfg := config.Envs

	// Kết nối đến cơ sở dữ liệu MySQL
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Addr:                 cfg.DBAddress,
		DBName:               cfg.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	initStorage(db)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Khởi chạy API server với cơ sở dữ liệu đã kết nối
	server := api.NewAPIServer(":"+cfg.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Data connect")
}
