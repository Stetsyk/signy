package signy

type Document struct {
	Id      uint   `gorm:"primaryKey"`
	Title   string `json:"title"`
	Status  string `json:"status"` // "open", "in_review", "rejected", "approved"
	OwnerId int    `json:"owner_id"`
}

type Signature struct {
	Id         uint   `gorm:"primaryKey"`
	UserId     int    `json:"userId"`
	DocumentId int    `json:"documentId"`
	Status     string `json:"status"` // "approved" or "unsigned", "rejected"
}
