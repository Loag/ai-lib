package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Prompt struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type OpenAICompletionRequest struct {
	Model       string   `json:"model"`
	Messages    []Prompt `json:"messages"`
	Temperature float64  `json:"temperature"`
	Top_P       float64  `json:"top_p"`
}

type OpenAICompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens       int `json:"prompt_tokens"`
		CompletionTokens   int `json:"completion_tokens"`
		TotalTokens        int `json:"total_tokens"`
		PromptTokenDetails struct {
			PromptTokens int `json:"prompt_tokens"`
			TotalTokens  int `json:"total_tokens"`
		} `json:"prompt_token_details"`
		CompletionTokenDetails struct {
			ReasoningTokens          int `json:"reasoning_tokens"`
			AudioTokens              int `json:"audio_tokens"`
			AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
			RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
		} `json:"completion_token_details"`
	} `json:"usage"`
	ServiceTier       string `json:"service_tier"`
	SystemFingerprint string `json:"system_fingerprint"`
}

type OpenAI struct {
	baseURL      string
	apiKey       string
	SystemPrompt string
	Model        string
	Temperature  float64
	TopP         float64
}

func NewOpenAI(baseURL, apiKey string) *OpenAI {
	return &OpenAI{
		baseURL:     baseURL,
		apiKey:      apiKey,
		Model:       "gpt-3.5-turbo",
		Temperature: 0.7,
		TopP:        1,
	}
}

func NewOpenAIWithSystemPrompt(baseURL, apiKey, systemPrompt string) *OpenAI {
	return &OpenAI{
		baseURL:      baseURL,
		apiKey:       apiKey,
		SystemPrompt: systemPrompt,
		Model:        "gpt-3.5-turbo",
		Temperature:  0.7,
		TopP:         1,
	}
}

func (o *OpenAI) SetSystemPrompt(prompt string) error {
	if prompt != "" {
		return errors.New("system prompt cannot be empty")
	}
	o.SystemPrompt = prompt
	return nil
}

func (o *OpenAI) SetModel(model string) error {
	if model == "" {
		return errors.New("model cannot be empty")
	}
	o.Model = model
	return nil
}

func (o *OpenAI) SetTemperature(temperature float64) error {
	if temperature < 0 || temperature > 2 {
		return errors.New("temperature must be between 0 and 2")
	}
	o.Temperature = temperature
	return nil
}

func (o *OpenAI) SetTopP(topP float64) error {
	if topP < 0 || topP > 1 {
		return errors.New("topP must be between 0 and 1")
	}
	o.TopP = topP
	return nil
}

func (o *OpenAI) GetCompletion(prompt string) (OpenAICompletionResponse, error) {

	prompts := []Prompt{
		{
			Role:    "developer",
			Content: o.SystemPrompt,
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	request := OpenAICompletionRequest{
		Model:       o.Model,
		Messages:    prompts,
		Temperature: o.Temperature,
		Top_P:       o.TopP,
	}

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		log.Err(err).Msg("Error marshalling request")
		return OpenAICompletionResponse{}, err
	}

	req, err := http.NewRequest("POST", o.baseURL, bytes.NewBuffer(jsonRequest))
	if err != nil {
		log.Err(err).Msg("Error creating request")
		return OpenAICompletionResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", o.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("Error making request")
		return OpenAICompletionResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("Error reading response")
		return OpenAICompletionResponse{}, err
	}

	var response OpenAICompletionResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Err(err).Msg("Error unmarshalling response")
		return OpenAICompletionResponse{}, err
	}

	return response, nil
}
