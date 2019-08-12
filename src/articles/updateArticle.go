package articles

import (
	"fmt"
)

func (mdle *Module) updateArticle(art Article) string {
	res, err := mdle.queries.UpdateArticle.Exec(art.Title, art.Content, art.IsPublished, art.ArticleID)
	checkErr(err, "Error running query updateArticle")

	if i, _ := res.RowsAffected(); i == int64(0) {
		return "No article updated"
	}

	return fmt.Sprintf("Article updated")
}
