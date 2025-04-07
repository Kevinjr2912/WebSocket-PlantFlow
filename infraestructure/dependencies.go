package infraestructure

import (
	"consumersData/application/services"
	"consumersData/infraestructure/adapters"
)

var (
	rabbit   *adapters.Rabbit
	gorrilla *adapters.Gorilla
)

func StartDependencies() {
	gorrilla = adapters.NewGorilla()
	rabbit = adapters.NewRabbit()

	go services.NewSuscriber(rabbit, gorrilla).Run("queue.monitoring", "data.monitoring", "monitoring")
}

// go func() {
// 	fmt.Println("Starting RabbitMQ suscriber...")
// 	err := services.NewSuscriber(rabbit, gorrilla).Run("queue.monitoring", "data.monitoring", "monitoring")
// 	if err != nil {
// 		log.Printf("Suscriber error: %v", err)
// 	}
// }()


func GetGorilla() *adapters.Gorilla {
	return gorrilla
}
