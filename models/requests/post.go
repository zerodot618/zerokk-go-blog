package requests

type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryID int    `json:"categoryId"`
	UserID     int    `json:"userId"`
	Type       int    `json:"type"`
}
