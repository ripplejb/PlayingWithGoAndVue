package models

type ClassifySearchResponse struct {
	Results []SearchResult `xml:"works>work"`
}
