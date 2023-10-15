package car

type Comment struct {
	AppName     string `json:"app_name"`
	TraceId     string `json:"trace_id"`
	SessionId   string `json:"session_id"`
	UserId      string `json:"user_id"`
	TypeChinese string `json:"type_chinese"`
	CreatedAt   string `json:"created_at"`
	Extra       string `json:"extra"`
}
