package post

type IPostRepository interface {
	Load(url string) (IResponse, error)
	SaveBatch(posts []Post) (int64, error)
	GetAll() ([]Post, error)
	Get(id int64) (Post, error)
	Update(post IPost) (int64, error)
	Delete(id int64) (int64, error)
	TableName() string
}

type IPost interface {
	GetId() int64
	GetUserId() int64
	GetTitle() string
	GetBody() string
}

type Post struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *Post) GetId() int64 {
	return p.Id
}
func (p *Post) GetUserId() int64 {
	return p.UserId
}
func (p *Post) GetTitle() string {
	return p.Title
}
func (p *Post) GetBody() string {
	return p.Body
}
