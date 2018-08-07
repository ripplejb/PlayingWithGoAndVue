package services

import (
	. "app/clients"
	"app/models"
	"encoding/xml"
	"net/url"
)

func SearchClassifyAPI(query string) (models.SearchResultView, error) {
	var body []byte
	var err error

	if body, err = RestClientGet("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query)); err != nil {
		return models.SearchResultView{}, err
	}

	var c models.ClassifySearchResponse
	err = xml.Unmarshal(body, &c)

	result := models.SearchResultView{Results: c.Results, Headers: []string{"Title", "Author", "Year", "ID"}, ColumnWidths: []string{"40%", "30%", "10%", "20%"}}

	return result, err
}
