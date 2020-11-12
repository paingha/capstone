// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"bitbucket.com/irb/api/controllers"
	"bitbucket.com/irb/api/middlewares"
	"github.com/gin-gonic/gin"
)

//ConversationRouter - conversation subroutes
func ConversationRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1/conversation")
	{
		v1.GET("/:id", middlewares.AuthenticationMiddleware(), controllers.ConversationControllers["getConversation"])
		v1.GET("/:id/my", middlewares.AuthenticationMiddleware(), controllers.ConversationControllers["getMyConversations"])
		v1.PUT("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.ConversationControllers["updateConversations"])
		v1.DELETE("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.ConversationControllers["deleteConversation"])
		v1.GET("/", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.ConversationControllers["getConversations"])
		v1.POST("/create", middlewares.AuthenticationMiddleware(), controllers.ConversationControllers["createConversation"])
	}
	return r
}
