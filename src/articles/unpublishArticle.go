package articles

import (
	"fmt"
	"strconv"
)

func (mdle *Module) unpublishArticle(text string) string {
	id, err := strconv.ParseInt(text, 10, 0)
	checkErr(err, "Failed to parse int on unpublishArticle")
	res, err := mdle.queries.UnpublishArticle.Exec(id)
	checkErr(err, "Error running query unpublishArticle")

	if i, _ := res.RowsAffected(); i == int64(0) {
		return "No article unpublished"
	}

	return fmt.Sprintf("Article with id%v unpublished", id)
}
