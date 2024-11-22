package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Activity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var activities = []Activity{
	{
		Id:          "1",
		Name:        "activity1",
		Description: "activity1 description",
	},
	{
		Id:          "2",
		Name:        "activity2",
		Description: "activity2 description",
	},
	{
		Id:          "3",
		Name:        "activity3",
		Description: "activity3 description",
	},
}

func getActivities(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, activities)
}

func main() {
	router := gin.Default()
	router.GET("/activities", getActivities)
	router.Run("localhost:8080")
}
