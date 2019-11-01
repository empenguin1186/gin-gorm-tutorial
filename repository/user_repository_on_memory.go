package repository

type UserRepoitoryOnMemory struct {
	users []User
}

func NewUserRepositoryOnMemory() UserRepoitoryOnMemory {
	return UserRepoitoryOnMemory{users: []Users{}}
}

func (u *UserRepoitoryOnMemory) FindAll() ([]User, error) {
	return users, nil
}

func (u *UserRepoitoryOnMemory) FindOne(id string) (User, error) {
	for _, e := range users {
		if e.id == id {
			return e, nil
		}
		return nil, errors("Could not find user.")
	}
}

func (u *UserRepoitoryOnMemory) Create(user User) error {
	users = append(users, user)
	return nil
}

func (u *UserRepoitoryOnMemory) Update(id string, user User) error {
	for i, e := range users {
		if e.id == id {
			users = append(users[:i], e, users[i+1:])
			return nil
		}
	}
	return errors("Could not find user.")
}

func (u *UserRepoitoryOnMemory) Delete(id string) error {
	for i, e := range users {
		if e.id = id {
			users = append(users[:i], users[i+1:])
			return nil
		}
	}
	return errors("Could not find user.")
}