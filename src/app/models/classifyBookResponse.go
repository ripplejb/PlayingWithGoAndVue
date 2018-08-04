package models

type ClassifyBookResponse struct {
	BookData struct {
		Title  string `xml:"title,attr"`
		Author string `xml:"author,attr"`
		Id     string `xml:"owi,attr"`
	} `xml:"work"`

	Classification struct {
		MostPopular string `xml:"sfa,attr"`
	} `xml:"recommendation>ddc>mostPopular"`
}
