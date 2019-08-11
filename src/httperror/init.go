package httperror

import (
	"log"
	"net/http"

	"github.com/ciplax/hacktiv8/finalproject/src/config"

	//ok
	_ "github.com/go-sql-driver/mysql"
)

//Module stores module
type Module struct {
}

//NewModule initiate NewModule
func NewModule(c *config.Config) *Module {
	m := &Module{}

	return m
}

//RegisterRouters register routers
func RegisterRouters(mdle *Module) {
	http.HandleFunc("/401", mdle.handler401Render)
}

func checkErr(err error, msg ...string) {
	if err != nil {
		log.Panicln(msg, "\n", err)
	}
}
