package articles

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ciplax/hacktiv8/finalproject/src/config"

	//ok
	_ "github.com/go-sql-driver/mysql"
)

//Module stores module
type Module struct {
	Conf     *config.Config
	queries  *Queries
	DBMaster *sql.DB
	DBSlave  *sql.DB
}

//NewModule initiate NewModule
func NewModule(c *config.Config) *Module {
	m := &Module{
		Conf: c,
	}

	dbMaster, err := sql.Open("mysql", c.Database.DBMaster)
	checkErr(err)

	if dbMaster == nil {
		log.Println("Error connecting to DBMaster")
		return nil
	}

	m.DBMaster = dbMaster

	dbSlave, err := sql.Open("mysql", c.Database.DBSlave)
	checkErr(err)

	if dbSlave == nil {
		log.Println("Error connecting to DBSlave")
		return nil
	}

	m.DBSlave = dbSlave

	m.queries = NewQueries(m.DBMaster, m.DBSlave)

	return m
}

//RegisterRouters register routers
func RegisterRouters(mdle *Module) {
	http.HandleFunc("/articles", mdle.handlerArticlesRender)
	http.HandleFunc("/articles/delete", mdle.handlerDeleteArticle)
	http.HandleFunc("/articles/publish", mdle.handlerPublishArticle)
	http.HandleFunc("/articles/unpublish", mdle.handlerUnpublishArticle)
	http.HandleFunc("/articles/new", mdle.handlerNewArticleRender)
	http.HandleFunc("/articles/add", mdle.handlerAddArticle)
}

func checkErr(err error, msg ...string) {
	if err != nil {
		log.Println(msg, "\n", err)
	}
}
