package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/krutzika/backend_golang/models"
)

//var api_key = os.Getenv("NEWS_API_KEY")

func FetchTopLines(country string) ([]models.Article, error) {
	api_key := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s", country, api_key)

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

func FetchCategoryArticles(category string) ([]models.Article, error) {
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

func FetchSource() ([]models.Source, error) {
	api_key := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines/sources?apiKey=%s", api_key)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result models.SourceResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Sources, nil
}
