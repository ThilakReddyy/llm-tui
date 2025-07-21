package helpers

import (
	"encoding/json"
	"os"
	"wizard-tutorial/internal/types"
)

func SaveMessagesToJSON(filename string, messages []types.Message) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(messages)
}

func SaveHistoryToJSON(filename string, messages []types.ChatHistory) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(messages)
}

func GetHistoryFromJSON(filename string) ([]types.ChatHistory, error) {
	var history []types.ChatHistory

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&history)

	return history, err
}

func GetMessagesFromJSON(filename string) ([]types.Message, error) {
	var messages []types.Message

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&messages)

	return messages, err
}
