package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func SendMessage(messageText string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", os.Getenv("TELEGRAM_BOT_TOKEN"))

	data := url.Values{}
	data.Set("chat_id", os.Getenv("TELEGRAM_USERID"))
	data.Set("text", messageText)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Telegram API error: %s (Status: %d)", string(bodyBytes), resp.StatusCode)
	}

	return nil
}
