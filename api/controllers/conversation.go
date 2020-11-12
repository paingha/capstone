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

//ConversationControllers - map of all the conversations controllers
var ConversationControllers = map[string]func(*gin.Context){
	"getConversations":   GetConversations,
	"getMyConversations": GetMyConversations,
	"createConversation": CreateConversation,
	"getConversation":    GetConversation,
	"updateConversation": UpdateConversation,
	"deleteConversation": DeleteConversation,
}

//GetConversations - List all Conversations
// @Summary List all Conversations
// @Tags Conversation
// @Produce json
// @Success 200 {object} models.Conversation
// @Router /conversation [get]
// @Security ApiKeyAuth
func GetConversations(c *gin.Context) {
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
	var conversation []models.Conversation
	count, err := models.GetAllConversations(&conversation, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error getting all conversations", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":      count,
			"data":       conversation,
			"statusCode": 200,
		})
	}
}

//GetMyConversations - List all my Conversations
// @Summary List all my Conversations
// @Tags Conversation
// @Produce json
// @Success 200 {object} models.Conversation
// @Router /conversation/:id/my [get]
// @Security ApiKeyAuth
func GetMyConversations(c *gin.Context) {
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
		var conversation []models.Conversation
		count, err := models.GetMyConversations(&conversation, id, offset, limit)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			plugins.LogError("API", "error getting your conversations", err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"count":      count,
				"data":       conversation,
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

//CreateConversation - Create a Conversation
// @Summary Creates a new Conversation
// @Description Creates a new Conversation
// @Tags Conversation
// @Accept  json
// @Produce json
// @Param conversation body models.Conversation true "Create Conversation"
// @Success 200 {object} models.Conversation
// @Router /conversation/create [post]
// @Security ApiKeyAuth
func CreateConversation(c *gin.Context) {
	var conversation models.Conversation
	c.BindJSON(&conversation)
	stats, err := models.CreateConversation(&conversation)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error creating conversation", err)
	} else {
		if stats != true {
			c.JSON(http.StatusConflict, gin.H{
				"message":    "Conversation already exists",
				"statusCode": 409,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":    "Conversation created successfully",
				"statusCode": 200,
			})
		}
	}
}

//GetConversation - Get a particular Conversation with id
// @Summary Retrieves conversation based on given ID
// @Tags Conversation
// @Produce json
// @Param id path integer true "Conversation ID"
// @Success 200 {object} models.Conversation
// @Router /conversation/{id} [get]
// @Security ApiKeyAuth
func GetConversation(c *gin.Context) {
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
		var conversation models.Conversation
		err := models.GetConversation(&conversation, id)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			plugins.LogError("API", "error getting conversation", err)
		} else {
			c.JSON(http.StatusOK, conversation)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Needs Elevation",
			"statusCode": 400,
		})
	}
}

//UpdateConversation - Update an existing Conversation
// @Summary Updates conversation based on given ID
// @Tags Conversation
// @Produce json
// @Param id path integer true "Conversation ID"
// @Success 200 {object} models.Conversation
// @Router /conversation/{id} [patch]
// @Security ApiKeyAuth
func UpdateConversation(c *gin.Context) {
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
		var conversation models.Conversation
		err := models.GetConversation(&conversation, id)
		if err != nil {
			c.JSON(http.StatusNotFound, conversation)
			plugins.LogError("API", "error getting conversation", err)
		}
		c.BindJSON(&conversation)
		err = models.UpdateConversation(&conversation, id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			plugins.LogError("API", "error updating conversation", err)
		} else {
			c.JSON(http.StatusOK, conversation)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Needs Elevation",
			"statusCode": 400,
		})
	}
}

//DeleteConversation - Deletes Conversation
// @Summary Deletes a conversation based on given ID
// @Tags Conversation
// @Produce json
// @Param id path integer true "Conversation ID"
// @Success 200 {object} models.Conversation
// @Router /conversation/{id} [delete]
// @Security ApiKeyAuth
func DeleteConversation(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "error id conv to int", err)
	}
	errs := models.DeleteConversation(id)
	if errs != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error deleting conversation", errs)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":    "Conversation Deleted successfully",
			"statusCode": 200,
		})
	}
}
