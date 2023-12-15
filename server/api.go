package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func create[T Company | Team | Route | Driver | RoadManager | Violation | Vehicle | Station]() func(c *gin.Context) {
	return func(c *gin.Context) {
		var d T
		if err := db.Create(&d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}

func get[T Company | Team | Route | Driver | RoadManager | Violation | Vehicle | Station]() func(c *gin.Context) {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		var d T
		if err := db.Where(params).Find(&d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}

func update[T Company | Team | Route | Driver | RoadManager | Violation | Vehicle | Station]() func(c *gin.Context) {
	return func(c *gin.Context) {
		var d T
		if err := db.Save(&d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}

func delete[T Company | Team | Route | Driver | RoadManager | Violation | Vehicle | Station]() func(c *gin.Context) {
	return func(c *gin.Context) {
		var d T
		params := c.Request.URL.Query()
		if err := db.Where(params).Delete(&d).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, d)
		}
	}
}
func init_router() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/company", get[Company]())
		api.GET("/team", get[Team]())
		api.GET("/route", get[Route]())
		api.GET("/driver", get[Driver]())
		api.GET("/road_manager", get[RoadManager]())
		api.GET("/violation", get[Violation]())
		api.GET("/vehicle", get[Vehicle]())
		api.GET("/station", get[Station]())
		api.PUT("/company", update[Company]())
		api.PUT("/team", update[Team]())
		api.PUT("/route", update[Route]())
		api.PUT("/driver", update[Driver]())
		api.PUT("/road_manager", update[RoadManager]())
		api.PUT("/violation", update[Violation]())
		api.PUT("/vehicle", update[Vehicle]())
		api.PUT("/station", create[Station]())
		api.POST("/team", create[Team]())
		api.POST("/route", create[Route]())
		api.POST("/driver", create[Driver]())
		api.POST("/road_manager", create[RoadManager]())
		api.POST("/violation", create[Violation]())
		api.POST("/vehicle", create[Vehicle]())
		api.POST("/station", create[Station]())
		api.DELETE("/team", delete[Team]())
		api.DELETE("/route", delete[Route]())
		api.DELETE("/driver", delete[Driver]())
		api.DELETE("/road_manager", delete[RoadManager]())
		api.DELETE("/violation", delete[Violation]())
		api.DELETE("/vehicle", delete[Vehicle]())
		api.DELETE("/station", delete[Station]())
	}

	r.Run(":8080")
}
