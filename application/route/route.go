package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}

	fileURL := "https://raw.githubusercontent.com/codeedu/imersao8/main/simulator/destinations/"
	resp, err := http.Get(fileURL + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	s := bufio.NewScanner(resp.Body)
	for s.Scan() {
		data := strings.Split(s.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}

	return nil
}

func (r *Route) ExportJSONPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}
		routeJSON, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(routeJSON))
	}

	return result, nil
}
