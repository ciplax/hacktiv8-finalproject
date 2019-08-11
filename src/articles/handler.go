package articles

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	ss "github.com/ciplax/hacktiv8/finalproject/src/sessions"
)

const (
	templateHTMLPath = "files/www/html/article/"
)

type dataHTML struct {
	Title  string
	Values []Article
	Sess   ss.Sessions
}

func (mdle *Module) handlerArticlesRender(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	tmpl, err := template.ParseFiles(templateHTMLPath + "articles.html")
	checkErr(err, "Failed to parse articles.html")

	art := mdle.getAllArticles()

	data := dataHTML{
		Title:  "Articles",
		Values: art,
		Sess:   ss.Sess,
	}

	err = tmpl.Execute(w, data)
	checkErr(err, "Failed to execute articles.html")

}

func (mdle *Module) handlerDeleteArticle(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	r.ParseForm()
	id := strings.Join(r.Form["article_id"], "")
	mdle.deleteArticle(id)

	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}

func (mdle *Module) handlerPublishArticle(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	r.ParseForm()
	id := strings.Join(r.Form["article_id"], "")
	log.Println(mdle.publishArticle(id))

	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}

func (mdle *Module) handlerUnpublishArticle(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	r.ParseForm()
	id := strings.Join(r.Form["article_id"], "")
	log.Println(mdle.unpublishArticle(id))

	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}

func (mdle *Module) handlerNewArticleRender(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	tmpl, err := template.ParseFiles(templateHTMLPath + "articles_new.html")
	checkErr(err, "Failed to parse articles_new.html")

	data := dataHTML{
		Title: "New Article",
		Sess:  ss.Sess,
	}

	err = tmpl.Execute(w, data)
	checkErr(err, "Failed to execute articles_new.html")
}

func (mdle *Module) handlerAddArticle(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	r.ParseForm()
	artF := ArticleForm{
		Title:       r.Form["title"],
		Content:     r.Form["content"],
		IsPublished: r.Form["published"],
	}

	art, err := artF.convertToArticle()
	checkErr(err)

	if err != nil {

		mdle.handlerNewArticleRender(w, r)
	} else {

		mdle.insertArticle(art)
		http.Redirect(w, r, "/articles", http.StatusSeeOther)
	}

}
