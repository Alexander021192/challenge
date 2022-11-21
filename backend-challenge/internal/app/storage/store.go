package storage

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	challenge "github.com/Alexander021192/challenge/backend-challenge/pkg"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

const (
	databaseURL = "host=localhost dbname=challenge_test sslmode=disable"
)

// New ...
func NewStorage() (*Storage, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) CreateUser(u *User) (int32, error) {
	if err := u.Validate(); err != nil {
		return 0, err
	}

	if err := u.BeforeCreate(); err != nil {
		return 0, err
	}

	if err := s.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (s *Storage) CreateComment(c *Comment) (int32, error) {
	// if err := c.Validate(); err != nil {
	// 	return 0, err
	// }

	if err := s.db.QueryRow(
		"INSERT INTO comments (author, author_img, date, comment, comment_img) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		c.Author,
		c.ProfileImg,
		c.Date,
		c.Comment,
		c.CommentImg,
	).Scan(&c.ID); err != nil {
		return 0, err
	}
	return c.ID, nil
}

func (s *Storage) GetComments() ([]*challenge.Comment, error) {
	var listComments []*challenge.Comment

	rows, err := s.db.Query("SELECT author,author_img,date,comment,comment_img FROM comments ORDER BY id DESC")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		object := &challenge.Comment{}
		errScan := rows.Scan(
			&object.Author,
			&object.AuthorImg,
			&object.Date,
			&object.Comment,
			&object.CommentImg)
		if errScan != nil {
			fmt.Println(errScan)
			return nil, errScan
		}
		listComments = append(listComments, object)
	}
	return listComments, nil
}

func (s *Storage) CreatePost(p *Post) (int32, error) {
	// if err := c.Validate(); err != nil {
	// 	return 0, err
	// }

	if err := s.db.QueryRow(
		"INSERT INTO posts (author, title, location, post_text, post_img) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		p.Author,
		p.Title,
		p.Location,
		p.PostText,
		p.PostImg,
	).Scan(&p.ID); err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (s *Storage) GetPosts() ([]*challenge.Post, error) {
	var listPosts []*challenge.Post

	rows, err := s.db.Query("SELECT author,title,location,post_text,post_img,id FROM posts ORDER BY id")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		object := &challenge.Post{}
		errScan := rows.Scan(
			&object.Author,
			&object.Title,
			&object.Location,
			&object.PostText,
			&object.PostImg,
			&object.Id)
		if errScan != nil {
			fmt.Println(errScan)
			return nil, errScan
		}
		listPosts = append(listPosts, object)
	}
	return listPosts, nil
}

func (s *Storage) GetPostById(postId string) (*challenge.Post, error) {
	var post *challenge.Post

	rows, err := s.db.Query("SELECT author,title,location,post_text,post_img,id FROM posts WHERE id = $1", postId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		object := &challenge.Post{}
		errScan := rows.Scan(
			&object.Author,
			&object.Title,
			&object.Location,
			&object.PostText,
			&object.PostImg,
			&object.Id)
		if errScan != nil {
			fmt.Println(errScan)
			return nil, errScan
		}
		post = object
	}
	return post, nil
}

func (s *Storage) FindByEmail(email string) (*User, error) {
	u := &User{}
	if err := s.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Storage) FindBySession(sessionId string) (*User, error) {
	u := &User{}
	if err := s.db.QueryRow(
		"SELECT profile_name, profile_img FROM users WHERE sessionId = $1",
		sessionId,
	).Scan(
		&u.ProfileName,
		&u.ProfileImg,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Storage) SessionSave(email string) (string, error) {
	sessionId, err := encryptString(email)
	if err != nil {
		return "", err
	}
	_, err = s.db.Exec("UPDATE users SET sessionId = $1 where email = $2",
		sessionId, email)
	if err != nil {
		return "", err
	}

	// need run goroutine witd deadline for update sessionId to none
	go func(email string) {
		fmt.Println(email)
		time.Sleep(time.Minute * 10)
		_, err = s.db.Exec("UPDATE users SET sessionId = 'none' where email = $1", email)
		// need log error
		if err != nil {
			fmt.Println(err)
		}
	}(email)

	return sessionId, nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil

}
