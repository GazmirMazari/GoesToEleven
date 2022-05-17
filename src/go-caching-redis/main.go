package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Starting server :5000")

	http.HandleFunc("/api", Handler)
	http.ListenAndServe(":5000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In the handler")

	q := r.URL.Query().Get("q")
	data, err := getDate(q)
	fmt.Println(data)
	if err != nil {
		fmt.Printf("error getting the data: %v/n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := ApiResponse{
		Cache: false,
		Data:  data,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Printf("error getting the data: %v/n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func getDate(q string) ([]NominatimResponse, error) {
	escapedQ := url.PathEscape(q)
	address := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", escapedQ)

	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}

	data := make([]NominatimResponse, 0)

	erro := json.NewDecoder(resp.Body).Decode(data)
	if erro != nil {
		return nil, err
	}
	return data, nil
}

type ApiResponse struct {
	Cache bool                `json:"cache"`
	Data  []NominatimResponse `json:data`
}

type NominatimResponse struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Boundingbox []string `json:"boundingbox"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
	Icon        string   `json:"icon"`
}
