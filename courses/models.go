package courses

type CourseModel struct {
	ID          uint   `gorm:"primary_key"`
	UUID        string `gorm:"column:uuid, index"`
	Name        string `gorm:"column:name"`
	Address     string `gorm:"column:adress"`
	Description string `gorm:"column:description"`
}
