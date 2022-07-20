package post

type links struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
	Next     string `json:"next"`
}

type pagination struct {
	Total int   `json:"total"`
	Pages int   `json:"pages"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Links links `json:"links"`
}

type meta struct {
	Pagination pagination `json:"pagination"`
}

type IResponse interface {
	GetNextLink() string
	GetPosts() []Post
}

type Response struct {
	Meta  meta   `json:"meta"`
	Posts []Post `json:"data"`
}

func (r Response) GetNextLink() string {
	return r.Meta.Pagination.Links.Next
}

func (r Response) GetPosts() []Post {
	return r.Posts
}
