package routes

import (
	"coitiondiary.tmbr.me/handlers"
)

var authRoutes = Routes{
	Route{
		"Register user",
		"POST",
		"/auth/register",
		handlers.Register,
	},
	Route{
		"Login user",
		"POST",
		"/auth/login",
		handlers.Login,
	},
	Route{
		"Logout user",
		"GET",
		"/auth/logout",
		handlers.Logout,
	},
}
