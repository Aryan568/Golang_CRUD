package routes

import (
	"Go_test/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/addData", controller.AddData)

	router.DELETE("/deleteAllData", controller.DeleteData)
	router.DELETE("/deleteDataById/:id", controller.DeleteDataByID)

	router.GET("getAllData", controller.GetData)
	router.GET("getDataById/:id", controller.GetDataByID)

	router.PUT("updateData/:id", controller.UpdateData)
}
