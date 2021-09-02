package registeration

type RegisterationModel struct {
	ID         uint   `gorm:"primary_key"`
	UserID     uint   `gorm:"column:user_id"`
	CourseID   uint   `gorm:"column:course_id"`
	Name       string `gorm:"column:name"`
	Date       string `gorm:"column:date"`
	Repeatable bool   `gorm:"column:repeatable"`
}
