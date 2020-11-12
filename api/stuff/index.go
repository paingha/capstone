// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stuff

import (
	"bitbucket.com/irb/api/service"
)

//MailService - pointer to hold initalized mail service
var MailService *service.QueueService

//SmsService - pointer to hold initalized sms service
var SmsService *service.QueueService

//PushService - pointer to hold initalized push service
var PushService *service.QueueService

//UploadService - pointer to hold initalized upload service
var UploadService *service.QueueService
