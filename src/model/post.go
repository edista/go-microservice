package model

type Post struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ViewCount int32  `json:"view_count"`
}
