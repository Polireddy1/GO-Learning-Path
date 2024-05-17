package main

import (
	"encoding/json" //For parsing JSON data from the API response
	"fmt"
	"io"       //For reading the response body
	"net/http" //for handling http requests and to implement web services
)

// Struct to hold the joke data
type Joke struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	ID        int    `json:"id"`
}

func main() {
	//index about and jokes are the handler functions
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/jokes", jokes)
	fmt.Println("Server Starting....on port 3000")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World how are you!</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About Page</h1>")
}

func jokes(w http.ResponseWriter, r *http.Request) {
	jokes, err := fetchJokes()
	if err != nil {
		http.Error(w, "Failed to fetch jokes", http.StatusInternalServerError)
		return
	}

	// Display the jokes
	fmt.Fprintf(w, `<h1 style="background-color:yellow; text-align:center;">JOKES</h1>`)
	for _, joke := range jokes {
		fmt.Fprintf(w, "<h2>%s</h2><p>%s</p>", joke.Setup, joke.Punchline)
	}
}

// Fetch data from the jokes API
func fetchJokes() ([]Joke, error) {
	apiURL := "https://official-joke-api.appspot.com/random_ten"
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //deferred function call is executed just before the surrounding function returns
	//resp.Body.Close() method is called when the fetchJokes function returns.

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data
	var jokes []Joke
	err = json.Unmarshal(body, &jokes)
	if err != nil {
		return nil, err
	}

	return jokes, nil
}

// https://official-joke-api.appspot.com/random_ten
