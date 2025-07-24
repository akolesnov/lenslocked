package main

import (
	"fmt"
	"net/http"

	"github.com/akolesnov/lenslocked/controllers"
	"github.com/akolesnov/lenslocked/templates"
	"github.com/akolesnov/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
