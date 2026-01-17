package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	githubEnums "discus.TelegramAlert/enum"
	"discus.TelegramAlert/models"
	"discus.TelegramAlert/telegram"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func processHealth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(Response{
		Status:  "ok",
		Message: "Success!",
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func processGitHub(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	eventType := githubEnums.GitHubEvent(r.Header.Get("x-github-event"))

	var body models.GitHubWebhook
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	switch eventType {
	case githubEnums.DISCUSSION:
		switch githubEnums.GitHubAction(body.Action) {
		case githubEnums.CREATED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_NEW_DISCUSSION,
				body.Discussion.User.Login,
				body.Discussion.Title,
				body.Discussion.HTMLURL,
				body.Discussion.Body))
			if err != nil {
				http.Error(w, "Failed send telegram message", http.StatusBadRequest)
				return
			}
		case githubEnums.EDITED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_EDITED_DISCUSSION,
				body.Discussion.User.Login,
				body.Discussion.Title,
				body.Discussion.HTMLURL,
				body.Discussion.Body))
			if err != nil {
				http.Error(w, "Failed send telegram message", http.StatusBadRequest)
				return
			}
		case githubEnums.DELETED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_DELETED_DISCUSSION,
				body.Discussion.User.Login,
				body.Discussion.Title,
				body.Discussion.HTMLURL,
				body.Discussion.Body))
			if err != nil {
				http.Error(w, "Failed send telegram message", http.StatusBadRequest)
				return
			}
		}
	case githubEnums.DISCUSSION_COMMENT:
		switch githubEnums.GitHubAction(body.Action) {
		case githubEnums.CREATED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_NEW_DISCUSSION_COMMENT,
				body.Comment.User.Login,
				body.Discussion.Title,
				body.Comment.HTMLURL,
				body.Comment.Body))
			if err != nil {
				http.Error(w, "Failed send telegram message", http.StatusBadRequest)
				return
			}
		case githubEnums.DELETED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_DELETED_DISCUSSION_COMMENT,
				body.Comment.User.Login,
				body.Discussion.Title,
				body.Comment.HTMLURL,
				body.Comment.Body))
			if err != nil {
				http.Error(w, "Failed send telegram message", http.StatusBadRequest)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook/github", processGitHub)
	http.HandleFunc("/health", processHealth)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	fmt.Printf("Server listening on port %s...\n", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil); err != nil {
		log.Fatal(err)
	}
}
