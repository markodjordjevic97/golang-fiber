package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-" gorm:"-"`
	RoleId    uint   `json:"roleId"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user.Password = hashedPassword

}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
