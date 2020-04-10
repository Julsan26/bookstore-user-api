package app

import "github.com/julsan26/bookstore-user-api/controllers"

func mapUrls(){
	router.GET("/ping",controllers.Ping)

}
