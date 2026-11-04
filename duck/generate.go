package duck

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const generateAddressUrl = "https://quack.duckduckgo.com/api/email/addresses"

// GenerateAddress generates a new @duck.com email address
func GenerateAddress(token string) (string, error) {
	client := GetClient()

	type Response struct {
		Address string `json:"address"`
		Error   string `json:"error"`
	}

	var response = Response{}

	r, err := client.R().
		SetHeader("authorization", "Bearer "+token).
		SetResult(&response).
		Post(generateAddressUrl)

	if err != nil {
		return "", fmt.Errorf("generateAddress failed to unmarshal response: %s, error: %w", r.String(), err)
	}

	if len(response.Error) > 0 {
		return "", errors.New("generateAddress failed to generate email: " + response.Error)
	}

	alias := response.Address + "@duck.com"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	write := []byte(timestamp + " - " + alias + "\n")

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(string(write))

	return response.Address + "@duck.com", nil
}
