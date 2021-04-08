package repository

import "gorm.io/gorm"

type ExampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
	return ExampleRepository{db: db}
}

/*
Implements Interface Here...
func (repo *ExampleRepository) something(...) (*entity, error){

}
*/
