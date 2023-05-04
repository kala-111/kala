package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	id      string   `json:"id"`
	name    string   `json:"name"`
	salary  int      `json:"salary"`
	project *Project `json:"id"`
	manager *Manager `json:"maneger"`
}

type Project struct {
	gorm.Model
	Name  string `json:"name" gorm:"text;not null;default:null`
	Price string `json:"price" gorm:"text;not null;default:null`
}

type Manager struct {
	gorm.Model
	id   string `json:"id"`
	name string `json:"name"`
}
