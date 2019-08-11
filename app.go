package main

import (
	"log"
	"net/http"

	"github.com/ciplax/hacktiv8/finalproject/src/articles"
	"github.com/ciplax/hacktiv8/finalproject/src/config"
	"github.com/ciplax/hacktiv8/finalproject/src/contactus"
	"github.com/ciplax/hacktiv8/finalproject/src/httperror"
	"github.com/ciplax/hacktiv8/finalproject/src/users"
)

var conf *config.Config

func init() {
	conf = config.ReadConfig()
}

func main() {
	log.Println("Serving Webserver http://localhost" + conf.Server.Port)

	mdlErrors := httperror.NewModule(conf)
	httperror.RegisterRouters(mdlErrors)

	mdlUsers := users.NewModule(conf)
	users.RegisterRouters(mdlUsers)

	mdlArticles := articles.NewModule(conf)
	articles.RegisterRouters(mdlArticles)

	mdleContactUs := contactus.NewModule(conf)
	contactus.RegisterRouters(mdleContactUs)

	err := http.ListenAndServe(conf.Server.Port, nil)
	if err != nil {
		log.Fatal("Error serving http")
	}
}
