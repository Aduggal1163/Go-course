package user

type User struct {
	Name string
}

type UserManager struct {
	Users []User
}

func (um *UserManager) AddUser(name string) {
	um.Users = append(um.Users, User{Name: name})
}

func (um *UserManager) FindUserByName(name string) (*User, bool) {
	for i := range um.Users {
		if name == um.Users[i].Name {
			return &um.Users[i], true
		}
	}
	return nil, false
}

// Delete user by name
func (um *UserManager) DeleteUser(name string) bool {
	for i := range um.Users {
		if um.Users[i].Name == name {
			um.Users = append(um.Users[:i], um.Users[i+1:]...)
			return true
		}
	}
	return false
}
