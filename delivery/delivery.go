package delivery

const (
	OkStatus    = "ok"
	ErrorStatus = "error"
)

type ApiRequest struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ApiResponse struct {
	Status  string
	Message string
	Data    interface{}
}
