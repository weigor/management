package model

type Live struct {
	*BaseModel
	Username string `json:"username"`
	photo  string `json:"photo"`
	head string `json:"head"`
}
