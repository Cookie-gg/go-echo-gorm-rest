package model

import "database/sql"

type Profile struct {
	ID       uint           `gorm:"primarykey" json:"id" param:"id"`
	Gender   string         `json:"gender"`
	Age      uint           `json:"age"`
	Twitter  sql.NullString `json:"twitter" swaggertype:"string"`  // nullable
	Facebook sql.NullString `json:"facebook" swaggertype:"string"` // nullable
	UserID   uint           `json:"user_id"`
}
