package Models

import (
	"../Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllAuthor(a *[]Author) (err error) {
	if err = Config.DB.Find(a).Error; err != nil {
		return err
	}
	return nil
}