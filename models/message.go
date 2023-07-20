package models

type Message struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id" gorm:"many2many, foreignKey:ID"`
	Body   string `json:"body" gorm:"type: text"`
}
