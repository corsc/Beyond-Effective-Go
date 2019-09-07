package _2_thin_interface

import (
	"net/http"
)

type CityByIDLoader interface {
	LoadByID(int) (*City, error)
}

type LoadCityByIDEndpoint struct {
	loader CityByIDLoader
}

func (l *LoadCityByIDEndpoint) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	id, err := l.extractIDFromRequest(req)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	city, err := l.loader.LoadByID(id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	l.renderCity(resp, city)
}

func (l *LoadCityByIDEndpoint) extractIDFromRequest(request *http.Request) (int, error) {
	// implementation removed
	return 1, nil
}

func (l *LoadCityByIDEndpoint) renderCity(writer http.ResponseWriter, city *City) {
	// implementation removed
}

type City struct {
	ID          int
	Description string
}
