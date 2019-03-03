package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// FailOnError function provided custom error strings based off the corresponding error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%v:  %v", msg, err)
	}
}

type Example struct {
	name string

	age int
}

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Could not connect to Rabbit MQ services!")

	// Close connection after all subsequent code in main has been executed
	defer connection.Close()

	channel, err := connection.Channel()
	FailOnError(err, "Failed to create channel")

	// Close channel after main function has been executed
	defer channel.Close()

	queue, err := channel.QueueDeclare("testQueue", true, false, false, false, nil)
	FailOnError(err, "Failed to declare queue")

	example := Example{"matthew", 12}
	messageContents, err := json.Marshal(&example)

	err = channel.Publish(
		"",         // Default exchange name
		queue.Name, // Name used as routing key to pass message to corresponding queue
		false,      // Mandatory parameter meaning that message must be delivered ... can run into error if queue doesn't exist
		false,      // Immediate parameter meaning that message needs to be delivered immediately can fail if corresponding queue cant hold any more messages in queue
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageContents), // Convert body to be byte slice to send through=
		})

	FailOnError(err, "Failed to publish a message")
	fmt.Printf("Sent Message %v", messageContents)

}
