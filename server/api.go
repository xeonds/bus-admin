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
		return db.Create(model).Error
	}

	return func(c *gin.Context) {
		var model T
		if err := c.ShouldBindJSON(&model); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": model})
			return
		}
		if err := _create(&model); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": model})
			return
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
				return
			}
			c.JSON(http.StatusOK, models)
		} else {
			model, err := _get(id)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
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
		return db.Save(model).Error
	}

	return func(c *gin.Context) {
		// id := c.Param("id")
		var model T
		// TODO: check exist
		if err := c.ShouldBindJSON(&model); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := _update(&model); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, model)
	}
}

func delete[T any]() func(c *gin.Context) {
	_delete := func(model any) error {
		if model == nil {
			return errors.New("model is nil")
		}
		return db.Delete(model).Error
	}
	_deleteById := func(id string, model any) error {
		if model == nil {
			return errors.New("model is nil")
		}
		return db.Where("id = ?", id).Delete(model).Error
	}

	return func(c *gin.Context) {
		id := c.Param("id")
		var model T
		if id == "" {
			// TODO: check exist
			if err := c.ShouldBindJSON(&model); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := _delete(&model); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			if err := _deleteById(id, &model); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, model)
	}
}

func CRUD[T any](r *gin.RouterGroup, relativePath string) {
	r.POST(relativePath, create[T]())
	r.GET(relativePath+"/:id", get[T]())
	r.GET(relativePath, get[T]())
	r.PUT(relativePath+"/:id", update[T]())
	r.DELETE(relativePath+"/:id", delete[T]())
}

func init_router() {
	r = gin.Default()
	api := r.Group("/api/v1/")
	{
		CRUD[Company](api, "company")
		CRUD[Team](api, "team")
		CRUD[Route](api, "route")
		CRUD[Driver](api, "driver")
		CRUD[RoadManager](api, "road_manager")
		CRUD[Violation](api, "violation")
		CRUD[Vehicle](api, "vehicle")
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
