package models

import (
	"time"
	//"errors"
	//"crypto/sha256"
	"encoding/json"
)

// the model for a basic User
type User struct {
	Id			int64 `form:"-" json:"id"`
	Name		string `form:"name" json:"name" binding:"required"`
	Email		string `form:"email" json:"email" binding:"required"`
	Password	string `form:"password" json:"password" binding:"required"`
	Location	string `form:"location" json:"location"`
	Created 	time.Time `form:"-" json:"created"`
	Updated		time.Time `form:"-" json:"updated"`
	CreatedWith	int8 `form:"-" json:"createdWith"`
	// basic user information

	GithubId	int64 `form:"-" json:"githubId"`
	BitbucketId	int64 `form:"-" json:"bitbucketId"`
	LinkedinId	int64 `form:"-" json:"linkedinId"`
	// other profile ids
}

func (user *User) ToJson() (error, string) {
	js, err := json.Marshal(user)

	return err, string(js[:])
} 

// create a new user record and store it in the database
func RegisterUser(user *User) (error, *User) {
	return nil, nil
}