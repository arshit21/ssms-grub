package main

import (
	"ssms_grub/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/grub/post", controller.CreateGrub)
	router.GET("/grub/getall", controller.GetAllGrubs)
	router.GET("/grub/:name", controller.GetGrubByName)
	router.POST("/grub/usergrubinfo", controller.GetUserGrubInfo)
	router.POST("/grub/scanner", controller.HandleScan)
	router.Run("localhost:8080")

}
