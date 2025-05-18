package models

type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type SourceResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

type Source struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	URL      string `json:"url"`
	Category string `json:"category"`
	Country  string `json:"country"`
}
