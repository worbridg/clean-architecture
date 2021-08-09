//go:generate mockgen -source repository.go -destination=mock_repository.go -package user
package user

type Repository interface {
	GetByID(id UserID) (*User, error)
	Store(user *User) error
	Delete(id UserID) error
}
