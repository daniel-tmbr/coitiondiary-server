package routes

import (
	"coitiondiary.tmbr.me/handlers"
)

var locationsRoutes = Routes{
	Route{
		"Find nearby locations",
		"GET",
		"/locations/{coords}",
		handlers.GetNearbyLocations,
	},
	Route{
		"Find locations by text",
		"GET",
		"/locations/{text}",
		handlers.FindLocations,
	},
}
