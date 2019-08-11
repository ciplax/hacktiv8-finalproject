package httperror

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
}

func (mdle *Module) handler401Render(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateHTMLPath + "401.html")
	checkErr(err, "Failed to parse 401.html")

	data := dataHTML{
		Title: "401",
		Sess:  ss.Sess,
	}

	err = tmpl.Execute(w, data)
	checkErr(err, "Failed to execute 401.html")

}
