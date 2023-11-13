package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type planets struct
{
	Id string `json:"id"`
	Name string `json:"name"`
	Diameter int `json:"Diameter"`
	Distance string `json:"distance"`
	Position int8 `json:"position"`
}

var planets_arr = [] planets{
	{Id: "1", Name: "Mercury", Diameter: 4879, Distance: "212 million KM", Position: 1},
	{Id: "2", Name: "Venus", Diameter: 12104, Distance: "126.89 million KM", Position: 2},
	{Id: "3", Name: "Earth", Diameter: 12742, Distance: "148.05 million KM from Sun", Position: 3},
	{Id: "4", Name: "Mars", Diameter: 6779, Distance: "380 million KM", Position: 4},
	{Id: "5", Name: "Jupiter", Diameter: 139820 , Distance: "598 million KM", Position: 5},
	{Id: "6", Name: "Saturn", Diameter: 116460, Distance: "1,425,285,860 Billion KM", Position: 6},
	{Id: "7", Name: "Uranus", Diameter: 50724, Distance: "1.7 Billion KM", Position: 7},
	{Id: "8", Name: "Naptune", Diameter: 49244, Distance: "2.703 billion KM", Position: 8},	
}

// Get Method
func getplanets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, planets_arr)
}

// Post Method
// func postplanets(c*gin.Context){
// 	var newplanet planets_arr

// 	if err := c.BindJSON(&newplanet); err! = nil{
// 		return
// 	}

// 	planets_arr = append(planets_arr, newplanet)
// 	c.IndentedJSON(http.StatusCreated, newplanet)
// }


// Main method
func main(){
	router := gin.Default()
	router.GET("/planets", getplanets)
	// router.POST("/planets", postplanets)
	router.Run("localhost:3000")
}