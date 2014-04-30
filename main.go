package main

import (
	"runtime"
	"labix.org/v2/mgo"
	"github.com/rickihastings/codeprofile/modules"
)

var (
	DATABASE *mgo.Database = modules.DbConnect()
)

const (
	VERSION = "0.1.0-alpha"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// set multicore
}

func main() {
	modules.ConfigLoad("/home/vagrant/shared/bin/config.json")
	// load config

	SetupWebServer(DATABASE)
	// create webserver, pass database into it so we can use it elsewhere
}