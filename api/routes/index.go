// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"bitbucket.com/irb/api/plugins"
	"github.com/gin-gonic/gin"
)

//SetupRouter - setup routes for api
func SetupRouter(inboxHub *plugins.Hub) *gin.Engine {
	r := gin.Default()
	{
		RootRouter(r)
		//WebsocketRouter(r, inboxHub)
		UserRouter(r)
		ConversationRouter(r)
		MessageRouter(r)
		ApplicationRouter(r)
	}
	return r
}
