# gpt Go Package

The `gpt` package is a Go library designed to facilitate easy communication with the GPT-3 API. This library enables you to perform text completion operations using OpenAI's GPT-3 model.

## Installation

To include this package as a Go module in your project:

```bash
go get github.com/ssibrahimbas/gpt
```

## Usage

Below is a basic usage example:

```go
package main

import (
 "fmt"
 "github.com/ssibrahimbas/gpt"
)

func main() {
 client := gpt.New(gpt.Config{
  ApiKey: "YOUR_API_KEY",
 })
 res, err := client.Complete(gpt.CompletionRequest{
  Messages: []gpt.Message{
   {
    Role:    gpt.RoleUser,
    Content: "Hello, how are you?",
   },
  },
  Model:            gpt.ModelGpt3Turbo,
  Temperature:      0.7,
  MaxTokens:        64,
  TopP:             1,
  FrequencyPenalty: 0,
  PresencePenalty:  0,
 })
 if err != nil {
  panic(err)
 }
 fmt.Println(res.Choices[0].Message.Content)
}
```

In this example, we're sending a completion request to the GPT-3 API and printing the generated text response from the model.

## Configuration

The `gpt.New` function takes a `Config` object containing your API key. You can obtain this key from OpenAI's official website.

## Documentation

For detailed documentation and further usage examples, please refer to the Go documentation comments within the package. You can also view the documentation online at [pkg.go.dev](https://pkg.go.dev/github.com/ssibrahimbas/gpt).

## License

This project is licensed under the Apache 2.0 license. Please refer to the [LICENSE](LICENSE) file for more details.
