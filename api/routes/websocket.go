// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"bitbucket.com/irb/api/plugins"
	"github.com/gin-gonic/gin"
)

//WebsocketRouter - root subroutes
func WebsocketRouter(r *gin.Engine, inboxHub *plugins.Hub) *gin.Engine {
	inbox := r.Group("/ws/inbox")
	{
		inbox.GET("/:room", func(c *gin.Context) {
			//plugins.ServeWs(inboxHub, c)
		})
	}
	return r
}
