// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controllers

import (
	"net/http"

	"bitbucket.com/irb/api/middlewares"
	"bitbucket.com/irb/api/models"
	"bitbucket.com/irb/api/plugins"
	"bitbucket.com/irb/api/utils"
	"github.com/gin-gonic/gin"
)

//ApplicationControllers - map of all the applications controllers
var ApplicationControllers = map[string]func(*gin.Context){
	"getApplications":   GetApplications,
	"getMyApplications": GetMyApplications,
	"createApplication": CreateApplication,
	"getApplication":    GetApplication,
	"updateApplication": UpdateApplication,
	"deleteApplication": DeleteApplication,
}

//GetApplications - List all Applications
// @Summary List all Applications
// @Tags Application
// @Produce json
// @Success 200 {object} models.Application
// @Router /application [get]
// @Security ApiKeyAuth
func GetApplications(c *gin.Context) {
	offsetString := c.Query("offset")
	limitString := c.Query("limit")
	offset, err := utils.ConvertStringToInt(offsetString)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Offset conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "Offset conv error", err)
	}
	limit, errs := utils.ConvertStringToInt(limitString)
	if errs != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Limit conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "Limit conv error", errs)
	}
	var application []models.Application
	count, err := models.GetAllApplications(&application, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error getting all applications", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":      count,
			"data":       application,
			"statusCode": 200,
		})
	}
}

//GetMyApplications - List all my Applications
// @Summary List all my Applications
// @Tags Application
// @Produce json
// @Success 200 {object} models.Application
// @Router /application/:id/my [get]
// @Security ApiKeyAuth
func GetMyApplications(c *gin.Context) {
	offsetString := c.Query("offset")
	limitString := c.Query("limit")
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	res, _ := middlewares.GetSession(c)
	offset, offsetErr := utils.ConvertStringToInt(offsetString)
	limit, errs := utils.ConvertStringToInt(limitString)
	if offsetErr != nil || idErr != nil || errs != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error offset conv to int", offsetErr)
		plugins.LogError("API", "error id conv to int", idErr)
		plugins.LogError("API", "error limit conv to int", errs)
	}
	if res.IsAdmin || id == res.UserID {
		var application []models.Application
		count, err := models.GetMyApplications(&application, id, offset, limit)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			plugins.LogError("API", "error getting your applications", err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"count":      count,
				"data":       application,
				"statusCode": 200,
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Needs Elevation",
			"statusCode": 400,
		})
	}
}

//CreateApplication - Create an Application
// @Summary Creates a new Application
// @Description Creates a new Application
// @Tags Application
// @Accept  json
// @Produce json
// @Param application body models.Application true "Create Application"
// @Success 200 {object} models.Application
// @Router /application/create [post]
// @Security ApiKeyAuth
func CreateApplication(c *gin.Context) {
	var application models.Application
	c.BindJSON(&application)
	stats, err := models.CreateApplication(&application)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error creating application", err)
	} else {
		if stats != true {
			c.JSON(http.StatusConflict, gin.H{
				"message":    "Application already exists",
				"statusCode": 409,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":    "Application created successfully",
				"statusCode": 200,
			})
		}
	}
}

//GetApplication - Get a particular Application with id
// @Summary Retrieves application based on given ID
// @Tags Application
// @Produce json
// @Param id path integer true "Application ID"
// @Success 200 {object} models.Application
// @Router /application/{id} [get]
// @Security ApiKeyAuth
func GetApplication(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	res, _ := middlewares.GetSession(c)
	if idErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error id conv to int", idErr)
	}
	if res.IsAdmin || id == res.UserID {
		var application models.Application
		err := models.GetApplication(&application, id)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			plugins.LogError("API", "error getting application", err)
		} else {
			c.JSON(http.StatusOK, application)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Needs Elevation",
			"statusCode": 400,
		})
	}
}

//UpdateApplication - Update an existing Application
// @Summary Updates application based on given ID
// @Tags Application
// @Produce json
// @Param id path integer true "Application ID"
// @Success 200 {object} models.Application
// @Router /application/{id} [patch]
// @Security ApiKeyAuth
func UpdateApplication(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	res, _ := middlewares.GetSession(c)
	if idErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error id conv to int", idErr)
	}
	if res.IsAdmin || id == res.UserID {
		var application models.Application
		err := models.GetApplication(&application, id)
		if err != nil {
			c.JSON(http.StatusNotFound, application)
			plugins.LogError("API", "error getting application", err)
		}
		c.BindJSON(&application)
		err = models.UpdateApplication(&application, id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			plugins.LogError("API", "error updating application", err)
		} else {
			c.JSON(http.StatusOK, application)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Needs Elevation",
			"statusCode": 400,
		})
	}
}

//DeleteApplication - Deletes Application
// @Summary Deletes a application based on given ID
// @Tags Application
// @Produce json
// @Param id path integer true "Application ID"
// @Success 200 {object} models.Application
// @Router /application/{id} [delete]
// @Security ApiKeyAuth
func DeleteApplication(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "error id conv to int", err)
	}
	errs := models.DeleteApplication(id)
	if errs != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error deleting application", errs)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":    "Application Deleted successfully",
			"statusCode": 200,
		})
	}
}
