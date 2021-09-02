package registeration

import "github.com/mail2liuming/mycourses/courses"

type RegisterationService interface {
	resgisterCourse(courseRegisterDTO *CourseRegisterDTO) (uint, error)
	listRegisteredCourses(userID uint) ([]CourseRegisterResponseDTO, error)
}

type RegisterationServiceImpl struct {
	courseRepo        courses.Repository
	registerationRepo Repository
}

func NewRegisterationService(courseRepo courses.Repository, registerationRepo Repository) RegisterationService {
	return &RegisterationServiceImpl{courseRepo: courseRepo, registerationRepo: registerationRepo}
}

func (r *RegisterationServiceImpl) resgisterCourse(courseRegisterDTO *CourseRegisterDTO) (uint, error) {
	courseID, err := r.courseRepo.Create((&courses.CourseModel{UUID: courseRegisterDTO.CourseUIID, Name: courseRegisterDTO.Name, Address: courseRegisterDTO.Address, Description: courseRegisterDTO.Description}))
	if err != nil {
		return courseID, err
	}
	registerID, err := r.registerationRepo.Create(&RegisterationModel{UserID: courseRegisterDTO.UserID, CourseID: courseID, Name: courseRegisterDTO.Name, Date: courseRegisterDTO.Date, Repeatable: courseRegisterDTO.Repeatable})
	return registerID, err
}

func (r *RegisterationServiceImpl) listRegisteredCourses(userID uint) ([]CourseRegisterResponseDTO, error) {
	registeredCourses, err := r.registerationRepo.ListByUser(userID)
	if err != nil {
		return nil, err
	}
	course := make([]CourseRegisterResponseDTO, len(registeredCourses))
	for i, v := range registeredCourses {
		course[i].CourseID = v.CourseID
		course[i].Date = v.Date
		course[i].Name = v.Name
		course[i].Repeatable = v.Repeatable
	}
	return course, err
}
