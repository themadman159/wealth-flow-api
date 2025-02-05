package examplesvc

import (
	"go-api/pkg/repository/examplerepo"

	"gorm.io/gorm"
)

type IExampleService interface {
	Example() string
}

type ExampleService struct {
	ExampleRepository examplerepo.IExampleRepository
}

func NewExampleServiced(dbconn *gorm.DB) IExampleService {
	return &ExampleService{
		ExampleRepository: examplerepo.NewExampleRepository(dbconn),
	}
}

func (s *ExampleService) Example() string {
	example := s.ExampleRepository.Example()
	return example
}
