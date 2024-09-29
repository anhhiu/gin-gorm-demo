package main

import (
	"log"
	"shoe_inventory/config"
	"shoe_inventory/models"
	"shoe_inventory/routes"
)

func main() {
	// Kết nối cơ sở dữ liệu
	config.ConnectDatabase()
	//log.Println("connec thanh cong\n")
	// Tự động tạo bảng Shoe
	config.DB.AutoMigrate(&models.Shoe{})
	//log.Println("tao bang thanh cong\n")
	// Khởi tạo router
	router := routes.SetupRouter()

	log.Println("Server is running at http://localhost:8080")
	router.Run(":8080")
}
