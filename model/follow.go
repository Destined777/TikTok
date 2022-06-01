package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	UserID		int64
	ToID		int64
	IsFollow	bool
}
