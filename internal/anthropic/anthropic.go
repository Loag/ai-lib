package anthropic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type AnthropicCompletionRequest struct {
	Model       string                    `json:"model"`
	Temperature float32                   `json:"temperature"`
	TopP        float32                   `json:"top_p"`
	MaxTokens   int                       `json:"max_tokens"`
	System      string                    `json:"system"`
	Messages    []AnthropicRequestMessage `json:"messages"`
}

type AnthropicRequestMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type AnthropicResponseMessage struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type AnthropicCompletionResponse struct {
	Id           string                     `json:"id"`
	Content      []AnthropicResponseMessage `json:"content"`
	Model        string                     `json:"model"`
	Role         string                     `json:"role"`
	StopReason   string                     `json:"stop_reason"`
	StopSequence string                     `json:"stop_sequence"`
	Type         string                     `json:"type"`
	Usage        struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
}

type Anthropic struct {
	apiKey      string
	baseURL     string
	Model       string
	Temperature float32
	TopP        float32
	MaxTokens   int
	System      string
}

func NewAnthropic(apiKey, baseURL string) *Anthropic {
	return &Anthropic{
		apiKey:      apiKey,
		baseURL:     baseURL,
		Model:       "claude-3-5-sonnet-20241022",
		Temperature: 0.7,
		TopP:        1,
		MaxTokens:   2480,
	}
}

func (a *Anthropic) SetModel(model string) error {
	if model == "" {
		return errors.New("model is required")
	}
	a.Model = model
	return nil
}

func (a *Anthropic) SetTemperature(temperature float32) error {
	if temperature < 0 || temperature > 1 {
		return errors.New("temperature must be between 0 and 1")
	}
	a.Temperature = temperature
	return nil
}

func (a *Anthropic) SetTopP(topP float32) error {
	if topP < 0 || topP > 1 {
		return errors.New("topP must be between 0 and 1")
	}
	a.TopP = topP
	return nil
}

func (a *Anthropic) SetMaxTokens(maxTokens int) error {
	if maxTokens < 1 {
		return errors.New("max tokens must be more than 1")
	}
	a.MaxTokens = maxTokens
	return nil
}

func (a *Anthropic) SetSystemPrompt(system string) error {
	if system == "" {
		return errors.New("system is required")
	}
	a.System = system
	return nil
}

func (a *Anthropic) getCompletion(prompt string) (AnthropicCompletionResponse, error) {
	request := AnthropicCompletionRequest{
		Model:       a.Model,
		Temperature: a.Temperature,
		TopP:        a.TopP,
		MaxTokens:   a.MaxTokens,
		System:      a.System,
		Messages: []AnthropicRequestMessage{
			{
				Content: prompt,
				Role:    "user",
			},
		},
	}

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		log.Err(err).Msg("Error marshalling request")
		return AnthropicCompletionResponse{}, err
	}

	req, err := http.NewRequest("POST", a.baseURL, bytes.NewBuffer(jsonRequest))
	if err != nil {
		log.Err(err).Msg("Error creating request")
		return AnthropicCompletionResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", a.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("Error making request")
		return AnthropicCompletionResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("Error reading response")
		return AnthropicCompletionResponse{}, err
	}

	var response AnthropicCompletionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Err(err).Msg("Error unmarshalling response")
		return AnthropicCompletionResponse{}, err
	}

	return response, nil
}

func (a *Anthropic) GetCompletion(prompt string) (string, error) {
	response, err := a.getCompletion(prompt)
	if err != nil {
		return "", err
	}
	return response.Content[0].Text, nil
}
