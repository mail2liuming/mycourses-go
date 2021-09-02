package courses

import "gorm.io/gorm"

type repo struct {
	DB *gorm.DB
}

func NewSqlliteRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) Create(course *CourseModel) (uint, error) {
	result := r.DB.Create(course)
	return course.ID, result.Error
}
