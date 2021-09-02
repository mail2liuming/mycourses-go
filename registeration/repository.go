package registeration

type Reader interface {
	ListByUser(userID uint) ([]RegisterationModel, error)
}

type Writer interface {
	Create(registeration *RegisterationModel) (uint, error)
}

type Repository interface {
	Reader
	Writer
}
