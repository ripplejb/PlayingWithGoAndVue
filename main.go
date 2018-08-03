package main

import (
	"html/template"
	"log"
	"net/http"

	_ "./src/github.com/mattn/go-sqlite3"
	"database/sql"
	"encoding/json"

	"./src/github.com/codegangsta/negroni"
	"encoding/xml"
	"io/ioutil"
	"net/url"
)

type Page struct {
	Name     string
	DBStatus bool
}

type SearchResult struct {
	Title  string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Year   string `xml:"hyr,attr"`
	ID     string `xml:"owi,attr"`
}

type SearchResultView struct {
	Results      []SearchResult
	Headers      []string
	ColumnWidths []string
}

type ClassifySearchResponse struct {
	Results []SearchResult `xml:"works>work"`
}

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

var db *sql.DB

func verifyDatabase(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if err := db.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	next(w, r)
}

func find(id string) (ClassifyBookResponse, error) {
	var body []byte
	var err error
	if body, err = classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&owi=" + url.QueryEscape(id)); err != nil {
		return ClassifyBookResponse{}, err
	}

	var classifyBookResponse ClassifyBookResponse
	err = xml.Unmarshal(body, &classifyBookResponse)

	return classifyBookResponse, err

}

func classifyAPI(url string) ([]byte, error) {
	var resp *http.Response
	var err error

	if resp, err = http.Get(url); err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := Page{Name: "Go Ripal"}
	if name := r.FormValue("name"); name != "" {
		p.Name = name
	}

	template := template.Must(template.ParseFiles("templates/main/index.html"))

	if err := template.ExecuteTemplate(w, "index.html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var results []SearchResult
	var err error

	if results, err = search(r.URL.Query().Get("searchInput")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	result := SearchResultView{Results: results, Headers: []string{"Title", "Author", "Year", "ID"}, ColumnWidths: []string{"40%", "30%", "10%", "20%"}}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func search(query string) ([]SearchResult, error) {
	var body []byte
	var err error

	if body, err = classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query)); err != nil {
		return []SearchResult{}, err
	}

	var c ClassifySearchResponse
	err = xml.Unmarshal(body, &c)

	return c.Results, err
}

func addBookToDB(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var result SearchResult
	err := decoder.Decode(&result)
	if err != nil {
		println(err.Error())
	} else {
		var book ClassifyBookResponse
		if book, err = find(result.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		_, err = db.Exec("insert into books (pk, title, author, id, classification) values "+
			"(?, ?, ?, ?, ?)", nil, book.BookData.Title, book.BookData.Author, book.BookData.Id, book.Classification.MostPopular)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	db, _ = sql.Open("sqlite3", "dev.db")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/books/add", addBookToDB)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates"))))

	ng := negroni.Classic()
	ng.Use(negroni.HandlerFunc(verifyDatabase))
	ng.UseHandler(mux)
	ng.Run(":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
