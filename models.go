package main

import "groupie-tracker/fetch"

type ArtistView struct {
	ID           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   string
	LocationsSum int
	LastActivity string
	Relations    map[string][]string
}

type ErrorData struct {
	Code    int
	Message string
}

func LoadArtists() error {

	artists, err := fetch.FetchArtists()
	if err != nil {
		return err
	}

	locations, err := fetch.FetchLocations()
	if err != nil {
		return err
	}

	dates, err := fetch.FetchDates()
	if err != nil {
		return err
	}

	relations, err := fetch.FetchRelation()
	if err != nil {
		return err
	}

	var result []ArtistView

	for i := range artists {

		var lastActivity string
		if len(dates[i].Dates) > 0 {
			lastActivity = dates[i].Dates[len(dates[i].Dates)-1]
		}

		view := ArtistView{
			ID:           artists[i].ID,
			Name:         artists[i].Name,
			Image:        artists[i].Image,
			Members:      artists[i].Members,
			CreationDate: artists[i].CreationDate,
			FirstAlbum:   artists[i].FirstAlbum,
			LocationsSum: len(locations[i].Locations),
			LastActivity: lastActivity,
			Relations:    relations[i].DatesLocations,
		}
		result = append(result, view)
	}
	artistsCache = result
	return nil
}
