package image

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"net/url"
)

func GiphyImage() {
	// make giphy call with api key
	// get a gif
	// get a gif based on a term
	// ZhpgEgAIcrjXG7pz9q3sun97lZF6IyXh
	apikey := "gtxwg0swllyvXaihOpVAxPdl8Av0ibby"
	req, err := url.Parse("http://api.giphy.com/v1/gifs/search")
	if err != nil {
		panic("balls")
	}

	query := req.Query()
	query.Set("api_key", apikey)
	query.Set("rating", "pg-13")
	query.Set("q", "kitty")
	query.Set("limit", "25")
	query.Set("lang", "en")

	req.RawQuery = query.Encode()

	// response, err := http.NewRequest("GET", req.String(),nil)
	fmt.Println(req.String())
	response, err := http.Get(req.String())
	if err != nil {
		panic(fmt.Sprintf("Error sending request: %s", err))
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(fmt.Sprintf("Error reading response body: %s", err))
	}

	giphyResponse := GiphyResponse{}

	_ = json.Unmarshal([]byte(responseData), &giphyResponse)

	// each data[] is a different image, just return the first one
	fmt.Println("preview gif: ", giphyResponse.Data[0].URL)
}
