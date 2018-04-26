package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/server-side-rendering-with-vue-demos/golang-demo/models"
)

func main() {
	render := template.Must(template.ParseFiles("views/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		random := models.Random{}
		random.Number = rand.Intn(100)

		if err := render.ExecuteTemplate(w, "index.html", random); err != nil {
			http.Error(w, "Something went wrong =(", http.StatusInternalServerError)
			fmt.Println("[main] Erro on parse template: ", err.Error())
		}
	})

	fmt.Println("Server running..")
	http.ListenAndServe(":8080", nil)
}
