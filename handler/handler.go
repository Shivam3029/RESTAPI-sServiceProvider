package handler

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	//"reflect"
	"log"
	"ServiceProvider/models"
	
)
var serviceproviders = []models.ServiceProvider{{ID: 1,Name :"SP1"},
				{ID:2,Name:"SP2"},
				{ID:3,Name:"SP3"},
				{ID:4,Name:"SP4"},
				{ID:5,Name:"SP5"}}



func GetServiceProviders(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": serviceproviders})
	
}

func GetServiceProvider(c *gin.Context){
	//var serviceprovider ServiceProvider
	serviceproviderID := c.Param("id")
	//log.Println(reflect.TypeOf(serviceproviderID))
	i,_:=strconv.Atoi(serviceproviderID)
	for _,serviceprovider:=range serviceproviders{
		if serviceprovider.ID==i{
			c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"data":serviceprovider})
			
		}
	}
	
	//i,_:=strconv.Atoi(serviceproviderID)
	//log.Println(reflect.TypeOf(i))
	
}
func AddServiceProvider(c *gin.Context) {
	var serviceprovider models.ServiceProvider
	c.BindJSON(&serviceprovider)
	log.Println(serviceprovider)
	//serviceproviders=append(serviceproviders,serviceprovider)
	for _,item:=range serviceproviders{
		if item.ID==serviceprovider.ID{
			c.JSON(409,gin.H{"status":409,"message":"service provider already exists"})
			return
		}
	}
	serviceproviders=append(serviceproviders,serviceprovider)
	c.JSON(http.StatusOK,gin.H{
		"id":serviceprovider.ID,
		"name":serviceprovider.Name,
	})
	//serviceproviders=append(serviceproviders,serviceprovider)
	//c.JSON(http.StatusOK,gin.H{
	//	"id":serviceprovider.ID,
	//	"name":serviceprovider.Name,
	//})
	
}
func UpdateServiceProvider(c *gin.Context){

	var serviceprovider models.ServiceProvider
	c.BindJSON(&serviceprovider)
	
	for i,item:=range serviceproviders{
		if item.ID==serviceprovider.ID{
			serviceproviders[i]=serviceprovider
		}
	}
	//serviceproviders=append(serviceproviders,serviceprovider)
	c.JSON(http.StatusOK,gin.H{
		"id":serviceprovider.ID,
		"name":serviceprovider.Name,
	})
	
	//c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"data":serviceproviders})
}
func DeleteServiceProvider(c *gin.Context){
	serviceproviderID := c.Param("id")
	id,_:=strconv.Atoi(serviceproviderID)
	//id, _:=strconv.Atoi(params["id"])
	for i,item:=range serviceproviders{
		if item.ID==id{
			serviceproviders=append(serviceproviders[:i],serviceproviders[i+1:]...)
		}
	}
	c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"data":serviceproviders})
	
	//json.NewEncoder(w).Encode(serviceproviders)
	//log.Println("deleting a service provider")
}