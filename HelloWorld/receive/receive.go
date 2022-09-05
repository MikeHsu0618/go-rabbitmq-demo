package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//conn, err := amqp.Dial("amqp://5VWDVq1abMqICFmRpt:5VWDVq1abMqICFmRpt@stgmq.betradar.com:5671/unifiedfeed/34959")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	FailOnError(err, "Failed to connect toç RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Queue 不用每次都建立，只要其中一方或已存在就好了
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
