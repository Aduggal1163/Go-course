package post

import (
	"fmt"

	"example.com/user"
)

type Post struct {
	title   string
	content string
	author  user.User
}

func NewPost(title string, content string, author user.User) *Post {
	return &Post{
		title:   title,
		content: content,
		author:  author,
	}
}

func (p *Post) OutputDetails() {
	fmt.Println("Post title:", p.title)
	fmt.Println("Content:", p.content)
	fmt.Println("Post by", p.author.FullName())
}
