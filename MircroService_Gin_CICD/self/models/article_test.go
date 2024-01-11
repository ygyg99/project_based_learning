package models

import "testing"

func TestArticle(t *testing.T) {
	alsit := contGetAllArticles()
	if len(alsit) != len(ArticleList){
		t.Fail()
	}
	for idx,list := range alsit{
		if list.Content != ArticleList[idx].Content ||
			list.ID != ArticleList[idx].ID ||
			list.Title != ArticleList[idx].Title{
				t.Fail()
		}
	}
}
