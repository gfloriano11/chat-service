package user

type UserRepository interface {
	Save(user *User) (User, error)
	FindUserByEmail(email string) (*User, error)
	FindUserById(id int) (*User, error)
}