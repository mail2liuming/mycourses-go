package users

type UserModel struct {
	ID       uint   `gorm:"primary_key"`
	UserName string `gorm:"column:username"`
	Email    string `gorm:"column:email;unique_index"`
	Token    string `gorm:"column:token;not null"`
}
