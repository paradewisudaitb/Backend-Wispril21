# Example of Entity File

```go
package example

import "github.com/paradewisudaitb/Backend/domain"

type Example struct {
	domain.EntityBase
	ExampleString string `json:"example"`
	ExampleInt    int    `json:"example_int"`
}

type ExampleUsecase interface {
	UpdateExample(exampleId string) error
}

type ExampleRepository interface {
    UpdateExample(exampleId string, example Example) error
}

type ExampleController interface {
    UpdateExample(echo.Context) error
}
```