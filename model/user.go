package model

type User struct {
	*BaseModel
	UserName   string `gorm:"not null;unique" json:"username" name:"用户名"`
	PassWord   string `gorm:"not null" json:"password" name:"密码"`
	Age        uint   ` json:"age" name:"年龄"`
	Tel        string `gorm:"not null"  json:"tel" name:"手机号"`
	Addr       string `gorm:"not null"  json:"addr" name:"地址"`
	Token      string `gorm:"type:text" json:"token"`
	ExpireTime int64  `json:"expire_time" name:"过期时间"`
}
