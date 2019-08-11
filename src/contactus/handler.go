package contactus

import (
	"net/http"
	"text/template"

	ss "github.com/ciplax/hacktiv8/finalproject/src/sessions"
)

const (
	templateHTMLPath = "files/www/html/contact/"
)

//Data stores html data
type Data struct {
	Title string
	Sess  ss.Sessions
	Error string
	Ok    string
}

var data Data

func (mdle *Module) handlerContactUsRender(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templateHTMLPath + "contactus.html")
	checkErr(err, "Failed to render contactus.html")
	data.Title = "Contact Us"
	data.Sess = ss.Sess

	err = tmpl.Execute(w, data)
	checkErr(err, "Failed to execute login.html")
}

func (mdle *Module) handlerContactUsPOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cuF := ContactUsForm{
		Fullname: r.Form["fullname"],
		Email:    r.Form["email"],
		Phone:    r.Form["phone"],
		Subject:  r.Form["subject"],
		Message:  r.Form["message"],
	}

	cu, err := cuF.convertToContactUs()
	checkErr(err)

	if err != nil {
		data.Error = err.Error()
		mdle.handlerContactUsRender(w, r)
	} else {
		mdle.insertContactUs(cu)
		data.Error = ""
		data.Ok = "Form submitted succesfully"
		http.Redirect(w, r, "/contactus", http.StatusSeeOther)
	}
}
