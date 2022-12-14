package main

import (
	"api-star-wars-golang/internal/provider/mongo/dao"
	"context"
	"fmt"
	"log"
	"net/http"

	"api-star-wars-golang/internal/router"
	"api-star-wars-golang/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Servidor esta rodando na porta 8080")

	client, database, err := getDatabase()
	if err != nil {
		log.Fatal("error get database")
	}

	dao := dao.NewMongoPlanet(client, database)
	swapi := service.NewSWAPI()
	service := service.NewPlanet(dao, swapi)
	handler := router.NewPlanetHandler(service)
	healthHandler := router.NewHealthHandler(dao)

	router := mux.NewRouter()
	router.HandleFunc("/planets/{id}", handler.DeleteById).Methods("DELETE")
	router.HandleFunc("/planets/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/planets", handler.SavePlanet).Methods("POST")
	router.HandleFunc("/planets/{id}", handler.FindById).Methods("GET")
	router.HandleFunc("/planets", handler.FindByParam).Methods("GET").Queries()
	router.HandleFunc("/health", healthHandler.Healthcheck).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getDatabase() (*mongo.Client, *mongo.Database, error) {
	// usar essa linha a baixo quando for utilizar o docker file e apagar a linha posterior
	//clientOptions := options.Client().ApplyURI("mongodb://mongo-star")

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, err
	}
	return client, client.Database("star-wars"), nil
}
