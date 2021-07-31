package signy

type User struct {
	Id       int    `gorm:"primaryKey"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
