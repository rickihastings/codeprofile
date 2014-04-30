package models

import (
	"time"
	"net/http"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"github.com/martini-contrib/binding"
)

var (
	DATABASE *mgo.Database = nil
)

func InjectDatabase(db *mgo.Database) {
	DATABASE = db
}

// the model for a basic Github OAuth response
type GithubOAuth struct {
	Token	string		`form:"-" bson:"tk,omitempty"`
	Code	string		`form:"code" bson:"c,omitempty"`
	State	string		`form:"state" binding:"required" bson:"s"`
	Timeout	time.Time	`form:"-" bson:"t"`
}

// validate the OAuth response
func (oauth *GithubOAuth) Validate(errors *binding.Errors, req *http.Request) {
	query, err := DATABASE.C("states").Find(bson.M{"s": oauth.State, "t": bson.M{"$lt": time.Now().Add(time.Hour)}}).Count()
	// validate the State that "github" has given us
	
	if err != nil || query == 0 {
		errors.Fields["state"] = "Invalid state"
	}
}