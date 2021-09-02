package courses

type Reader interface{}

type Writer interface {
	Create(course *CourseModel) (uint, error)
}

type Repository interface {
	Reader
	Writer
}
