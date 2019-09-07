package _1_fat_interface

import (
	"net/http"
)

type CityModel interface {
	Save(*City) (int, error)
	Update(*City) error
	LoadByID(int) (*City, error)
	LoadAll() ([]*City, error)
}

type LoadCityByIDEndpoint struct {
	cityModel CityModel
}

func (l *LoadCityByIDEndpoint) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	id, err := l.extractIDFromRequest(req)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	city, err := l.cityModel.LoadByID(id)
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
