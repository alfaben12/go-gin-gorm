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

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Author struct {
	gorm.Model
	Name     string `json:"name" bson:"_name" query:"name" form:"name"`
}

func (b *Book) TableName() string {
	return "book"
}

func (b *Author) TableName() string {
	return "author"
}

func (b *Credential) TableName() string {
	return "credential"
}


