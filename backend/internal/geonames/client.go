package geonames

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type GeonamesClient struct {
	baseURL    string
	username   string
	httpClient *http.Client
}

type GeoPlace struct {
	Name string `json:"name"`
}

type SearchResponse struct {
	TotalResultsCount int        `json:"totalResultsCount"`
	Geonames          []GeoPlace `json:"geonames"`
}

func NewGeonamesClient(
	baseURL string,
	username string,
	httpClient *http.Client,
) *GeonamesClient {
	return &GeonamesClient{
		baseURL:    baseURL,
		username:   username,
		httpClient: httpClient,
	}
}

func (c *GeonamesClient) WithBaseParams(
	maxRows int,
	nameStartsWith *string,
) url.Values {
	params := url.Values{}

	params.Set("featureClass", "P")
	params.Set("orderby", "population")
	params.Set("username", c.username)
	params.Set("maxRows", strconv.Itoa(maxRows))

	if nameStartsWith != nil {
		params.Set("name_startsWith", *nameStartsWith)
	}

	return params
}

func (client *GeonamesClient) Search(
	ctx context.Context,
	params url.Values,
) (SearchResponse, error) {
	requestURL := client.baseURL + "?" + params.Encode()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		requestURL,
		nil,
	)

	if err != nil {
		return SearchResponse{}, fmt.Errorf(
			"create GeoNames request: %w",
			err,
		)
	}

	response, err := client.httpClient.Do(req)

	if err != nil {
		return SearchResponse{}, fmt.Errorf(
			"request GeoNames: %w",
			err,
		)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return SearchResponse{}, fmt.Errorf(
			"GeoNames returned status %d",
			response.StatusCode,
		)
	}

	var result SearchResponse

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return SearchResponse{}, fmt.Errorf(
			"decode GeoNames response: %w",
			err,
		)
	}

	return result, nil
}
