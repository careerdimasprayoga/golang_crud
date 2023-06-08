package models

import (
	"database/sql"
	"fmt"
	"github.com/careerdimasprayoga/golang_crud/config"
	"github.com/careerdimasprayoga/golang_crud/entities"
	"github.com/mitchellh/mapstructure"
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

func (p *PostModel) Update(post entities.Post) error {

	_, err := p.conn.Exec(
		"update posts set title = ?, content = ?, category = ?, status = ? where id = ?",
		post.Title, post.Content, post.Category, post.Status, post.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *PostModel) GetPaginatedPosts(offset, limit int) []entities.Post {
	query := "SELECT * FROM posts LIMIT ? OFFSET ?"
	rows, err := p.conn.Query(query, limit, offset)
	if err != nil {
		fmt.Println(err)
		return []entities.Post{}
	}
	defer rows.Close()

	var posts []entities.Post
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return []entities.Post{}
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(valuePtrs...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		rowData := make(map[string]interface{})
		for i, column := range columns {
			val := values[i]
			switch v := val.(type) {
			case []byte:
				rowData[column] = string(v)
			default:
				rowData[column] = v
			}
		}
		// fmt.Println(rowData)
		var post entities.Post
		err = mapstructure.Decode(rowData, &post)
		if err != nil {
			fmt.Println(err)
			continue
		}
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

func (p *PostModel) Find(id int64, post *entities.Post) error {

	return p.conn.QueryRow("select * from posts where Id = ?", id).Scan(
		&post.Id,
		&post.Title,
		&post.Content,
		&post.Category,
		&post.Status)
}

func (p *PostModel) FindAll() ([]entities.Post, error) {

	rows, err := p.conn.Query("select * from posts")
	if err != nil {
		return []entities.Post{}, err
	}
	defer rows.Close()

	var dataPost []entities.Post
	for rows.Next() {
		var post entities.Post
		rows.Scan(&post.Id,
			&post.Title,
			&post.Content,
			&post.Category,
			&post.Status)

		dataPost = append(dataPost, post)
	}

	return dataPost, nil

}