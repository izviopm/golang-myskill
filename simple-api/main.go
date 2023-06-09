package main

import (
	"database/sql"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type newStudent struct {
	Student_id			uint64 `json:"student_id" binding:"required"`
	Student_name 		string `json:"student_name" binding:"required"`
	Student_age			uint64 `json:"student_age" binding:"required"`
	Student_address		string `json:"student_address" binding:"required"`
	Student_phone_no	string `json:"student_phone_no" binding:"required"`
}

func rowToStruct(rows *sql.Rows, dest interface{}) error {
	destv := reflect.ValueOf(dest).Elem()

	args := make([]interface{}, destv.Type().Elem().NumField())

	for rows.Next() {
		rowp := reflect.New(destv.Type().Elem())
		rowv := rowp.Elem()

		for i := 0; i < rowv.NumField(); i++ {
			args[i] = rowv.Field(i).Addr().Interface()
		}

		if err := rows.Scan(args...); err != nil {
			return err
		}

		destv.Set(reflect.Append(destv, rowv))
	}

	return nil
}

func postHandler(c *gin.Context, db *sql.DB) {
	var newStudent newStudent

	if c.Bind(&newStudent) == nil {
		_, err := db.Exec("INSERT INTO students VALUES ($1, $2, $3, $4, $5)", newStudent.Student_id, newStudent.Student_name, newStudent.Student_age, newStudent.Student_address, newStudent.Student_phone_no)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Success Insert Data to the Database.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	}
}

func getAllHandler(c *gin.Context, db *sql.DB) {
	var newStudent []newStudent

	row, err := db.Query("SELECT * FROM students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	rowToStruct(row, &newStudent)

	if newStudent == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newStudent,
	})
}

func getHandler(c *gin.Context, db *sql.DB) {
	var newStudent []newStudent

	studentId := c.Param("student_id")
	row, err := db.Query("SELECT * FROM students WHERE student_id = $1", studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	rowToStruct(row, &newStudent)

	if newStudent == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newStudent,
	})
}

func setupRouter() *gin.Engine {
	conn := "postgres://postgres:user001@localhost/iprijaya?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

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

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}