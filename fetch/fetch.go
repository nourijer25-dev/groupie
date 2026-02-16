package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const Url = "https://groupietrackers.herokuapp.com/api/"

type Artist struct {
	ID           int
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
type LocationResponse struct {
	Index []Locations `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DatesResponse struct {
	Index []Dates `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationResponse struct {
	Index []Relation `json:"index"`
}

func fetch(endpoint string, target interface{}) error {
	resp, err := http.Get(Url + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	// Decode the JSON response into the artists slice
	return json.NewDecoder(resp.Body).Decode(target)
}

// FetchArtists fetches the list of artists from the external API.
func FetchArtists() ([]Artist, error) {

	var artists []Artist
	err := fetch("artists", &artists)
	return artists, err
}

func FetchLocations() ([]Locations, error) {

	var location LocationResponse
	err := fetch("locations", &location)
	return location.Index, err
}

func FetchDates() ([]Dates, error) {

	var dates DatesResponse
	err := fetch("dates", &dates)
	return dates.Index, err
}

func FetchRelation() ([]Relation, error) {

	var relation RelationResponse
	err := fetch("relation", &relation)
	return relation.Index, err
}
