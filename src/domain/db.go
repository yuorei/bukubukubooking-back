package domain

import (
	"time"
)

// USERS テーブルに対応する構造体
type DBUser struct {
	ID        string    `json:"id"`         // slackのuserID
	Name      string    `json:"name"`       // not null
	Email     string    `json:"email"`      // unique
	CreatedAt time.Time `json:"created_at"` // not null
	UpdatedAt time.Time `json:"updated_at"` // not null
	DeletedAt time.Time `json:"deleted_at"`
}

// BOOKS テーブルに対応する構造体
type DBBook struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"` // not null
	Publisher string    `json:"publisher"`
	Owner     string    `json:"owner"`      // not null, UserのID
	Status    string    `json:"status"`     // not null
	CreatedAt time.Time `json:"created_at"` // not null
	UpdatedAt time.Time `json:"updated_at"` // not null
	DeletedAt time.Time `json:"deleted_at"`
}

// TAGS テーブルに対応する構造体
type DBTag struct {
	ID   string `json:"id"`
	Name string `json:"name"` // not null, unique
}

// BOOKS_TAGS テーブルに対応する構造体
type DBBookTagd struct {
	ID     int    `json:"id"`
	BookID string `json:"book_id"` // FK, BookのID
	TagID  string `json:"tag_id"`  // FK, TagのID
}

// AUTHORS テーブルに対応する構造体
type DBAuthor struct {
	ID   string `json:"id"`
	Name string `json:"name"` // not null, unique
}

// BOOKS_AUTHORS テーブルに対応する構造体
type DBBookAuthor struct {
	ID       int    `json:"id"`
	BookID   string `json:"book_id"`   // FK, BookのID
	AuthorID string `json:"author_id"` // FK, AuthorのID
}

// BOOK_LIKES テーブルに対応する構造体
type DBBookLike struct {
	ID        string    `json:"id"`
	BookID    string    `json:"book_id"` // FK, BookのID
	UserID    string    `json:"user_id"` // FK, UserのID
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"` // not null
	UpdatedAt time.Time `json:"updated_at"` // not null
}

// MESSAGES テーブルに対応する構造体
type DBMessage struct {
	ID        string    `json:"id"`
	RequestID string    `json:"request_id"` // FK, RequestのID
	SenderID  string    `json:"sender_id"`  // FK, UserのID
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"` // not null
	UpdatedAt time.Time `json:"updated_at"` // not null
}

// REQUESTS テーブルに対応する構造体
type DBRequest struct {
	ID             string    `json:"id"`
	BookID         string    `json:"book_id"`          // FK, BookのID
	SenderUserID   string    `json:"sender_user_id"`   // FK, UserのID
	ReceiverUserID string    `json:"receiver_user_id"` // FK, UserのID
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"` // not null
	UpdatedAt      time.Time `json:"updated_at"` // not null
	FinishedAt     time.Time `json:"finished_at"`
}
