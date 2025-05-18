package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/krutzika/backend_golang/services"
)

func GetNews(w http.ResponseWriter, r *http.Request) {
	country := "us"
	articles, err := services.FetchTopLines(country)
	if err != nil {
		http.Error(w, "Error fetching top lines", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(articles)

}

func GetNewsCategory(w http.ResponseWriter, r *http.Request) {
	category := "technology"
	category_articles, err := services.FetchCategoryArticles(category)
	if err != nil {
		http.Error(w, "Error fetching top lines", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(category_articles)
}

func GetSource(w http.ResponseWriter, r *http.Request) {
	source, err := services.FetchSource()
	if err != nil {
		http.Error(w, "Error fetching top lines", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(source)
}
