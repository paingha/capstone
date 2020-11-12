// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"bitbucket.com/irb/api/config"
)

//Messages - message data struct
type Messages struct {
	ID             int        `json:"id,omitempty" sql:"primary_key"`
	Content        string     `gorm:"not null" json:"content"`
	ConversationID int        `gorm:"not null" json:"conversationId"`
	SenderID       int        `gorm:"not null" json:"senderId"`
	ReceiverID     int        `gorm:"not null" json:"receiverId"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

//TableName - messages db table name override
func (Messages) TableName() string {
	return "messages"
}

//GetAllMessages - fetch all messages at once
func GetAllMessages(message *[]Messages, offset int, limit int) (int, error) {
	var count = 0
	if err := config.DB.Model(&Messages{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(message).Error; err != nil {
		return count, err
	}
	return count, nil
}

//GetConversationMessages - fetch conversation messages
func GetConversationMessages(message *[]Messages, id int, offset int, limit int) (int, error) {
	var count = 0
	if err := config.DB.Model(&Messages{}).Where("conversation_id = ?", id).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(message).Error; err != nil {
		return count, err
	}
	return count, nil
}

//CreateMessage - create a message
func CreateMessage(message *Messages) (bool, error) {
	//Chore: return the verification token
	if err := config.DB.Create(message).Error; err != nil {
		return false, err
	}
	return false, nil
}

//GetMessage - fetch one message
func GetMessage(message *Messages, id int) error {
	if err := config.DB.Where("id = ?", id).First(message).Error; err != nil {
		return err
	}
	return nil
}

//UpdateMessage - update a message
func UpdateMessage(message *Messages, id int) error {
	if err := config.DB.Model(&message).Where("id = ?", id).Updates(message).Error; err != nil {
		return err
	}
	return nil
}

//DeleteMessage - delete a message
func DeleteMessage(id int) error {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Messages{}).Error; err != nil {
		return err
	}
	return nil
}
