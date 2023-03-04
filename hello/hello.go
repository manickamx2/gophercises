package main

import "fmt"
import "example.com/greetings"
import "log"

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	names := []string{"Manickam", "Adam", "Tom"}
	messages, err := greetings.Hellos(names)

	if err != nil {
	    log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
