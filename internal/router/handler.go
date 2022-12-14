package router

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"api-star-wars-golang/internal/model"
	"api-star-wars-golang/internal/service"
	"time"

	validator "github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

type PlanetHandler struct {
	service service.Planet
}

type HealthChecker interface {
	Check(ctx context.Context) error
}

type HealthHandler struct {
	hc HealthChecker
}

// retornei a struct pq dentro dela esta quem implementa a interface e o metodo check da interface e iplementado no dao
func NewHealthHandler(hc HealthChecker) *HealthHandler {
	return &HealthHandler{hc: hc}
}

func NewPlanetHandler(service service.Planet) *PlanetHandler {
	return &PlanetHandler{service: service}
}

func (p *PlanetHandler) SavePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	//Unmarsshal vc transforma o json em objeto
	// Marshal vc trasforma o objeto em json
	var in model.PlanetIn
	err = json.Unmarshal(body, &in)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	error := ValidateStruct(&in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	hexId, err := p.service.Save(context.Background(), &in)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	url := "http://localhost:8080/planets/" + hexId
	w.Header().Add("location", url)
	w.WriteHeader(http.StatusCreated)

}

func (p *PlanetHandler) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	planet, err := p.service.FindById(context.Background(), vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(planet)
	w.WriteHeader(http.StatusOK)
}

func (p *PlanetHandler) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var planetIn model.PlanetIn
	json.NewDecoder(r.Body).Decode(&planetIn)

	//if err != nil {
	//log.Println("Error Decoding the planet", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	error := ValidateStruct(&planetIn)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	planet := p.service.Update(context.Background(), planetIn, vars["id"])

	if planet != nil {
		log.Println("Error updating the planet", planet)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *PlanetHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := p.service.DeleteById(context.Background(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
func (p *PlanetHandler) FindByParam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	planetName := r.URL.Query().Get("name")
	planetClimate := r.URL.Query().Get("climate")
	planetTerrain := r.URL.Query().Get("terrain")

	planets, err := p.service.FindByParam(context.Background(), &model.PlanetIn{
		Name:    planetName,
		Climate: planetClimate,
		Terrain: planetTerrain,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(planets)
	w.WriteHeader(http.StatusOK)
	return

}
func (p *HealthHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	error := p.hc.Check(ctx)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func ValidateStruct(v *model.PlanetIn) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}

func GetDatabase() error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	errr := client.Ping(ctx, nil)
	if errr != nil {
		return errr
	}
	return nil

}
