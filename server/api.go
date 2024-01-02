package main

import (
	"errors"
	"net/http"
	"time"

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

type ReqeustQuery struct {
	Name string
	Time []string
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
		api.POST("query/driver", get[Driver]())           //查询某车队下司机基本信息
		api.POST("query/violation/driver", func(c *gin.Context) {
			var data ReqeustQuery
			c.ShouldBindJSON(&data)
			query := db.Model(&Violation{}).Joins("Vehicle").Joins("Team").Joins("Route").Joins("Driver")
			start, _ := time.Parse(time.RFC3339, data.Time[0])
			end, _ := time.Parse(time.RFC3339, data.Time[1])
			query = query.Where("Driver.Name = ? AND occurred_at BETWEEN ? AND ?", data.Name, start, end)
			var violations []Violation
			query.Find(&violations)
			c.JSON(200, violations)
		})
		api.POST("query/violation/team", func(c *gin.Context) {
			var data ReqeustQuery
			c.ShouldBindJSON(&data)
			query := db.Model(&Violation{}).Joins("Vehicle").Joins("Team").Joins("Route").Joins("Driver")
			start, _ := time.Parse(time.RFC3339, data.Time[0])
			end, _ := time.Parse(time.RFC3339, data.Time[1])
			query = query.Where("Team.Name = ? AND occurred_at BETWEEN ? AND ?", data.Name, start, end)
			var violations []Violation
			query.Find(&violations)
			c.JSON(200, violations)
		})
	}
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("./dist/"))))
}
