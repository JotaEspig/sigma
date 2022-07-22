package admin

import (
	"errors"
	"sigma/models/user"
)

type Admin struct {
	UID  uint `gorm:"primary_key;column:id"`
	Role string
	User *user.User `gorm:"foreignKey:UID"`
}

func InitAdmin(u *user.User) (*Admin, error) {
	if u.ID == 0 {
		return nil, errors.New("admin: UserID cannot be 0")
	}

	a := &Admin{
		UID:  u.ID,
		User: u,
	}

	return a, nil
}

// Adds admin value to map contaning user info
func (a *Admin) ToMap() map[string]interface{} {
	adminMap := make(map[string]interface{})
	if a.User != nil {
		adminMap = a.User.ToMap()
	}

	adminMap["role"] = a.Role

	return adminMap
}
