package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func create[T Company | Team | Route | Driver | RoadManager | Violation | Vehicle]() func(c *gin.Context) {
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

func get[T Company | Team | Route | Driver | RoadManager | Violation | Vehicle]() func(c *gin.Context) {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		var d T
		if params != nil {
			if err := db.Where(params).Find(&d).Error; err != nil {
				c.AbortWithStatus(404)
				fmt.Println(err)
			} else {
				c.JSON(200, d)
			}
		} else {
			c.AbortWithStatus(404)
			fmt.Println("No params")
		}
	}
}

func update[T Company | Team | Route | Driver | RoadManager | Violation | Vehicle]() func(c *gin.Context) {
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

func delete[T Company | Team | Route | Driver | RoadManager | Violation | Vehicle]() func(c *gin.Context) {
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
	r = gin.Default()
	api := r.Group("/api/v1/")
	{
		api.GET("company", get[Company]())
		api.GET("team", get[Team]())
		api.GET("route", get[Route]())
		api.GET("driver", get[Driver]())
		api.GET("road_manager", get[RoadManager]())
		api.GET("violation", get[Violation]())
		api.GET("vehicle", get[Vehicle]())
		api.PUT("company", update[Company]())
		api.PUT("team", update[Team]())
		api.PUT("route", update[Route]())
		api.PUT("driver", update[Driver]())
		api.PUT("road_manager", update[RoadManager]())
		api.PUT("violation", update[Violation]())
		api.PUT("vehicle", update[Vehicle]())
		api.POST("team", create[Team]())
		api.POST("route", create[Route]())
		api.POST("driver", create[Driver]())
		api.POST("road_manager", create[RoadManager]())
		api.POST("violation", create[Violation]())
		api.POST("vehicle", create[Vehicle]())
		api.DELETE("team", delete[Team]())
		api.DELETE("route", delete[Route]())
		api.DELETE("driver", delete[Driver]())
		api.DELETE("road_manager", delete[RoadManager]())
		api.DELETE("violation", delete[Violation]())
		api.DELETE("vehicle", delete[Vehicle]())
	}
	{
		//录入司机基本信息
		api.PUT("record/driver", create[Driver]())
		//录入汽车基本信息
		api.PUT("record/vehicle", create[Vehicle]())
		//录入司机违章信息
		api.PUT("record/violation", create[Violation]())
		//查询某车队下司机基本信息
		api.GET("query/driver", get[Driver]())
		//查询某司机在某时间段落的违章详细信息
		api.GET("query/validation/driver", get[Driver]())
		//查询某车队在某时间段的违章统计信息
		api.GET("query/validation/team", get[Team]())
	}
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("./dist/"))))
}
