package articles

import (
	"fmt"
)

func (mdle *Module) insertArticle(art Article) string {
	res, err := mdle.queries.InsertArticle.Exec(art.ArticleID, art.Title, art.Content, art.IsPublished)
	checkErr(err, "Error running query insertArticle")

	if i, _ := res.RowsAffected(); i == int64(0) {
		return "No article inserted"
	}

	return fmt.Sprintf("Article inserted")
}
