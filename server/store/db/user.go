package db

import (
	"github.com/pocketbase/pocketbase/core"
)

var _ core.RecordProxy = (*User)(nil)

type User struct {
	core.BaseRecordProxy
}

func (u *User) GetFirstname() string {
	return u.GetString("firstname")
}

func (u *User) SetFirstname(firstname *string) {
	u.Set("firstname", firstname)
}

func (u *User) GetLastname() string {
	return u.GetString("lastname")
}

func (u *User) SetLastname(lastname *string) {
	u.Set("lastname", lastname)
}

func (u *User) GetBio() string {
	return u.GetString("bio")
}

func (u *User) SetBio(bio *string) {
	u.Set("bio", bio)
}

func (u *User) GetLockoutEnabled() bool {
	return u.GetBool("lockoutenabled")
}

func (u *User) SetLockoutEnabled(lockout bool) {
	u.Set("lockoutenabled", lockout)
}
