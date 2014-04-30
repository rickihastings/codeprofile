package models

import (
	"time"
	"net/http"
	"net/mail"
	"encoding/json"
	"github.com/martini-contrib/binding"
)

// the model for a basic User
type User struct {
	Id			int64		`form:"-" json:"id"`
	Name		string		`form:"name" json:"name" binding:"required"`
	Email		string		`form:"email" json:"email" binding:"required"`
	Password	string		`form:"password" json:"password" binding:"required"`
	Location	string		`form:"location" json:"location"`
	Created		time.Time	`form:"-" json:"created"`
	Updated		time.Time	`form:"-" json:"updated"`
	CreatedWith	string		`form:"createdWith" json:"createdWith" binding:"required"`
	// basic user information

	GithubId	int64		`form:"githubId" json:"githubId"`
	BitbucketId	int64		`form:"bitbucketId" json:"bitbucketId"`
	LinkedinId	int64		`form:"linkedinId" json:"linkedinId"`
	// other profile ids
}

func (user *User) ToJson() (error, string) {
	js, err := json.Marshal(user)

	return err, string(js[:])
} 

// validate the user record
func (user *User) Validate(errors *binding.Errors, req *http.Request) {
	if len(user.Name) > 100 {
		errors.Fields["name"] = "Too long; maximum 100 characters"
	}

	if email, err := mail.ParseAddress(user.Email); email == nil || err != nil {
		errors.Fields["email"] = "Invalid email address"
	}
	// validate name, email, password etc

	if user.CreatedWith != "github" && user.CreatedWith != "bitbucket" && user.CreatedWith != "linkedin" {
		errors.Fields["createdWith"] = "Can only be github, bitbucket or linkedin"
	}
	// is createdWith valid?

	if user.CreatedWith == "github" && user.GithubId <= 0 {
		errors.Fields["githubId"] = "Required"
	}
	// if it's github, check if there is an id too?

	if user.CreatedWith == "bitbucket" && user.BitbucketId <= 0 {
		errors.Fields["bitbucketId"] = "Required"
	}
	// bitbucket

	if user.CreatedWith == "linkedin" && user.LinkedinId <= 0 {
		errors.Fields["linkedinId"] = "Required"
	}
	// linkedin
}

// pull profile information
func (user *User) PullProfile(error) {
	// here we need to determine what kind of profile they have and request it from the provider
	// we should be already past OAuth at this stage and the provider will have successfully 
	// authenticated us
	// XXX: TODO
}