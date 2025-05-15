package main

import (
  "fmt"
  "log"
  // "os"
  // "path"
  zen "github.com/gorules/zen-go"
  "github.com/gofiber/fiber/v2"
  "encoding/json"
)

func readTestFile(key string) ([]byte, error) {
  return []byte(key), nil
}

func eval(code string, input string) (string, error) {
  engine := zen.NewEngine(zen.EngineConfig{Loader: readTestFile})
  defer engine.Dispose() // Call to avoid leaks

  var jsonMap map[string]any
  err := json.Unmarshal([]byte(input), &jsonMap)
  if err != nil {
    fmt.Printf("Error unmarshaling JSON: %v\n", err)
    return "", err
  }

  output, err := engine.Evaluate(code, jsonMap)
  if err != nil {
    return "", err
  }

  return string(output.Result), nil
}

type RequestBody struct {
  Code  string `json:"code"`
  Input string `json:"input"`
}

func main() {
  // Initialize a new Fiber app
  app := fiber.New()

  // Define a route for the GET method on the root path '/'
  app.Post("/eval", func(c *fiber.Ctx) error {
    var body RequestBody

    if err := c.BodyParser(&body); err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error": "Cannot parse JSON",
      })
    }

    fmt.Printf("Received code: %s\n", body.Code)
    fmt.Printf("Received input: %s\n", body.Input)

    result, err := eval(body.Code, body.Input)
    if err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error": "deu ruim vei",
      })
    }

    return c.JSON(fiber.Map{
      "result": result,
    })
  })

  // Start the server on port 3000
  log.Fatal(app.Listen(":3000"))
}
