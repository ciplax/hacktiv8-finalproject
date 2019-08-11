package articles

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	ss "github.com/ciplax/hacktiv8/finalproject/src/sessions"
)

const (
	templateHTMLPath = "files/www/html/"
)

type dataHTML struct {
	Title  string
	Values []Article
	Sess   ss.Sessions
	Ok     string
	Error  string
}

var data dataHTML

func (mdle *Module) handlerArticlesRender(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	tmpl, err := template.ParseFiles(templateHTMLPath+"article/articles.html", templateHTMLPath+"header.html", templateHTMLPath+"navbar.html")
	checkErr(err, "Failed to parse articles.html")

	art := mdle.getAllArticles()

	data.Title = "Articles"
	data.Values = art
	data.Sess = ss.Sess

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

	data.Ok = "Article deleted succesfully"
	mdle.handlerArticlesRender(w, r)
	// http.Redirect(w, r, "/articles", http.StatusSeeOther)
	data.Ok = ""
}

func (mdle *Module) handlerPublishArticle(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	r.ParseForm()
	id := strings.Join(r.Form["article_id"], "")
	log.Println(mdle.publishArticle(id))

	data.Ok = "Article published succesfully"
	mdle.handlerArticlesRender(w, r)
	// http.Redirect(w, r, "/articles", http.StatusSeeOther)
	data.Ok = ""
}

func (mdle *Module) handlerUnpublishArticle(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	r.ParseForm()
	id := strings.Join(r.Form["article_id"], "")
	log.Println(mdle.unpublishArticle(id))

	data.Ok = "Article unpublished succesfully"
	mdle.handlerArticlesRender(w, r)
	// http.Redirect(w, r, "/articles", http.StatusSeeOther)
	data.Ok = ""
}

func (mdle *Module) handlerNewArticleRender(w http.ResponseWriter, r *http.Request) {
	if ss.Sess.SessionID == 0 {
		http.Redirect(w, r, "/401", http.StatusSeeOther)
	}

	tmpl, err := template.ParseFiles(templateHTMLPath+"article/articles_new.html", templateHTMLPath+"header.html", templateHTMLPath+"navbar.html")
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
		data.Error = err.Error()
		mdle.handlerNewArticleRender(w, r)
		data.Error = ""
	} else {
		mdle.insertArticle(art)
		http.Redirect(w, r, "/articles", http.StatusSeeOther)
	}

}
