package aboutus

import (
	"net/http"
	"text/template"

	ss "github.com/ciplax/hacktiv8/finalproject/src/sessions"
)

const (
	templateHTMLPath = "files/www/html/"
)

type dataHTML struct {
	Title string
	Sess  ss.Sessions
	Ok    string
	Error string
}

var data dataHTML

func (mdle *Module) handlerAboutUsRender(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateHTMLPath+"about/about.html", templateHTMLPath+"header.html", templateHTMLPath+"navbar.html")
	checkErr(err, "Failed to parse about.html")

	err = tmpl.Execute(w, data)
	checkErr(err, "Failed to execute about.html")

}
