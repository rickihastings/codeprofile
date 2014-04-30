package modules

import (
	"fmt"
	"time"
	"net/http"
	"crypto/rand"
	"labix.org/v2/mgo"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/rickihastings/codeprofile/models"
)

const (
	Codeprofile	string = "http://10.0.33.34:3000/oauth/"
	Github		string = "https://github.com/login/oauth/authorize"
	Bitbucket	string = ""
	Linkedin	string = ""
)

func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b % byte(len(alphanum))]
	}
	return string(bytes)
}

func Register(db *mgo.Database, params martini.Params, res http.ResponseWriter, req *http.Request) string {
	if params["type"] == "github" {
		state := randString(10)
		url := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s%s&scope=user,repo:status,read:org&state=%s", Github, Config.GithubApiKey, Codeprofile, "github", state)
		// generate a random string and a valid github oauth url

		var newTime = time.Now()
		var stateStruct = models.GithubOAuth{"", "", state, newTime}

		err := db.C("states").Insert(&stateStruct)
		if err != nil {
			panic(err)
		}
		// we need to save the state we used so we can verify it later, we'll save it with a small timeout window aswell

		return url
		//http.Redirect(res, req, url, 301)
	}

	return ""
}

func GithubOAuth(oauth models.GithubOAuth, err binding.Errors, res http.ResponseWriter) string {
	fmt.Println(oauth, "wow");
	return ""
}