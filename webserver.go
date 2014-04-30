package main

import (
	"labix.org/v2/mgo"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/rickihastings/codeprofile/models"
	"github.com/rickihastings/codeprofile/modules"
	"github.com/rickihastings/codeprofile/modules/api"
)

func SetupWebServer(db *mgo.Database) (*martini.ClassicMartini) {
	m := martini.Classic()
	// create a classic martini webserver

	m.Map(db)
	// inject the database

	models.InjectDatabase(db)
	SetupRoutes(db, m)
	// setup our routes

	m.Run()
	// run the server

	return m;
}

func SetupRoutes(db *mgo.Database, m *martini.ClassicMartini) {
	m.Group("/oauth", func(r martini.Router) {
		r.Get("/register/:type", modules.Register)
		r.Get("/github", binding.Bind(models.GithubOAuth{}), modules.GithubOAuth)
	})

	m.Group("/api/user", func(r martini.Router) {
		r.Put("/", binding.Bind(models.User{}), api.RegisterUser)
	})
}