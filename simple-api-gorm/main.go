package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type newStudent struct {
	Student_id			uint64 `json:"student_id" binding:"required"`
	Student_name 		string `json:"student_name" binding:"required"`
	Student_age			uint64 `json:"student_age" binding:"required"`
	Student_address		string `json:"student_address" binding:"required"`
	Student_phone_no	string `json:"student_phone_no" binding:"required"`
}

func postHandler(c *gin.Context, db *gorm.DB) {
	var newStudent newStudent

	c.Bind(&newStudent)
	db.Create(&newStudent)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Insert Data",
		"data": newStudent,
	})
}

func getAllHandler(c *gin.Context, db *gorm.DB) {
	var newStudent []newStudent

	db.Find(&newStudent)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Get All Data",
		"data": newStudent,
	})
}

func getHandler(c *gin.Context, db *gorm.DB) {
	var newStudent newStudent
	studentId := c.Param("student_id")
	
	if db.Find(&newStudent, "student_id=?", studentId).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Get Specified Data using student_id",
		"data": newStudent,
	})
}

func putHandler(c *gin.Context, db *gorm.DB) {
	var newStudent newStudent
	studentId := c.Param("student_id")
	if db.Find(&newStudent, "student_id=?", studentId).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	}

	var reqStudent = newStudent
	c.Bind(&reqStudent)
	db.Model(&newStudent).Where("student_id=?", studentId).Updates(reqStudent)
	c.JSON(http.StatusOK, gin.H{
		"message": "success update data",
		"data": reqStudent,
	})
}

func delHandler(c *gin.Context, db *gorm.DB) {
	var newStudent newStudent
	studentId := c.Param("student_id")
	db.Delete(&newStudent, "student_id=?", studentId)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Deleted Data",
	})
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&newStudent{})

	data := newStudent{}
	if db.Find(&data).RowsAffected == 0 {
		fmt.Println("-------Run Seeder User-------")
		seederUser(db)
	}
}

func seederUser(db *gorm.DB) {
	data := newStudent{
		Student_id: 1,
		Student_name: "Zexceed",
		Student_age: 24,
		Student_address: "Bandung",
		Student_phone_no: "0821231234",
	}

	db.Create(&data)
}

func setupRouter() *gin.Engine {
	conn := "host=localhost user=postgres password=user001 dbname=iprijaya port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	r := gin.Default()

	r.POST("/student", func(ctx *gin.Context) {
		postHandler(ctx, db)
	})

	r.GET("/student", func(ctx *gin.Context) {
		getAllHandler(ctx, db)
	})

	r.GET("/student/:student_id", func(ctx *gin.Context) {
		getHandler(ctx, db)
	})

	r.PUT("/student/:student_id", func(ctx *gin.Context) {
		putHandler(ctx, db)
	})

	r.DELETE("/student/:student_id", func(ctx *gin.Context) {
		delHandler(ctx, db)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}