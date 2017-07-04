package routes

import (
	"coitiondiary.tmbr.me/handlers"
)

var usersRoutes = Routes{
	Route{
		"Update user",
		"PUT",
		"/users",
		handlers.UpdateUser,
	},
	Route{
		"Get user by id",
		"GET",
		"/users/{id:[0-9]+}",
		handlers.GetUser,
	},
	Route{
		"Delete user by id",
		"DELETE",
		"/users/{id:[0-9]+}",
		handlers.DeleteUser,
	},
	Route{
		"Upload avatar for user by id",
		"POST",
		"/users/{id:[0-9]+}/uploadImage",
		handlers.UploadUserAvatar,
	},
	Route{
		"Find user by username",
		"GET",
		"/users/find/{username}",
		handlers.FindUser,
	},
	Route{
		"Create dummy partner",
		"POST",
		"/users/dummy",
		handlers.CreateDummyPartner,
	},
	Route{
		"Update dummy partner",
		"PUT",
		"/users/dummy",
		handlers.UpdateDummyPartner,
	},
	Route{
		"Merge user with dummy partner",
		"POST",
		"/users/dummy/merge",
		handlers.MergeDummyPartner,
	},
}
