package routes

import (
	"coitiondiary.tmbr.me/handlers"
)

var tagsRoutes = Routes{
	Route{
		"Create tag",
		"POST",
		"/tags",
		handlers.CreateTag,
	},
	Route{
		"Find tags by name",
		"GET",
		"/tags/{name}",
		handlers.FindTags,
	},
}
