package model

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Text    string `json:"text"`
	Likes   int    `json:"likes"`
	User_ID uint   `json:"user_id"`
}

type Comment struct {
	gorm.Model
	Text     string `json:"text"`
	Tweet_ID string `json:"tweet_id"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
