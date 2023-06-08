package models

import (
	"database/sql"
	"fmt"
	"github.com/careerdimasprayoga/golang_crud/config"
	"github.com/careerdimasprayoga/golang_crud/entities"
)

type PostModel struct {
	conn *sql.DB
}

func NewPostModel() *PostModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PostModel{
		conn: conn,
	}
}

func (p *PostModel) Create(post entities.Post) bool {
	result, err := p.conn.Exec("INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?)",
		post.Title, post.Content, post.Category, post.Status)
	if err != nil {
		fmt.Println(err)
		return false
	}
	lastInsertID, _ := result.LastInsertId()
	return lastInsertID > 0
}

func (p *PostModel) GetPaginatedPosts(offset, limit int) []entities.Post {
	query := "SELECT * FROM posts LIMIT ?, ?"
	fmt.Println("Query:", query, offset, limit)

	rows, err := p.conn.Query(query, offset, limit)
	if err != nil {
		fmt.Println(err)
		return []entities.Post{}
	}
	defer rows.Close()

	var posts []entities.Post
	for rows.Next() {
		var post entities.Post
		rows.Scan(&post.Title, &post.Content, &post.Category, &post.Status)
		posts = append(posts, post)
	}

	return posts
}

func (p *PostModel) CountPosts() int {
	var count int
	err := p.conn.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return count
}
