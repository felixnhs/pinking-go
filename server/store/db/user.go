package db

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/filesystem"
)

var _ core.RecordProxy = (*User)(nil)

type User struct {
	core.BaseRecordProxy
}

const User_Firstname = "firstname"
const User_Lastname = "lastname"
const User_Bio = "bio"
const User_LockoutEnabled = "bio"
const User_Avatar = "avatar"

func (u *User) GetFirstname() string {
	return u.GetString(User_Firstname)
}

func (u *User) SetFirstname(firstname *string) {
	u.Set(User_Firstname, firstname)
}

func (u *User) GetLastname() string {
	return u.GetString(User_Lastname)
}

func (u *User) SetLastname(lastname *string) {
	u.Set(User_Lastname, lastname)
}

func (u *User) GetBio() string {
	return u.GetString(User_Bio)
}

func (u *User) SetBio(bio *string) {
	u.Set(User_Bio, bio)
}

func (u *User) GetLockoutEnabled() bool {
	return u.GetBool(User_LockoutEnabled)
}

func (u *User) SetLockoutEnabled(lockout bool) {
	u.Set(User_LockoutEnabled, lockout)
}

func (p *Image) SetAvatar(f *filesystem.File) {
	p.Set(User_Avatar, f)
}
