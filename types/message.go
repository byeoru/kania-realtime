package types

type Request struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}
type Response struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
	Body  []byte `json:"body"`
}
