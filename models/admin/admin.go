package admin

import (
	"errors"
	"sigma/models/user"
)

// Admin represents the admin role in sigma
type Admin struct {
	UID  uint `gorm:"primary_key;column:id"`
	Role string
	User *user.User `gorm:"foreignKey:UID"`
}

// InitAdmin initializes an admin struct
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

// ToMap adds admin value to map contaning user info
func (a *Admin) ToMap() map[string]interface{} {
	adminMap := make(map[string]interface{})
	if a.User != nil {
		adminMap = a.User.ToMap()
	}

	adminMap["role"] = a.Role

	return adminMap
}
