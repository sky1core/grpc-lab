package legacy

type MemoResponseV0 struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Priority  int32  `json:"priority"`
	CreatedAt int64  `json:"created_at"`
}
