package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/krutzika/backend_golang/models"
)

func FetchTopLines(category string) ([]models.Article, error) {
	api_key := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?category=%s&country=us&apiKey=%s", category, api_key)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result models.NewsResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Articles, nil
}
