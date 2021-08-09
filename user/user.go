package user

type User struct {
	id   UserID
	name UserName
}

func NewUser(id UserID, name UserName) *User {
	return &User{
		id:   id,
		name: name,
	}
}

func (user *User) ID() UserID {
	return user.id
}

func (user *User) Name() UserName {
	return user.name
}
