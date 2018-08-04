package services

import (
	. "app/clients"
	"app/models"
	"encoding/xml"
	"net/url"
)

func Search(query string) ([]models.SearchResult, error) {
	var body []byte
	var err error

	if body, err = ClassifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query)); err != nil {
		return []models.SearchResult{}, err
	}

	var c models.ClassifySearchResponse
	err = xml.Unmarshal(body, &c)

	return c.Results, err
}
