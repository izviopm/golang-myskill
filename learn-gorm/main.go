package main

import (
	"fmt"
	"lear-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=user001 dbname=iprijaya port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal konek ke db")
	}
	
	fmt.Println("Sukses konek ke db")

	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&models.Student{})

	data := models.Student{}
	if DB.Find(&data).RowsAffected == 0 {
		fmt.Println("-------Run Seeder User-------")
		seederUser()
	}
}

func seederUser() {
	data := models.Student{
		Student_id: 1,
		Student_name: "Zexceed",
		Student_age: 24,
		Student_address: "Bandung",
		Student_phone_no: "0821231234",
	}

	DB.Create(&data)
}