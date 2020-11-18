package model

type Live struct {
	*BaseModel
	Username string `gorm:"not null;"json:"username"`
	Photo    string `json:"photo"`
	Head     string `gorm:"not null;"json:"head"`
}
