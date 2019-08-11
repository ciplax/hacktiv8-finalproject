package contactus

import (
	"errors"
	"strings"
)

//ContactUs store contactUs
type ContactUs struct {
	ContactUsID string `json:"contactus_id" db:"contactus_id"`
	Fullname    string `json:"fullname" db:"fullname"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	Subject     string `json:"subject" db:"subject"`
	Message     string `json:"message" db:"message"`
}

//ContactUsForm stores contactUsForm
type ContactUsForm struct {
	Fullname []string
	Email    []string
	Phone    []string
	Subject  []string
	Message  []string
}

func (cuF *ContactUsForm) convertToContactUs() (ContactUs, error) {
	cu := ContactUs{}
	cu.ContactUsID = "0"

	if len(cuF.Fullname) == 0 || strings.Join(cuF.Fullname, "") == "" {
		return ContactUs{}, errors.New("Fullname cannot be empty")
	}
	cu.Fullname = strings.Join(cuF.Fullname, "")

	if len(cuF.Email) == 0 || strings.Join(cuF.Email, "") == "" {
		return ContactUs{}, errors.New("Email cannot be empty")
	}
	cu.Email = strings.Join(cuF.Email, "")

	if len(cuF.Phone) == 0 || strings.Join(cuF.Phone, "") == "" {
		return ContactUs{}, errors.New("Phone cannot be empty")
	}
	cu.Phone = strings.Join(cuF.Phone, "")

	if len(cuF.Subject) == 0 || strings.Join(cuF.Subject, "") == "" {
		return ContactUs{}, errors.New("Subject cannot be empty")
	}
	cu.Subject = strings.Join(cuF.Subject, "")

	if len(cuF.Message) == 0 || strings.Join(cuF.Message, "") == "" {
		return ContactUs{}, errors.New("Message cannot be empty")
	}
	cu.Message = strings.Join(cuF.Message, "")

	return cu, nil
}
