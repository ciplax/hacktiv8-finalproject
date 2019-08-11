package contactus

import (
	"fmt"
)

func (mdle *Module) insertContactUs(cu ContactUs) string {
	res, err := mdle.queries.InsertContactUs.Exec(cu.ContactUsID, cu.Fullname, cu.Email, cu.Phone, cu.Subject, cu.Message)
	checkErr(err, "Error running query InsertContactUs")

	if i, _ := res.RowsAffected(); i == int64(0) {
		return "No ContactUs inserted"
	}

	return fmt.Sprintf("ContactUs inserted")
}
