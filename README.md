# ai-lib
go client for various chat ai apis

## Usage

### OpenAI

the openai client returns a `OpenAICompletionResponse` object that contains all of the meta data and the choices.

``` go
baseurl := "https://api.openai.com/v1/chat/completions"
apiKey := os.Getenv("OPENAI_KEY")
openai := client.NewOpenAI(baseurl, apiKey)

openai.SetSystemPrompt("You are a helpful assistant.")

res, err := openai.GetCompletion("Hello, how are you?")
if err != nil {
    log.Err(err).Msg("Error sending message")
    return
}

fmt.Println(res.Choices[0].Message.Content)
```

### Anthropic
``` go
apiKey := os.Getenv("ANTHROPIC_KEY")
baseURL := "https://api.anthropic.com/v1/messages"

anthropic := client.NewAnthropic(apiKey, baseURL)

anthropic.SetSystemPrompt("You are a helpful assistant.")

res, err := anthropic.GetCompletion("Hello, how are you?")
if err != nil {
    log.Err(err).Msg("Error sending message")
    return
}

fmt.Println(res.Content[0].Text)
```
