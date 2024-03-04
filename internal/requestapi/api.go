package requestapi

import (
	"encoding/json"
	"errors"
	"groupie-tracker/nyeltay/algaliyev/models"
	"io/ioutil"
	"net/http"
)

type RequestAPI struct{}

func New() *RequestAPI {
	return &RequestAPI{}
}

func (r *RequestAPI) Api() error {
	artistsAPI := "https://groupietrackers.herokuapp.com/api/artists"
	err := r.unmarshalJSON(artistsAPI, &models.Artists)
	if err != nil {
		return err
	}

	return nil
}

func (r *RequestAPI) ApiArtist(s string) error {
	artistAPI := "https://groupietrackers.herokuapp.com/api/artists/" + s
	err := r.unmarshalJSON(artistAPI, &models.ArtistOne)
	if err != nil {
		return err
	}

	relationAPI := models.ArtistOne.RelationAPI
	err = r.unmarshalJSON(relationAPI, &models.RelationOne)
	if err != nil {
		return err
	}

	models.ArtistOne.RelationPerArtist = models.RelationOne.DatesLocations

	return nil
}

func (r *RequestAPI) unmarshalJSON(s string, v interface{}) error {
	response, err := http.Get(s)
	if err != nil {
		return errors.New("Error getting API: " + err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("API request fails with status: " + response.Status)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("Error reading API: " + err.Error())
	}

	err = json.Unmarshal(responseData, v)
	if err != nil {
		return errors.New("Error unmarshaling JSON: " + err.Error())
	}

	return nil
}
