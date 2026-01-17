package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	githubEnums "discus.TelegramAlert/enum"
)

func processGitHub(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	eventType := githubEnums.GitHubEvent(r.Header.Get("x-github-event"))

	if eventType == githubEnums.DISCUSSION {
		//telegram.SendMessage(fmt.Sprintf(LANG_NEW_DISCUSSION))
	}

	if eventType == githubEnums.DISCUSSION_COMMENT {

	}

	var payload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Lower is test

	// 2. Pretty-print the JSON to the console
	prettyJSON, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println("eventType:", eventType)
	fmt.Printf("Received Webhook:\n%s\n", string(prettyJSON))

	// 3. Respond to GitHub
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Register the handler function for the "/items" path
	http.HandleFunc("/webhook/github", processGitHub)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	fmt.Printf("Server listening on port %s...\n", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil); err != nil {
		log.Fatal(err)
	}
}
