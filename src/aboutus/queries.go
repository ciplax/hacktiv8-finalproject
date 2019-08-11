package aboutus

import "database/sql"

//Queries store queries
type Queries struct {
}

func prepare(query string, db *sql.DB) *sql.Stmt {
	s, err := db.Prepare(query)
	checkErr(err, "Failed to prepare statement")

	return s
}

//NewQueries returns prepared queries
func NewQueries(dbMaster, dbSlave *sql.DB) *Queries {
	q := &Queries{}

	return q
}

const ()
