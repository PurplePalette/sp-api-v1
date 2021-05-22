package server

import (
	"firebase.google.com/go/db"
	"golang.org/x/net/context"
)

// userIdList stores
var userIdList map[string]bool

// IsUserAlive checks the user is exist and not removed
func IsUserAlive(userId string) bool {
	_, ok := userIdList[userId]
	return ok
}

// LoadUserList gets the list of users from firebase for caching user status
func LoadUserIdList(db *db.Client) {
	ref := db.NewRef("user").OrderByChild("isDeleted").EqualTo(false)
	ref.Get(context.Background(), &userIdList)
}
