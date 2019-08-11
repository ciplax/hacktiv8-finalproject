package users

import "database/sql"

//Queries store queries
type Queries struct {
	ValidateUser *sql.Stmt
}

func prepare(query string, db *sql.DB) *sql.Stmt {
	s, err := db.Prepare(query)
	checkErr(err, "Failed to prepare statement")

	return s
}

//NewQueries returns prepared queries
func NewQueries(dbMaster, dbSlave *sql.DB) *Queries {
	q := &Queries{
		ValidateUser: prepare(cValidateUser, dbMaster),
	}

	return q
}

const (
	cValidateUser = `
		SELECT 
		user_id
		, username
		, password
		, name
		, birthdate
		, address
		, gender 
		, role_id
		, last_login 
	FROM 
		users
	WHERE
		username = ?
		AND password = ?;`
)
