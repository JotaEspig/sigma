package login

type UserInfo struct {
	username string
	password string
}

func (u *UserInfo) CheckLogin(userInput, passInput string) bool {
	return u.username == userInput && u.password == passInput
}
