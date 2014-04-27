package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/rickihastings/codeprofile/models"
	"github.com/rickihastings/codeprofile/modules/api"
)

func SetupWebServer() *martini.ClassicMartini {
	m := martini.Classic()
	// create a classic martini webserver

	SetupRoutes(m)
	// setup our routes

	m.Run()
	// run the server

	return m;
}

func SetupRoutes(m *martini.ClassicMartini) {
	m.Group("/api/user", func(r martini.Router) {
		r.Put("/", binding.Bind(models.User{}), api.RegisterUser)
	})
}