package model

type User struct {
	*BaseModel
	UserName string `gorm:"not null" json:"username" name:"用户名"`
	PassWord string `gorm:"not null" json:"password" name:"密码"`
	Age      uint   ` json:"age" name:"年龄"`
	Tel      string `gorm:"not null"  json:"tel" name:"手机号"`
	Addr     string `gorm:"not null"  json:"addr" name:"地址"`
	Card     string `gorm:"not null;index:idx_no;unique"  json:"card" name:"身份证"`
}
