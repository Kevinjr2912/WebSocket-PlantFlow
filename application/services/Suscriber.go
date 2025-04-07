package services

import (
	"consumersData/application/repositories"
	"consumersData/infraestructure/adapters"
	"log"
)

type Suscriber struct {
	suscriber repositories.ISuscriber
	ws *adapters.Gorilla
}

func NewSuscriber(suscriber repositories.ISuscriber, ws *adapters.Gorilla) *Suscriber {
	return &Suscriber{suscriber: suscriber, ws: ws}
}

func (s *Suscriber) Run(nameQueue string, routingKey string, messageType string) {

	msgs := s.suscriber.ReceiveContent(nameQueue, routingKey)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)

			s.ws.BroadCast(messageType, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
