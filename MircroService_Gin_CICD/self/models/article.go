package models

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var ArticleList = []Article{
	{ID: 1, Title: "Atc 1", Content: "111"},
	{ID: 2, Title: "Atc 2", Content: "222"},
}
