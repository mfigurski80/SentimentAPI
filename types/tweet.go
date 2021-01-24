package types

type Tweet struct {
	ID         int    `json:"id"`
	Time       int64  `json:"time"`
	CreatedAt  int64  `json:"createdAt"`
	Sentiment  string `json:"sentiment"`
	Confidence int    `json:"confidence"`
	Text       string `json:"text"`
	Username   string `json:"username"`
	Link       string `json:"link"`
}

type Point struct {
	Time     int64   `json:"time"`
	Positive int     `json:"positive"`
	Negative int     `json:"negative"`
	Retweets int     `json:"retweets"`
	Total    int     `json:"total"`
	Tweets   []Tweet `json:"tweet"`
}
