package articles

import (
	"strconv"
)

//GetAllArticles get all articles
func (mdle *Module) getArticleByID(text string) (ret Article) {
	id, err := strconv.ParseInt(text, 10, 0)
	checkErr(err, "Failed to parse int on getArticleByID")

	rows, err := mdle.queries.GetArticleByID.Query(id)
	checkErr(err, "Error running query getArticleByID")

	for rows.Next() {
		a := Article{}
		err := rows.Scan(&a.ArticleID, &a.Title, &a.Content, &a.IsPublished)
		if err == nil {
			return a
		}
	}
	return Article{}
}
