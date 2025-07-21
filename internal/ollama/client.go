package ollama

import (
	"bytes"
	"encoding/json"
	"net/http"
	"wizard-tutorial/internal/types"
)

const defaultOllamaURL = "http://localhost:11434/api/chat"

func GetBotResponse(messages []types.Message) string {
	req := types.Request{
		Model:    "llama3.2:1b",
		Stream:   false,
		Messages: messages,
	}
	resp, err := talkToOllama(defaultOllamaURL, req)
	if err != nil {
		return err.Error()
	}
	return resp.Message.Content
}

func talkToOllama(url string, ollamaReq types.Request) (*types.Response, error) {
	js, err := json.Marshal(&ollamaReq)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(js))
	if err != nil {
		return nil, err
	}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	ollamaResp := types.Response{}
	err = json.NewDecoder(httpResp.Body).Decode(&ollamaResp)
	return &ollamaResp, err
}
