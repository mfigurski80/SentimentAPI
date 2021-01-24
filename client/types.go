package client

type Tweet struct {
	ID          int
	Time        int
	TimePosted  int
	Sentiment   string
	Confidence  int
	Text        string
	Username    string
	TwitterLink string
}

type Point struct {
	Time     []uint8
	Positive int
	Negative int
	Retweets int
	Total    int
}
