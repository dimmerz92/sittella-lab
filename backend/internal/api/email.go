package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
)

func EnquiryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// read the template txt file
		template, err := os.ReadFile("txt/contact_template.txt")
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintln(err))
			return
		}

		// send it to the client
		RespondWithJSON(w, http.StatusOK, ClientResponse{
			Message: string(template),
		})

	case http.MethodPost:
		// extract the client email
		email := ClientEmail{}
		if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
			RespondWithError(w, http.StatusBadRequest, fmt.Sprintln(err))
			return
		}

		// send the email
		if err := sendEmail(email); err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// respond to client
		RespondWithJSON(w, http.StatusOK, ClientResponse{
			Message: "Thank you, your enquiry has been sent",
		})

	default:
		http.Error(
			w,
			fmt.Sprintf("%v not allowed", r.Method),
			http.StatusMethodNotAllowed,
		)
	}
}

func sendEmail(email ClientEmail) error {
	// create an auth object
	auth := smtp.PlainAuth(
		"",
		os.Getenv("ICLOUD_USER"),
		os.Getenv("ICLOUD_PASSWORD"),
		os.Getenv("EMAIL_SMTP"),
	)

	// interpolate client email with email template
	tmpl, err := os.ReadFile("txt/email_template.txt")
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(string(tmpl),
		os.Getenv("EMAIL"),
		os.Getenv("EMAIL"),
		email.FirstName+" "+email.LastName, email.Email,
		email.Message,
	)

	// sent email to sittella and return response to client
	if err := smtp.SendMail(
		os.Getenv("EMAIL_SMTP")+":"+os.Getenv("EMAIL_PORT"),
		auth,
		os.Getenv("EMAIL"),
		[]string{os.Getenv("EMAIL")},
		[]byte(msg),
	); err != nil {
		return err
	}

	return nil
}
