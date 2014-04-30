package api

import (
	"log"
	"net/http"
	"github.com/rickihastings/codeprofile/models"
)

func RegisterUser(user models.User, res http.ResponseWriter, req *http.Request) {
	log.Print("yay")
}