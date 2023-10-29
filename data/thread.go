package data

import "time"

type Thread struct {
	ID        int64
	Uuid      string
	Topic     string
	CreatedBy int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Post struct {
	ID        int64
	Uuid      string
	Body      string
	CreatedBy int64
	ThreadID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func Threads() (threads []*Thread, err error) {
	rows, err := Db.Query(`SELECT id, uuid, topic, created_by, created_at, updated_at FROM threads ORDER BY created_at DESC`)
	if err != nil {
		return
	}
	for rows.Next() {
		conv := &Thread{}
		if err = rows.Scan(conv); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	return
}
