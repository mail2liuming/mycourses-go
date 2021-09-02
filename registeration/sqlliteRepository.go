package registeration

import "gorm.io/gorm"

type repo struct {
	DB *gorm.DB
}

func NewRegisterationRepository(db *gorm.DB) *repo {
	return &repo{
		DB: db,
	}
}

func (r *repo) Create(registeration *RegisterationModel) (uint, error) {
	result := r.DB.Create(registeration)
	return registeration.ID, result.Error
}

func (r *repo) ListByUser(userID uint) ([]RegisterationModel, error) {
	var registeredCourses []RegisterationModel
	result := r.DB.Where("user_id = ?", userID).Find(&registeredCourses)

	return registeredCourses, result.Error
}
