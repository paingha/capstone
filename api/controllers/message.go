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

//MessageControllers - map of all the messages controllers
var MessageControllers = map[string]func(*gin.Context){
	"getMessages":             GetMessages,
	"getConversationMessages": GetConversationMessages,
	"createMessage":           CreateMessage,
	"getMessage":              GetMessage,
	"updateMessage":           UpdateMessage,
	"deleteMessage":           DeleteMessage,
}

//GetMessages - List all Messages
// @Summary List all Messages
// @Tags Message
// @Produce json
// @Success 200 {object} models.Message
// @Router /message [get]
// @Security ApiKeyAuth
func GetMessages(c *gin.Context) {
	offsetString := c.Query("offset")
	limitString := c.Query("limit")
	offset, err := utils.ConvertStringToInt(offsetString)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Offset conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error offset conv to int", err)
	}
	limit, errs := utils.ConvertStringToInt(limitString)
	if errs != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Limit conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error limit conv to int", errs)
	}
	var message []models.Messages
	count, err := models.GetAllMessages(&message, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error getting messages", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":      count,
			"data":       message,
			"statusCode": 200,
		})
	}
}

//GetConversationMessages - List all Conversation Messages
// @Summary List all my Messages
// @Tags Message
// @Produce json
// @Success 200 {object} models.Message
// @Router /message/:id/my [get]
// @Security ApiKeyAuth
func GetConversationMessages(c *gin.Context) {
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
		var message []models.Messages
		count, err := models.GetConversationMessages(&message, id, offset, limit)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			plugins.LogError("API", "error getting conversation messages", err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"count":      count,
				"data":       message,
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

//CreateMessage - Create a Message
// @Summary Creates a new Message
// @Description Creates a new Message
// @Tags Message
// @Accept  json
// @Produce json
// @Param message body models.Message true "Create Message"
// @Success 200 {object} models.Message
// @Router /message/create [post]
// @Security ApiKeyAuth
func CreateMessage(c *gin.Context) {
	var message models.Messages
	c.BindJSON(&message)
	stats, err := models.CreateMessage(&message)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error creating message", err)
	} else {
		if stats != true {
			c.JSON(http.StatusConflict, gin.H{
				"message":    "Message already exists",
				"statusCode": 409,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":    "Message created successfully",
				"statusCode": 200,
			})
		}
	}
}

//GetMessage - Get a particular Message with id
// @Summary Retrieves message based on given ID
// @Tags Message
// @Produce json
// @Param id path integer true "Message ID"
// @Success 200 {object} models.Message
// @Router /message/{id} [get]
// @Security ApiKeyAuth
func GetMessage(c *gin.Context) {
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
		var message models.Messages
		err := models.GetMessage(&message, id)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			plugins.LogError("API", "error getting message", err)
		} else {
			c.JSON(http.StatusOK, message)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Needs Elevation",
			"statusCode": 400,
		})
	}
}

//UpdateMessage - Update an existing Message
// @Summary Updates message based on given ID
// @Tags Message
// @Produce json
// @Param id path integer true "Message ID"
// @Success 200 {object} models.Message
// @Router /message/{id} [patch]
// @Security ApiKeyAuth
func UpdateMessage(c *gin.Context) {
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
		var message models.Messages
		err := models.GetMessage(&message, id)
		if err != nil {
			c.JSON(http.StatusNotFound, message)
			plugins.LogError("API", "error getting message", err)
		}
		c.BindJSON(&message)
		err = models.UpdateMessage(&message, id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			plugins.LogError("API", "error updating message", err)
		} else {
			c.JSON(http.StatusOK, message)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Needs Elevation",
			"statusCode": 400,
		})
	}
}

//DeleteMessage - Deletes Message
// @Summary Deletes a message based on given ID
// @Tags Message
// @Produce json
// @Param id path integer true "Message ID"
// @Success 200 {object} models.Message
// @Router /message/{id} [delete]
// @Security ApiKeyAuth
func DeleteMessage(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "error id conv to int", err)
	}
	errs := models.DeleteMessage(id)
	if errs != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error deleting message. message not found", errs)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":    "Message Deleted successfully",
			"statusCode": 200,
		})
	}
}
