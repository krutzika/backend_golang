package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/krutzika/backend_golang/services"
)

func GetNews(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	if category == "" {
		http.Error(w, "missing category", http.StatusBadRequest)
		return
	}
	articles, err := services.FetchTopLines(category)
	if err != nil {
		http.Error(w, "Error fetching top lines", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(articles)

}
