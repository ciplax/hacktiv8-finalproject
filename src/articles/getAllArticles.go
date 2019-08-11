package articles

//GetAllArticles get all articles
func (mdle *Module) getAllArticles() (ret []Article) {
	rows, err := mdle.queries.GetAllArticles.Query()
	checkErr(err, "Error running query validateUser")
	ret = []Article{}
	for rows.Next() {
		a := Article{}
		err := rows.Scan(&a.ArticleID, &a.Title, &a.Content, &a.IsPublished)
		if err == nil {
			ret = append(ret, a)
		}
	}
	return
}
