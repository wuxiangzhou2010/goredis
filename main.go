package main

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)
var client *redis.Client
var templates *template.Template



func main()  {
	client = redis.NewClient(&redis.Options{
		Addr:"172.16.3.116:6379",
	})
	templates = template.Must(template.ParseGlob("template/*.html"))
	r:=mux.NewRouter()
	r.HandleFunc("/hello", indexHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

	
}

func indexHandler( w http.ResponseWriter, r * http.Request)  {
	//fmt.Fprintf(w, "Hello world!")
	comments, err := client.LRange("comments", 0, 10).Result()
	if err != nil {
		panic(err)
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}
