package signy

type Document struct {
	Id      uint   `gorm:"primaryKey"`
	Title   string `json:"title"`
	Status  string `json:"status"` // "open", "in_review", "rejected", "approved"
	OwnerId int    `json:'owner`
}

type Signature struct {
	Id         uint `gorm:"primaryKey"`
	UserId     int
	DocumentId int
	Status     string // "approved" or "unsigned", "rejected"
}
