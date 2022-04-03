package login

type UserInfo struct {
	username string
	password string
}

func DefaultUserInfo() *UserInfo {
	return &UserInfo{
		username: "teste",
		password: "password",
	}
}

func (u *UserInfo) CheckLogin(userInput, passInput string) bool {
	return u.username == userInput && u.password == passInput
}
