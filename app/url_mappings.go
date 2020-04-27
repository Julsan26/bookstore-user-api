package app

import (
	"github.com/julsan26/bookstore-user-api/controllers/ping"
	"github.com/julsan26/bookstore-user-api/controllers/users"
)

func mapUrls(){
	router.GET("/ping",ping.Ping)
	router.POST("/users",users.CreateUser)
	router.GET("/user/:user_id",users.GetUser)


}
