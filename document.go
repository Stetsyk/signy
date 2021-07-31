package signy

type Document struct {
	Id      uint `gorm:"primaryKey"`
	Title   string
	Status  string // "open", "in_review", "rejected", "approved"
	OwnerId int
}

type Signature struct {
	Id         uint `gorm:"primaryKey"`
	UserId     int
	DocumentId int
	Status     string // "approved" or "unsigned", "rejected"
}
