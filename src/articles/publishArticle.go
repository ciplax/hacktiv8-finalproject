package articles

import (
	"fmt"
	"strconv"
)

func (mdle *Module) publishArticle(text string) string {
	id, err := strconv.ParseInt(text, 10, 0)
	checkErr(err, "Failed to parse int on publishArticle")
	res, err := mdle.queries.PublishArticle.Exec(id)
	checkErr(err, "Error running query publishArticle")

	if i, _ := res.RowsAffected(); i == int64(0) {
		return "No article published"
	}

	return fmt.Sprintf("Article with id%v published", id)
}
