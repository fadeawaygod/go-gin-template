package model

import "time"

type Tag struct {
	TagId        int       `gorm:"primary_key"`
	Name         string    `gorm:"column:Name"`
	IsEnabled    bool      `gorm:"column:IsEnabled"`
	InsertedTime time.Time `gorm:"column:InsertedTime"`
	UpdatedTime  time.Time `gorm:"column:UpdatedTime"`
}

func GetTags(skip int, limit int, maps interface{}) (tags []Tag) {
	query := db.Where(maps).Offset(skip)
	if limit > 0 {
		query = query.Limit(limit)
	}
	query.Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.TagId > 0
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name: name,
	})

	return true
}
