package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		log.Println(email)

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
