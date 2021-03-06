package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
)


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
	UpdateExample(gin.Context) error
}
