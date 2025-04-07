package adapters

import (
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

type Rabbit struct {
	Broker  *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbit() *Rabbit {

	// Cargamos las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Obtenemos la URL de Rabbit de nuestras variables de entorno
	rabbitUrl := os.Getenv("RABBIT_URL")

	// Conexión a RabbitMQ
	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		log.Fatal("Error al abrir una conexión hacia rabbitmq")
	}

	// Abrimos un canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error al abrir un canal")
	}

	return &Rabbit{Broker: conn, Channel: ch}

}

func (r *Rabbit) ReceiveContent(nameQueue, routingKey string) <-chan amqp.Delivery {

	err := r.Channel.ExchangeDeclare(
		"data.system", // name
		"topic",       // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	r.FailOnError(err, "Failed to declare an exchange")

	q, err := r.Channel.QueueDeclare(
		nameQueue, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	r.FailOnError(err, "Failed to declare a queue")

	err = r.Channel.QueueBind(
		q.Name,        // queue name
		routingKey,    // routing key
		"data.system", // exchange
		false,
		nil,
	)
	r.FailOnError(err, "Failed to bind a queue")

	msgs, err := r.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	r.FailOnError(err, "Failed to register a consumer")

	return msgs
}

func (r *Rabbit) FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
