package gemini

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Gemini struct {
	model   string
	apiKey  string
	baseURL string
	system  string
}

func NewGemini(apiKey string, baseURL string) *Gemini {
	return &Gemini{apiKey: apiKey, baseURL: baseURL}
}

func (g *Gemini) SetModel(model string) {
	g.model = model
}

func (g *Gemini) SetSystemPrompt(system string) {
	g.system = system
}

/*

response format:

{
  "candidates": [
    {
      "content": {
        "parts": [
          {
            "text": "Artificial intelligence (AI) is a broad field encompassing many different techniques, but at its core, it's about creating systems that can perform tasks that typically require human intelligence. These tasks include things like:\n\n* **Learning:** Acquiring information and rules for using the information.\n* **Reasoning:** Using rules to reach approximate or definite conclusions.\n* **Problem-solving:** Finding solutions to complex situations.\n* **Perception:** Interpreting sensory information (like images, sound, and text).\n* **Language understanding:** Comprehending and generating human language.\n\nAI systems achieve these capabilities through a variety of approaches, the most prominent being:\n\n**1. Machine Learning (ML):** This is the most prevalent approach.  Instead of explicitly programming rules, ML algorithms learn patterns from data.  They identify relationships and correlations within large datasets to make predictions or decisions.  Different types of ML exist:\n\n* **Supervised Learning:**  The algorithm is trained on labeled data (data where the correct answer is already known).  For example, showing an algorithm many pictures of cats and dogs, labeled accordingly, so it learns to distinguish them.\n* **Unsupervised Learning:** The algorithm is trained on unlabeled data and tries to find structure or patterns on its own.  For example, grouping customers based on their purchasing behavior without pre-defined categories.\n* **Reinforcement Learning:** The algorithm learns through trial and error by interacting with an environment. It receives rewards for good actions and penalties for bad ones, learning to maximize its cumulative reward.  Think of a robot learning to walk – it gets rewarded for taking steps and penalized for falling.\n\n**2. Deep Learning (DL):** A subset of machine learning that uses artificial neural networks with multiple layers (hence \"deep\").  These networks are inspired by the structure and function of the human brain, allowing them to learn complex patterns from vast amounts of data.  Deep learning excels at tasks like image recognition, natural language processing, and speech recognition.\n\n**3. Expert Systems:** These systems encode human expertise into a computer program, using a set of rules and facts to make decisions in a specific domain.  For example, a medical diagnosis system might use rules based on symptoms and test results to suggest possible diagnoses.\n\n**4. Natural Language Processing (NLP):** This focuses on enabling computers to understand, interpret, and generate human language.  Tasks include translation, sentiment analysis, text summarization, and chatbot development.\n\n**5. Computer Vision:** This allows computers to \"see\" and interpret images and videos.  Applications include object recognition, facial recognition, and image segmentation.\n\n\n**How it works in a simplified way:**\n\nImagine you want to build an AI that identifies cats in pictures.  Using machine learning:\n\n1. **Data Collection:** You gather a massive dataset of images, some containing cats and some not.\n2. **Training:** You feed this data to a machine learning algorithm (maybe a deep learning model).  The algorithm identifies features in the images (e.g., pointy ears, whiskers, fur patterns) that are associated with cats.\n3. **Testing:** You test the algorithm on new images it hasn't seen before to evaluate its accuracy in identifying cats.\n4. **Refinement:** You adjust the algorithm's parameters or add more data to improve its performance.\n\nThis is a simplified explanation, but it highlights the fundamental principles behind many AI systems.  The specific techniques and algorithms used depend heavily on the particular task and the available data.  The field is constantly evolving, with new approaches and breakthroughs being made regularly.\n"
          }
        ],
        "role": "model"
      },
      "finishReason": "STOP",
      "citationMetadata": {
        "citationSources": [
          {
            "startIndex": 257,
            "endIndex": 390,
            "uri": "https://www.deskera.com/blog/ai-automation-procurement/"
          }
        ]
      },
      "avgLogprobs": -0.17730267498944258
    }
  ],
  "usageMetadata": {
    "promptTokenCount": 4,
    "candidatesTokenCount": 740,
    "totalTokenCount": 744,
    "promptTokensDetails": [
      {
        "modality": "TEXT",
        "tokenCount": 4
      }
    ],
    "candidatesTokensDetails": [
      {
        "modality": "TEXT",
        "tokenCount": 740
      }
    ]
  },
  "modelVersion": "gemini-1.5-flash"
}
*/

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	UsageMetadata struct {
		PromptTokenCount     int `json:"promptTokenCount"`
		CandidatesTokenCount int `json:"candidatesTokenCount"`
		TotalTokenCount      int `json:"totalTokenCount"`
		PromptTokensDetails  []struct {
			Modality   string `json:"modality"`
			TokenCount int    `json:"tokenCount"`
		} `json:"promptTokensDetails"`
		CandidatesTokensDetails []struct {
			Modality   string `json:"modality"`
			TokenCount int    `json:"tokenCount"`
		} `json:"candidatesTokensDetails"`
	} `json:"usageMetadata"`
}

/*

// instruction format:

{ "system_instruction": {
    "parts":
      { "text": "You are a cat. Your name is Neko."}},
    "contents": {
      "parts": {
        "text": "Hello there"}}}
*/

func (g *Gemini) GetCompletion(prompt string) (string, error) {

	url := g.baseURL + "/" + g.model + ":generateContent" + "?key=" + g.apiKey

	reqBody := map[string]interface{}{
		"system_instruction": map[string]interface{}{
			"parts": map[string]interface{}{
				"text": g.system,
			},
		},
		"contents": map[string]interface{}{
			"parts": map[string]interface{}{
				"text": prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response GeminiResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return "", err
	}

	return response.Candidates[0].Content.Parts[0].Text, nil
}
