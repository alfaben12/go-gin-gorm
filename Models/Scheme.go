package Models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name     string `json:"name" bson:"_name" query:"name" form:"name"`
	Author   string `json:"author" bson:"_author" query:"author" form:"author"`
	Category string `json:"category" bson:"_category" query:"category" form:"category"`
}

func (b *Book) TableName() string {
	return "book"
}
