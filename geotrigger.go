// Package geotrigger provides an API to interact with ESRI's geotrigger service
package geotrigger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"code.google.com/p/goauth2/oauth"
)

var baseURL = "https://geotrigger.arcgis.com"

type API struct {
	t *oauth.Transport
}

func NewAPI(clientID, clientSecret string) (*API, error) {
	a := &API{
		&oauth.Transport{
			Config: &oauth.Config{
				ClientId:     clientID,
				ClientSecret: clientSecret,
				TokenURL:     "https://www.arcgis.com/sharing/oauth2/token",
			},
		}}
	return a, a.getToken()
}

func (api *API) getToken() error {
	v := url.Values{}
	v.Set("client_id", api.t.ClientId)
	v.Set("client_secret", api.t.ClientSecret)
	v.Set("grant_type", "client_credentials")
	v.Set("f", "json")

	resp, err := http.PostForm(api.t.TokenURL, v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var tokenResponse struct {
		AccessToken string  `json:"access_token"`
		ExpiresIn   float64 `json:"expires_in"`
	}

	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return err
	}
	api.t.Token = &oauth.Token{
		AccessToken: tokenResponse.AccessToken,
		Expiry:      time.Now().Add(time.Second),
	}
	return nil
}

func (api *API) get(endpointURL string, target interface{}) error {
	r, err := api.t.Client().Get(baseURL + endpointURL)
	if err != nil {
		return err
	}
	// @todo do refreshing, other error checking?
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&target)
}

func (api *API) post(endpointURL string, data interface{}) error {
	payload, merr := json.Marshal(data)
	if merr != nil {
		return merr
	}
	r, err := api.t.Client().Post(baseURL+endpointURL, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	log.Println("post:", string(body), r.Status)

	// attempt error unmarshal
	var apiErr errorResponse
	err = json.Unmarshal(body, apiErr)
	if apiErr.Error.Code != 0 {
		return apiErr.Error
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("geotrigger: Unexpected return code: %d. Body: %s", r.StatusCode, string(body))
	}
	return nil
}
