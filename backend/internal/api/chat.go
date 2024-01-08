package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GPTChatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// extract the client messages
		messages := []GPTMessage{}
		if err := json.NewDecoder(r.Body).Decode(&messages); err != nil {
			RespondWithError(w, http.StatusBadRequest, fmt.Sprintln(err))
			return
		}

		// prepend the system prompt
		prompt, err := os.ReadFile("txt/prompt.txt")
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintln(err))
			return
		}

		messages = append([]GPTMessage{{Role: "system", Content: string(prompt)}}, messages...)

		// prepare the payload and request
		payload, err := json.Marshal(GPTRequest{
			Messages: messages,
			Model:    "gpt-3.5-turbo",
		})
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintln(err))
			return
		}

		request, err := http.NewRequest(
			http.MethodPost,
			os.Getenv("GPT_API_URL"),
			bytes.NewBuffer(payload),
		)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintln(err))
			return
		}

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GPT_API_KEY")))

		// send request to GPT api
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintln(err))
			return
		}

		// extract GPT response and return to client
		defer res.Body.Close()
		response := GPTResponse{}
		if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintln(err))
			return
		}

		RespondWithJSON(w, http.StatusOK, GPTMessage{
			Role:    response.Choices[len(response.Choices)-1].Message.Role,
			Content: response.Choices[len(response.Choices)-1].Message.Content,
		})

	default:
		http.Error(
			w,
			fmt.Sprintf("%v not allowed", r.Method),
			http.StatusMethodNotAllowed,
		)
	}
}
