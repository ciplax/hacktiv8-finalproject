package articles

import "database/sql"

//Queries store queries
type Queries struct {
	GetAllArticles   *sql.Stmt
	DeleteArticle    *sql.Stmt
	PublishArticle   *sql.Stmt
	UnpublishArticle *sql.Stmt
	InsertArticle    *sql.Stmt
	GetArticleByID   *sql.Stmt
	UpdateArticle    *sql.Stmt
}

func prepare(query string, db *sql.DB) *sql.Stmt {
	s, err := db.Prepare(query)
	checkErr(err, "Failed to prepare statement")

	return s
}

//NewQueries returns prepared queries
func NewQueries(dbMaster, dbSlave *sql.DB) *Queries {
	q := &Queries{
		GetAllArticles:   prepare(cGetAllArticles, dbMaster),
		DeleteArticle:    prepare(cDeleteArticle, dbMaster),
		PublishArticle:   prepare(cPublishArticle, dbMaster),
		UnpublishArticle: prepare(cUnpublishArticle, dbMaster),
		InsertArticle:    prepare(cInsertArticle, dbMaster),
		GetArticleByID:   prepare(cGetArticleByID, dbMaster),
		UpdateArticle:    prepare(cUpdateArticle, dbMaster),
	}

	return q
}

const (
	cGetAllArticles = `
	SELECT 
		article_id
		, title
		, content
		, flag
	FROM
		articles;`
	cGetArticleByID = `
	SELECT 
		article_id
		, title
		, content
		, flag
	FROM
		articles
	WHERE
		article_id = ?;`
	cUpdateArticle = `
	UPDATE
		articles
	SET 
		title = ?
		, content = ?
		, flag = ?
	WHERE
		article_id = ?;`
	cDeleteArticle = `
	DELETE
	FROM
			articles
	WHERE
		article_id = ?;`
	cPublishArticle = `
	UPDATE
		articles
	SET
		flag = 1
	WHERE
	article_id = ?;`
	cUnpublishArticle = `
	UPDATE
		articles
	SET
		flag = 0
	WHERE
	article_id = ?;`
	cInsertArticle = `
	INSERT INTO
		articles
	VALUES (
		?, ?, ?, ?
	);`
)
