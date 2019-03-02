package main

import (
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

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Could not connect to Rabbit MQ services!")

	// Close connection after all subsequent code in main has been executed 
	defer connection.Close()

	channel, err := connection.Channel()
	FailOnError(err, "Failed to create channel")

	// Close channel after main function has been executed
	defer channel.Close()
}