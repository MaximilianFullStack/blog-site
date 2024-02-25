package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/a-h/templ"

	"github.com/MaximilianFullStack/htmx-blog/components"
)

func main() {
	random := string(rune(rand.Int()))

	component := components.Hello(random)

	http.Handle("/", templ.Handler(component))

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Fatal(http.ListenAndServe(":3000", nil))

	fmt.Println("Server running on port [::]:3000 sd")
}
