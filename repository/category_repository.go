package repository

import "belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(name string) *entity.Category
}
