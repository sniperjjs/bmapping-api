package models

import (
	"time"
)

type ElandStore struct {
	Id          int64
	GroupId     int64
	Code        string
	Name        string
	LogoUrl     string
	QrUrl       string
	Description string
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}
