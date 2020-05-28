package model

import (
	"go-gin-template/exception"
	"log"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)"`
}
type RawTag struct {
	Name string `gorm:"type:varchar(100)"`
}

func GetTags(skip int, limit int, maps interface{}) (tags []Tag, err *exception.Exception) {
	query := db.Where(maps).Offset(skip)
	if limit > 0 {
		query = query.Limit(limit)
	}
	dbState := query.Find(&tags)
	if dbState.Error != nil {
		err = &exception.QueryDatabaseError
		log.Println(dbState.Error)
	}
	return tags, err
}

func GetTagTotal(maps interface{}) (count int, err *exception.Exception) {
	dbState := db.Model(&Tag{}).Where(maps).Count(&count)
	if dbState.Error != nil {
		err = &exception.QueryDatabaseError
		log.Println(dbState.Error)
	}
	return count, err
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0
}

func AddTag(requestTag RawTag) (err *exception.Exception) {
	dbState := db.Create(&Tag{
		Name: requestTag.Name,
	})
	if dbState.Error != nil {
		err = &exception.QueryDatabaseError
		log.Println(dbState.Error)
	}
	return err
}
func EditTag(id int, requestTag RawTag) (err *exception.Exception) {
	// Update with struct only works with none zero values, or use map[string]interface{}
	dbState := db.Model(&Tag{}).Where("id = ?", id).Update(&Tag{
		Name: requestTag.Name,
	})
	if dbState.Error != nil {
		err = &exception.QueryDatabaseError
		log.Println(dbState.Error)
	}
	return err
}

func DeleteTag(id int) (err *exception.Exception) {
	dbState := db.Where("id = ?", id).Delete(&Tag{})

	if dbState.Error != nil {
		err = &exception.QueryDatabaseError
		log.Println(dbState.Error)
	}
	return err
}
