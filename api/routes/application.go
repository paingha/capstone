// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"bitbucket.com/irb/api/controllers"
	"bitbucket.com/irb/api/middlewares"
	"github.com/gin-gonic/gin"
)

//ApplicationRouter - application subroutes
func ApplicationRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1/application")
	{
		v1.GET("/:id", middlewares.AuthenticationMiddleware(), controllers.ApplicationControllers["getApplication"])
		v1.GET("/:id/my", middlewares.AuthenticationMiddleware(), controllers.ApplicationControllers["getMyApplications"])
		v1.PUT("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.ApplicationControllers["updateApplications"])
		v1.DELETE("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.ApplicationControllers["deleteApplication"])
		v1.GET("/", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.ApplicationControllers["getApplications"])
		v1.POST("/create", middlewares.AuthenticationMiddleware(), controllers.ApplicationControllers["createApplication"])
	}
	return r
}
