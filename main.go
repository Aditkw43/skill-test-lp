package main

import (
	"skill-test-lp/config"
	"skill-test-lp/internal"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	db := config.InitGorm()
	data := internal.GetData()
	controller := internal.NewController(db, data)
	r.POST("/insert-payroll", controller.InsertPayroll)

	r.Run(":8080")
}
