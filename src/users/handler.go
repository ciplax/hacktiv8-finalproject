package users

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/ciplax/hacktiv8/finalproject/src/articles"
	ss "github.com/ciplax/hacktiv8/finalproject/src/sessions"
)

const (
	templateHTMLPath = "files/www/html/user/"
)

//Data stores html data
type Data struct {
	Title    string
	Username string
	Message  string
	Art      []articles.Article
	Sess     ss.Sessions
}

var data Data

func (mdle *Module) handlerLoginRender(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateHTMLPath + "login.html")
	checkErr(err, "Failed to render login.html")

	data.Title = "Login"

	err = tmpl.Execute(w, data)
	checkErr(err, "Failed to execute login.html")
}

func (mdle *Module) handlerLogout(w http.ResponseWriter, r *http.Request) {
	data = Data{}
	ss.Sess.ClearSessions()
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (mdle *Module) handlerLoginDo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		u := User{
			Username: strings.Join(r.Form["username"], ""),
			Password: strings.Join(r.Form["password"], ""),
		}

		u = mdle.validateUser(u)
		if u == (User{}) {
			data.Message = "Username and password doesnt match"
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			ss.Sess.GenerateSessions()
			ss.Sess.InsertUsernameToSessions(u.Username)
			mdle.handlerHomeRender(w, r)
		}
	}
}

func (mdle *Module) handlerHomeRender(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateHTMLPath + "home.html")
	checkErr(err, "Failed to render home.html")

	data.Art = mdle.getAllPublishedArticles()
	data.Title = "Home"
	data.Sess = ss.Sess

	err = tmpl.Execute(w, data)
	checkErr(err, "Failed to execute home.html")
}
