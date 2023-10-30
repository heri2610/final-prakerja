package models

type User struct {
	Id          int           `json:"id" gorm:"primaryKey"`
	Name        string        `json:"name" gorm:"not null"`
	Email       string        `json:"email" gorm:"unique"`
	CreatedAt   string        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   string        `json:"updated_at" gorm:"autoUpdateTime"`
	SosialMedia []SosialMedia `json:"sosial_media" gorm:"foreignKey:UserID"`
}
