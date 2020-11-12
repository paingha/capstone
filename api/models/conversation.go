// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"bitbucket.com/irb/api/config"
)

//Conversation - conversation data struct
type Conversation struct {
	ID           int          `json:"id,omitempty" sql:"primary_key"`
	Participants Participants `gorm:"not null" json:"participants"`
	Messages     []Messages   `gorm:"foreignkey:ConversationID" json:"messages"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    *time.Time   `json:"deleted_at"`
}

//Participants - participants data struct
type Participants struct {
	PersonOne int `gorm:"not null" json:"personOne"`
	PersonTwo int `gorm:"not null" json:"personTwo"`
}

//TableName - conversations db table name override
func (Conversation) TableName() string {
	return "conversations"
}

//GetAllConversations - fetch all conversations at once
func GetAllConversations(conversation *[]Conversation, offset int, limit int) (int, error) {
	var count = 0
	if err := config.DB.Model(&Conversation{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(conversation).Error; err != nil {
		return count, err
	}
	return count, nil
}

//GetMyConversations - fetch my conversations
func GetMyConversations(conversation *[]Conversation, id int, offset int, limit int) (int, error) {
	var count = 0
	if err := config.DB.Model(&Conversation{}).Where("id = ?", id).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(conversation).Error; err != nil {
		return count, err
	}
	return count, nil
}

//CreateConversation - create a conversation
func CreateConversation(conversation *Conversation) (bool, error) {
	if err := config.DB.Create(conversation).Error; err != nil {
		return false, err
	}
	return false, nil
}

//GetConversation - fetch one conversation
func GetConversation(conversation *Conversation, id int) error {
	if err := config.DB.Where("id = ?", id).First(conversation).Error; err != nil {
		return err
	}
	return nil
}

//UpdateConversation - update a conversation
func UpdateConversation(conversation *Conversation, id int) error {
	if err := config.DB.Model(&conversation).Where("id = ?", id).Updates(conversation).Error; err != nil {
		return err
	}
	return nil
}

//DeleteConversation - delete a conversation
func DeleteConversation(id int) error {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Conversation{}).Error; err != nil {
		return err
	}
	return nil
}
