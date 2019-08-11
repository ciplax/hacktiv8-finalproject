package users

import (
	"github.com/ciplax/hacktiv8/finalproject/src/articles"
)

func (mdle *Module) getAllPublishedArticles() (ret []articles.Article) {
	rows, err := mdle.queries.GetAllPublishedArticles.Query()
	checkErr(err, "Error running query GetAllPublishedArticles")
	ret = []articles.Article{}
	for rows.Next() {
		a := articles.Article{}
		err := rows.Scan(&a.ArticleID, &a.Title, &a.Content, &a.IsPublished)
		if err == nil {
			ret = append(ret, a)
		}
	}
	return
}
