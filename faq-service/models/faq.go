package models

// FAQ represents the structure of a frequently asked question entity
type FAQ struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
