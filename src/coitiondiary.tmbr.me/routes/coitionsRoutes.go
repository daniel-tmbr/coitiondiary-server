package routes

import (
	"coitiondiary.tmbr.me/handlers"
)

var coitionsRoutes = Routes{
	Route{
		"List coitions",
		"GET",
		"/coitions",
		handlers.ListCoitions,
	},
	Route{
		"Create coition",
		"POST",
		"/coitions",
		handlers.CreateCoition,
	},
	Route{
		"Update coition",
		"PUT",
		"/coitions",
		handlers.UpdateCoition,
	},
	Route{
		"Get coition by id",
		"GET",
		"/coitions/{id:[0-9]+}",
		handlers.GetCoition,
	},
	Route{
		"Delete coition by id",
		"DELETE",
		"/coitions/{id:[0-9]+}",
		handlers.DeleteCoition,
	},
}
