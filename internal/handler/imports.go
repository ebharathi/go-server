package handler

import (
	"server/internal/handler/oauth"
	"server/internal/handler/users"
)

var CreateUser = users.CreateUser
var LoginUser = users.LoginUser
var GetMe = users.GetMe
var GoogleCallback = oauth.GoogleCallback
