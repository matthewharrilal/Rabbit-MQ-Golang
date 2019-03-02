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