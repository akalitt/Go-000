package repository

import (
	"github.com/jinzhu/gorm"
	//"github/akalitt/go-errors-example/internal/model"
)

type Dao struct {
	Db *gorm.DB
}
