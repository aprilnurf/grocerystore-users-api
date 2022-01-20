package app

import (
	"github.com/aprilnurf/grocerystore_users-api/controllers/ping"
	"github.com/aprilnurf/grocerystore_users-api/controllers/users"
)

func mapUrls(){
	router.GET("/ping", ping.Ping)
	//router.GET("/users", controllers.SearchUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.POST("/users", users.CreateUser)
}
