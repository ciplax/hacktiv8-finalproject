package articles

import (
	"fmt"
	"strconv"
)

func (mdle *Module) deleteArticle(text string) string {
	id, err := strconv.ParseInt(text, 10, 0)
	checkErr(err, "Failed to parse int on deleteArticle")
	res, err := mdle.queries.DeleteArticle.Exec(id)
	checkErr(err, "Error running query deleteArticle")

	if i, _ := res.RowsAffected(); i == int64(0) {
		return "No article deleted"
	}

	return fmt.Sprintf("Article with id%v deleted", id)
}
