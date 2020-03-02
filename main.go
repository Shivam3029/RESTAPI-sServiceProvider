package main

import (
	   "github.com/gin-gonic/gin"
	   //"net/http"
	   //"strconv"
	   //"reflect"
	   //"log"
	   //"ServiceProvider/models"
	   "ServiceProvider/handler"
	   
)
//type ServiceProvider struct{
//	ID int `json:"id"`
//	Name string `json:"name"`
//}


//var serviceproviders []models.ServiceProvider


func main() {
	
router := gin.Default()
// serviceproviders=append(serviceproviders, models.ServiceProvider{ID: 1,Name :"SP1"},
// 						models.ServiceProvider{ID:2,Name:"SP2"},
// 						models.ServiceProvider{ID:3,Name:"SP3"},
// 						models.ServiceProvider{ID:4,Name:"SP4"},
// 						models.ServiceProvider{ID:5,Name:"SP5"})
v1 := router.Group("/serviceproviders")
{
	v1.POST("/", handler.AddServiceProvider)
	v1.GET("/", handler.GetServiceProviders)
	v1.GET("/:id", handler.GetServiceProvider)
	v1.PATCH("/", handler.UpdateServiceProvider)
	v1.DELETE("/:id", handler.DeleteServiceProvider)
}
 router.Run()
}
