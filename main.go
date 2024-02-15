package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	api_url string
	title   string
	year    int
	debug   bool
)

type Movie struct {
	Title      string
	Year       string
	Rated      string
	Release    string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Ratings    interface{}
	Metascore  string
	imdbRating string
	imdbId     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

func init() {
	api_key := os.Getenv("OMDB_KEY")
	if api_key != "" {
		api_url = fmt.Sprintf("http://www.omdbapi.com/?apikey=%s", api_key)
	} else {
		fmt.Println("The env var OMDB_KEY must be set to use the api")
		os.Exit(1)
	}
}

func main() {
	flag.StringVar(&title, "title", "", "The movie title to search the OMDB for")
	flag.IntVar(&year, "year", 0, "Optional year to provide in the movie search")
	flag.BoolVar(&debug, "debug", false, "Output extra debug info")
	flag.Parse()

	if title != "" {
		var movie Movie
		search_url := api_url + fmt.Sprintf("&t=%s", strings.Replace(title, " ", "+", -1))
		if year != 0 {
			search_url += fmt.Sprintf("&y=%d", year)
		}
		if debug {
			fmt.Printf("url: %s\n", search_url)
		}
		response, err := http.Get(search_url)
		if err != nil {
			fmt.Printf("There was an error searching the OMDB api: %s\n", err)
			os.Exit(1)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Printf("Did not receive a HTTP 200 OK: %d\n", response.StatusCode)
			fmt.Printf("The film %s, %d might not exist", title, year)
			os.Exit(1)
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			os.Exit(1)
		}

		err = json.Unmarshal(body, &movie)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			os.Exit(1)
		}

		fmt.Println("Title: ", movie.Title)
		fmt.Println("Year: ", movie.Year)
		fmt.Println("Director: ", movie.Director)
		fmt.Println("Writer: ", movie.Writer)
		fmt.Println("Genre: ", movie.Genre)
		fmt.Println("Plot: ", movie.Plot)
		fmt.Println("IMDB Rating: ", movie.imdbRating)

	} else {
		fmt.Println("A title must be provided")
		os.Exit(1)
	}

}
