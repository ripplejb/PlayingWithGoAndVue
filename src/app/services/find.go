package services

import (
	. "app/clients"
	"app/models"
	"encoding/xml"
	"net/url"
)

func Find(id string) (models.ClassifyBookResponse, error) {
	var body []byte
	var err error
	if body, err = ClassifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&owi=" + url.QueryEscape(id)); err != nil {
		return models.ClassifyBookResponse{}, err
	}

	var classifyBookResponse models.ClassifyBookResponse
	err = xml.Unmarshal(body, &classifyBookResponse)

	return classifyBookResponse, err

}
