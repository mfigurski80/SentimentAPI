package schema

type Tweet struct {
	ID         int    `json:"id"`
	Time       int    `json:"time"`
	CreatedAt  int    `json:"createdAt"`
	Sentiment  string `json:"sentiment"`
	Confidence int    `json:"confidence"`
	Text       string `json:"text"`
	Username   string `json:"username"`
	Link       string `json:"link"`
}

type Point struct {
	Time     int     `json:"time"`
	Positive int     `json:"positive"`
	Negative int     `json:"negative"`
	Retweets int     `json:"retweets"`
	Total    int     `json:"total"`
	Tweets   []Tweet `json:"tweet"`
}