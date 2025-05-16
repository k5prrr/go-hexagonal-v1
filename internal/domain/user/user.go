package user

import (
	"app/internal/ports"
	"time"
)

type User struct {
	Uid string

	FamilyName string
	Name       string
	MiddleName string

	BirthDate time.Time
	Phone     string
	Email     string

	LastLogin time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	key  string
}

func (u *User) FullName() string {
	return fmt.Sprintf(
		"%s %s %s",
		FamilyName,
		Name,
		MiddleName,
	)
}





func NewUser() *User {
	return &User{}
}

func (u *User) Create() {
	u.Uid = 
}
