package game

import (
	"net/http"
	"fmt"
	"encoding/json"
	"time"
	"math/rand"
	"log"
)

func FindGames(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %v", r.URL)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	time.Sleep(time.Duration(generateRandom(3000)) * time.Millisecond)

	title := r.URL.Query().Get("title")
	searchResults := generateRandomGameResponse(title)
	if err := json.NewEncoder(w).Encode(searchResults); err != nil {
		panic(err)
	}
}

func generateRandomGameResponse(title string) []SearchResult {
	var searchResults []SearchResult
	for i := 0; i < generateRandom(100); i++ {
		element := SearchResult{time.Now().Nanosecond(), fmt.Sprintf("%v %v",title, i)}
		searchResults = append(searchResults, element)
	}
	return searchResults;
}

func generateRandom(i int) int {
	h := rand.Int31n(1)
	if h == 0 {
		return int(rand.Int31n(int32(i/2)))
	} else {
		return int(rand.Int31n(int32(i)))
	}
}
