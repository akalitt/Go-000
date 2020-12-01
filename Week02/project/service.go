package main

type UserService struct {
	ID int
}

// service 负责传递
func (u *UserService) GetUserByIdService() (string, error) {
	return GetUserById(u.ID)
}
