package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"

	"github.com/MaximilianFullStack/htmx-blog/components"
)

func main() {
	component := components.Hello()

	http.Handle("/", templ.Handler(component))

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Server running on [::]:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
