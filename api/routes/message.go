// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"bitbucket.com/irb/api/controllers"
	"bitbucket.com/irb/api/middlewares"
	"github.com/gin-gonic/gin"
)

//MessageRouter - message subroutes
func MessageRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1/message")
	{
		v1.GET("/:id", middlewares.AuthenticationMiddleware(), controllers.MessageControllers["getMessage"])
		v1.GET("/:id/conversation", middlewares.AuthenticationMiddleware(), controllers.MessageControllers["getConversationMessages"])
		v1.PUT("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.MessageControllers["updateMessages"])
		v1.DELETE("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.MessageControllers["deleteMessage"])
		v1.GET("/", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.MessageControllers["getMessages"])
		v1.POST("/create", middlewares.AuthenticationMiddleware(), controllers.MessageControllers["createMessage"])
	}
	return r
}
