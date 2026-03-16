package user

type UserRepository interface {
	Save(user User) (User, error)
}