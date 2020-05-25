package model

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(skip int, limit int, maps interface{}) (tags []Tag) {
	query := db.Where(maps).Offset(skip)
	if limit > 0 {
		query = query.Limit(limit)
	}
	query = query.Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
