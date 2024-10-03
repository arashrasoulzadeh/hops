package renderer

import (
	"os/user"
	"strconv"
	"syscall"
)

// UserInfo provides information about the current user and their permissions.
type UserInfo struct{}

// NewUserInfo returns an instance of UserInfo.
func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

// Username returns the current username.
func (u *UserInfo) Username() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.Username, nil
}

// Group returns the primary group of the user.
func (u *UserInfo) Group() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	group, err := user.LookupGroupId(usr.Gid)
	if err != nil {
		return "", err
	}
	return group.Name, nil
}

// Permissions returns the current user's permissions.
func (u *UserInfo) Permissions() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	// Get permissions of the user's home directory as an example.
	stat := syscall.Stat_t{}
	err = syscall.Stat(usr.HomeDir, &stat)
	if err != nil {
		return "", err
	}
	mode := strconv.Itoa(int(stat.Mode))
	return mode, nil
}
