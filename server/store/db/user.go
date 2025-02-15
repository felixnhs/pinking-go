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
const User_LockoutEnabled = "lockoutenabled"
const User_Avatar = "avatar"
const User_Followers = "followers"
const User_Following = "following"

const User_Followers_Count = "followercount"
const User_Following_Count = "followingcount"

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

func (p *User) GetAvatar() string {
	return p.GetString(User_Avatar)
}

func (p *User) SetAvatar(f *filesystem.File) {
	p.Set(User_Avatar, f)
}

func (p *User) SetAvatarBase64(base64Str *string) {
	p.Set(User_Avatar, *base64Str)
}

func (p *User) ClearAvatar() {
	p.Set(User_Avatar, nil)
}

func (p *User) GetFollowers() []string {
	return p.GetStringSlice(User_Followers)
}

func (p *User) AddFollower(id string) {
	p.Set(User_Followers+"+", id)
}

func (p *User) RemoveFollower(id string) {
	p.Set(User_Followers+"-", id)
}

func (p *User) GetFollowing() []string {
	return p.GetStringSlice(User_Following)
}

func (p *User) AddFollowing(id string) {
	p.Set(User_Following+"+", id)
}

func (p *User) RemoveFollowing(id string) {
	p.Set(User_Following+"-", id)
}
