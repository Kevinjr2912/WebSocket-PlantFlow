package routes

import (
	"consumersData/infraestructure/controllers"
	"net/http"
)

func RegisterRoutes() {

	http.HandleFunc("/ws", controllers.NewWsController().Run)

}