package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	githubEnums "github.com/FoerMaster/GitHubDiscussionTelegramAlert/enum"
	"github.com/FoerMaster/GitHubDiscussionTelegramAlert/models"
	"github.com/FoerMaster/GitHubDiscussionTelegramAlert/telegram"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func escapeTelegramMarkdownV2(text string) string {
	specialChars := []string{
		"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!",
	}

	result := text
	for _, char := range specialChars {
		result = strings.ReplaceAll(result, char, "\\"+char)
	}

	return result
}

func validateGitHubSignature(signature string, body []byte) bool {
	if signature == "" {
		return false
	}

	if !strings.HasPrefix(signature, "sha256=") {
		return false
	}
	signatureHash := strings.TrimPrefix(signature, "sha256=")

	mac := hmac.New(sha256.New, []byte(os.Getenv("GITHUB_WEBHOOK_SECRET")))
	mac.Write(body)
	expectedHash := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signatureHash), []byte(expectedHash))
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

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Check signature for security (Secret in GitHub webhook settings)
	signature := r.Header.Get("X-Hub-Signature-256")
	if !validateGitHubSignature(signature, bodyBytes) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	eventType := githubEnums.GitHubEvent(r.Header.Get("x-github-event"))

	var body models.GitHubWebhook
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	switch eventType {
	case githubEnums.DISCUSSION:
		switch githubEnums.GitHubAction(body.Action) {
		case githubEnums.CREATED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_NEW_DISCUSSION,
				escapeTelegramMarkdownV2(body.Discussion.User.Login),
				escapeTelegramMarkdownV2(body.Discussion.Title),
				body.Discussion.HTMLURL,
				escapeTelegramMarkdownV2(body.Discussion.Body)))
			if err != nil {
				http.Error(w, "Failed send telegram message", http.StatusBadRequest)
				return
			}
		case githubEnums.EDITED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_EDITED_DISCUSSION,
				escapeTelegramMarkdownV2(body.Discussion.User.Login),
				escapeTelegramMarkdownV2(body.Discussion.Title),
				body.Discussion.HTMLURL,
				escapeTelegramMarkdownV2(body.Discussion.Body)))
			if err != nil {
				http.Error(w, "Failed send telegram message", http.StatusBadRequest)
				return
			}
		case githubEnums.DELETED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_DELETED_DISCUSSION,
				escapeTelegramMarkdownV2(body.Discussion.User.Login),
				escapeTelegramMarkdownV2(body.Discussion.Title),
				body.Discussion.HTMLURL,
				escapeTelegramMarkdownV2(body.Discussion.Body)))
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
				escapeTelegramMarkdownV2(body.Comment.User.Login),
				escapeTelegramMarkdownV2(body.Discussion.Title),
				body.Comment.HTMLURL,
				escapeTelegramMarkdownV2(body.Comment.Body)))
			if err != nil {
				http.Error(w, "Failed send telegram message", http.StatusBadRequest)
				return
			}
		case githubEnums.DELETED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_DELETED_DISCUSSION_COMMENT,
				escapeTelegramMarkdownV2(body.Comment.User.Login),
				escapeTelegramMarkdownV2(body.Discussion.Title),
				body.Discussion.HTMLURL,
				escapeTelegramMarkdownV2(body.Comment.Body)))
			if err != nil {
				http.Error(w, "Failed send telegram message", http.StatusBadRequest)
				return
			}
		case githubEnums.EDITED:
			err := telegram.SendMessage(fmt.Sprintf(
				LANG_EDITED_DISCUSSION_COMMENT,
				escapeTelegramMarkdownV2(body.Comment.User.Login),
				escapeTelegramMarkdownV2(body.Discussion.Title),
				body.Comment.HTMLURL,
				escapeTelegramMarkdownV2(body.Comment.Body)))
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
