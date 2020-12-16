package model

type Tag struct {
	ID    int    `json:"id" gorm:"primary_key;not null"`
	Name  string `json:"name"`
	State int    `json:"state"`
	Base
}


