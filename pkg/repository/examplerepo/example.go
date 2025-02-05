package examplerepo

import "gorm.io/gorm"

type IExampleRepository interface {
	Example() string
}

type ExampleRepository struct {
	Database *gorm.DB
}

func NewExampleRepository(dbconn *gorm.DB) IExampleRepository {
	return &ExampleRepository{
		Database: dbconn,
	}
}

func (r *ExampleRepository) Example() string {
	return "Hello, World from ExampleRepository!"
}
