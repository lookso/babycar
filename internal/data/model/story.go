package model

// Comment: 故事
type Story struct {
	ID         int    `db:"id"`
	Tag        int    `db:"tag"`
	Title      string `db:"title"`
	Status     int    `db:"status"`
	Content    string `db:"content"`
	SourceURL  string `db:"source_url"`
	CreateTime int    `db:"create_time"`
	UpdateTime int    `db:"update_time"`
	CurlTime   int    `db:"curl_time"`
}

// TableName: story

func (s *Story) TableName() string {
	return "storys"
}

