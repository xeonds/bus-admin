package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func create[T any]() func(c *gin.Context) {
	_create := func(model any) error {
		if model == nil {
			return errors.New("model is nil")
		}
		return db.Create(&model).Error
	}

	return func(c *gin.Context) {
		var model T
		if err := c.ShouldBindJSON(&model); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := _create(&model); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, model)
	}
}

func get[T any]() func(c *gin.Context) {
	_get := func(id string) (T, error) {
		var model T
		if err := db.First(&model, id).Error; err != nil {
			return model, err
		}
		return model, nil
	}
	_getAll := func() ([]T, error) {
		var models []T
		if err := db.Find(&models).Error; err != nil {
			return models, err
		}
		return models, nil
	}

	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			models, err := _getAll()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, models)
		} else {
			model, err := _get(id)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, model)
		}
	}
}

func update[T any]() func(c *gin.Context) {
	_update := func(model any) error {
		if model == nil {
			return errors.New("model is nil")
		}
		return db.Save(&model).Error
	}

	return func(c *gin.Context) {
		// id := c.Param("id")
		var model T
		// TODO: check exist
		if err := c.ShouldBindJSON(&model); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := _update(&model); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, model)
	}
}

func delete[T any]() func(c *gin.Context) {
	_delete := func(model any) error {
		if model == nil {
			return errors.New("model is nil")
		}
		return db.Delete(&model).Error
	}
	_deleteById := func(id string, model any) error {
		if model == nil {
			return errors.New("model is nil")
		}
		return db.Where("id = ?", id).Delete(&model).Error
	}

	return func(c *gin.Context) {
		id := c.Param("id")
		var model T
		if id == "" {
			// TODO: check exist
			if err := c.ShouldBindJSON(&model); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			if err := _delete(&model); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		} else {
			if err := _deleteById(id, &model); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
		c.JSON(http.StatusOK, model)
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
		api.POST("company", create[Company]())
		api.POST("team", create[Team]())
		api.POST("route", create[Route]())
		api.POST("driver", create[Driver]())
		api.POST("road_manager", create[RoadManager]())
		api.POST("violation", create[Violation]())
		api.POST("vehicle", create[Vehicle]())
		api.DELETE("company", delete[Company]())
		api.DELETE("team", delete[Team]())
		api.DELETE("route", delete[Route]())
		api.DELETE("driver", delete[Driver]())
		api.DELETE("road_manager", delete[RoadManager]())
		api.DELETE("violation", delete[Violation]())
		api.DELETE("vehicle", delete[Vehicle]())
		// service apis
		api.POST("record/driver", create[Driver]())       //录入司机基本信息
		api.POST("record/vehicle", create[Vehicle]())     //录入汽车基本信息
		api.POST("record/violation", create[Violation]()) //录入司机违章信息
		api.GET("query/driver", get[Driver]())            //查询某车队下司机基本信息
		api.GET("query/violation/driver", get[Driver]())  //查询某司机在某时间段落的违章详细信息
		api.GET("query/violation/team", get[Team]())      //查询某车队在某时间段的违章统计信息
	}
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("./dist/"))))
}
