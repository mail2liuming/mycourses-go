package registeration

type CourseRegisterDTO struct {
	UserID      uint   `json:"userid"`
	CourseUIID  string `json:"courseuuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Repeatable  bool   `json:"repeatable"`
	Address     string `json:"address"`
}
