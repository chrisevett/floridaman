package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	// add context
	// "github.com/aws/aws-lambda-go/lambda"
)

/*
type GifEvent struct {
	Query string `json:"query"`
}*/

/*
func HandleRequest(ctx context.Context, name GifEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name ), nil
}*/

func main() {
	// make giphy call with api key
	// get a gif
	// get a gif based on a term
	// ZhpgEgAIcrjXG7pz9q3sun97lZF6IyXh
	apikey := "ZhpgEgAIcrjXG7pz9q3sun97lZF6IyXh"
	req, err := url.Parse("http://api.giphy.com/v1/gifs/search")
	if err != nil {
		panic("balls")
	}

	query := req.Query()
	query.Set("api_key", apikey)
	query.Set("rating", "pg-13")
	query.Set("limit", "25")
	query.Set("lang", "en")

	req.RawQuery = query.Encode()

	// response, err := http.NewRequest("GET", req.String(),nil)
	response, err := http.Get(req.String())
	if err != nil {
		panic(fmt.Sprintf("Error sending request: %s", err))
	}

	responseData, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseData))
}
