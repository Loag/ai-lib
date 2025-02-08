# ai-lib

a single interface lib for interacting with ai completion apis.

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
