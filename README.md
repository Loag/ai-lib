# ai-lib

a single interface lib for interacting with ai completion apis.

The purpose of this lib is to get up and running fast with ai completion apis for go 
It only includes exactly what is necessary to get text responses.

## Supported Providers

- [Anthropic (Claude)](https://docs.anthropic.com/en/docs/api)
- [OpenAI (ChatGPT)](https://platform.openai.com/docs/api-reference)
- [Google (Gemini)](https://ai.google.dev/gemini-api/docs)

## Usage

```go

conf := client.AIConfig{
    ApiKey:       apiKey,
    Model:        "claude-3-5-sonnet-20240620",
    SystemPrompt: "You are a helpful assistant.",
    Temperature:  0.7,
    TopP:         1.0,
    MaxTokens:    2480,
}

ai, err := client.NewAI(conf)
if err != nil {
    log.Err(err).Msg("Error creating AI")
    return
}

res, err := ai.GetCompletion("Hello, how are you?")
if err != nil {
    log.Err(err).Msg("Error getting completion")
    return
}

fmt.Println(res)

```
