package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MYGBM/goTesting/dependency"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Salut, "
)

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix

	}
	return
}
func main() {
	// fmt.Println(hello("Juan", "Spanish"))
	dependency.Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependency.MyGreeterHandler)))

}
