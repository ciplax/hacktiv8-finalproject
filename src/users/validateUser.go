package users

func (mdle *Module) validateUser(u User) User {
	rows, err := mdle.queries.ValidateUser.Query(u.Username, u.Password)
	checkErr(err, "Error running query validateUser")
	u = User{}
	for rows.Next() {
		rows.Scan(&u.UserID, &u.Username, &u.Password, &u.Name, &u.BirthDate, &u.Address, &u.Gender, &u.RoleID, &u.LastLogin)
	}
	return u
}
