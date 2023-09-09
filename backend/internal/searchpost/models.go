package searchpost

type Post struct {
	ID       int64  `json:"id_post"`
	UserID   int64  `json:"id_user"`
	Content  string `json:"content"`
	PostHour string `json:"postHour"`
	PostDate string `json:"postDate"`
	Location string `json:"location"`
	Labels   string `json:"labels"`
}
