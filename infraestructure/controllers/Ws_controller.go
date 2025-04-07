package controllers

import (
	"consumersData/infraestructure"
	"consumersData/infraestructure/adapters"
	"net/http"
)

type WsController struct {
	service *adapters.Gorilla
}

func NewWsController() *WsController {

	gorilla := infraestructure.GetGorilla()
	
	return &WsController{service: gorilla}

}

func (sg_c *WsController) Run(w http.ResponseWriter, r *http.Request) {

	if sg_c.service == nil {
		http.Error(w, "WebSocket service not initialized", http.StatusInternalServerError)
		return
	}

	sg_c.service.Handler(w, r)
}