package signy

type Document struct {
	Title   string
	Status  string // "open", "in_review", "rejected", "approved"
	OwnerId int
}

type Signature struct {
	UserId     int
	DocumentId int
	Status     string // "approved" or "unsigned", "rejected"
}
