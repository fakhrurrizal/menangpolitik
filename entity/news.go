package entity

type NewsDetail struct {
	Data    NewsData `json:"data"`
	Message string   `json:"message"`
	Status  int64    `json:"status"`
}

type NewsData struct {
	ID          int64        `json:"id"`
	Author      string       `json:"author"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Body        string       `json:"body"`
	Keyword     string       `json:"keyword"`
	Image       string       `json:"image"`
	Status      int64        `json:"status"`
	Slug        string       `json:"slug"`
	Tags        []string     `json:"tags"`
	CreatedAt   string       `json:"created_at"`
	Category    NewsCategory `json:"category"`
	UpdatedAt   string       `json:"updated_at"`
	PublishAt   string       `json:"publish_at"`
	Shared      int64        `json:"number_of_shares"`
}

type NewsCategory struct {
	ID    int64  `json:"id"`
	Label string `json:"string"`
}
