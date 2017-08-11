package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/ghodss/yaml"
	"github.com/rs/cors"
	"github.com/vincentvanderweele/swagger-todo-mvc/generated/router"
	"github.com/vincentvanderweele/swagger-todo-mvc/todo"
)

func main() {
	storage := todo.NewStorage()
	errorHandler := todo.NewErrorTransformer()

	handler := http.NewServeMux()
	handler.HandleFunc("/swagger", respondSwagger)
	handler.Handle("/", router.NewServer(storage, errorHandler))

	log.Fatal(nil, http.ListenAndServe("0.0.0.0:9999", addCors(handler)))
}

func addCors(handler http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete},
	}).Handler(handler)
}

func respondSwagger(w http.ResponseWriter, r *http.Request) {
	var (
		yamlFile, jsonFile []byte
		err                error
	)

	if yamlFile, err = ioutil.ReadFile(path.Join(os.Getenv("GOPATH"), "/src/github.com/vincentvanderweele/swagger-todo-mvc/swagger.yaml")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if jsonFile, err = yaml.YAMLToJSON(yamlFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonFile)
}
