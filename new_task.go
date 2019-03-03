package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/streadway/amqp"
)

// FailOnError function provided custom error strings based off the corresponding error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%v:  %v", msg, err)
	}
}

// Hlper function to extract text from the arguments
func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "Placeholder Message"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}


func main() {
	body := bodyFrom(os.Args)
}
